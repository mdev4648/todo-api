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
	"github.com/mdev4648/todo-api/internal/database"
	"log"


)

func main() {

	cfg := config.Load() // Load the configuration from environment variables or .env file. This function returns a pointer to a Config struct, which contains the application configuration.
	database.Connect(cfg)
    log.Printf("Loaded config: %+v\n", cfg) // Print the loaded configuration to the console for debugging purposes. The %+v format verb is used to print the struct with field names.
	router := gin.Default()

	log.Printf("%s is starting on port %s (%s)",
		cfg.AppName,
		cfg.Port,
		cfg.AppEnv,
	)
	routes.RegisterRoutes(router)
	router.Run(":" + cfg.Port)

	

	// router.Run(":8000")
}