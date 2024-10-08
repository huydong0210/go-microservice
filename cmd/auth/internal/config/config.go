package config

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"os"
)

type Config struct {
	SecretKey          string
	DatabaseUrl        string
	AuthServiceAddress string
}

func LoadConfig(path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}
	config := &Config{
		SecretKey:          os.Getenv("SECRET_KEY"),
		DatabaseUrl:        os.Getenv("DATABASE_URL"),
		AuthServiceAddress: os.Getenv("AUTH_SERVICE_ADDRESS"),
	}

	return config, nil
}
