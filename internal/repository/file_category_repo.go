package repository

import (
	"encoding/json"
	"os"
	"practice/gocast/todo-cli/internal/model"
	"strings"
	"sync"
)

type FileCategoryRepo struct {
	file  string
	mutex sync.Mutex
}

func NewFileCategoryRepo(file string) *FileCategoryRepo {
	return &FileCategoryRepo{file: file}
}

func (c *FileCategoryRepo) Create(category *model.Category) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	fileContent, err := os.ReadFile(c.file)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	lines := strings.Split(string(fileContent), "\n")
	maxID := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		var existingCategory model.Category
		if err := json.Unmarshal([]byte(line), &existingCategory); err == nil {
			if existingCategory.Id > maxID {
				maxID = existingCategory.Id
			}
		}
	}
	category.Id = maxID + 1

	file, err := os.OpenFile(c.file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	data, _ := json.Marshal(category)
	_, err = file.Write(append(data, '\n'))
	return err
}

func (c *FileCategoryRepo) GetByUser(userId int) ([]*model.Category, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	file, err := os.ReadFile(c.file)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(file), "\n")
	categories := []*model.Category{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		category := model.Category{}
		err := json.Unmarshal([]byte(line), &category)
		if err == nil && category.UserID == userId {
			categories = append(categories, &category)
		}
	}
	return categories, nil
}
