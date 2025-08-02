package store

import "practice/gocast/todo-cli/internal/model"

type CategoryStore interface {
	Create(category *model.Category) error
	GetByUser(userId int) ([]*model.Category, error)
}
