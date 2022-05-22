package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

func NewRedis() *redis.Client {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("addr-redis"),
		Password: os.Getenv("password-redis"), // no password set
		DB:       0,                           // use default DB
	})

	pong, err := rdb.Ping(context.Background()).Result()
	fmt.Println(pong)
	if err != nil {
		log.Println(err)
	}

	return rdb
}
