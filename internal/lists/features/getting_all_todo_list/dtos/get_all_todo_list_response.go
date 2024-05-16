package dtos

import (
	"demoapp/internal/lists/dtos"
	"demoapp/internal/message/models"
)

type GetAllTodoListQueryResponse struct {
	TodoListsDto []*dtos.TodoListDto   `json:"todolistdto"`
	Messages     []*models.TodoMessage `json:"messages"`
}
