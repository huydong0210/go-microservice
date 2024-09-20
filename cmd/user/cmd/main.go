package main

import (
	"github.com/gin-gonic/gin"
	"go-microservices/cmd/user/internal/api"
	"go-microservices/cmd/user/internal/config"
	"go-microservices/internal/database"
	"go-microservices/pkg/logger"
)

func main() {
	log := logger.NewLogger()

	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal().Err(err).Msg("User service: Load config file failed")
	}

	db, err := database.Initialize(cfg.DatabaseUrl)

	if err != nil {
		log.Fatal().Err(err).Msgf("User service: Failed to initialize database: %s", cfg.DatabaseUrl)
	}

	router := gin.Default()

	err = router.Run(cfg.ServerAddress)
	api.SetUpRoutes(router, db, cfg)

	if err != nil {
		log.Fatal().Err(err).Msg("User service: Failed to start server")
	}
}
