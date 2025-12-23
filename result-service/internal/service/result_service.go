package service

import (
	"encoding/json"
	"log"

	"result-service/internal/config"
	"result-service/internal/event"
	"result-service/internal/model"
)

func CalculateTestResult(msgBody []byte) {

	var testResult model.TestResult
	if err := json.Unmarshal(msgBody, &testResult); err != nil {
		log.Println(" TestResult parse error:", err)
		return
	}

	redisValue, err := config.RedisClient.
		Get(config.Ctx, testResult.TestID).
		Result()

	
	if err != nil {
		event.PublishErrorEvent(model.ErrorEvent{
			TelegramID: testResult.TelegramID,
			TestID:     testResult.TestID,
			Error:      "test_id not found in redis",
		})
		return
	}

	var testData model.TestData
	if err := json.Unmarshal([]byte(redisValue), &testData); err != nil {
		event.PublishErrorEvent(model.ErrorEvent{
			TelegramID: testResult.TelegramID,
			TestID:     testResult.TestID,
			Error:      "test data parse error",
		})
		return
	}


	correctMap := make(map[string]string)
	for _, q := range testData.Questions {
		correctMap[q.ID] = q.Correct
	}

	score := 0
	for _, ans := range testResult.Answers {
		if correctMap[ans.QuestionID] == ans.CorrectAnswer {
			score++
		}
	}

	event.PublishUserResultEvent(model.UserResultEvent{
		TelegramID:    testResult.TelegramID,
		UserFirstName: testResult.UserFirstName,
		TeacherTelegramID: testData.TeacherTelegramID,
		TestID:        testResult.TestID,
		Score:         score,
		Total:         len(testData.Questions),
	})
}
