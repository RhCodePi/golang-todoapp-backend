package main

import (
	"context"
	list_api "demoapp/internal/lists/api"
	list_commands "demoapp/internal/lists/features/creating_todo_list/commands"
	create_dtos "demoapp/internal/lists/features/creating_todo_list/dtos"
	delete_commands "demoapp/internal/lists/features/deleting_todo_list/commands"
	"demoapp/internal/lists/features/deleting_todo_list/dtos"
	get_all_dtos "demoapp/internal/lists/features/getting_all_todo_list/dtos"
	get_all_queries "demoapp/internal/lists/features/getting_all_todo_list/queries"
	by_id_dtos "demoapp/internal/lists/features/getting_todo_list_by_id/dtos"
	by_id_queries "demoapp/internal/lists/features/getting_todo_list_by_id/queries"
	repo_list "demoapp/internal/lists/repository"
	message_api "demoapp/internal/messages/api"
	message_commands "demoapp/internal/messages/features/creating_todo_message/commands"
	message_create_dtos "demoapp/internal/messages/features/creating_todo_message/dtos"
	message_get_dtos "demoapp/internal/messages/features/getting_todo_message_by_id/dtos"
	get_by_todo_list_id_dtos "demoapp/internal/messages/features/getting_todo_messages_by_todo_list_id/dtos"

	message_queries "demoapp/internal/messages/features/getting_todo_message_by_id/queries"
	get_by_todo_list_id_queries "demoapp/internal/messages/features/getting_todo_messages_by_todo_list_id/queries"

	"demoapp/internal/messages/repository"
	main_repository "demoapp/internal/repository"
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
	mockRepository := main_repository.NewMockRepository(*repo_list.NewMockTodoListRepository(), *repository.NewMockTodoMessageRepository())

	// Register request handlers to the mediatr

	createTodoListCommandHandler := list_commands.NewCreateTodoListCommandHandler(&mockRepository.MockTodoListRepository)
	getByIdQueryHandler := by_id_queries.NewGetTodoListByIdQueryHandler(&mockRepository.MockTodoListRepository)
	getAllQueryHandler := get_all_queries.NewGetAllTodoListQueryHandler(&mockRepository.MockTodoListRepository)
	deleteCommandHandler := delete_commands.NewDeleteTodoListCommandHandler(&mockRepository.MockTodoListRepository)

	createTodoMessageCommandHandler := message_commands.NewCreateTodoMessageCommandHandler(
		&mockRepository.MockTodoMessageRepository,
		&mockRepository.MockTodoListRepository,
	)
	getMessageByIdQueryHandle := message_queries.NewGetTodoMessageByIdQueryHandler(&mockRepository.MockTodoMessageRepository)

	getMessageByTodoListIdQueryHandle := get_by_todo_list_id_queries.NewGetTodoMessagesByTodoListIdQueryHandler(&mockRepository.MockTodoMessageRepository)

	err := mediatr.RegisterRequestHandler[*list_commands.CreateTodoListCommand, *create_dtos.CreateTodoListCommandResponse](createTodoListCommandHandler)
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

	err = mediatr.RegisterRequestHandler[*message_commands.CreateTodoMessageCommand, *message_create_dtos.CreateTodoMessageCommandResponse](createTodoMessageCommandHandler)
	if err != nil {
		log.Fatal(err)
	}

	err = mediatr.RegisterRequestHandler[*delete_commands.DeleteTodoListCommand, *dtos.DeleteTodoListCommandResponse](deleteCommandHandler)
	if err != nil {
		log.Fatal(err)
	}

	err = mediatr.RegisterRequestHandler[*message_queries.GetTodoMessageByIdQuery, *message_get_dtos.GetTodoMessageByIdQueryResponse](getMessageByIdQueryHandle)
	if err != nil {
		log.Fatal(err)
	}

	err = mediatr.RegisterRequestHandler[*get_by_todo_list_id_queries.GetTodoMessagesByTodoListIDQuery, *get_by_todo_list_id_dtos.GetTodoMessagesByTodoListIDQueryResponse](getMessageByTodoListIdQueryHandle)
	if err != nil {
		log.Fatal(err)
	}

	// Controllers setup
	list_controller := list_api.NewTodoListsController(echo)
	message_controller := message_api.NewTodoMessageController(echo)

	list_api.MapProductsRoutes(echo, list_controller)

	message_api.MapProductsRoutes(echo, message_controller)

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
