package app

import (
	"database/sql"
	"fmt"
	"sync"
)

var (
	dbOnce     sync.Once
	dbInstance *sql.DB
)

func NewPostgresConnection() (*sql.DB, error) {
	dbOnce.Do(func() {
		// Connect to the database only once
		connStr := "postgres://user:password@localhost/dbname?sslmode=disable"
		var err error
		dbInstance, err = sql.Open("postgres", connStr)
		if err != nil {
			fmt.Println(err)
		}
	})
	return dbInstance, nil
}
