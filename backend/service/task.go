package service

import (
	"errors"

	"github.com/tkuldeep/todo-backend/database"
	"github.com/tkuldeep/todo-backend/models"
)

const (
	todoStatus     = "To Do"
	progressStatus = "In Progress"
	doneStatus     = "Done"
)

// TaskManager is service to handle task CRUD related business.
type TaskManager interface {
	Create(*models.Task) error
	List() ([]*models.Task, error)
}

type taskManager struct {
	repo database.TaskRepo
}

func NewTaskManager(repo database.TaskRepo) TaskManager {
	return taskManager{
		repo: repo,
	}
}

func (tm taskManager) Create(task *models.Task) error {
	status, err := getStatusValByStr(todoStatus)
	if err != nil {
		return err
	}
	task.Status = status
	tm.repo.Create(task)
	return nil
}

func (tm taskManager) List() ([]*models.Task, error) {
	tasks, err := tm.repo.List()
	if err != nil {
		return nil, err
	}
	for _, task := range tasks {
		status, _ := getStatusValByInt(task.Status)
		task.TaskStatus = status
	}

	return tasks, nil
}

// helper function to get task status by index or int repersentation
func getStatusValByInt(index int) (string, error) {
	status := map[int]string{
		1: todoStatus,
		2: progressStatus,
		3: doneStatus,
	}
	if _, ok := status[index]; ok {
		return status[index], nil
	}

	return "", errors.New("not valid status")
}

// helper function to get task int repersentation by string value.
// 0 value mean invalid task status
func getStatusValByStr(val string) (int, error) {
	status := map[string]int{
		todoStatus:     1,
		progressStatus: 2,
		doneStatus:     3,
	}

	if _, ok := status[val]; ok {
		return status[val], nil
	}

	return 0, errors.New("not valid status index")
}
