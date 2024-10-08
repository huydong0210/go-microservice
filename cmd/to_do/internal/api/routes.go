package api

import (
	"github.com/gin-gonic/gin"
	"go-microservices/cmd/to_do/internal/api/handler"
	"go-microservices/cmd/to_do/internal/config"
	"go-microservices/cmd/to_do/internal/repository"
	"go-microservices/cmd/to_do/internal/service"
	"go-microservices/internal/midleware"
	"gorm.io/gorm"
)

func SetUpRoutes(router *gin.Engine, db *gorm.DB, cfg *config.Config) {

	todoRepo := repository.NewTodoItemRepository(db)
	toDoService := service.NewTodoItemService(todoRepo)
	todoHandler := handler.NewTodoItemHandler(toDoService)

	jwtMiddleWare := midleware.JwtMiddleWare()

	todoRoutes := router.Group("/api/todo-item")
	todoRoutes.Use(jwtMiddleWare)

	{
		todoRoutes.GET("/:id", todoHandler.GetTodoItemById)
		todoRoutes.GET("/", todoHandler.GetListTodoItem)
		todoRoutes.POST("/", todoHandler.CreateTodoItem)
		todoRoutes.PUT("/:id", todoHandler.UpdateTodoItem)
		todoRoutes.DELETE("/:id", todoHandler.DeleteTodoItem)

	}

}
