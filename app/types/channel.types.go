package types

type AddChannel struct {
	TelegramID int64 `json:"telegramId" validate:"required"`
}

type Channel struct {
	ID       uint       `json:"id"`
	Telegram int64      `json:"telegramId"`
	Strategy []Strategy `json:"strategies"`
}
