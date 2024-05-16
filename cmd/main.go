package main

import (
	"context"
	"demoapp/internal/lists/api"
	"demoapp/internal/lists/features/creating_todo_list/commands"
	create_dtos "demoapp/internal/lists/features/creating_todo_list/dtos"
	get_all_dtos "demoapp/internal/lists/features/getting_all_todo_list/dtos"
	get_all_queries "demoapp/internal/lists/features/getting_all_todo_list/queries"
	by_id_dtos "demoapp/internal/lists/features/getting_todo_list_by_id/dtos"
	by_id_queries "demoapp/internal/lists/features/getting_todo_list_by_id/queries"
	"demoapp/internal/lists/repository"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	echo := echo.New()
	mockRepository := repository.NewMockTodoListRepository()

	// Register request handlers to the mediatr

	createTodoListCommandHandler := commands.NewCreateTodoListCommandHandler(mockRepository)
	getByIdQueryHandler := by_id_queries.NewGetTodoListByIdQueryHandler(mockRepository)
	getAllQueryHandler := get_all_queries.NewGetAllTodoListQueryHandler(mockRepository)

	err := mediatr.RegisterRequestHandler[*commands.CreateTodoListCommand, *create_dtos.CreateTodoListCommandResponse](createTodoListCommandHandler)
	if err != nil {
		log.Fatal(err)
	}

	err = mediatr.RegisterRequestHandler[*by_id_queries.GetTodoListByIdQuery, *by_id_dtos.GetTodoListByIdQueryResponse](getByIdQueryHandler)
	if err != nil {
		log.Fatal(err)
	}

	err = mediatr.RegisterRequestHandler[*get_all_queries.GetAllTodoListQuery, *get_all_dtos.GetAllTodoListQueryResponse](getAllQueryHandler)
	if err != nil {
		log.Fatal(err)
	}

	// Controllers setup
	controller := api.NewTodoListsController(echo)

	api.MapProductsRoutes(echo, controller)

	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "Catalogs Write-Service Api"
	docs.SwaggerInfo.Description = "Catalogs Write-Service Api."

	echo.GET("/swagger/*", echoSwagger.WrapHandler)

	go func() {
		if err := echo.Start(":9080"); err != nil {
			log.Fatal("Error starting Server", err)
		}
	}()

	<-ctx.Done()

	if err := echo.Shutdown(ctx); err != nil {
		log.Fatal("(Shutdown) err", err)
	}

}
