package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"todolist/app"
)

func TestSingletonPostgres(t *testing.T) {
	db1, _ := app.NewPostgresConnection()
	db2, _ := app.NewPostgresConnection()

	assert.Equal(t, db1, db2, "must have same result")
}
func TestSingletonRedis(t *testing.T) {
	redis1, _ := app.NewRedisConnection()
	redis2, _ := app.NewRedisConnection()

	assert.Equal(t, redis1, redis2, "must have same result")
}
