package response

import "practice/gocast/todo-cli/server/model"

type CreateTask struct {
	Task model.Task `json:"task"`
}
