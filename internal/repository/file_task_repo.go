package repository

import (
	"encoding/json"
	"errors"
	"os"
	"practice/gocast/todo-cli/internal/model"
	"strings"
	"sync"
)

type FileTaskRepo struct {
	file  string
	mutex sync.Mutex
}

func NewFileTaskRepo(file string) *FileTaskRepo {
	return &FileTaskRepo{file: file}
}

func (t *FileTaskRepo) Create(task *model.Task) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	fileContent, err := os.ReadFile(t.file)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	lines := strings.Split(string(fileContent), "\n")
	maxID := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		var existingTask model.Task
		if err := json.Unmarshal([]byte(line), &existingTask); err == nil {
			if existingTask.Id > maxID {
				maxID = existingTask.Id
			}
		}
	}
	task.Id = maxID + 1

	file, err := os.OpenFile(t.file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	data, _ := json.Marshal(task)
	_, err = file.Write(append(data, '\n'))
	return err
}

func (t *FileTaskRepo) Update(task *model.Task) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	file, err := os.ReadFile(t.file)
	if err != nil {
		return err
	}
	lines := strings.Split(string(file), "\n")
	newLines := []string{}
	found := false
	for _, line := range lines {
		if line == "" {
			continue
		}
		t := model.Task{}
		json.Unmarshal([]byte(line), &t)
		if t.Id == task.Id {
			data, _ := json.Marshal(task)
			newLines = append(newLines, string(data))
			found = true
		} else {
			newLines = append(newLines, line)
		}
	}
	if !found {
		return errors.New("task not found")
	}
	return os.WriteFile(t.file, []byte(strings.Join(newLines, "\n")), 0644)
}

func (t *FileTaskRepo) GetByUser(userId int) ([]*model.Task, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	file, err := os.ReadFile(t.file)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(file), "\n")
	tasks := []*model.Task{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		task := model.Task{}
		err := json.Unmarshal([]byte(line), &task)
		if err == nil && task.UserID == userId {
			tasks = append(tasks, &task)
		}
	}
	return tasks, nil
}

func (t *FileTaskRepo) GetById(id int) (*model.Task, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	file, err := os.ReadFile(t.file)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		task := model.Task{}
		err := json.Unmarshal([]byte(line), &task)
		if err == nil && task.Id == id {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}
