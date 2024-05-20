package dtos

import (
	"demoapp/internal/lists/dtos"
	"demoapp/internal/messages/models"
)

type GetTodoListByIdQueryResponse struct {
	TodoListDto *dtos.TodoListDto   `json:"todolist"`
	TodoMessage *models.TodoMessage `json:"todomessage"`
}
