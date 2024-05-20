package parse

import (
	"demoapp/internal/lists/models"
	"encoding/json"
	"fmt"
)

func ParseTodoListFromJson(response []byte) ([]models.TodoList, error) {
	var responseStruct struct {
		TodoList []models.TodoList `json:"todolists"`
	}
	err := json.Unmarshal(response, &responseStruct)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}
	return responseStruct.TodoList, nil
}
