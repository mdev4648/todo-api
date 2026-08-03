package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mdev4648/todo-api/internal/handlers"
)

func RegisterRoutes(router *gin.Engine) {

	router.GET("/health", handlers.Health) //Because we're passing the function, not calling it. so we don't use parentheses. This registers a route that responds to HTTP GET requests. 

}