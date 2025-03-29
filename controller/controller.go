package controller

import (
	"fmt"
	"todolist/logic"
)

type TodoController struct {
	Service logic.TodoService
}

func NewTodoController(service logic.TodoService) *TodoController {
	return &TodoController{Service: service}
}

func (c *TodoController) GetTodos(id string) {
	todos, err := c.Service.GetUserTodos(id)
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println(todos)
}
