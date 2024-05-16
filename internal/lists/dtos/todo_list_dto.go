package dtos

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TodoListDto struct {
	TodoListID          uuid.UUID `json:"todolistid"`
	CreationTime        time.Time `json:"creation_time"`
	RefactoritionTime   time.Time `json:"refactorition_time"`
	DeletionTime        time.Time `json:"deletion_time"`
	CompletionOfPercent float64   `json:"completion_of_percent"`
}
