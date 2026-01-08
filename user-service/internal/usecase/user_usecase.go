package usecase

import (
	domain "user-service/internal/domen"
)

// Usecase interface
type UserUsecase interface {
	GetUserByTelegramID(telegramID int64) (*domain.User, error)
	CreateUser(input *domain.User) (*domain.User, error)
	DeleteUserByTelegramID(telegramID int64) (bool, error) 
}

// usecase struct
type userUsecase struct {
	repo domain.UserRepository
}

// Constructor
func NewUserUsecase(repo domain.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

// Get user
func (u *userUsecase) GetUserByTelegramID(telegramID int64) (*domain.User, error) {
	return u.repo.GetByTelegramID(telegramID)
}

// Create user
func (u *userUsecase) CreateUser(input *domain.User) (*domain.User, error) {
	if input.TelegramID == 0 || input.FirstName == "" || input.Role == "" {
		return nil, ErrInvalidUserInput
	}
	return u.repo.Create(input)
}

// ✅ Delete user by TelegramID
func (u *userUsecase) DeleteUserByTelegramID(telegramID int64) (bool, error) {
	deleted, err := u.repo.DeleteByTelegramID(telegramID)
	if err != nil {
		return false, err
	}
	return deleted, nil
}

// Error definition
var ErrInvalidUserInput = &UserError{Message: "invalid user input"}

type UserError struct {
	Message string
}

func (e *UserError) Error() string {
	return e.Message
}
