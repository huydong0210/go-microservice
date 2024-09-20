package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DatabaseUrl   string
	ServerAddress string
	LogLevel      string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	config := &Config{
		DatabaseUrl:   os.Getenv("DATABASE_URL"),
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		LogLevel:      os.Getenv("LOG_LEVEL"),
	}

	return config, nil
}
