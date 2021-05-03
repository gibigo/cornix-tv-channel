package types

type Strategy struct {
	ID             uint
	AllowCounter   bool
	Symbol         string
	TargetStrategy *TargetStrategy
	ZoneStrategy   *ZoneStrategy
	ChannelID      uint
}

type ZoneStrategy struct {
	ID         uint
	EntryStart float64
	EntryStop  float64
	TPs        []*TP
	SL         *SL
	IsBreakout bool
	StrategyID uint
}

type TargetStrategy struct {
	ID         uint
	Entries    []*Entry
	TPs        []*TP
	SL         *SL
	IsBreakout bool
	StrategyID uint
}

type Entry struct {
	ID               uint
	Diff             float64
	TargetStrategyID uint
	ZoneStrategyID   uint
}

type TP struct {
	ID               uint
	Diff             float64
	TargetStrategyID uint
	ZoneStrategyID   uint
}

type SL struct {
	ID               uint
	Diff             float64
	TargetStrategyID uint
	ZoneStrategyID   uint
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
