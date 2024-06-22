package models

type TodoManager interface {
	Create(Todo) error
}

type todoManager struct {
}

func (tm todoManager) Create(todo Todo) error {
	return nil
}
