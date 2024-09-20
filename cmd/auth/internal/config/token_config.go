package config

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"os"
)

type TokenConfig struct {
	SecretKey string
}

func LoadTokenConfig(path string) (*TokenConfig, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}
	config := &TokenConfig{
		SecretKey: os.Getenv("SECRET_KEY"),
	}

	return config, nil
}
