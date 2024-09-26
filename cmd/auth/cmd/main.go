package main

import (
	"github.com/gin-gonic/gin"
	"go-microservices/cmd/auth/internal/api"
	"go-microservices/cmd/auth/internal/config"
	"go-microservices/internal/database"
	"go-microservices/pkg/logger"
)

func main() {
	log := logger.NewLogger()

	cfg, err := config.LoadConfig(".env")

	if err != nil {
		log.Fatal().Err(err).Msg("User service: Load config file failed")
	}
	db, err := database.Initialize(cfg.DatabaseUrl)

	router := gin.Default()
	api.SetUpRoutes(router, db, cfg)

	err = router.Run(cfg.AuthServiceAddress)

	if err != nil {
		log.Fatal().Err(err).Msg("User service: Failed to start server")
	}
}
