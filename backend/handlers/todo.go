package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tkuldeep/todo-backend/database"
	"github.com/tkuldeep/todo-backend/models"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Welcome to TODO App!")
}

type TodoApp struct {
	Repo database.TodoRepo
}

type TodoHandler interface {
	CreateTodo(c *fiber.Ctx) error
	ListTodos(c *fiber.Ctx) error
}

func (ta TodoApp) CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := ta.Repo.Create(todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(todo)
}

func (ta TodoApp) ListTodos(c *fiber.Ctx) error {

	todos, err := ta.Repo.List()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(todos)
}
