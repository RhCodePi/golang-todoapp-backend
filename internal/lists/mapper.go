package lists

import (
	"demoapp/internal/lists/dtos"
	"demoapp/internal/lists/models"
)

func MapTodoListToTodoListDto(todoList *models.TodoList) *dtos.TodoListDto {
	return &dtos.TodoListDto{
		TodoListID:          todoList.TodoListID,
		CreationTime:        todoList.CreationTime,
		RefactoritionTime:   todoList.RefactoritionTime,
		DeletionTime:        todoList.DeletionTime,
		CompletionOfPercent: todoList.CompletionOfPercent,
	}
}

func MapTodoListsToTodoListDto(todoLists []*models.TodoList) []*dtos.TodoListDto {

	var newTodoListDtos []*dtos.TodoListDto

	for _, t := range todoLists {
		newTodoListDtos = append(newTodoListDtos, &dtos.TodoListDto{
			TodoListID:          t.TodoListID,
			CreationTime:        t.CreationTime,
			RefactoritionTime:   t.RefactoritionTime,
			DeletionTime:        t.DeletionTime,
			CompletionOfPercent: t.CompletionOfPercent,
		})
	}

	return newTodoListDtos
}
