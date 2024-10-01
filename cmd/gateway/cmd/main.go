package main

import (
	"github.com/gin-gonic/gin"
	"go-microservices/cmd/gateway/internal/api"
	"go-microservices/cmd/gateway/internal/config"
	"go-microservices/pkg/logger"
)

func main() {

	log := logger.NewLogger()

	cfg, err := config.LoadConfig(".env")

	if err != nil {
		log.Fatal().Err(err).Msg("Todo service: Load config file failed")
	}

	router := gin.Default()

	api.SetUpRoutes(router, cfg)

	err = router.Run(cfg.GatewayAddress)

	if err != nil {
		log.Fatal().Err(err).Msg("Todo service: Failed to start server")
	}
}
