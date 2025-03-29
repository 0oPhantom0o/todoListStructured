package app

import (
	"database/sql"
)

func NewPostGresConnection() (*sql.DB, error) {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
