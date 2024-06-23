package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tkuldeep/todo-backend/database"
	"github.com/tkuldeep/todo-backend/models"
	"github.com/tkuldeep/todo-backend/service"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Welcome to TODO App!")
}

// TodoApp manages handlers info for task service
type TodoApp struct {
	taskManager service.TaskManager
}

type TaskHandler interface {
	Create(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

// Return instance of todo app
func NewTodoApp(repo database.TaskRepo) TaskHandler {
	return TodoApp{
		taskManager: service.NewTaskManager(repo),
	}
}

func (ta TodoApp) Create(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Call task manager service to create task
	err := ta.taskManager.Create(task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(task)
}

func (ta TodoApp) List(c *fiber.Ctx) error {
	statusQuery := c.Query("status")
	query := map[string]string{}
	if statusQuery != "" {
		query["status"] = statusQuery
	}

	// Call task manager servie to fetch list of tasks.
	tasks, err := ta.taskManager.List(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(tasks)
}

func (ta TodoApp) Update(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	taskIDStr := c.Params("id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	task.ID = uint(taskID)

	// Call task manager servie to fetch list of tasks.
	err = ta.taskManager.Update(task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(task)
}

func (ta TodoApp) Delete(c *fiber.Ctx) error {
	taskIDStr := c.Params("id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Call task manager servie to delete of taks given ID.
	err = ta.taskManager.Delete(taskID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	c.Status(http.StatusNoContent)

	return nil
}
