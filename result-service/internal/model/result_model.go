package model

// user request

type Answer struct {
	QuestionID    string `json:"question_id"`
	CorrectAnswer string `json:"correct_answer"`
}

type TestResult struct {
	TelegramID    int64    `json:"telegram_id"`
	TestID        string   `json:"test_id"`
	Answers       []Answer `json:"answers"`
	UserFirstName string   `json:"user_first_name"`
}

// test model



type Option struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type Question struct {
	ID       string   `json:"id"`
	Question string   `json:"question"`
	Options  []Option `json:"options"`
	Correct  string   `json:"correct"`
}

type TestData struct {
	TeacherTelegramID int64      `json:"telegram_id"`
	TestID            string     `json:"test_id"`
	Questions         []Question `json:"questions"`
}

// user event



type UserResultEvent struct {
	TelegramID    int64  `json:"telegram_id"`
	TeacherTelegramID int64 `json:"teacher_telegram_id"`
	UserFirstName string `json:"user_first_name"`
	TestID        string `json:"test_id"`
	Score         int    `json:"score"`
	Total         int    `json:"total"`
}

type ErrorEvent struct {
	TelegramID int64  `json:"telegram_id"`
	TestID     string `json:"test_id"`
	Error      string `json:"error"`
}

