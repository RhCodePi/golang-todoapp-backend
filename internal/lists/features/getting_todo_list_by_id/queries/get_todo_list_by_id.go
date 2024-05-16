package queries

import uuid "github.com/satori/go.uuid"

type GetTodoListByIdQuery struct {
	TodoListId uuid.UUID `validate:"required"`
}

func NewGetTodoListByIdQuery(todoListId uuid.UUID) *GetTodoListByIdQuery {
	return &GetTodoListByIdQuery{TodoListId: todoListId}
}
