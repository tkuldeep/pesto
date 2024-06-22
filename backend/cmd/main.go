package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tkuldeep/todo-backend/database"
	"github.com/tkuldeep/todo-backend/handlers"
)

func main() {

	// initialize todo app's depedencies like handlers and DB Connection
	todoApp := new(TodoAppContext)
	todoApp.FiberApp = fiber.New()
	todoApp.TodoHandler = handlers.TodoApp{
		Repo: database.NewDBInstance(),
	}

	setupRoutes(todoApp)

	todoApp.FiberApp.Listen(":3000")
}
