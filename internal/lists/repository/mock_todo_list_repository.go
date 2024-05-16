package repository

import (
	"context"
	"demoapp/internal/lists/models"
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type MockTodoListRepository struct {
}

var todoLists []*models.TodoList

func init() {

	var newId = uuid.NewV1()
	var newId2 = uuid.NewV1()

	todoLists = []*models.TodoList{
		{
			TodoListID:          newId,
			CreationTime:        time.Now(),
			RefactoritionTime:   time.Time{},
			DeletionTime:        time.Time{},
			CompletionOfPercent: 0.5,
		},
		{
			TodoListID:          newId2,
			CreationTime:        time.Now().Add(-24 * time.Hour),
			RefactoritionTime:   time.Time{},
			DeletionTime:        time.Time{},
			CompletionOfPercent: 0.75,
		},
	}
}

func NewMockTodoListRepository() *MockTodoListRepository {
	return &MockTodoListRepository{}
}

func (p *MockTodoListRepository) CreateTodoList(ctx context.Context, todoList *models.TodoList) (*models.TodoList, error) {

	todoLists = append(todoLists, todoList)

	return todoList, nil
}

func (p *MockTodoListRepository) GetTodoListById(ctx context.Context, uuid uuid.UUID) (*models.TodoList, error) {

	var todoList *models.TodoList

	for _, t := range todoLists {

		if t.TodoListID == uuid {
			todoList = t
		}
	}

	return todoList, nil
}

func (p *MockTodoListRepository) GetAllTodoList() ([]*models.TodoList, error) {

	if len(todoLists) <= 0 {
		return []*models.TodoList{}, errors.New("todo lists are empty ")
	}

	return todoLists, nil
}
