package database

import "github.com/tkuldeep/todo-backend/models"

type TodoRepo interface {
	Create(*models.Todo) error
	List() ([]models.Todo, error)
}

func (pd PostgreInstance) Create(todo *models.Todo) error {
	tx := pd.Db.Create(todo)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (pd PostgreInstance) List() ([]models.Todo, error) {
	todos := []models.Todo{}
	tx := pd.Db.Find(&todos)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return todos, nil
}
