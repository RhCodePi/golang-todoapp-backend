package queries

import (
	"context"
	"demoapp/internal/lists"
	"demoapp/internal/lists/features/getting_todo_list_by_id/dtos"
	"demoapp/internal/lists/repository"
	"fmt"

	"github.com/pkg/errors"
)

type GetTodoListByIdQueryHandler struct {
	mockRepository *repository.MockTodoListRepository
}

func NewGetTodoListByIdQueryHandler(mockRepository *repository.MockTodoListRepository) *GetTodoListByIdQueryHandler {
	return &GetTodoListByIdQueryHandler{mockRepository: mockRepository}
}

func (q *GetTodoListByIdQueryHandler) Handle(ctx context.Context, query *GetTodoListByIdQuery) (*dtos.GetTodoListByIdQueryResponse, error) {
	todoList, err := q.mockRepository.GetTodoListById(ctx, query.TodoListId)

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Todo List with id %s not found", query.TodoListId))
	}

	todoListDto := lists.MapTodoListToTodoListDto(todoList, nil)

	return &dtos.GetTodoListByIdQueryResponse{TodoListDto: todoListDto}, nil
}
