package types

type TVSignal struct {
	Ticker    string  `json:"ticker"`
	Price     float64 `json:"price"`
	ChannelID int64   `json:"TGChannel"`
	Direction string  `json:"direction,omitempty"`
	User      string  `json:"user"`
	UUID      string  `json:"uuid"`
}
