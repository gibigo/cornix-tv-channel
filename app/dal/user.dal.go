package dal

import (
	"github.com/gibigo/cornix-tv-channel/config/database"
	"gorm.io/gorm"
)

type User struct {
	Name       string `gorm:"primaryKey"`
	Password   string
	Channels   []Channel
	Strategies []Strategy
}

func FindUser(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Model(&User{}).Take(dest, conds...)
}

func FindUserByName(dest interface{}, username interface{}) *gorm.DB {
	return FindUser(dest, "name = ?", username)
}

func CreateUser(user *User) *gorm.DB {
	return database.DB.Create(user)
}

func DeleteUser(username interface{}) *gorm.DB {
	return database.DB.Where("name = ?", username).Delete(&User{})
}
