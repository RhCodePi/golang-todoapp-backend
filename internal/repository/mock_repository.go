package repository

import (
	repo "demoapp/internal/lists/repository"
	"demoapp/internal/messages/repository"
)

type MockRepository struct {
	MockTodoListRepository    repo.MockTodoListRepository
	MockTodoMessageRepository repository.MockTodoMessageRepository
}

func NewMockRepository(mockTodoList repo.MockTodoListRepository, mockTodoMessage repository.MockTodoMessageRepository) *MockRepository {
	return &MockRepository{MockTodoListRepository: mockTodoList, MockTodoMessageRepository: mockTodoMessage}
}

func (m *MockRepository) SaveChanges() {
}
