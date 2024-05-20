package commands

import (
	"context"
	"demoapp/internal/lists/repository"
	"demoapp/internal/messages/features/creating_todo_message/dtos"
	message_dtos "demoapp/internal/messages/models"
	message_repo "demoapp/internal/messages/repository"
	"fmt"
)

type CreateTodoMessageCommandHandler struct {
	todoMessageRepository *message_repo.MockTodoMessageRepository
	todoListRepository    *repository.MockTodoListRepository
}

func NewCreateTodoMessageCommandHandler(todoMessageRepository *message_repo.MockTodoMessageRepository, todolistRepository *repository.MockTodoListRepository) *CreateTodoMessageCommandHandler {
	return &CreateTodoMessageCommandHandler{todoMessageRepository: todoMessageRepository, todoListRepository: todolistRepository}
}

func (c *CreateTodoMessageCommandHandler) Handle(ctx context.Context, command *CreateTodoMessageCommand) (*dtos.CreateTodoMessageCommandResponse, error) {
	// isLoggerPipelineEnabled := ctx.Value("logger_pipeline").(bool)
	// if isLoggerPipelineEnabled {
	// 	fmt.Println("[CreateTodoListCommandHandler]: logging pipeline is enabled")
	// }

	todoList, err := c.todoListRepository.GetTodoListById(ctx, command.TodoListID)

	if err != nil {
		return nil, err
	}

	c.todoListRepository.UpdateTodoList(todoList.TodoListID)

	todoMessage := &message_dtos.TodoMessage{
		MessageID:         command.MessageID,
		TodoListID:        command.TodoListID,
		CreationTime:      command.CreationTime,
		RefactoritionTime: command.RefactoritionTime,
		DeletionTime:      command.DeletionTime,
		Content:           command.Content,
		CompletionStatus:  command.CompletionStatus,
	}

	fmt.Println("todoMessage", todoMessage.MessageID)

	createdTodoMessage, err := c.todoMessageRepository.CreateTodoMessage(ctx, todoMessage)

	if err != nil {
		return nil, err
	}

	//todoListsDto := lists.MapTodoListToTodoListDto(todoList, nil)

	response := &dtos.CreateTodoMessageCommandResponse{TodoMessage: createdTodoMessage.MessageID, TodolistId: createdTodoMessage.TodoListID}

	// // Publish notification event to the mediatr for dispatching to the notification handlers

	// productCreatedEvent := events.NewProductCreatedEvent(createdProduct.ProductID, createdProduct.Name, createdProduct.Description, createdProduct.Price, createdProduct.CreatedAt)
	// err = mediatr.Publish[*events.ProductCreatedEvent](ctx, productCreatedEvent)
	// if err != nil {
	// 	return nil, err
	// }

	return response, nil
}
