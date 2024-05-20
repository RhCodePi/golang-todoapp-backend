package queries

import uuid "github.com/satori/go.uuid"

type GetTodoMessageByIdQuery struct {
	TodoMessageID uuid.UUID `validate:"required"`
}

func NewGetTodoMessageByIdQuery(todoMessageID uuid.UUID) *GetTodoMessageByIdQuery {
	return &GetTodoMessageByIdQuery{TodoMessageID: todoMessageID}
}
