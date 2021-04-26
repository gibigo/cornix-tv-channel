package dal

import (
	"github.com/gibigo/cornix-tv-channel/config/database"
	"gorm.io/gorm"
)

type Channel struct {
	ID         int64
	UserID     int64
	Signals    []*TVSignal `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StrategyID int64
	Strategy   *Strategy
}

type Strategy struct {
	ID           int64
	UserID       int64
	AllowCounter bool
	Entries      []*Entry `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" validate:"gt=0,dive,required"`
	TPs          []*TP    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" validate:"gt=0,dive,required"`
	SL           *SL      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" validate:"required,dive,required"`
}

type Entry struct {
	gorm.Model `swaggerignore:"true"`
	Diff       float32 `validate:"required"`
	StrategyID int64
}

type TP struct {
	gorm.Model `swaggerignore:"true"`
	Diff       float32 `validate:"required"`
	StrategyID int64
}

type SL struct {
	gorm.Model `swaggerignore:"true"`
	Diff       float32 `validate:"required"`
	StrategyID int64
}

func FindStrategy(dest interface{}, query interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Table("strategies").Where(query, conds...).
		Joins("JOIN entries ON entries.strategy_id = strategies.id").
		Joins("JOIN tps ON tps.strategy_id = strategies.id").
		Joins("JOIN sls ON sls.strategy_id = strategies.id").Select("strategies.id", "strategies.allow_counter", "entries.diff", "tps.diff", "sls.diff")
}

func FindAllStrategiesFromUser(dest interface{}, userID int64) *gorm.DB {
	return FindStrategy(dest, "strategies.user_id = ?", userID)
}

func FindStrategyByID(dest interface{}, id int64) *gorm.DB {
	return FindStrategy(dest, "strategies.id = ?", id)
}

func CreateStrategy(strategy *Strategy) *gorm.DB {
	return database.DB.Create(strategy)
}
