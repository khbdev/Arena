package service

import (
	"ai-service/internal/config"
	"ai-service/internal/event"
	"ai-service/internal/model"

	"ai-service/internal/util"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

type TestService struct {
	RabbitCh *amqp.Channel
}



func NewTestService(ch *amqp.Channel) *TestService {
	return &TestService{
		RabbitCh: ch,
	}
}

func (s *TestService) CreateTest(TelegramID int64, questions interface{}) error {
	testID := util.GenerateTestID()

	data := model.TestData{
		TelegramID: TelegramID,
		TestID:            testID,
		Questions:         questions,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ttlHours := 2
	if v := os.Getenv("REDIS_TTL"); v != "" {
		if t, err := strconv.Atoi(v); err == nil && t > 0 {
			ttlHours = t
		}
	}

	if err := config.RedisClient.Set(
		config.Ctx,
		testID,
		jsonData,
		time.Duration(ttlHours)*time.Hour,
	).Err(); err != nil {
		return err
	}

	fmt.Println(" Redis saqlandi, key:", testID)

	
	return event.PublishNotification(
		s.RabbitCh,
		TelegramID,
		testID,
	)
}

