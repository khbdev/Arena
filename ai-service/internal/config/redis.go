package config

import (
	"context"
	"fmt"
	"log"
	

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)
 
func InitRedis() {


	addr := getEnv("REDIS_HOST", "localhost:6379")
	pass := getEnv("REDIS_PASSWORD", "")

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})


	pong, err := RedisClient.Ping(Ctx).Result()
if err != nil {
	log.Fatalf("Failed to connect to Redis: %v", err)
}
fmt.Println("Connected to Redis:", pong)

}


func getEnv(key, fallback string) string {
	if val, ok := lookupEnv(key); ok {
		return val
	}
	return fallback
}


func lookupEnv(key string) (string, bool) {
	envs, _ := godotenv.Read()
	val, ok := envs[key]
	return val, ok
}