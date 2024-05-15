package models

import "time"

type TodoList struct {
	ID                  string    `json:"id"`
	CreationTime        time.Time `json:"creation_time"`
	RefactoritionTime   time.Time `json:"refactorition_time"`
	DeletionTime        time.Time `json:"deletion_time"`
	CompletionOfPercent float64   `json:"completion_of_percent"`
}
