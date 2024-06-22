package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tkuldeep/todo-backend/handlers"
)

// TodoAppContext is entity for managing todo features.
type TodoAppContext struct {
	FiberApp    *fiber.App
	TaskHandler handlers.TaskHandler
}

func setupRoutes(app *TodoAppContext) {
	app.FiberApp.Get("/", handlers.Home)

	app.FiberApp.Post("/todos", app.TaskHandler.Create)
	app.FiberApp.Get("/todos", app.TaskHandler.List)
}
