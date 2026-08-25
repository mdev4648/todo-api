package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mdev4648/todo-api/internal/dto"
	"github.com/mdev4648/todo-api/internal/services"
)

func CreateTodo(
	todoService *services.TodoService,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		var req dto.CreateTodoRequest

		if err := c.ShouldBindJSON(&req); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		value, exists := c.Get("user_id")

		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User not authenticated",
			})

			return
		}

		userIDFloat, ok := value.(float64)

		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid user ID",
			})

			return
		}

		userID := uint(userIDFloat)

		todo, err := todoService.CreateTodo(
			req,
			userID,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create todo",
			})

			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Todo created successfully",
			"todo":    todo,
		})
	}
}
