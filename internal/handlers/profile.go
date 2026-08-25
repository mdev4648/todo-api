package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mdev4648/todo-api/internal/database"
	"github.com/mdev4648/todo-api/internal/models"
)

func Profile(c *gin.Context) {

	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})

		return
	}

	var user models.User

	result := database.DB.First(
		&user,
		userID,
	)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}
