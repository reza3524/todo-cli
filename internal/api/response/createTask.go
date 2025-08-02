package response

import "practice/gocast/todo-cli/internal/model"

type CreateTask struct {
	Task model.Task `json:"task"`
}
