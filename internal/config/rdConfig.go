package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func RedisConfig() (client *redis.Client) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Redis config error:", err)
			client = nil
		}
	}()

	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	username := os.Getenv("REDIS_USERNAME")
	password := os.Getenv("REDIS_PASSWORD")

	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Username: username,
		Password: password,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return client
}
