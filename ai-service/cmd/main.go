package main

import (
	"ai-service/internal/config"
	"ai-service/internal/rabbitmq"
	"ai-service/internal/service"
	"ai-service/pkg"
)

func main() {
	
	pkg.Init()
	config.InitRedis()

	
	rabbit := config.NewRabbitMQ() 


	testService := service.NewTestService(rabbit.Channel)

	messageProcress := service.NewMessageProcessor(testService, rabbit.Channel)
	
 consumer := 	rabbitmq.NewConsumer(rabbit.Channel, messageProcress)

consumer.Start()
	
	select {} 
}
