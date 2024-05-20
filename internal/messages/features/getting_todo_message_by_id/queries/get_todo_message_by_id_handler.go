package queries

import (
	"context"
	"demoapp/internal/messages/features/getting_todo_message_by_id/dtos"
	"demoapp/internal/messages/repository"
	"fmt"

	"github.com/pkg/errors"
)

type GetTodoMessageByIdQueryHandler struct {
	mockRepository *repository.MockTodoMessageRepository
}

func NewGetTodoMessageByIdQueryHandler(mockRepository *repository.MockTodoMessageRepository) *GetTodoMessageByIdQueryHandler {
	return &GetTodoMessageByIdQueryHandler{mockRepository: mockRepository}
}

func (q *GetTodoMessageByIdQueryHandler) Handle(ctx context.Context, query *GetTodoMessageByIdQuery) (*dtos.GetTodoMessageByIdQueryResponse, error) {
	todoMessage, err := q.mockRepository.GetTodoMessageById(ctx, query.TodoMessageID)

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Todo List with id %s not found", query.TodoMessageID))
	}

	return &dtos.GetTodoMessageByIdQueryResponse{TodoMessage: *todoMessage}, nil
}
