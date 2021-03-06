package dal

import (
	"github.com/gibigo/cornix-tv-channel/internal/config/database"
	"github.com/gibigo/cornix-tv-channel/internal/utils/password"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
	UUID     string
	Channels []*Channel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func CreateUser(user *User) *gorm.DB {
	return database.DB.Create(user)
}

func FindUser(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Model(&User{}).Take(dest, conds...)
}

func FindUserByName(dest interface{}, username interface{}) *gorm.DB {
	return FindUser(dest, "name = ?", username)
}

func FindUserByUUID(dest interface{}, username interface{}, uuid interface{}) *gorm.DB {
	return FindUser(dest, "name = ? AND uuid = ?", username, uuid)
}

func DeleteUser(username interface{}) *gorm.DB {
	return database.DB.Unscoped().Where("name = ?", username).Delete(&User{})
}

func ChangeUserPassword(username, newPassword string) *gorm.DB {
	return UpdateUser(&User{}, map[string]interface{}{"password": password.Generate(newPassword)}, username)
}

func ChangeUserName(oldName, newName string) *gorm.DB {
	return UpdateUser(&User{}, map[string]interface{}{"name": newName}, oldName)
}

func ChangeUserAndPassword(oldName, newName, newPassword string) *gorm.DB {
	return UpdateUser(&User{}, map[string]interface{}{"name": newName, "password": password.Generate(newPassword)}, oldName)
}

func UpdateUser(dest interface{}, update map[string]interface{}, user string) *gorm.DB {
	return database.DB.Model(dest).Where("name = ?", user).Updates(update)
}
