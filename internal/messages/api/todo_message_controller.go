package api

import (
	create_command "demoapp/internal/messages/features/creating_todo_message/commands"
	create_dtos "demoapp/internal/messages/features/creating_todo_message/dtos"
	get_by_id_dtos "demoapp/internal/messages/features/getting_todo_message_by_id/dtos"
	get_by_id_queries "demoapp/internal/messages/features/getting_todo_message_by_id/queries"
	get_by_todo_list_id_dtos "demoapp/internal/messages/features/getting_todo_messages_by_todo_list_id/dtos"
	get_by_todo_list_id_queries "demoapp/internal/messages/features/getting_todo_messages_by_todo_list_id/queries"

	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
)

type TodoMessageController struct {
	echo      *echo.Echo
	validator *validator.Validate
}

func NewTodoMessageController(echo *echo.Echo) *TodoMessageController {
	return &TodoMessageController{echo: echo, validator: validator.New()}
}

func (td *TodoMessageController) createTodoMessage() echo.HandlerFunc {

	return func(ctx echo.Context) error {
		request := &create_dtos.CreateTodoMessageRequestDto{}
		if err := ctx.Bind(request); err != nil {
			return err
		}

		if err := td.validator.StructCtx(ctx.Request().Context(), request); err != nil {
			return err
		}

		command := create_command.NewCreateTodoListCommand(request.TodolistId, request.Content)
		result, err := mediatr.Send[*create_command.CreateTodoMessageCommand, *create_dtos.CreateTodoMessageCommandResponse](ctx.Request().Context(), command)

		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusCreated, result)
	}
}

func (td *TodoMessageController) getTodoMessageByID() echo.HandlerFunc {
	return func(ctx echo.Context) error {

		request := &get_by_id_dtos.GetTodoMessageByIdRequestDto{}
		if err := ctx.Bind(request); err != nil {
			return err
		}

		if err := td.validator.StructCtx(ctx.Request().Context(), request); err != nil {
			return err
		}

		query := get_by_id_queries.NewGetTodoMessageByIdQuery(request.TodoMessageID)

		queryResult, err := mediatr.Send[*get_by_id_queries.GetTodoMessageByIdQuery, *get_by_id_dtos.GetTodoMessageByIdQueryResponse](ctx.Request().Context(), query)

		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, queryResult)
	}
}

func (td *TodoMessageController) getTodoMessagesByTodoListID() echo.HandlerFunc {
	return func(ctx echo.Context) error {

		request := &get_by_todo_list_id_dtos.GetTodoMessagesByTodoListIDRequestDto{}
		if err := ctx.Bind(request); err != nil {
			return err
		}

		if err := td.validator.StructCtx(ctx.Request().Context(), request); err != nil {
			return err
		}

		query := get_by_todo_list_id_queries.NewGetTodoMessagesByTodoListIDQuery(request.TodoListID)

		queryResult, err := mediatr.Send[*get_by_todo_list_id_queries.GetTodoMessagesByTodoListIDQuery, *get_by_todo_list_id_dtos.GetTodoMessagesByTodoListIDQueryResponse](ctx.Request().Context(), query)

		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, queryResult)
	}
}
