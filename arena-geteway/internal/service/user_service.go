package service

import (
	"geteway/internal/client"
	"log"

	"github.com/khbdev/arena-startup-proto/proto/user"
)

type UserService struct {
	client *client.UserClient
}


func NewUserService(c *client.UserClient) *UserService {
	return &UserService{
		client: c,
	}
}


func (s *UserService) GetByTelegramID(telegramID int64) (*userpb.User, error) {
	user, err := s.client.GetUserByTelegramID(telegramID)
	if err != nil {
		log.Println("UserService GetByTelegramID error:", err)
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return user, nil
}

func (s *UserService) CreateUser(telegramID int64, role, firstName string) (*userpb.User, error) {
	user, err := s.client.CreateUser(telegramID, role, firstName)
	if err != nil {
		log.Println("UserService CreateUser error:", err)
		return nil, err
	}

	return user, nil
}
