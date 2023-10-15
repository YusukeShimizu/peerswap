package lnd

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"math/rand"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/elementsproject/peerswap/lightning"
	"github.com/elementsproject/peerswap/onchain"
	"github.com/elementsproject/peerswap/swap"
	"github.com/elementsproject/peerswap/test"
	"github.com/elementsproject/peerswap/testframework"
	"github.com/stretchr/testify/require"
)

type Testthing struct {
	confirmedChan chan string
}

func (t *Testthing) callback(swapId string) error {
	log.Printf("callback caleld")
	t.confirmedChan <- swapId
	return nil
}

// randomString returns a random 32 byte random string
func randomString() string {
	idBytes := make([]byte, 32)
	_, _ = rand.Read(idBytes[:])
	return hex.EncodeToString(idBytes)
}

type EstimatorMock struct {
	EstimateFeePerKWCalled int
	EstimateFeePerKWReturn btcutil.Amount
	EstimateFeePerKWError  *error
}

func (e *EstimatorMock) EstimateFeePerKW(targetBlocks uint32) (btcutil.Amount, error) {
	e.EstimateFeePerKWCalled++
	if e.EstimateFeePerKWError != nil {
		return e.EstimateFeePerKWReturn, *e.EstimateFeePerKWError
	}
	return e.EstimateFeePerKWReturn, nil
}

func (e *EstimatorMock) Start() error {
	panic("not implemented") // We dont need this function.
}

func Test_LndSystemsPreimage(t *testing.T) {
	testDir := t.TempDir()
	// Get PeerSwap plugin path and test dir
	_, filename, _, _ := runtime.Caller(0)
	pathToPlugin := filepath.Join(filename, "..", "..", "out", "test-builds", "peerswapd")
	// Setup nodes (1 bitcoind, 2 lightningd, 2 peerswapd)
	bitcoind, err := testframework.NewBitcoinNode(testDir, 1)
	if err != nil {
		t.Fatalf("could not create bitcoind %v", err)
	}
	t.Cleanup(bitcoind.Kill)
	extraConfig := map[string]string{"protocol.wumbo-channels": "true"}
	lightningd, err := testframework.NewLndNode(testDir, bitcoind, 1, extraConfig)
	if err != nil {
		t.Fatalf("could not create liquidd %v", err)
	}
	t.Cleanup(lightningd.Kill)
	peerswapd, err := test.NewPeerSwapd(testDir, pathToPlugin, &test.LndConfig{LndHost: fmt.Sprintf("localhost:%d", lightningd.RpcPort), TlsPath: lightningd.TlsPath, MacaroonPath: lightningd.MacaroonPath}, nil, 1)
	if err != nil {
		t.Fatalf("could not create peerswapd %v", err)
	}
	t.Cleanup(peerswapd.Kill)

	err = bitcoind.Run(true)
	if err != nil {
		t.Fatalf("bitcoind.Run() got err %v", err)
	}
	err = lightningd.Run(true, true)
	require.NoError(t, err)
	err = peerswapd.Run(true)
	if err != nil {
		t.Fatalf("peerswapd.Run() got err %v", err)
	}
	err = peerswapd.WaitForLog("peerswapd grpc listening on", testframework.TIMEOUT)
	if err != nil {
		t.Fatalf("peerswapd.WaitForLog() got err %v", err)
	}
	err = bitcoind.GenerateBlocks(1)
	if err != nil {
		t.Fatal(err)
	}
	_, err = lightningd.FundWallet(uint64(math.Pow10(9)), true)
	if err != nil {
		t.Fatalf("lightningd.FundWallet() %v", err)
	}

	btcOnchain := onchain.NewBitcoinOnChain(
		&EstimatorMock{},
		btcutil.Amount(253),
		&chaincfg.RegressionNetParams,
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Setup lnd client.
	swapLnd, err := NewClient(
		ctx,
		lightningd.LndRpcClient.Conn,
		nil,
		nil,
		btcOnchain,
	)
	if err != nil {
		t.Fatal(err)
	}
	txParams := NewTxParams(uint32(100), 10000)
	txParams.SwapAmount = 10000
	openingParams := &swap.OpeningParams{
		TakerPubkey:      hex.EncodeToString(txParams.AliceKey.PubKey().SerializeCompressed()),
		MakerPubkey:      hex.EncodeToString(txParams.BobKey.PubKey().SerializeCompressed()),
		ClaimPaymentHash: hex.EncodeToString(txParams.PaymentHash),
		Amount:           txParams.SwapAmount,
	}

	unfinishedTxHex, _, _, err := swapLnd.CreateOpeningTransaction(openingParams)
	if err != nil {
		t.Fatal(err)
	}
	txId, openingTxHex, err := swapLnd.BroadcastOpeningTx(unfinishedTxHex)
	if err != nil {
		t.Fatal(err)
	}

	err = bitcoind.GenerateBlocks(1)
	if err != nil {
		t.Fatal(err)
	}
	claimParams := &swap.ClaimParams{
		Preimage:     txParams.Preimage.String(),
		Signer:       &swap.Secp256k1Signer{Key: txParams.AliceKey},
		OpeningTxHex: openingTxHex,
	}

	wantScript, err := btcOnchain.GetOutputScript(openingParams)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("scriptpubkey %s", hex.EncodeToString(wantScript))

	log.Printf("opening txid: %s", txId)
	err = bitcoind.GenerateBlocks(1)
	if err != nil {
		t.Fatal(err)
	}

	valid, err := btcOnchain.ValidateTx(openingParams, openingTxHex)
	if err != nil {
		t.Fatal(err)
	}
	if !valid {
		t.Fatal("tx not valid")
	}
	spendingTxId, _, err := swapLnd.CreatePreimageSpendingTransaction(openingParams, claimParams)
	if err != nil {
		t.Fatal(err)
	}
	swapLnd.CreateCoopSpendingTransaction(openingParams, claimParams, claimParams.Signer)
	swapLnd.CreateCsvSpendingTransaction(openingParams, claimParams)

	log.Printf("spending txid: %s", spendingTxId)

}

type TxParams struct {
	AliceKey    *btcec.PrivateKey
	BobKey      *btcec.PrivateKey
	Preimage    lightning.Preimage
	PaymentHash []byte
	SwapAmount  uint64
	Csv         uint32
}

func getRandomPrivkey() *btcec.PrivateKey {
	privkey, err := btcec.NewPrivateKey()
	if err != nil {
		return nil
	}
	return privkey
}

func NewTxParams(csv uint32, swapAmount uint64) *TxParams {
	preimage, _ := lightning.GetPreimage()
	pHash := preimage.Hash()
	return &TxParams{
		AliceKey:    getRandomPrivkey(),
		BobKey:      getRandomPrivkey(),
		Preimage:    preimage,
		PaymentHash: pHash[:],
		Csv:         csv,
		SwapAmount:  swapAmount,
	}
}
