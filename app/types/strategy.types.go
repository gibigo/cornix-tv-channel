package types

type Strategy struct {
	ID             uint `json:"-"`
	AllowCounter   bool
	Symbol         string
	TargetStrategy *TargetStrategy
	ZoneStrategy   *ZoneStrategy
	ChannelID      uint `json:"-"`
}

type ZoneStrategy struct {
	ID         uint `json:"-"`
	EntryStart float64
	EntryStop  float64
	TPs        []*TP
	SL         *SL
	IsBreakout bool
	StrategyID uint `json:"-"`
}

type TargetStrategy struct {
	ID         uint `json:"-"`
	Entries    []*Entry
	TPs        []*TP
	SL         *SL
	IsBreakout bool
	StrategyID uint `json:"-"`
}

type Entry struct {
	ID               uint `json:"-"`
	Diff             float64
	TargetStrategyID uint `json:"-"`
}

type TP struct {
	ID               uint `json:"-"`
	Diff             float64
	TargetStrategyID uint `json:"-"`
	ZoneStrategyID   uint `json:"-"`
}

type SL struct {
	ID               uint `json:"-"`
	Diff             float64
	TargetStrategyID uint `json:"-"`
	ZoneStrategyID   uint `json:"-"`
}

/* type Strategy struct {
	AllowCounter   bool
	Symbol         string
	TargetStrategy *TargetStrategy
	ZoneStrategy   *ZoneStrategy
}

type ZoneStrategy struct {
	EntryStart float64
	EntryStop  float64
	TPs        []*TP
	SL         *SL
	IsBreakout bool
}

type TargetStrategy struct {
	ID         uint
	Entries    []*Entry
	TPs        []*TP
	SL         *SL
	IsBreakout bool
}

type Entry struct {
	Diff             float64
	TargetStrategyID uint
}

type TP struct {
	Diff float64
}

type SL struct {
	Diff float64
} */

type AddStrategy struct {
	Symbol         string          `json:"symbol" validate:"required"`
	AllowCounter   bool            `json:"allowCounter"`
	TargetStrategy *TargetStrategy `json:"targetStrategy,omitempty"`
	ZoneStrategy   *ZoneStrategy   `json:"zoneStrategy,omitempty"`
}
