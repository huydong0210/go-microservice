package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	SecretKey          string
	DatabaseUrl        string
	TodoServiceAddress string
}

func LoadConfig(path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}
	config := &Config{
		DatabaseUrl:        os.Getenv("DATABASE_URL"),
		TodoServiceAddress: os.Getenv("TODO_SERVICE_ADDRESS"),
	}

	return config, nil
}
