package database

import "github.com/tkuldeep/todo-backend/models"

// TaskRepo contains methods related to task database operation like CRUD
type TaskRepo interface {
	Create(*models.Task) error
	List() ([]*models.Task, error)
}

func (pd PostgreInstance) Create(task *models.Task) error {
	tx := pd.Db.Create(task)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (pd PostgreInstance) List() ([]*models.Task, error) {
	tasks := []*models.Task{}
	tx := pd.Db.Find(&tasks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tasks, nil
}
