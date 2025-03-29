package repository

import (
	"fmt"
	"time"
	"todolist/domain"
)

func (r *TodoRepo) GetTodos(ID string) (domain.Todo, error) {
	fmt.Println(ID, time.Now(), "selected")
	return domain.Todo{
		Id:        ID,
		Title:     "2",
		Completed: false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}, nil
}
