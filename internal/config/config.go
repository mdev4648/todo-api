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
	DBHost    string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func Load() *Config { // this function is pointer it returns a pointer to a Config struct. This is useful because it allows us to modify the struct in place, rather than returning a copy of it.

	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found, using system environment variables")
	}

	return &Config{ // return a address of a Config struct. The & operator is used to get the address of the struct, which is then returned as a pointer.
		Port:    os.Getenv("PORT"),
		AppName: os.Getenv("APP_NAME"),
		AppEnv:  os.Getenv("APP_ENV"),

		DBHost:    os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
	}
}