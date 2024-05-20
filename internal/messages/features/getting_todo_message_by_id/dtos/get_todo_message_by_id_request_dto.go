package dtos

import uuid "github.com/satori/go.uuid"

type GetTodoMessageByIdRequestDto struct {
	TodoMessageID uuid.UUID `json:"todomessageid"`
}
