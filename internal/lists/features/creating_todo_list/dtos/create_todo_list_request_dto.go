package dtos

import uuid "github.com/satori/go.uuid"

type CreateTodoListRequestDto struct {
	TodoListID uuid.UUID `json:"todolistid"`
}
