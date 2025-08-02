package service

import (
	"errors"
	"practice/gocast/todo-cli/internal/model"
	"practice/gocast/todo-cli/internal/security"
	"practice/gocast/todo-cli/internal/store"
)

type UserService struct {
	userStore store.UserStore
}

func NewUserService(userStore store.UserStore) *UserService {
	return &UserService{userStore: userStore}
}

func (userService *UserService) Register(username, password string) (*model.User, error) {
	if username == "" || password == "" {
		return nil, errors.New("username or password is empty")
	}
	exist, _ := userService.userStore.GetByUsername(username)
	if exist != nil {
		return nil, errors.New("username exist")
	}
	user := &model.User{
		Username: username,
		Password: security.HashPassword(password),
	}
	err := userService.userStore.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userService *UserService) Login(username, password string) (*model.User, error) {
	if username == "" || password == "" {
		return nil, errors.New("username or password is empty")
	}
	user, err := userService.userStore.GetByUsername(username)
	if err != nil || !security.CheckPassword(password, user.Password) {
		return nil, errors.New("invalid user")
	}
	return user, nil
}
