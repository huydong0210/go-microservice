package api

import (
	"github.com/gin-gonic/gin"
	"go-microservices/cmd/auth/internal/api/handler"
	"go-microservices/cmd/auth/internal/api/handler/http"
	"go-microservices/cmd/auth/internal/config"
	"go-microservices/cmd/auth/internal/service"
	"go-microservices/internal/midleware"
	repository2 "go-microservices/internal/repository"
	service2 "go-microservices/internal/service"
	"gorm.io/gorm"
)

func SetUpRoutes(router *gin.Engine, db *gorm.DB, cfg *config.Config) {

	httpHandler := http.NewHttpHandler()
	authService := service.NewAuthService(&cfg.SecretKey, httpHandler)
	authHandler := handler.NewAuthHandler(authService)

	whiteIpRepo := repository2.NewWhiteIpRepository(db)
	whiteIpService := service2.NewWhiteService(whiteIpRepo)
	whiteIpMiddleWare := midleware.WhiteIpMiddleware(whiteIpService)

	authRoutes := router.Group("api/auth")
	authRoutes.Use(whiteIpMiddleWare)

	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}

}
