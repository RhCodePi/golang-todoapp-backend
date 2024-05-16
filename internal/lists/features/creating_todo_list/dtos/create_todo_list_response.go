package dtos

import uuid "github.com/satori/go.uuid"

type CreateTodoListCommandResponse struct {
	TodoListID uuid.UUID `json:"todolistid"`
}
