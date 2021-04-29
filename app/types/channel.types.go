package types

type AddChannel struct {
	TelegramID int64 `json:"telegramId" validate:"required"`
}
