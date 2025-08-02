package service

import (
	"errors"
	"practice/gocast/todo-cli/server/model"
	"practice/gocast/todo-cli/server/store"
)

type TaskService struct {
	taskStore store.TaskStore
}

func NewTaskService(taskStore store.TaskStore) *TaskService {
	return &TaskService{taskStore: taskStore}
}

func (taskService *TaskService) AddTask(title string, userId, categoryId int) (*model.Task, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}
	task := &model.Task{
		Title:      title,
		UserID:     userId,
		CategoryID: categoryId,
		Completed:  false,
	}
	err := taskService.taskStore.Create(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (taskService *TaskService) ListTasks(userId int) ([]*model.Task, error) {
	return taskService.taskStore.GetByUser(userId)
}

func (taskService *TaskService) ToggleTask(taskId int, completed bool) error {
	task, err := taskService.taskStore.GetById(taskId)
	if err != nil {
		return err
	}
	task.Completed = completed
	return taskService.taskStore.Update(task)
}
