package model


type TestData struct {
	TelegramID int64       `json:"telegram_id"`
	TestID            string      `json:"test_id"`
	Questions         interface{} `json:"questions"` 
}