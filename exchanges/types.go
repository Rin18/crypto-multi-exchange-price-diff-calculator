package exchanges

type Ticker struct {
	Exchange  string
	Symbol    string `json:"symbol"`
	Price     string `json:"markPrice"`
	Timestamp int64  `json:"time"`
}
