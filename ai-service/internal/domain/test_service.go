package domain



type TestCreator interface {
	CreateTest(telegramID int64, questions interface{}) error
}