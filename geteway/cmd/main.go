package main

import (
	"geteway/internal/bot"
	"geteway/internal/client"
	"geteway/internal/config"
	"geteway/internal/service"
)

func main() {
	config.InitEnv()

	telegramBot := config.InitTelegramBot()
 rabitt :=	config.InitRabbitMQ() 

	userClient := client.NewUserClient()
	
	userService := service.NewUserService(userClient)
resultService :=	 service.NewResultService()
telegramClient := client.NewTelegramClient()


	
	botApp := bot.New(telegramBot, userService, resultService, rabitt.Channel, telegramClient)

	botApp.Start()
}