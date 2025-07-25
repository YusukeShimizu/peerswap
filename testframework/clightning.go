package testframework

import (
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"time"

	"github.com/elementsproject/glightning/glightning"
	"github.com/elementsproject/peerswap/clightning"
	"github.com/elementsproject/peerswap/lightning"
)

type CLightningNode struct {
	*DaemonProcess
	*CLightningProxy

	DataDir    string
	ConfigFile string
	Port       int
	Info       *glightning.NodeInfo

	bitcoin *BitcoinNode
}

func NewCLightningNode(testDir string, bitcoin *BitcoinNode, id int) (*CLightningNode, error) {
	port, err := GetFreePort()
	if err != nil {
		return nil, fmt.Errorf("GetFreePort() %w", err)
	}

	// Use node ID for directory name instead of random string (shorter and more predictable)
	dataDir := filepath.Join(testDir, fmt.Sprintf("c%d", id))
	networkDir := filepath.Join(dataDir, "regtest")

	err = os.MkdirAll(networkDir, os.ModeDir|os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("os.MkdirAll() %w", err)
	}

	bitcoinConf, err := ReadConfig(bitcoin.configFile)
	if err != nil {
		return nil, fmt.Errorf("ReadConfig() %w", err)
	}

	var bitcoinRpcPass string
	if pass, ok := bitcoinConf["rpcpassword"]; ok {
		bitcoinRpcPass = pass
	} else {
		return nil, fmt.Errorf("bitcoin rpcpassword not found in config %s", bitcoin.configFile)
	}

	var bitcoinRpcUser string
	if user, ok := bitcoinConf["rpcuser"]; ok {
		bitcoinRpcUser = user
	} else {
		return nil, fmt.Errorf("bitcoin rpcuser not found in config %s", bitcoin.configFile)
	}

	var bitcoinRpcPort string
	if port, ok := bitcoinConf["rpcport"]; ok {
		bitcoinRpcPort = port
	} else {
		return nil, fmt.Errorf("bitcoin rpcuser not found in config %s", bitcoin.configFile)
	}

	cmdLine := []string{
		"lightningd",
		fmt.Sprintf("--lightning-dir=%s", dataDir),
		fmt.Sprintf("--log-level=%s", "debug"),
		fmt.Sprintf("--addr=127.0.0.1:%d", port),
		fmt.Sprintf("--network=%s", "regtest"),
		fmt.Sprintf("--ignore-fee-limits=%s", "true"),
		fmt.Sprintf("--bitcoin-rpcuser=%s", bitcoinRpcUser),
		fmt.Sprintf("--bitcoin-rpcpassword=%s", bitcoinRpcPass),
		fmt.Sprintf("--bitcoin-rpcport=%s", bitcoinRpcPort),
		fmt.Sprintf("--bitcoin-datadir=%s", bitcoin.DataDir),
		fmt.Sprintf("--developer"),
		fmt.Sprintf("--allow-deprecated-apis=true"),
	}

	// Check socket path length before proceeding
	socketPath := filepath.Join(networkDir, "lightning-rpc")
	if len(socketPath) > 104 { // Unix domain socket path limit
		return nil, fmt.Errorf("socket path too long (%d chars): %s. Unix domain sockets are limited to 104-108 characters. Try setting TMPDIR to a shorter path.", len(socketPath), socketPath)
	}

	proxy, err := NewCLightningProxy("lightning-rpc", networkDir)
	if err != nil {
		return nil, fmt.Errorf("NewCLightningProxy() %w", err)
	}

	// Create seed file with a deterministic but unique seed
	// Use dataDir path and node ID to generate a 32-byte seed
	seedSource := fmt.Sprintf("%s-node-%d-seed-padding", dataDir, id)
	// Ensure we have at least 32 bytes
	for len(seedSource) < 32 {
		seedSource += "0"
	}
	seed := []byte(seedSource)[:32]
	seedFile := filepath.Join(networkDir, "hsm_secret")
	err = os.WriteFile(seedFile, seed, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("WriteFile() %w", err)
	}

	return &CLightningNode{
		DaemonProcess:   NewDaemonProcess(cmdLine, fmt.Sprintf("cln-%d", id)),
		CLightningProxy: proxy,
		DataDir:         dataDir,
		Port:            port,
		bitcoin:         bitcoin,
	}, nil
}

