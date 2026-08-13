package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mdev4648/todo-api/internal/handlers" //Then every internal import starts with: github.com/mdev4648/todo-api/
	"github.com/mdev4648/todo-api/internal/services"
)

func RegisterRoutes(router *gin.Engine, authService *services.AuthService) {

	router.GET("/health", handlers.Health) //Because we're passing the function, not calling it. so we don't use parentheses. This registers a route that responds to HTTP GET requests.
	router.POST("/api/register", handlers.Register)
	router.POST(
		"/api/login",
		handlers.Login(authService),
	)

}
