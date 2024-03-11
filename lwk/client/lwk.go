package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/elementsproject/glightning/jrpc2"
	"github.com/elementsproject/peerswap/lwk/api"
)

type Lwk struct {
	api api.API
}

func NewLwk(endpoint string) *Lwk {
	return &Lwk{
		api: *api.NewAPI(endpoint),
	}
}

func (l *Lwk) request(ctx context.Context, m jrpc2.Method, resp interface{}) error {
	id := l.api.NextId()
	mr := &jrpc2.Request{Id: id, Method: m}
	jbytes, err := json.Marshal(mr)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, l.api.BaseURL, bytes.NewBuffer(jbytes))
	if err != nil {
		return err
	}
	rezp, err := l.api.Do(req)
	if err != nil {
		return err
	}
	defer l.api.Drain(rezp)
	switch rezp.StatusCode {
	case http.StatusUnauthorized:
		return errors.New("Authorization failed: Incorrect user or password")
	case http.StatusBadRequest, http.StatusNotFound, http.StatusInternalServerError:
		// do nothing
	default:
		if rezp.StatusCode > http.StatusBadRequest {
			return errors.New(fmt.Sprintf("server returned HTTP error %d", rezp.StatusCode))
		} else if rezp.ContentLength == 0 {
			return errors.New("no response from server")
		}
	}

	var rawResp jrpc2.RawResponse

	decoder := json.NewDecoder(rezp.Body)
	err = decoder.Decode(&rawResp)
	if err != nil {
		return err
	}

	if rawResp.Error != nil {
		return rawResp.Error
	}
	return json.Unmarshal(rawResp.Raw, resp)
}

type AddressRequest struct {
	Index      *uint32 `json:"index,omitempty"`
	WalletName string  `json:"name"`
	Signer     *string `json:"signer,omitempty"`
}

func (r *AddressRequest) Name() string {
	return "address"
}

type AddressResponse struct {
	Address string  `json:"address"`
	Index   *uint32 `json:"index,omitempty"`
}

func (l *Lwk) Address(ctx context.Context, req *AddressRequest) (*AddressResponse, error) {
	var resp AddressResponse
	err := l.request(ctx, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UnvalidatedAddressee struct {
	Address string `json:"address"`
	Asset   string `json:"asset"`
	Satoshi uint64 `json:"satoshi"`
}

type SendRequest struct {
	Addressees []*UnvalidatedAddressee `json:"addressees"`
	FeeRate    *float64                `json:"fee_rate,omitempty"`
	WalletName string                  `json:"name"`
}

type SendResponse struct {
	Pset string `json:"pset"`
}

func (s *SendRequest) Name() string {
	return "send_many"
}

func (l *Lwk) Send(ctx context.Context, s *SendRequest) (*SendResponse, error) {
	var resp SendResponse
	err := l.request(ctx, s, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SignRequest struct {
	SignerName string `json:"name"`
	Pset       string `json:"pset"`
}

type SignResponse struct {
	Pset string `json:"pset"`
}

func (s *SignRequest) Name() string {
	return "sign"
}

func (l *Lwk) Sign(ctx context.Context, s *SignRequest) (*SignResponse, error) {
	var resp SignResponse
	err := l.request(ctx, s, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type BroadcastRequest struct {
	DryRun     bool   `json:"dry_run"`
	WalletName string `json:"name"`
	Pset       string `json:"pset"`
}

type BroadcastResponse struct {
	Txid string `json:"txid"`
}

func (b *BroadcastRequest) Name() string {
	return "broadcast"
}

func (l *Lwk) Broadcast(ctx context.Context, b *BroadcastRequest) (*BroadcastResponse, error) {
	var resp BroadcastResponse
	err := l.request(ctx, b, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type BalanceRequest struct {
	WalletName  string `json:"name"`
	WithTickers bool   `json:"with_tickers"`
}

func (b *BalanceRequest) Name() string {
	return "balance"
}

type BalanceResponse struct {
	Balance map[string]int64 `json:"balance"`
}

func (l *Lwk) Balance(ctx context.Context, b *BalanceRequest) (*BalanceResponse, error) {
	var resp BalanceResponse
	err := l.request(ctx, b, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
