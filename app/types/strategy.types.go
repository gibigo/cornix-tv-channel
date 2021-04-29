package types

import "github.com/gibigo/cornix-tv-channel/app/dal"

type Entry struct {
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

type AddStrategy struct {
	AllowCounter bool         `json:"allowCounter"`
	Entries      []*dal.Entry `json:"entires"`
	TPs          []*dal.TP    `json:"tps"`
	SL           *dal.SL      `json:"sl"`
}