func (n *CLightningNode) Run(waitForReady, waitForBitcoinSynced bool) error {
	n.DaemonProcess.Run()

	// Establish RPC connection first
	var counter int
	var err error
	for {
		if counter > 20 {
			return fmt.Errorf("too many retries establishing RPC connection: %w", err)
		}

		err = n.StartProxy()
		if err != nil {
			counter++
			time.Sleep(500 * time.Millisecond)
			continue
		}

		break
	}

	if waitForReady {
		// Wait for CLN to be ready via RPC
		err = WaitFor(func() bool {
			info, err := n.Rpc.GetInfo()
			if err != nil {
				// RPC might not be fully ready yet
				return false
			}
			// CLN is ready when it has a valid node ID
			if info.Id != "" {
				log.Printf("CLN node ready with ID: %s", info.Id)
				return true
			}
			return false
		}, TIMEOUT)

		if err != nil {
			return fmt.Errorf("CLightningNode.Run() startup detection failed: %w", err)
		}
	}

	// Cache info
	n.Info, err = n.Rpc.GetInfo()
	if err != nil {
		return fmt.Errorf("rpc.GetInfo() %w", err)
	}

	if waitForBitcoinSynced {
		// Wait for sync with bitcoin network
		return WaitFor(func() bool {
			info, err := n.Rpc.GetInfo()
			if err != nil {
				log.Printf("rpc.GetInfo() %v", err)
				return false
			}

			isHeightSync, err := n.IsBlockHeightSynced()
			if err != nil {
				log.Printf("rpc.GetInfo() %v", err)
				return false
			}

			return info.IsBitcoindSync() && info.IsLightningdSync() && isHeightSync
		}, TIMEOUT)
	}

	return nil
}

func (n *CLightningNode) Stop() error {
	n.Rpc.Stop()
	return n.WaitForLog("hsmd: Shutting down", TIMEOUT)
}

func (n *CLightningNode) Shutdown() error {
	n.Stop()
	return os.Remove(filepath.Join(n.dataDir, "lightning-rpc"))
}

func (n *CLightningNode) Id() string {
	return n.Info.Id
}

func (n *CLightningNode) Address() string {
	return fmt.Sprintf("%s@127.0.0.1:%d", n.Info.Id, n.Port)
}

func (n *CLightningNode) GetBtcBalanceSat() (sats uint64, err error) {
	r, err := n.Rpc.ListFunds()
	if err != nil {
		return 0, fmt.Errorf("ListFunds() %w", err)
	}

	var sum uint64
	for _, output := range r.Outputs {
		// AmountMilliSatoshi is in msat.
		sum += output.AmountMilliSatoshi.MSat() / 1000
	}
	return sum, nil
}

func (n *CLightningNode) GetChannelBalanceSat(scid string) (sats uint64, err error) {
	funds, err := n.Rpc.ListFunds()
	if err != nil {
		return 0, fmt.Errorf("rpc.ListFunds() %w", err)
	}

	for _, ch := range funds.Channels {
		if ch.ShortChannelId == scid {
			return ch.OurAmountMilliSatoshi.MSat() / 1000, nil
		}
	}

	return 0, fmt.Errorf("no channel found with scid %s", scid)
}

func (n *CLightningNode) GetScid(remote LightningNode) (string, error) {
	scid := ""
	err := WaitForWithErr(func() (bool, error) {
		var res clightning.ListPeerChannelsResponse
		err := n.Rpc.Request(clightning.ListPeerChannelsRequest{}, &res)
		if err != nil {
			return false, fmt.Errorf("ListPeers() %w", err)
		}
		for _, c := range res.Channels {
			if c.PeerId == remote.Id() {
				scid = c.ShortChannelId
				return scid != "", nil
			}
		}
		return false, fmt.Errorf("peer not found")
	}, TIMEOUT)
	if err != nil {
		return "", err
	}
	return scid, nil
}

