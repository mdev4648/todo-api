package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mdev4648/todo-api/internal/config"
	"github.com/mdev4648/todo-api/internal/database"
	"github.com/mdev4648/todo-api/internal/repositories"
	"github.com/mdev4648/todo-api/internal/routes"
	"github.com/mdev4648/todo-api/internal/services"
)

func main() {

	cfg := config.Load() // Load the configuration from environment variables or .env file. This function returns a pointer to a Config struct, which contains the application configuration.
	database.Connect(cfg)
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(
		userRepository,
	)

	todoRepository := repositories.NewTodoRepository()

	todoService := services.NewTodoService(
		todoRepository,
	)
	router := gin.Default()
	authService := services.NewAuthService(cfg, userRepository)

	log.Printf("%s is starting onnn port %s (%s)",
		cfg.AppName,
		cfg.Port,
		cfg.AppEnv,
	)

	routes.RegisterRoutes(
		router,
		authService,
		userService,
		todoService,
		cfg,
	)
	router.Run(":" + cfg.Port)

	// router.Run(":8000")
}
