package lists

import (
	"demoapp/internal/lists/dtos"
	"demoapp/internal/lists/models"
	messages_models "demoapp/internal/messages/models"
)

func MapTodoListToTodoListDto(todoList *models.TodoList, todoMesages []*messages_models.TodoMessage) *dtos.TodoListDto {
	return &dtos.TodoListDto{
		TodoListID:          todoList.TodoListID,
		CreationTime:        todoList.CreationTime,
		RefactoritionTime:   todoList.RefactoritionTime,
		DeletionTime:        todoList.DeletionTime,
		CompletionOfPercent: todoList.CompletionOfPercent,
		TodoMessages:        todoMesages,
	}
}

func MapTodoListsToTodoListDto(todoLists []*models.TodoList, todoMessages []*messages_models.TodoMessage) []*dtos.TodoListDto {

	var newTodoListDtos []*dtos.TodoListDto

	for _, t := range todoLists {
		newTodoListDtos = append(newTodoListDtos, &dtos.TodoListDto{
			TodoListID:          t.TodoListID,
			CreationTime:        t.CreationTime,
			RefactoritionTime:   t.RefactoritionTime,
			DeletionTime:        t.DeletionTime,
			CompletionOfPercent: t.CompletionOfPercent,
			TodoMessages:        todoMessages,
		})
	}

	return newTodoListDtos
}
