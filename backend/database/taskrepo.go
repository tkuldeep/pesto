package database

import (
	"fmt"

	"github.com/tkuldeep/todo-backend/models"
	"gorm.io/gorm"
)

// TaskRepo contains methods related to task database operation like CRUD
type TaskRepo interface {
	Create(*models.Task) error
	List(query map[string]int) ([]*models.Task, error)
	Update(*models.Task) error
	Delete(taskID uint) error
}

func (pd PostgreInstance) Create(task *models.Task) error {
	tx := pd.Db.Create(task)

	return tx.Error
}

func (pd PostgreInstance) List(query map[string]int) ([]*models.Task, error) {
	tasks := []*models.Task{}
	fmt.Println(query)
	var tx *gorm.DB
	if _, ok := query["status"]; ok {
		tx = pd.Db.Where("status = ?", query["status"]).Find(&tasks)
	} else {
		tx = pd.Db.Find(&tasks)
	}

	if tx.Error != nil {
		return nil, tx.Error
	}
	return tasks, nil
}

func (pd PostgreInstance) Update(task *models.Task) error {
	oldTask := new(models.Task)
	oldTask.ID = task.ID
	tx := pd.Db.First(oldTask)
	if tx.Error != nil {
		return tx.Error
	}

	task.Status = oldTask.Status
	task.CreatedAt = oldTask.CreatedAt
	tx = pd.Db.Save(task)

	return tx.Error
}

func (pd PostgreInstance) Delete(taskID uint) error {
	task := new(models.Task)
	task.ID = taskID
	tx := pd.Db.Delete(task)

	return tx.Error
}
