package types

type TVSignal struct {
	Symbol    string  `json:"ticker" validate:"required"`
	Price     float64 `json:"price" validate:"required"`
	ChannelID int64   `json:"TGChannel" validate:"required"`
	Direction string  `json:"direction,omitempty"`
	User      string  `json:"user" validate:"required"`
	UUID      string  `json:"uuid" validate:"required"`
}
