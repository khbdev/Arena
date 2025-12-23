package model

type ArenaRequest struct {
	TelegramID int64  `json:"telegram_id"`
	Prompt     string `json:"prompt"`
	Count      int    `json:"count"`
}

type TestData struct {
    TeacherTelegramID int64 `json:"teacher_telegram_id"`
    TestID            string `json:"test_id"`
    Questions         []Question `json:"questions"`
}

type Question struct {
    ID       string   `json:"id"`
    Question string   `json:"question"`
    Options  []Option `json:"options"`
    Correct  string   `json:"correct"`
}

type Option struct {
    ID   string `json:"id"`
    Text string `json:"text"`
}


type TestAnswerRequest struct {
	TelegramID    int64          `json:"telegram_id"`
	TestID        string         `json:"test_id"`
	Answers       []UserAnswer   `json:"answers"`
	UserFirstName string         `json:"user_first_name"`
}

type UserAnswer struct {
	QuestionID    string `json:"question_id"`
	CorrectAnswer string `json:"correct_answer"`
}