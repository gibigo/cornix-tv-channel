package dal

import "gorm.io/gorm"

type TVSignal struct {
	gorm.Model
	Ticker    string
	Price     float64
	Direction string
	ChannelID uint
}
