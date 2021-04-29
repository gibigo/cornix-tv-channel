package types

type Strategy struct {
	ID               uint
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
	ZoneStrategyID   uint `json:",omitempty"`
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

/* type Entry struct {
	Diff       float32 `json:"diff"`
	StrategyID int64   `json:",omitempty"` // maybe remove this
}

type TP struct {
	Diff       float32 `json:"diff"`
	StrategyID int64   `json:",omitempty"`
}

type SL struct {
	Diff       float32 `json:"diff"`
	StrategyID int64   `json:",omitempty"`
}

type Strategy struct {
	ID           int64    `json:"id,omitempty"`
	AllowCounter bool     `json:"allowCounter"`
	Entries      []*Entry `json:"entires"`
	TPs          []*TP    `json:"tps"`
	SL           *SL      `json:"sl"`
}
*/
type AddStrategy struct {
	AllowCounter bool     `json:"allowCounter"`
	Entries      []*Entry `json:"entires"`
	TPs          []*TP    `json:"tps"`
	SL           *SL      `json:"sl"`
}
