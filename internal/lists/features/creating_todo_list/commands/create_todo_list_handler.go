package commands

import (
	"context"
	"demoapp/internal/lists/features/creating_todo_list/dtos"
	"demoapp/internal/lists/models"
	"demoapp/internal/lists/repository"
	"fmt"
)

type CreateTodoListCommandHandler struct {
	todoListRepository *repository.MockTodoListRepository
}

func NewCreateTodoListCommandHandler(todoListRepository *repository.MockTodoListRepository) *CreateTodoListCommandHandler {
	return &CreateTodoListCommandHandler{todoListRepository: todoListRepository}
}

func (c *CreateTodoListCommandHandler) Handle(ctx context.Context, command *CreateTodoListCommand) (*dtos.CreateTodoListCommandResponse, error) {
	// isLoggerPipelineEnabled := ctx.Value("logger_pipeline").(bool)
	// if isLoggerPipelineEnabled {
	// 	fmt.Println("[CreateTodoListCommandHandler]: logging pipeline is enabled")
	// }

	todoList := &models.TodoList{
		TodoListID:          command.ID,
		CreationTime:        command.CreationTime,
		RefactoritionTime:   command.RefactoritionTime,
		DeletionTime:        command.DeletionTime,
		CompletionOfPercent: command.CompletionOfPercent,
	}

	fmt.Println("todolist", todoList.TodoListID)

	createdTodoList, err := c.todoListRepository.CreateTodoList(ctx, todoList)
	if err != nil {
		return nil, err
	}

	response := &dtos.CreateTodoListCommandResponse{TodoListID: createdTodoList.TodoListID}

	// // Publish notification event to the mediatr for dispatching to the notification handlers

	// productCreatedEvent := events.NewProductCreatedEvent(createdProduct.ProductID, createdProduct.Name, createdProduct.Description, createdProduct.Price, createdProduct.CreatedAt)
	// err = mediatr.Publish[*events.ProductCreatedEvent](ctx, productCreatedEvent)
	// if err != nil {
	// 	return nil, err
	// }

	return response, nil
}
