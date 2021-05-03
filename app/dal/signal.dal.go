package dal

import "gorm.io/gorm"

type TVSignal struct {
	gorm.Model
	Ticker     string
	TelegramID int64
	Direction  string
	ChannelID  uint
}
