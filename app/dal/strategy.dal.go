package dal

import (
	"github.com/gibigo/cornix-tv-channel/config/database"
	"gorm.io/gorm"
)

type Strategy struct {
	gorm.Model
	AllowCounter   bool
	Symbol         string
	TargetStrategy *TargetStrategy `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ZoneStrategy   *ZoneStrategy   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ChannelID      uint
}

type ZoneStrategy struct {
	gorm.Model
	EntryStart float64
	EntryStop  float64
	TPs        []*TP `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SL         *SL   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	IsBreakout bool
	StrategyID uint
}

type TargetStrategy struct {
	gorm.Model
	Entries    []*Entry `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TPs        []*TP    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SL         *SL      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	IsBreakout bool
	StrategyID uint
}

type Entry struct {
	gorm.Model
	Diff             float64
	TargetStrategyID uint
}

type TP struct {
	gorm.Model
	Diff             float64
	TargetStrategyID *uint
	ZoneStrategyID   *uint
}

type SL struct {
	gorm.Model
	Diff             float64
	TargetStrategyID *uint
	ZoneStrategyID   *uint
}

/* // maybe rewrite this, could be an old issue, not sure tho
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
*/

func FindStrategyBySymbol(dest interface{}, symbol string, channelID interface{}) *gorm.DB {
	return FindStrategy(dest, "symbol = ? AND channel_id = ?", symbol, channelID)
}

func FindStrategy(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Model(&Strategy{}).Take(dest, conds...)
}

func CreateStrategy(strategy interface{}) *gorm.DB {
	return database.DB.Create(strategy)
}

func FindAllStrategiesFromChannel(dest interface{}, channelID interface{}) *gorm.DB {
	//return database.DB.Debug().Select("*").Joins("JOIN target_strategies ON target_strategies.strategy_id = strategies.id").Where("channel_id = ?", channelID).Find(dest)
	return database.DB.Debug().
		Preload("TargetStrategy").
		Preload("TargetStrategy.Entries").Preload("TargetStrategy.TPs").Preload("TargetStrategy.SL").
		Preload("ZoneStrategy").
		Preload("ZoneStrategy.TPs").Preload("ZoneStrategy.SL").
		Find(dest)
}
