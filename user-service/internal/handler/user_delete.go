package handler

import (
	"context"
	"user-service/internal/usecase"

	pb "github.com/khbdev/arena-startup-proto/proto/user-delete"
)

// TelegramHandler - gRPC handler
type TelegramHandler struct {
	pb.UnimplementedTelegramServiceServer
	usecase usecase.UserUsecase
}

// Constructor
func NewTelegramHandler(u usecase.UserUsecase) *TelegramHandler {
	return &TelegramHandler{
		usecase: u,
	}
}

// CheckTelegramID implements TelegramService gRPC
func (h *TelegramHandler) CheckTelegramID(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	telegramID := req.GetTelegramId()

	deleted, err := h.usecase.DeleteUserByTelegramID(telegramID)
	if err != nil {
		// Xatolik bo'lsa false qaytaramiz
		return &pb.CheckResponse{
			Message: "false",
		}, nil
	}

	if deleted {
		return &pb.CheckResponse{
			Message: "muvaffaqiyatli o'chirildi /start bosib tekshring", // muvaffaqiyatli o'chirildi
		}, nil
	} else {
		return &pb.CheckResponse{
			Message: "user topilmadi", // user topilmadi
		}, nil
	}
}
