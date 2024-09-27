package main

import (
	"github.com/gin-gonic/gin"
	"go-microservices/cmd/to_do/internal/api"
	"go-microservices/cmd/to_do/internal/config"
	"go-microservices/internal/database"
	"go-microservices/pkg/logger"
)

func main() {
	log := logger.NewLogger()

	cfg, err := config.LoadConfig(".env")

	if err != nil {
		log.Fatal().Err(err).Msg("Todo service: Load config file failed")
	}
	db, err := database.Initialize(cfg.DatabaseUrl)

	router := gin.Default()
	api.SetUpRoutes(router, db, cfg)

	err = router.Run(cfg.TodoServiceAddress)

	if err != nil {
		log.Fatal().Err(err).Msg("Todo service: Failed to start server")
	}
}
