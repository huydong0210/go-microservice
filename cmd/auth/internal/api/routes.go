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
	authService := service.NewAuthService(&cfg.SecretKey)
	httpHandler := http.NewHttpHandler()
	authHandler := handler.NewAuthHandler(authService, httpHandler)

	whiteIpRepo := repository2.NewWhiteIpRepository(db)
	whiteIpService := service2.NewWhiteService(whiteIpRepo)
	whiteIpMiddleWare := midleware.WhiteIpMiddleware(whiteIpService)

	authRoutes := router.Group("api/auth")
	authRoutes.Use(whiteIpMiddleWare)

	{
		authRoutes.POST("/login", authHandler.Login)
	}

}
