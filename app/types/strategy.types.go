package types

type Strategy struct {
	ID             uint            `json:"-"`
	AllowCounter   bool            `json:"allowCounter"`
	Symbol         string          `json:"symbol"`
	TargetStrategy *TargetStrategy `json:"targetStrategy,omitempty"`
	ZoneStrategy   *ZoneStrategy   `json:"zoneStrategy,omitempty"`
	ChannelID      uint            `json:"-"`
}

type ZoneStrategy struct {
	ID         uint    `json:"-"`
	EntryStart float64 `json:"entryStart"`
	EntryStop  float64 `json:"entryStop"`
	TPs        []*TP   `json:"tps"`
	SL         *SL     `json:"sl"`
	IsBreakout bool    `json:"isBreakout"`
	StrategyID uint    `json:"-"`
}

type TargetStrategy struct {
	ID         uint     `json:"-"`
	Entries    []*Entry `json:"entries"`
	TPs        []*TP    `json:"tps"`
	SL         *SL      `json:"sl"`
	IsBreakout bool     `json:"isBreakout"`
	StrategyID uint     `json:"-"`
}

type Entry struct {
	ID               uint    `json:"-"`
	Diff             float64 `json:"diff"`
	TargetStrategyID uint    `json:"-"`
}

type TP struct {
	ID               uint    `json:"-"`
	Diff             float64 `json:"diff"`
	TargetStrategyID uint    `json:"-"`
	ZoneStrategyID   uint    `json:"-"`
}

type SL struct {
	ID               uint    `json:"-"`
	Diff             float64 `json:"diff"`
	TargetStrategyID uint    `json:"-"`
	ZoneStrategyID   uint    `json:"-"`
}

type AddStrategy struct {
	Symbol         string          `json:"symbol" validate:"required"`
	AllowCounter   bool            `json:"allowCounter"`
	TargetStrategy *TargetStrategy `json:"targetStrategy,omitempty"`
	ZoneStrategy   *ZoneStrategy   `json:"zoneStrategy,omitempty"`
}
