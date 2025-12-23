package client

import (
	"context"
	grpc_connection "geteway/internal/grpcConnect"

	"time"

	"github.com/khbdev/arena-startup-proto/proto/user"
)

type UserClient struct {
	client userpb.UserServiceClient
}

func NewUserClient() *UserClient {
	conn := grpc_connection.Connect("USER_SERVICE")
	return &UserClient{
		client: userpb.NewUserServiceClient(conn),
	}
}


func (u *UserClient) GetUserByTelegramID(telegramID int64) (*userpb.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := u.client.GetUserByTelegramId(ctx, &userpb.GetUserRequest{
		TelegramId: telegramID,
	})
	if err != nil {
		return nil, err
	}


	if resp.User == nil {
		return nil, nil
	}

	return resp.User, nil
}


func (u *UserClient) CreateUser(telegramID int64, role, firstName string) (*userpb.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := u.client.CreateUser(ctx, &userpb.CreateUserRequest{
		TelegramId: telegramID,
		Role:       role,
		FirstName:  firstName,
	})
	if err != nil {
		return nil, err
	}

	return resp.User, nil
}
