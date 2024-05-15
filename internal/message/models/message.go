package models

import "time"

type TodoMessage struct {
	ID                string    `json:"id"`
	TodoListID        string    `json:"todo_list_id"`
	CreationTime      time.Time `json:"creation_time"`
	RefactoritionTime time.Time `json:"refactorition_time"`
	DeletionTime      time.Time `json:"deletion_time"`
	Content           string    `json:"content"`
	CompletionStatus  bool      `json:"completion_status"`
}
