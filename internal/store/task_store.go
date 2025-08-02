package store

import "practice/gocast/todo-cli/internal/model"

type TaskStore interface {
	Create(task *model.Task) error
	Update(task *model.Task) error
	GetByUser(userId int) ([]*model.Task, error)
	GetById(id int) (*model.Task, error)
}
