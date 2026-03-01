package main

import (
	"ai-service/internal/config"
	"ai-service/internal/rabbitmq"
	"ai-service/internal/service"
)

func main() {
	
	config.Init()
	config.InitRedis()

	
	rabbit := config.NewRabbitMQ() 


	testService := service.NewTestService(rabbit.Channel)

	messageProcress := service.NewMessageProcessor(testService, rabbit.Channel)
	
	rabbitmq.NewConsumer(rabbit.Channel, messageProcress)


	select {} 
}
