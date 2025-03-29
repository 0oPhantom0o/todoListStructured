package app

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
)

func DatabaseConnection() (*sql.DB, *redis.Client) {
	redis, err := NewRedisConnection()
	if err != nil {
		panic(err)
	}
	postgres, err := NewPostGresConnection()
	if err != nil {
		panic(err)
	}
	return postgres, redis

}
