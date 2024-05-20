package api

import (
	"github.com/labstack/echo/v4"
)

func MapProductsRoutes(echo *echo.Echo, controller *TodoMessageController) {
	v1 := echo.Group("/api/v1/crud")
	message := v1.Group("/message")

	message.POST("/post", controller.createTodoMessage())
	message.GET("/get/id", controller.getTodoMessageByID())
	message.GET("/get/todolistid", controller.getTodoMessagesByTodoListID())
	// todoList.POST("/post", controller.createTodoList())
	// todoList.GET("/:id", controller.getTodoListByID())
	// todoList.GET("/getall", controller.getAllTodoList())
}
