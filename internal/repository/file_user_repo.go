package repository

import (
	"encoding/json"
	"errors"
	"os"
	"practice/gocast/todo-cli/internal/model"
	"practice/gocast/todo-cli/internal/store"
	"strings"
	"sync"
)

type FileUserRepo struct {
	file  string
	mutex sync.Mutex
}

func NewFileUserRepo(file string) store.UserStore {
	return &FileUserRepo{file: file}
}

func (u *FileUserRepo) Create(user *model.User) error {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	fileContent, err := os.ReadFile(u.file)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	lines := strings.Split(string(fileContent), "\n")
	maxID := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		var existingUser model.User
		if err := json.Unmarshal([]byte(line), &existingUser); err == nil {
			if existingUser.Id > maxID {
				maxID = existingUser.Id
			}
		}
	}

	user.Id = maxID + 1
	file, err := os.OpenFile(u.file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	data, _ := json.Marshal(user)
	_, err = file.Write(append(data, '\n'))
	return err
}

func (u *FileUserRepo) GetByUsername(username string) (*model.User, error) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	file, err := os.ReadFile(u.file)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		var user model.User
		json.Unmarshal([]byte(line), &user)
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (u *FileUserRepo) GetById(id int) (*model.User, error) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	file, err := os.ReadFile(u.file)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		var user model.User
		json.Unmarshal([]byte(line), &user)
		if user.Id == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}
