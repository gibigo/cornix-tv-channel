package dal

import (
	"github.com/gibigo/cornix-tv-channel/config/database"
	"gorm.io/gorm"
)

type Strategy struct {
	gorm.Model       `swaggerignore:"true"`
	AllowCounter     bool
	Coin             string
	IsTargetStrategy bool
	TargetStrategy   *TargetStrategy
	IsZoneStrategy   bool
	ZoneStrategy     *ZoneStrategy
	UserID           uint
	ChannelID        uint
}

type ZoneStrategy struct {
	gorm.Model `swaggerignore:"true"`
	EntryStart float64
	EntryStop  float64
	TPs        []*TP
	SL         *SL
	IsBreakout bool
	StrategyID uint
}

type TargetStrategy struct {
	gorm.Model `swaggerignore:"true"`
	Entries    []*Entry `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" validate:"gt=0,dive,required"`
	TPs        []*TP    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" validate:"gt=0,dive,required"`
	SL         *SL      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" validate:"required,dive,required"`
	IsBreakout bool
	StrategyID uint
}

type Entry struct {
	gorm.Model       `swaggerignore:"true"`
	Diff             float64 `validate:"required"`
	TargetStrategyID uint
	ZoneStrategyID   uint
}

type TP struct {
	gorm.Model       `swaggerignore:"true"`
	Diff             float64 `validate:"required"`
	TargetStrategyID uint
	ZoneStrategyID   uint
}

type SL struct {
	gorm.Model       `swaggerignore:"true"`
	Diff             float64 `validate:"required"`
	TargetStrategyID uint
	ZoneStrategyID   uint
}

// maybe rewrite this, could be an old issue, not sure tho
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
