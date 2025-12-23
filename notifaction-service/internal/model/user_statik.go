package model

type TestResult struct {
	TelegramID       int64  `json:"telegram_id"`
	TeacherTelegramID int64  `json:"teacher_telegram_id"`
	UserFirstName    string `json:"user_first_name"`
	TestID           string `json:"test_id"`
	Score            int    `json:"score"`
	Total            int    `json:"total"`
}
