package dal

import (
	"github.com/gibigo/cornix-tv-channel/config/database"
	"gorm.io/gorm"
)

type Channel struct {
	ID         int64
	UserName   string // foreign Key
	Signals    []TVSignal
	StrategyID int64
	Strategy   Strategy
}

type Strategy struct {
	ID           int64
	UserName     string // foreign Key
	AllowCounter bool
	Entries      []Entry `validate:"gt=0,dive,required"`
	TPs          []TP    `validate:"gt=0,dive,required"`
	SL           SL      `validate:"required,dive,required"`
}

type Entry struct {
	gorm.Model
	Diff       float32 `validate:"required"`
	StrategyID int64
}

type TP struct {
	gorm.Model
	Diff       float32 `validate:"required"`
	StrategyID int64
}

type SL struct {
	gorm.Model
	Diff       float32 `validate:"required"`
	StrategyID int64
}

func FindStrategy(dest interface{}, query interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Table("strategies").Where(query, conds...).
		Joins("JOIN entries ON entries.strategy_id = strategies.id").
		Joins("JOIN tps ON tps.strategy_id = strategies.id").
		Joins("JOIN sls ON sls.strategy_id = strategies.id").Select("strategies.id", "strategies.allow_counter", "entries.diff", "tps.diff", "sls.diff")
}

func FindAllStrategiesFromUser(dest interface{}, username interface{}) *gorm.DB {
	return FindStrategy(dest, "strategies.user_name = ?", username)
}

func FindStrategyByID(dest interface{}, id int64) *gorm.DB {
	return FindStrategy(dest, "strategies.id = ?", id)
}

func CreateStrategy(strategy *Strategy) *gorm.DB {
	return database.DB.Create(strategy)
}
