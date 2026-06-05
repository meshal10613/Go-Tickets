package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	Dsn          string
	JwtSecretKey string
}

func LoadEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Port:         os.Getenv("PORT"),
		Dsn:          os.Getenv("DSN"),
		JwtSecretKey: os.Getenv("JWT_SECRET_KEY"),
	}, nil
}
