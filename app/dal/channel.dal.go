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
	ID       int64
	UserName string // foreign Key
	Entries  []Entry
	TPs      []TP
	SL       SL
}

type Entry struct {
	gorm.Model
	Diff       float32
	StrategyID int64
}

type TP struct {
	gorm.Model
	Diff       float32
	StrategyID int64
}

type SL struct {
	gorm.Model
	Diff       float32
	StrategyID int64
}

func FindStrategy(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Model(&Strategy{}).Find(dest, conds...)
}

func FindAllStrategiesFromUser(dest interface{}, username interface{}) *gorm.DB {
	return FindStrategy(dest, "user_name = ?", username)
}

func CreateStrategy(strategy *Strategy) *gorm.DB {
	return database.DB.Create(strategy)
}
