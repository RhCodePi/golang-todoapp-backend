package api

import (
	"github.com/labstack/echo/v4"
)

func MapProductsRoutes(echo *echo.Echo, controller *TodoListsController) {
	v1 := echo.Group("/api/v1")
	todoList := v1.Group("/todolists")

	todoList.POST("/post", controller.createTodoList())
	todoList.GET("/:id", controller.getTodoListByID())
	todoList.GET("/getall", controller.getAllTodoList())
}
