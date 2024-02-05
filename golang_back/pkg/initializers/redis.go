package initializers

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	*redis.Client
}

func NewRedisClient() (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	res, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	println(res)

	return &RedisClient{client}, nil
}
