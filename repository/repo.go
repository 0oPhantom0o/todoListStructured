package repository

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
	"todolist/domain"
)

type TodoRepository interface {
	GetTodos(ID string) (domain.Todo, error)
	AddTodo(ID string) error
}
type TodoRepo struct {
	PostgresDB  *sql.DB
	RedisClient *redis.Client
}

func NewRepo(postgresDB *sql.DB, redisClient *redis.Client) *TodoRepo {
	return &TodoRepo{
		PostgresDB:  postgresDB,
		RedisClient: redisClient,
	}
}
