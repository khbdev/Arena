package model

type TestRequest struct {
	TelegramID int64  `json:"telegram_id"`
	Prompt     string `json:"prompt"`
	Count      int    `json:"count"`
}