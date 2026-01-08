package main

import (
	"net"
	"os"

	"user-service/internal/config"
	"user-service/internal/handler"
	"user-service/internal/repostroy/mysql"
	"user-service/internal/usecase"

	"github.com/joho/godotenv"
	userpb "github.com/khbdev/arena-startup-proto/proto/user"
	userdeletepb "github.com/khbdev/arena-startup-proto/proto/user-delete"
	"google.golang.org/grpc"
)

func main() {
	// Load .env (error bilan shug‘ullanmaymiz)
	_ = godotenv.Load()

	// MySQL connection
	db := config.NewMySQLConnection()
      
	   
	// Repository
	userRepo := mysql.NewUserRepository(db)

	// Usecase
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Handler
	userHandler := handler.NewUserHandler(userUsecase)
	userDeleteHandler := handler.NewTelegramHandler(userUsecase)

	// gRPC server
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50051"
	}

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, userHandler)
	userdeletepb.RegisterTelegramServiceServer(grpcServer, userDeleteHandler)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
