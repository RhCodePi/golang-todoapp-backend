package queries

import (
	"context"
	"demoapp/internal/messages/features/getting_todo_messages_by_todo_list_id/dtos"
	"demoapp/internal/messages/repository"
	"fmt"

	"github.com/pkg/errors"
)

type GetTodoMessagesByTodoListIDQueryHandler struct {
	mockRepository *repository.MockTodoMessageRepository
}

func NewGetTodoMessagesByTodoListIdQueryHandler(mockRepository *repository.MockTodoMessageRepository) *GetTodoMessagesByTodoListIDQueryHandler {
	return &GetTodoMessagesByTodoListIDQueryHandler{mockRepository: mockRepository}
}

func (q *GetTodoMessagesByTodoListIDQueryHandler) Handle(ctx context.Context, query *GetTodoMessagesByTodoListIDQuery) (*dtos.GetTodoMessagesByTodoListIDQueryResponse, error) {
	todoMessage, err := q.mockRepository.GetMessagesByTodoListId(ctx, query.TodoListID)

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Todo List with id %s not found", query.TodoListID))
	}

	return &dtos.GetTodoMessagesByTodoListIDQueryResponse{TodoMessages: todoMessage}, nil
}
