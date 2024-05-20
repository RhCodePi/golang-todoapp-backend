package dtos

import "demoapp/internal/messages/models"

type GetTodoMessagesByTodoListIDQueryResponse struct {
	TodoMessages []*models.TodoMessage `json:"todomessages"`
}
