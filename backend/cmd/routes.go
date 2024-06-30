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

	app.FiberApp.Post("/tasks", app.TaskHandler.Create)
	app.FiberApp.Get("/tasks/:id", app.TaskHandler.Get)
	app.FiberApp.Get("/tasks", app.TaskHandler.List)
	app.FiberApp.Delete("/tasks/:id", app.TaskHandler.Delete)
	app.FiberApp.Put("/tasks/:id", app.TaskHandler.Update)
	app.FiberApp.Post("/tasks/:id/status", app.TaskHandler.ChangeStatus)
}
