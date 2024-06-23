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
	List(map[string]string) ([]*models.Task, error)
	Update(*models.Task) error
	Delete(taskID int) error
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

	return tm.repo.Create(task)
}

func (tm taskManager) List(query map[string]string) ([]*models.Task, error) {
	dbQuery := map[string]int{}
	if _, ok := query["status"]; ok {
		status, err := getStatusValByStr(query["status"])
		if err != nil {
			return nil, err
		}
		dbQuery["status"] = status
	}
	tasks, err := tm.repo.List(dbQuery)
	if err != nil {
		return nil, err
	}
	for _, task := range tasks {
		status, _ := getStatusValByInt(task.Status)
		task.TaskStatus = status
	}

	return tasks, nil
}

func (tm taskManager) Update(task *models.Task) error {
	err := tm.repo.Update(task)
	if err != nil {
		return err
	}
	status, _ := getStatusValByInt(task.Status)
	task.TaskStatus = status

	return nil
}

func (tm taskManager) Delete(taskID int) error {
	return tm.repo.Delete(uint(taskID))
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
