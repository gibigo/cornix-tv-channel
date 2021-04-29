package types

type Channel struct {
	ID       int64     `json:"id"`
	UserName string    `json:"user"`
	Strategy *Strategy `json:"strategy"`
}
