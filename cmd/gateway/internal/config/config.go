package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	TodoServiceAddress string
	AuthServiceAddress string
	GatewayAddress     string
	UserServiceAddress string
	Host               string
}

func LoadConfig(path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}
	config := &Config{
		TodoServiceAddress: os.Getenv("TODO_SERVICE_ADDRESS"),
		AuthServiceAddress: os.Getenv("AUTH_SERVICE_ADDRESS"),
		GatewayAddress:     os.Getenv("GATEWAY_ADDRESS"),
		UserServiceAddress: os.Getenv("USER_SERVICE_ADDRESS"),
		Host:               os.Getenv("HOST"),
	}

	return config, nil
}
