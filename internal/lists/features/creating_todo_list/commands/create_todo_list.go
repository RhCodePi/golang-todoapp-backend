package commands

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type CreateTodoListCommand struct {
	ID                  uuid.UUID `json:"id"`
	CreationTime        time.Time `json:"creation_time"`
	RefactoritionTime   time.Time `json:"refactorition_time"`
	DeletionTime        time.Time `json:"deletion_time"`
	CompletionOfPercent float64   `json:"completion_of_percent"`
}

func NewCreateTodoListCommand() *CreateTodoListCommand {
	return &CreateTodoListCommand{ID: uuid.NewV1(), CreationTime: time.Now(), RefactoritionTime: time.Time{}, DeletionTime: time.Time{}, CompletionOfPercent: 1}
}