func (n *CLightningNode) Connect(peer LightningNode, waitForConnection bool) error {
	id, host, port, err := SplitLnAddr(peer.Address())
	if err != nil {
		return fmt.Errorf("SplitLnAddr() %w", err)
	}

	_, err = n.Rpc.Connect(id, host, uint(port))
	if err != nil {
		return fmt.Errorf("Connect() %w", err)
	}

	if waitForConnection {
		return WaitForWithErr(func() (bool, error) {
			localIsConnected, err := n.IsConnected(peer)
			if err != nil {
				return false, fmt.Errorf("IsConnected() %w", err)
			}
			peerIsConnected, err := peer.IsConnected(n)
			if err != nil {
				return false, fmt.Errorf("IsConnected() %w", err)
			}
			return localIsConnected && peerIsConnected, nil
		}, TIMEOUT)
	}

	return nil
}

func (n *CLightningNode) FundWallet(sats uint64, mineBlock bool) (string, error) {
	addr, err := n.Rpc.NewAddr()
	if err != nil {
		return "", fmt.Errorf("rpc.NewAddress() %w", err)
	}

	r, err := n.bitcoin.Call("sendtoaddress", addr, float64(sats)/math.Pow(10., 8))
	if err != nil {
		return "", fmt.Errorf("sendtoaddress %w", err)
	}

	var a struct {
		TxId   string
		Abc    error
		Reason int
	}

	var txId string
	err = r.GetObject(&a)
	if err != nil {
		txId, err = r.GetString()
		if err != nil {
			return "", err
		}
	}

	if mineBlock {
		err = n.bitcoin.GenerateBlocks(1)
		if err != nil {
			return "", fmt.Errorf("bitcoin.GenerateBlocks() %w", err)
		}
		err = n.WaitForLog(fmt.Sprintf("Owning output .* txid %s CONFIRMED", txId), TIMEOUT)
		if err != nil {
			return "", err
		}
	}

	return addr, nil
}

func (n *CLightningNode) OpenChannel(remote LightningNode, capacity, pushAmt uint64, connect, confirm, waitForActiveChannel bool) (string, error) {
	_, err := n.FundWallet(uint64(1.5*float64(capacity)), true)
	if err != nil {
		return "", fmt.Errorf("FundWallet() %w", err)
	}

	isConnected, err := n.IsConnected(remote)
	if err != nil {
		return "", fmt.Errorf("IsConnected() %w", err)
	}

	if !isConnected && connect {
		err = n.Connect(remote, true)
		if err != nil {
			return "", fmt.Errorf("Connect() %w", err)
		}
	}

	pushAmtSat := &glightning.Sat{Value: pushAmt}
	fr, err := n.Rpc.FundChannelExt(remote.Id(), &glightning.Sat{Value: capacity}, nil, true, nil, pushAmtSat.ConvertMsat())
	if err != nil {
		return "", fmt.Errorf("FundChannel() %w", err)
	}

	// Wait for tx in mempool
	err = WaitFor(func() bool {
		r, err := n.bitcoin.Call("getrawmempool")
		if err != nil {
			log.Println("getrawmempool: %w", err)
			return false
		}

		var rawMempool []string
		err = r.GetObject(&rawMempool)
		if err != nil {
			log.Println("can not unmarshal object: %w", err)
			return false
		}

		return slices.Contains(rawMempool, fr.FundingTxId)
	}, TIMEOUT)
	if err != nil {
		return "", fmt.Errorf("error waiting for tx in mempool: %w", err)
	}

	if waitForActiveChannel || confirm {
		n.bitcoin.GenerateBlocks(10)
	}

	if waitForActiveChannel {
		err = WaitForWithErr(func() (bool, error) {
			scid, err := n.GetScid(remote)
			if err != nil {
				return false, fmt.Errorf("GetScid() %w", err)
			}
			if scid == "" {
				return false, nil
			}

			localActive, err := n.IsChannelActive(scid)
			if err != nil {
				return false, fmt.Errorf("IsChannelActive() %w", err)
			}
			remoteActive, err := remote.IsChannelActive(scid)
			if err != nil {
				return false, fmt.Errorf("IsChannelActive() %w", err)
			}
			hasRoute, err := n.HasRoute(remote.Id(), scid)
			if err != nil {
				return false, nil
			}
			return remoteActive && localActive && hasRoute, nil
		}, TIMEOUT)
		if err != nil {
			return "", fmt.Errorf("error waiting for active channel: %w", err)
		}
	}

	scid, err := n.GetScid(remote)
	if err != nil {
		return "", fmt.Errorf("GetScid() %w", err)
	}

	return scid, nil
}

