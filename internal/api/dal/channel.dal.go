package dal

import (
	"github.com/gibigo/cornix-tv-channel/internal/config/database"
	"gorm.io/gorm"
)

type Channel struct {
	gorm.Model
	Telegram int64
	TVSignal []*TVSignal `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Strategy []*Strategy `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID   uint
}

func FindChannelByTelegramId(dest interface{}, telegramID int64) *gorm.DB {
	return FindChannel(dest, "telegram = ?", telegramID)
}

func FindChannel(dest interface{}, query string, args ...interface{}) *gorm.DB {
	return database.DB.Model(&Channel{}).Where(query, args...).Take(dest)
}

func CreateChannel(channel *Channel) *gorm.DB {
	return database.DB.Create(channel)
}

func FindAllChannelsFromUser(dest interface{}, username interface{}) *gorm.DB {
	subQuery := database.DB.Select("id").Where("name = ?", username).Table("users")
	return database.DB.Model(&Channel{}).Where("user_id = (?)", subQuery).Find(dest)
}

func FindChannelFromUser(dest interface{}, username interface{}, channel interface{}) *gorm.DB {
	subQuery := database.DB.Select("id").Where("name = ?", username).Table("users")
	return database.DB.Model(&Channel{}).Where("user_id = (?) AND id = ?", subQuery, channel).Find(dest)
}

func DeleteChannelFromUser(username interface{}, channel interface{}) *gorm.DB {
	subQuery := database.DB.Select("id").Where("name = ?", username).Table("users")
	return database.DB.Unscoped().Where("user_id = (?) AND id = ?", subQuery, channel).Delete(&Channel{})
}

func UpdateChannel(username interface{}, update map[string]interface{}, channel interface{}) *gorm.DB {
	subQuery := database.DB.Select("id").Where("name = ?", username).Table("users")
	return database.DB.Model(&Channel{}).Where("user_id = (?) AND id = ?", subQuery, channel).Updates(update)
}

func ChangeChannelTelegram(username interface{}, telegram int64, channel interface{}) *gorm.DB {
	return UpdateChannel(username, map[string]interface{}{"telegram": telegram}, channel)
}
