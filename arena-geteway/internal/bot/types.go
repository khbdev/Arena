// internal/bot/types.go
package bot

import "geteway/internal/model"

type UserState struct {
	Role      string // teacher | student
	Step      string // waiting_topic | waiting_count | waiting_test_id | solving_test
	Topic     string
	TestID    string
	Questions []model.Question
	Current   int
	Answers   []model.UserAnswer
	FirstName string
}