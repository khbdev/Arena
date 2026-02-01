package model



type NotificationEvent struct {
	TelegramID int64  `json:"telegram_id"`
	TestID     string `json:"test_id"`
}
