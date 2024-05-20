package repository

import (
	"context"
	"demoapp/internal/messages/models"
	"errors"

	uuid "github.com/satori/go.uuid"
)

type MockTodoMessageRepository struct {
}

var todoMessages []*models.TodoMessage

func init() {

	// var newId = uuid.NewV1()
	// var newId2 = uuid.NewV1()

	// todoLists = []*models.TodoMessage{
	// 	{
	// 		TodoListID:          newId,
	// 		CreationTime:        time.Now(),
	// 		RefactoritionTime:   time.Time{},
	// 		DeletionTime:        time.Time{},
	// 		CompletionOfPercent: 0.5,
	// 	},
	// 	{
	// 		TodoListID:          newId2,
	// 		CreationTime:        time.Now().Add(-24 * time.Hour),
	// 		RefactoritionTime:   time.Time{},
	// 		DeletionTime:        time.Time{},
	// 		CompletionOfPercent: 0.75,
	// 	},
	// }
}

func NewMockTodoMessageRepository() *MockTodoMessageRepository {
	return &MockTodoMessageRepository{}
}

func (p *MockTodoMessageRepository) CreateTodoMessage(ctx context.Context, todoMessage *models.TodoMessage) (*models.TodoMessage, error) {

	todoMessages = append(todoMessages, todoMessage)

	return todoMessage, nil
}

func (p *MockTodoMessageRepository) GetTodoMessageById(ctx context.Context, uuid uuid.UUID) (*models.TodoMessage, error) {

	var todoMessage *models.TodoMessage

	for _, t := range todoMessages {

		if t.MessageID == uuid {
			todoMessage = t
			return todoMessage, nil
		}
	}

	return nil, errors.New("message not found")
}

func (p *MockTodoMessageRepository) GetMessagesByTodoListId(ctx context.Context, todoListID uuid.UUID) ([]*models.TodoMessage, error) {
	var result []*models.TodoMessage

	for _, t := range todoMessages {

		if t.TodoListID == todoListID {
			result = append(result, t)
		}
	}

	// if len(result) <= 0 {
	// 	return nil, errors.New("no found data with")
	// }

	return result, nil
}
