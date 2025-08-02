package store

import "practice/gocast/todo-cli/server/model"

type UserStore interface {
	Create(user *model.User) error
	GetByUsername(username string) (*model.User, error)
	GetById(id int) (*model.User, error)
}
