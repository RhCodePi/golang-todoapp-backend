package commands

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type CreateTodoMessageCommand struct {
	MessageID         uuid.UUID `json:"id"`
	TodoListID        uuid.UUID `json:"todo_list_id"`
	CreationTime      time.Time `json:"creation_time"`
	RefactoritionTime time.Time `json:"refactorition_time"`
	DeletionTime      time.Time `json:"deletion_time"`
	Content           string    `json:"content"`
	CompletionStatus  bool      `json:"completion_status"`
}

func NewCreateTodoListCommand(todoListID uuid.UUID, content string) *CreateTodoMessageCommand {
	return &CreateTodoMessageCommand{
		MessageID:         uuid.NewV1(),
		TodoListID:        todoListID,
		CreationTime:      time.Now(),
		RefactoritionTime: time.Time{},
		DeletionTime:      time.Time{},
		Content:           content,
		CompletionStatus:  false,
	}
}
