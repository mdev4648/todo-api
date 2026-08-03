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
)

func main() {

	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run(":8000")
}