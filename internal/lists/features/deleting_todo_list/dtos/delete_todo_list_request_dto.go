package dtos

import uuid "github.com/satori/go.uuid"

type DeleteTodoListRequestDto struct {
	TodoListId uuid.UUID `json:"todolistid"`
}
