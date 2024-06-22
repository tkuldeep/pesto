package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tkuldeep/todo-backend/handlers"
)

type TodoAppContext struct {
	FiberApp    *fiber.App
	TodoHandler handlers.TodoHandler
}

func setupRoutes(app *TodoAppContext) {
	app.FiberApp.Get("/", handlers.Home)

	app.FiberApp.Post("/todos", app.TodoHandler.CreateTodo)
	app.FiberApp.Get("/todos", app.TodoHandler.ListTodos)
}
