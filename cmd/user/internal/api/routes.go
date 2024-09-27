package api

import (
	"github.com/gin-gonic/gin"
	"go-microservices/cmd/user/internal/api/handler"
	"go-microservices/cmd/user/internal/config"
	"go-microservices/cmd/user/internal/repository"
	"go-microservices/cmd/user/internal/service"
	"go-microservices/internal/midleware"
	repository2 "go-microservices/internal/repository"
	service2 "go-microservices/internal/service"
	"gorm.io/gorm"
)

func SetUpRoutes(router *gin.Engine, db *gorm.DB, config *config.Config) {

	roleRepo := repository.NewRoleRepository(db)
	rolService := service.NewRoleService(roleRepo)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, rolService)
	userHandler := handler.NewUserHandler(userService, rolService)

	whiteIpRepo := repository2.NewWhiteIpRepository(db)
	whiteIpService := service2.NewWhiteService(whiteIpRepo)
	whiteIpMiddleWare := midleware.WhiteIpMiddleware(whiteIpService)

	userRoute := router.Group("/api/user")
	userRoute.Use(whiteIpMiddleWare)
	{
		userRoute.GET("/:username", userHandler.FindUserLoginByUserName)
		userRoute.POST("/", userHandler.CreateUser)
	}

}
