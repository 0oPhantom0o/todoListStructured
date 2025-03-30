package app

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

var (
	client *redis.Client
)

func NewRedisConnection() (*redis.Client, error) {
	dbOnce.Do(func() {

		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		ctx := context.Background()
		result, err := client.Ping(ctx).Result()
		if err != nil {
			log.Err(err).Msgf("Redis connection failed  %s", result)
		}
	})
	return client, nil
}
