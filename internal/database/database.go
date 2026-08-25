package database

import (
	"fmt"
	"log"

	"github.com/mdev4648/todo-api/internal/config"

	"github.com/mdev4648/todo-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *config.Config) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	DB = db

	err = DB.AutoMigrate(
		&models.User{}, //creates a pointer to an empty user. GORM expects pointers so it can inspect and work with the model efficiently.
		&models.Todo{},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected successfully")
}
