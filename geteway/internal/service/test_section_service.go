package service

import (
	"encoding/json"
	"fmt"

	"geteway/internal/client"
	"geteway/internal/model"
)

type ResultService struct {
	resultClient *client.ResultClient
}

func NewResultService() *ResultService {
	return &ResultService{
		resultClient: client.NewResultClient(),
	}
}


func (s *ResultService) GetUserTestResult(
	telegramID int64,
	testID string,
) (*model.TestData, error) {

	jsonData, err := s.resultClient.GetUserTestResult(telegramID, testID)
	if err != nil {
		return nil, err
	}

	if jsonData == "" {
		return nil, fmt.Errorf("test result not found")
	}

	var data model.TestData
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		return nil, err
	}

	return &data, nil
}
