package logic

import (
	"todolist/domain"
	"todolist/repository"
)

type TodoService interface {
	GetUserTodos(ID string) (domain.Todo, error)
	AddUserTodos(ID string) error
}
type TodoServiceImpl struct {
	Repo *repository.TodoRepo
}

func NewTodoService(repo *repository.TodoRepo) *TodoServiceImpl {
	return &TodoServiceImpl{Repo: repo}
}
