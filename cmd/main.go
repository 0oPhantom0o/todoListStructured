package main

import (
	"fmt"
	"todolist/app"
	"todolist/controller"
	"todolist/logger"
	"todolist/logic"
	"todolist/repository"
)

func main() {
	logger.Log()
	postgres, redis := app.DatabaseConnection()
	todoRepo := repository.NewRepo(postgres, redis)

	todoService := logic.NewTodoService(todoRepo)
	todoController := controller.NewTodoController(todoService)
	todoController.GetTodos("1")
	fmt.Println("Hello World")
}
