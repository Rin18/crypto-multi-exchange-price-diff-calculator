package exchanges

type Ticker struct {
	Symbol    string `json:"symbol"`
	Price     string `json:"markPrice"`
	Timestamp int64  `json:"time"`
}
