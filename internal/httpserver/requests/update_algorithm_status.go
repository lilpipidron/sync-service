package requests

type UpdateAlgorithmStatusRequest struct {
	ID       int64 `json:"id"`
	ClientID int64 `json:"client_id,omitempty"`
	VWAP     bool  `json:"vwap,omitempty"`
	TWAP     bool  `json:"twap,omitempty"`
	HFT      bool  `json:"hft,omitempty"`
}
