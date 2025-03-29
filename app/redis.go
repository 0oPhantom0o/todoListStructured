package app

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func NewRedisConnection() (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
