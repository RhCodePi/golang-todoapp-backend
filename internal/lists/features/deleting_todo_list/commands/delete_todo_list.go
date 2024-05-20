package commands

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type DeleteTodoListCommand struct {
	TodoListId   uuid.UUID `json:"todolistid"`
	DeletionTime time.Time `json:"deletiontime"`
}

func NewDeleteTodoListCommand(todoListID uuid.UUID) *DeleteTodoListCommand {
	return &DeleteTodoListCommand{TodoListId: todoListID, DeletionTime: time.Now()}
}
