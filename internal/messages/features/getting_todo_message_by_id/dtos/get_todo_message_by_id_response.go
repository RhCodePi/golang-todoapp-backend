package dtos

import (
	"demoapp/internal/messages/models"
)

type GetTodoMessageByIdQueryResponse struct {
	TodoMessage models.TodoMessage `json:"todomessage"`
}
