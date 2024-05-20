package queries

import uuid "github.com/satori/go.uuid"

type GetTodoMessagesByTodoListIDQuery struct {
	TodoListID uuid.UUID `valide:"required"`
}

func NewGetTodoMessagesByTodoListIDQuery(todoListID uuid.UUID) *GetTodoMessagesByTodoListIDQuery {
	return &GetTodoMessagesByTodoListIDQuery{TodoListID: todoListID}
}
