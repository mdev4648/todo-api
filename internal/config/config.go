package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	AppName string
	AppEnv  string
}

func Load() *Config {

	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found, using system environment variables")
	}

	return &Config{
		Port:    os.Getenv("PORT"),
		AppName: os.Getenv("APP_NAME"),
		AppEnv:  os.Getenv("APP_ENV"),
	}
}