// HasRoute check the route is constructed
func (n *CLightningNode) HasRoute(remote, scid string) (bool, error) {
	routes, err := n.Rpc.GetRoute(remote, 1, 1, 0, n.Info.Id, 0, nil, 1)
	if err != nil {
		return false, fmt.Errorf("GetRoute() %w", err)
	}
	return len(routes) > 0, nil
}

func (n *CLightningNode) IsBlockHeightSynced() (bool, error) {
	r, err := n.bitcoin.Rpc.Call("getblockcount")
	if err != nil {
		return false, fmt.Errorf("bitcoin.rpc.Call(\"getblockcount\") %w", err)
	}

	chainHeight, err := r.GetFloat()
	if err != nil {
		return false, fmt.Errorf("GetFloat() %w", err)
	}

	nodeInfo, err := n.Rpc.GetInfo()
	if err != nil {
		return false, fmt.Errorf("GetInfo() %w", err)
	}
	return nodeInfo.Blockheight >= uint(chainHeight), nil
}

func (n *CLightningNode) IsChannelActive(scid string) (bool, error) {
	funds, err := n.Rpc.ListFunds()
	if err != nil {
		return false, fmt.Errorf("ListChannels() %w", err)
	}

	for _, ch := range funds.Channels {
		if ch.ShortChannelId == scid {
			return ch.State == "CHANNELD_NORMAL" && ch.Connected, nil
		}
	}

	return false, nil
}

func (n *CLightningNode) IsConnected(remote LightningNode) (bool, error) {
	peers, err := n.Rpc.ListPeers()
	if err != nil {
		return false, fmt.Errorf("rpc.ListPeers() %w", err)
	}

	for _, peer := range peers {
		if remote.Id() == peer.Id {
			return peer.Connected, nil
		}
	}
	return false, nil
}

func (n *CLightningNode) AddInvoice(amtSat uint64, desc, label string) (payreq string, err error) {
	if label == "" {
		var labelBytes = make([]byte, 5)
		_, err := rand.Read(labelBytes)
		if err != nil {
			return "", err
		}
		label = string(labelBytes)
	}

	inv, err := n.Rpc.Invoice(amtSat*1000, label, desc)
	if err != nil {
		return "", nil
	}
	return inv.Bolt11, nil
}

func (n *CLightningNode) PayInvoice(payreq string) error {
	_, err := n.Rpc.PayBolt(payreq)
	return err
}

func (n *CLightningNode) SendPay(bolt11, scid string) error {
	decodedBolt11, err := n.Rpc.DecodeBolt11(bolt11)
	if err != nil {
		return err
	}

	_, err = n.Rpc.SendPay([]glightning.RouteHop{
		{
			Id:             decodedBolt11.Payee,
			ShortChannelId: scid,
			// MilliSatoshi:   decodedBolt11.MilliSatoshis,
			AmountMsat: decodedBolt11.AmountMsat,
			Delay:      uint32(decodedBolt11.MinFinalCltvExpiry + 1),
			Direction:  0,
		},
	},
		decodedBolt11.PaymentHash,
		"",
		decodedBolt11.AmountMsat.MSat(),
		bolt11,
		decodedBolt11.PaymentSecret,
		0,
	)
	return err
}

