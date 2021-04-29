package dal

import "gorm.io/gorm"

type TVSignal struct {
	gorm.Model `swaggerignore:"true"`
	Ticker     string
	TelegramID int64
	Direction  string
}
