package dal

import (
	"github.com/gibigo/cornix-tv-channel/config/database"
	"gorm.io/gorm"
)

type TVSignal struct {
	gorm.Model
	Symbol    string
	Price     float64
	Direction string
	ChannelID uint
}

func FindSignalBySymbol(dest interface{}, channelID uint, symbol string) *gorm.DB {
	return database.DB.Where("channel_id = ? AND symbol = ?", channelID, symbol).First(dest)
}
