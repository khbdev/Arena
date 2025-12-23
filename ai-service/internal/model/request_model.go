package model

type TestRequest struct {
	TelegramID int64  `json:"telegram_id"`
	Prompt     string `json:"prompt"`
	Count      int    `json:"count"`
}

type TestData struct {
	TelegramID int64       `json:"telegram_id"`
	TestID            string      `json:"test_id"`
	Questions         interface{} `json:"questions"` 


}

type NotificationEvent struct {
	TelegramID int64  `json:"telegram_id"`
	TestID     string `json:"test_id"`
}