func (n *CLightningNode) GetDataDir() string {
	return n.dataDir
}

func (n *CLightningNode) GetLatestInvoice() (string, error) {
	r, err := n.Rpc.ListInvoices()
	if err != nil {
		return "", err
	}

	if len(r) > 0 {
		return r[len(r)-1].Bolt11, nil
	}

	return "", nil
}

func (n *CLightningNode) GetMemoFromPayreq(bolt11 string) (string, error) {
	r, err := n.Rpc.DecodeBolt11(bolt11)
	if err != nil {
		return "", err
	}

	return r.Description, nil
}

func (n *CLightningNode) GetLatestPayReqOfPayment() (string, error) {
	ps, err := n.Rpc.ListPays()
	if err != nil {
		return "", err
	}
	if len(ps) > 0 {
		return ps[len(ps)-1].Bolt11, nil
	}
	return "", fmt.Errorf("payments list is nil")
}

func (n *CLightningNode) GetFeeInvoiceAmtSat() (sat uint64, err error) {
	rx := regexp.MustCompile(`^peerswap .* fee .*`)
	var feeInvoiceAmt uint64
	r, err := n.Rpc.ListInvoices()
	if err != nil {
		return 0, err
	}
	for _, i := range r {
		if rx.MatchString(i.Description) {
			feeInvoiceAmt += i.AmountMilliSatoshi.MSat() / 1000
		}
	}
	return feeInvoiceAmt, nil
}

type SetChannel struct {
	Id                       string `json:"id"`
	HtlcMaximumMilliSatoshis string `json:"htlcmax,omitempty"`
}

type ChannelInfo struct {
	PeerID                    string            `json:"peer_id"`
	ChannelID                 string            `json:"channel_id"`
	ShortChannelID            string            `json:"short_channel_id"`
	FeeBaseMsat               glightning.Amount `json:"fee_base_msat"`
	FeeProportionalMillionths glightning.Amount `json:"fee_proportional_millionths"`
	MinimumHtlcOutMsat        glightning.Amount `json:"minimum_htlc_out_msat"`
	MaximumHtlcOutMsat        glightning.Amount `json:"maximum_htlc_out_msat"`
}

type SetChannelResponse struct {
	Channels []ChannelInfo `json:"channels"`
}

func (r *SetChannel) Name() string {
	return "setchannel"
}

func (n *CLightningNode) SetHtlcMaximumMilliSatoshis(scid string, maxHtlcMsat uint64) (msat uint64, err error) {
	var res SetChannelResponse
	err = n.Rpc.Request(&SetChannel{
		Id:                       lightning.Scid(scid).ClnStyle(),
		HtlcMaximumMilliSatoshis: fmt.Sprint(maxHtlcMsat),
	}, &res)
	if err != nil {
		return 0, err
	}
	return maxHtlcMsat, err
}

// ForceFeeUpdate force a fee update by restarting the node with the new feerate.
// Up to 5 values, separated by '/' can be provided for feerates
// if fewer are provided, then the final value is used for the
// remainder. The values are in per-kw (roughly 1/4 of bitcoind's per-kb
// values), and the order is "opening", "mutual_close", "unilateral_close",
// "delayed_to_us", "htlc_resolution", and "penalty".
// ref: https://docs.corelightning.org/reference/lightningd-config
func (n *CLightningNode) ForceFeeUpdate(scid, feerates string) error {
	err := n.Stop()
	if err != nil {
		return err
	}
	n.AppendCmdLine([]string{fmt.Sprintf("--force-feerates=%s", feerates)})
	err = n.Run(true, true)
	if err != nil {
		return err
	}
	err = n.WaitForLog("peerswap initialized", TIMEOUT)
	if err != nil {
		return err
	}
	return WaitForWithErr(func() (bool, error) {
		return n.IsChannelActive(scid)
	}, TIMEOUT)
}
