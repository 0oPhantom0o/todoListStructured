package app

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func DatabaseConnection() (*sql.DB, *redis.Client) {
	newRedis, err := NewRedisConnection()
	if err != nil {
		log.Fatal().Err(err).Msg("redis connection failed")
	}
	mewPostgres, err := NewPostgresConnection()
	if err != nil {
		log.Fatal().Err(err).Msg("postgres connection failed")
	}
	return mewPostgres, newRedis

}
