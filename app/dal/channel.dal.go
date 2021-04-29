package dal

import (
	"gorm.io/gorm"
)

type Channel struct {
	gorm.Model
	TelegramID int64
	Signals    []*TVSignal `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Strategy   []*Strategy `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID     uint
}
