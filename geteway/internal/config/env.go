package config

import (
	"log"

	"github.com/joho/godotenv"
)


func InitEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system env")
	}
}
