// package main

// import "github.com/gin-gonic/gin"

// func main() {

// 	router := gin.Default() //creates a new HTTP server with useful defaults: This creates a router.

// 	router.GET("/health", func(c *gin.Context) {  //This registers a route that responds to HTTP GET requests.
// 		//in django like:  def health(request):
// 		c.JSON(200, gin.H{
// 			"status": "OK",
// 		})

// 	})

// 	router.Run(":8000")
// }

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mdev4648/todo-api/internal/routes"
	"github.com/mdev4648/todo-api/internal/config"
	"log"


)

func main() {

	cfg := config.Load()

	router := gin.Default()

	log.Printf("%s is starting on port %s (%s)",
		cfg.AppName,
		cfg.Port,
		cfg.AppEnv,
	)

	router.Run(":" + cfg.Port)

	routes.RegisterRoutes(router)

	// router.Run(":8000")
}