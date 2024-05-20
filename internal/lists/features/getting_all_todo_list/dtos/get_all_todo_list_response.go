package dtos

import (
	"demoapp/internal/lists/dtos"
)

type GetAllTodoListQueryResponse struct {
	TodoListsDto []*dtos.TodoListDto `json:"todolistdto"`
}
