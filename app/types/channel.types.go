package types

type Channel struct {
	ID       int64     `json:"id"`
	UserName string    `json:"user"`
	Strategy *Strategy `json:"strategy"`
}

type Entry struct {
	Diff       float32 `json:"diff"`
	StrategyID int64   `json:",omitempty"`
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
	ID      int64    `json:"id,omitempty"`
	Entries []*Entry `json:"entires"`
	TPs     []*TP    `json:"tps"`
	SL      *SL      `json:"sl"`
}
