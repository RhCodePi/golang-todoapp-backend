package dtos

import uuid "github.com/satori/go.uuid"

type GetTodoMessagesByTodoListIDRequestDto struct {
	TodoListID uuid.UUID `json:"todolistid"`
}
