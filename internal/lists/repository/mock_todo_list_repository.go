package repository

import (
	"context"
	lists_models "demoapp/internal/lists/models"
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type MockTodoListRepository struct {
}

var todoLists []*lists_models.TodoList

func init() {

	var newId = uuid.NewV1()
	var newId2 = uuid.NewV1()

	todoLists = []*lists_models.TodoList{
		{
			TodoListID:          newId,
			CreationTime:        time.Now(),
			RefactoritionTime:   time.Time{},
			DeletionTime:        time.Time{},
			CompletionOfPercent: 0,
		},
		{
			TodoListID:          newId2,
			CreationTime:        time.Now(),
			RefactoritionTime:   time.Time{},
			DeletionTime:        time.Time{},
			CompletionOfPercent: 0,
		},
	}
}

func NewMockTodoListRepository() *MockTodoListRepository {
	return &MockTodoListRepository{}
}

func (p *MockTodoListRepository) CreateTodoList(ctx context.Context, todoList *lists_models.TodoList) (*lists_models.TodoList, error) {

	todoLists = append(todoLists, todoList)

	return todoList, nil
}

func (p *MockTodoListRepository) GetTodoListById(ctx context.Context, uuid uuid.UUID) (*lists_models.TodoList, error) {

	var todoList *lists_models.TodoList

	for _, t := range todoLists {

		if t.TodoListID == uuid {
			todoList = t
		}
	}

	return todoList, nil
}

func (p *MockTodoListRepository) GetAllTodoList() ([]*lists_models.TodoList, error) {

	if len(todoLists) <= 0 {
		return nil, errors.New("todo lists are empty ")
	}

	return todoLists, nil
}

func (p *MockTodoListRepository) UpdateTodoList(todoListId uuid.UUID) {

	UpdateTodoListWithID(todoListId)
}

func UpdateTodoListWithID(todoListId uuid.UUID) {
	for _, t := range todoLists {

		if t.TodoListID == todoListId {
			t.RefactoritionTime = time.Now()

		}
	}
}

func (p *MockTodoListRepository) DeleteTodoList(todoListId uuid.UUID, deletionTime time.Time) error {

	for _, t := range todoLists {

		if t.TodoListID == todoListId {
			t.DeletionTime = time.Now()
			return nil
		}
	}
	return errors.New("todo List not found")
}
