package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tkuldeep/todo-backend/database"
	"github.com/tkuldeep/todo-backend/handlers"
)

func main() {

	// initialize todo app's depedencies like handlers and DB Connection
	todoApp := new(TodoAppContext)
	app := fiber.New()

	// Initialize default config
	app.Use(cors.New())

	// Configure cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://192.168.1.12:8081/",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	todoApp.FiberApp = app
	todoApp.TaskHandler = handlers.NewTodoApp(database.NewDBInstance())

	// setting up routes for todo app
	setupRoutes(todoApp)

	todoApp.FiberApp.Listen(":3000")
}
