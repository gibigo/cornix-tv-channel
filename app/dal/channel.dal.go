package dal

import (
	"github.com/gibigo/cornix-tv-channel/config/database"
	"gorm.io/gorm"
)

type Channel struct {
	gorm.Model `swaggerignore:"true"`
	Telegram   int64
	TVSignal   []*TVSignal `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Strategy   []*Strategy `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID     uint
}

func FindChannelByTelegramId(dest interface{}, telegramID int64) *gorm.DB {
	return FindChannel(dest, "telegram = ?", telegramID)
}

func FindChannel(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Model(&Channel{}).Take(dest, conds...)
}

func CreateChannel(channel *Channel) *gorm.DB {
	return database.DB.Create(channel)
}
