package txwatcher

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/checksum0/go-electrum/electrum"
)

type TxWatcher interface {
	AddWaitForConfirmationTx(swapId, txId string, vout, startingHeight uint32, scriptpubkey []byte)
	AddWaitForCsvTx(swapId, txId string, vout uint32, startingHeight uint32, scriptpubkey []byte)
	AddConfirmationCallback(func(swapId string, txHex string, err error) error)
	AddCsvCallback(func(swapId string) error)
	GetBlockHeight() (uint32, error)
	StartWatchingTxs() error
}

type ElectrumTxWatcher struct {
	electrumClient       *electrum.Client
	blockheight          *electrum.SubscribeHeadersResult
	waitingConfirmation  []waitingConfirmation
	requiredConfs        uint32
	confirmationCallback func(swapId string, txHex string, err error) error
	csvCallback          func(swapId string) error
}

func NewElectrumTxWatcher(electrumClient *electrum.Client) (*ElectrumTxWatcher, error) {
	bcr := &ElectrumTxWatcher{electrumClient: electrumClient}
	return bcr, nil
}

var ticker = time.NewTicker(2 * time.Second)

func (r *ElectrumTxWatcher) StartWatchingTxs() error {
	ctx := context.Background()
	headerSubscription, err := r.electrumClient.SubscribeHeaders(ctx)
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Checking for confirmations")
				for _, w := range r.waitingConfirmation {
					fmt.Println("Checking for confirmation of tx:", w.txId, "vout:", w.vout, "startingHeight:", w.startingHeight, "scriptHash:", w.scriptHash)
					hs, err := r.electrumClient.GetHistory(ctx, w.scriptHash)
					if err != nil {
						fmt.Println("failed to get history:", err)
					}
					fmt.Println("History:", hs)
					for _, h := range hs {
						fmt.Println("Checking history:", h.Hash, h.Height, r.blockheight.Height+int32(r.requiredConfs))
						if h.Hash == w.txId {
							if r.blockheight.Height > h.Height+int32(r.requiredConfs) {
								rawTx, err := r.electrumClient.GetRawTransaction(ctx, h.Hash)
								if err != nil {
									fmt.Println("failed to get history:", err)
								}
								r.confirmationCallback(w.swapId, rawTx, nil)
							}
						}
					}
				}
			case blockHeader := <-headerSubscription:
				r.blockheight = blockHeader
				fmt.Println("New block header received:", blockHeader)
			case <-ctx.Done():
				fmt.Println("Context cancelled, stopping loop")
				return
			}
		}
	}()
	return nil
}

type waitingConfirmation struct {
	swapId         string
	txId           string
	vout           uint32
	startingHeight uint32
	scriptHash     string
}

func (r *ElectrumTxWatcher) AddWaitForConfirmationTx(swapId, txId string, vout, startingHeight uint32, scriptpubkey []byte) {
	hash := sha256.Sum256(scriptpubkey)
	reversedHash := make([]byte, len(hash))
	for i, b := range hash {
		reversedHash[len(hash)-1-i] = b
	}
	h := fmt.Sprintf("%X", reversedHash)
	r.waitingConfirmation = append(r.waitingConfirmation, waitingConfirmation{
		swapId:         swapId,
		txId:           txId,
		vout:           vout,
		startingHeight: startingHeight,
		scriptHash:     h,
	})
}

func (r *ElectrumTxWatcher) AddConfirmationCallback(f func(swapId string, txHex string, err error) error) {
	r.confirmationCallback = f
}
func (r *ElectrumTxWatcher) AddCsvCallback(f func(swapId string) error) {
	r.csvCallback = f
}

func (r *ElectrumTxWatcher) GetBlockHeight() (uint32, error) {
	if r.blockheight == nil {
		return 0, fmt.Errorf("blockheight not available")
	}
	return uint32(r.blockheight.Height), nil
}

func (r *ElectrumTxWatcher) AddWaitForCsvTx(swapId, txId string, vout uint32, startingHeight uint32, scriptpubkey []byte) {

}
