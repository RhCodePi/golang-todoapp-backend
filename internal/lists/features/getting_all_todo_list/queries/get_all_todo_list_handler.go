package queries

import (
	"context"
	"demoapp/internal/lists"
	"demoapp/internal/lists/features/getting_all_todo_list/dtos"
	"demoapp/internal/lists/repository"
)

type GetAllTodoListQueryHandler struct {
	mockRepository *repository.MockTodoListRepository
}

func NewGetAllTodoListQueryHandler(mockRepository *repository.MockTodoListRepository) *GetAllTodoListQueryHandler {
	return &GetAllTodoListQueryHandler{mockRepository: mockRepository}
}

func (q *GetAllTodoListQueryHandler) Handle(ctx context.Context, query *GetAllTodoListQuery) (*dtos.GetAllTodoListQueryResponse, error) {
	todoLists, err := q.mockRepository.GetAllTodoList()

	if err != nil {
		return nil, err
	}

	todoListsDto := lists.MapTodoListsToTodoListDto(todoLists)

	return &dtos.GetAllTodoListQueryResponse{TodoListsDto: todoListsDto}, nil
}
