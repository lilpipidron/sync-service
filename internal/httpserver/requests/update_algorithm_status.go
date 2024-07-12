package requests

type UpdateAlgorithmStatusRequest struct {
	ID   int64 `json:"id"`
	VWAP bool  `json:"vwap"`
	TWAP bool  `json:"twap"`
	HFT  bool  `json:"hft"`
}
