package main

import (
	"fmt"
	"todolist/app"
	"todolist/controller"
	"todolist/logic"
	"todolist/repository"
)

func main() {
	postgres, redis := app.DatabaseConnection()
	todoRepo := repository.Repo(postgres, redis)

	todoService := logic.NewTodoService(todoRepo)
	todoController := controller.NewTodoController(todoService)
	todoController.GetTodos("1")
	fmt.Println("Hello World")
}
