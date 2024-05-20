package commands

import (
	"context"
	"demoapp/internal/lists/features/deleting_todo_list/dtos"

	//"demoapp/internal/lists/models"
	"demoapp/internal/lists/repository"
	"fmt"
)

type DeleteTodoListCommandHandler struct {
	todoListRepository *repository.MockTodoListRepository
}

func NewDeleteTodoListCommandHandler(todoListRepository *repository.MockTodoListRepository) *DeleteTodoListCommandHandler {
	return &DeleteTodoListCommandHandler{todoListRepository: todoListRepository}
}

func (c *DeleteTodoListCommandHandler) Handle(ctx context.Context, command *DeleteTodoListCommand) (*dtos.DeleteTodoListCommandResponse, error) {
	// isLoggerPipelineEnabled := ctx.Value("logger_pipeline").(bool)
	// if isLoggerPipelineEnabled {
	// 	fmt.Println("[CreateTodoListCommandHandler]: logging pipeline is enabled")
	// }

	fmt.Println("todod", command.TodoListId)

	err := c.todoListRepository.DeleteTodoList(command.TodoListId, command.DeletionTime)
	if err != nil {
		return nil, err
	}

	response := &dtos.DeleteTodoListCommandResponse{DeletionTime: command.DeletionTime}

	// // Publish notification event to the mediatr for dispatching to the notification handlers

	// productCreatedEvent := events.NewProductCreatedEvent(createdProduct.ProductID, createdProduct.Name, createdProduct.Description, createdProduct.Price, createdProduct.CreatedAt)
	// err = mediatr.Publish[*events.ProductCreatedEvent](ctx, productCreatedEvent)
	// if err != nil {
	// 	return nil, err
	// }

	return response, nil
}
