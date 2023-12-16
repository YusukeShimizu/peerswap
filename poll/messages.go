package poll

import "github.com/elementsproject/peerswap/messages"

type PollMessage struct {
	Version            uint64   `json:"version"`
	Assets             []string `json:"assets"`
	PeerAllowed        bool     `json:"peer_allowed"`
	SwapInPremiumRate  int64    `json:"swap_in_premium_rate"`
	SwapOutPremiumRate int64    `json:"swap_out_premium_rate"`
}

func (PollMessage) MessageType() messages.MessageType {
	return messages.MESSAGETYPE_POLL
}

type RequestPollMessage struct {
	Version            uint64   `json:"version"`
	Assets             []string `json:"assets"`
	PeerAllowed        bool     `json:"peer_allowed"`
	SwapInPremiumRate  int64    `json:"swap_in_premium_rate"`
	SwapOutPremiumRate int64    `json:"swap_out_premium_rate"`
}

func (RequestPollMessage) MessageType() messages.MessageType {
	return messages.MESSAGETYPE_REQUEST_POLL
}
