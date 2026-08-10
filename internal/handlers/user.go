package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/mdev4648/todo-api/internal/database"
	"github.com/mdev4648/todo-api/internal/dto"
	"github.com/mdev4648/todo-api/internal/services"
	// "github.com/mdev4648/todo-api/internal/models"
)

// func Register(c *gin.Context) {

// 	var req dto.RegisterRequest // This line declares a variable req of type dto.RegisterRequest. This struct is used to bind and validate the incoming JSON payload from the client. It contains fields for Name, Email, and Password, along with validation rules specified in the struct tags.
// 	// err := ...

// 	// if err != nil {

// 	// }
// 	if err := c.ShouldBindJSON(&req); err != nil { // ShouldBindJSON is a method provided by the Gin framework that attempts to bind the incoming JSON payload to the specified struct (in this case, req). It also performs validation based on the struct tags. If the binding or validation fails, it returns an error.

// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})

// 		return
// 	}

// 	user := models.User{
// 		Name:     req.Name,
// 		Email:    req.Email,
// 		Password: req.Password,
// 	}

// 	if err := database.DB.Create(&user).Error; err != nil {

// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})

// 		return
// 	}

//		c.JSON(http.StatusCreated, gin.H{
//			"message": "User created successfully",
//		})
//	}
func Register(c *gin.Context) {

	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	user, err := services.RegisterUser(req)

	if err != nil {

		if err.Error() == "email already exists" {

			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create user",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
