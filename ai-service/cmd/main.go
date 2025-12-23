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

	
	rabbitmq.StartConsumer(rabbit.Channel, testService)


	select {} 
}
