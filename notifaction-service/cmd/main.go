package main

import (
	"log"
	"os"
	"os/signal"
	"notifaction-service/internal/config"
	"notifaction-service/internal/handler"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env: %s", err)
	}

	config.InitRabbitMQ()
	defer config.CloseRabbitMQ()

	config.InitTelegramBot()


	go handler.StartConsumer("notifications_queue", 5)
	go handler.StartUserStaticConsumer("user_statik_queue", 5)


	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println(" Shutting down...")
}
