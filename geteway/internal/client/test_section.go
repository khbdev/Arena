package client

import (
	"context"
	"time"

	grpc_connection "geteway/internal/grpcConnect"

	testsectionpb "github.com/khbdev/arena-startup-proto/proto/test-section"
)

type ResultClient struct {
	client testsectionpb.ResultServiceClient
}

func NewResultClient() *ResultClient {
	conn := grpc_connection.Connect("RESULT_SERVICE")
	return &ResultClient{
		client: testsectionpb.NewResultServiceClient(conn),
	}
}

func (r *ResultClient) GetUserTestResult(
	telegramID int64,
	testID string,
) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := r.client.GetUserTestResult(ctx, &testsectionpb.GetUserTestResultRequest{
		TelegramId: telegramID,
		TestId:     testID,
	})
	if err != nil {
		return "", err
	}

	return resp.JsonData, nil
}
