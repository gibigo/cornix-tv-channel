package dal

import "gorm.io/gorm"

type TVSignal struct {
	gorm.Model
	Ticker    string
	ChannelID *int64
	Direction string
}
