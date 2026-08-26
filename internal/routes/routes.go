package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mdev4648/todo-api/internal/config"
	"github.com/mdev4648/todo-api/internal/handlers" //Then every internal import starts with: github.com/mdev4648/todo-api/
	"github.com/mdev4648/todo-api/internal/middleware"
	"github.com/mdev4648/todo-api/internal/services"
)

func RegisterRoutes(router *gin.Engine, authService *services.AuthService, userService *services.UserService, todoService *services.TodoService, cfg *config.Config) {

	router.GET("/health", handlers.Health) //Because we're passing the function, not calling it. so we don't use parentheses. This registers a route that responds to HTTP GET requests.
	// router.POST("/api/register", handlers.Register)
	router.POST(
		"/api/register",
		handlers.Register(userService),
	)
	router.POST(
		"/api/login",
		handlers.Login(authService),
	)
	protected := router.Group("/api")
	protected.Use(
		middleware.AuthMiddleware(cfg),
	)
	protected.GET(
		"/profile",
		handlers.Profile,
	)

	protected.POST(
		"/todos",
		handlers.CreateTodo(todoService),
	)

	protected.GET(
		"/todos",
		handlers.GetTodos(todoService),
	)

}
