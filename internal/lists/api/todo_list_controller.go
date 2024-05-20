package api

import (
	create_command "demoapp/internal/lists/features/creating_todo_list/commands"
	create_dtos "demoapp/internal/lists/features/creating_todo_list/dtos"
	"demoapp/internal/lists/features/deleting_todo_list/commands"
	"demoapp/internal/lists/features/deleting_todo_list/dtos"
	get_all_dtos "demoapp/internal/lists/features/getting_all_todo_list/dtos"
	get_all_queries "demoapp/internal/lists/features/getting_all_todo_list/queries"
	get_by_id_dtos "demoapp/internal/lists/features/getting_todo_list_by_id/dtos"
	get_by_id_queries "demoapp/internal/lists/features/getting_todo_list_by_id/queries"

	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
)

type TodoListsController struct {
	echo      *echo.Echo
	validator *validator.Validate
}

func NewTodoListsController(echo *echo.Echo) *TodoListsController {
	return &TodoListsController{echo: echo, validator: validator.New()}
}

func (td *TodoListsController) createTodoList() echo.HandlerFunc {

	return func(ctx echo.Context) error {
		request := &create_dtos.CreateTodoListRequestDto{}
		if err := ctx.Bind(request); err != nil {
			return err
		}

		if err := td.validator.StructCtx(ctx.Request().Context(), request); err != nil {
			return err
		}

		command := create_command.NewCreateTodoListCommand()

		result, err := mediatr.Send[*create_command.CreateTodoListCommand, *create_dtos.CreateTodoListCommandResponse](ctx.Request().Context(), command)

		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusCreated, result)
	}
}

func (td *TodoListsController) getTodoListByID() echo.HandlerFunc {
	return func(ctx echo.Context) error {

		request := &get_by_id_dtos.GetTodoListByIdRequestDto{}
		if err := ctx.Bind(request); err != nil {
			return err
		}

		if err := td.validator.StructCtx(ctx.Request().Context(), request); err != nil {
			return err
		}
		query := get_by_id_queries.NewGetTodoListByIdQuery(request.TodoListID)

		queryResult, err := mediatr.Send[*get_by_id_queries.GetTodoListByIdQuery, *get_by_id_dtos.GetTodoListByIdQueryResponse](ctx.Request().Context(), query)

		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, queryResult)
	}
}

func (td *TodoListsController) getAllTodoList() echo.HandlerFunc {
	return func(ctx echo.Context) error {

		request := &get_all_dtos.GetAllTodoListRequestDto{}
		if err := ctx.Bind(request); err != nil {
			return err
		}

		if err := td.validator.StructCtx(ctx.Request().Context(), request); err != nil {
			return err
		}
		query := get_all_queries.NewGetAllTodoListQuery()

		queryResult, err := mediatr.Send[*get_all_queries.GetAllTodoListQuery, *get_all_dtos.GetAllTodoListQueryResponse](ctx.Request().Context(), query)

		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, queryResult)
	}
}

func (td *TodoListsController) deleteTodoList() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := &dtos.DeleteTodoListRequestDto{}
		if err := ctx.Bind(request); err != nil {
			return err
		}

		if err := td.validator.StructCtx(ctx.Request().Context(), request); err != nil {
			return err
		}
		command := commands.NewDeleteTodoListCommand(request.TodoListId)

		result, err := mediatr.Send[*commands.DeleteTodoListCommand, *dtos.DeleteTodoListCommandResponse](ctx.Request().Context(), command)

		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, result)
	}
}
