package dtos

import "time"

type DeleteTodoListCommandResponse struct {
	DeletionTime time.Time `json:"deletiontime"`
}
