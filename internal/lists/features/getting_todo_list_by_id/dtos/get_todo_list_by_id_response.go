package dtos

import (
	"demoapp/internal/lists/dtos"
	"demoapp/internal/message/models"
)

type GetTodoListByIdQueryResponse struct {
	TodoListDto *dtos.TodoListDto   `json:"todolist"`
	Message     *models.TodoMessage `json:"message"`
}
