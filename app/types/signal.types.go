package types

type TVSignal struct {
	Ticker    string `json:"ticker"`
	ChannelID int64  `json:"TGChannel"`
	Direction string `json:"direction,omitempty"`
	User      string `json:"user"`
	Password  string `json:"password"`
}
