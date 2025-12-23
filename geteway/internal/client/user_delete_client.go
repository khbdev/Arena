package client

import (
	"context"
	"time"

	pb "github.com/khbdev/arena-startup-proto/proto/user-delete"
	grpc_connection "geteway/internal/grpcConnect"
)

type TelegramClient struct {
	client pb.TelegramServiceClient
}

func NewTelegramClient() *TelegramClient {
	conn := grpc_connection.Connect("USER_SERVICE") 
	return &TelegramClient{
		client: pb.NewTelegramServiceClient(conn),
	}
}


func (c *TelegramClient) CheckTelegramID(telegramID int64) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := c.client.CheckTelegramID(ctx, &pb.CheckRequest{
		TelegramId: telegramID,
	})
	if err != nil {
		return "", err
	}

	return resp.GetMessage(), nil
}
