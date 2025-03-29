package repository

import (
	"fmt"
	"time"
)

func (r *TodoRepo) AddTodo(ID string) error {
	fmt.Println(ID, time.Now(), "inserted")
	return nil
}
