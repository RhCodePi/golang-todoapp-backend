package dtos

import uuid "github.com/satori/go.uuid"

type GetTodoListByIdRequestDto struct {
	TodoListID uuid.UUID `param:"id" json:"-"`
}
