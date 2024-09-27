package api

import (
	"github.com/gin-gonic/gin"
	"go-microservices/cmd/to_do/internal/api/handler"
	"go-microservices/cmd/to_do/internal/config"
	"go-microservices/cmd/to_do/internal/repository"
	"go-microservices/cmd/to_do/internal/service"
	"gorm.io/gorm"
)

func SetUpRoutes(router *gin.Engine, db *gorm.DB, cfg *config.Config) {

	todoRepo := repository.NewTodoItemRepository(db)
	toDoService := service.NewTodoItemService(todoRepo)
	todoHandler := handler.NewTodoItemHandler(toDoService)

	todoRoutes := router.Group("/api/todo-item")
	{
		todoRoutes.GET("/:id", todoHandler.GetTodoItemById)
		todoRoutes.GET("/", todoHandler.GetListTodoItem)
		todoRoutes.POST("/", todoHandler.CreateTodoItem)
		todoRoutes.PUT("/:id", todoHandler.UpdateTodoItem)
		todoRoutes.DELETE("/:id", todoHandler.DeleteTodoItem)

	}

}
