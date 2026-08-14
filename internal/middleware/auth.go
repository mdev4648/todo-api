package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/mdev4648/todo-api/internal/config"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
			})

			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization format",
			})

			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) { //This is a callback function.  jwt.Parse calls this function when it needs to know:  "What secret/key should I use to verify this token?"

				if token.Method != jwt.SigningMethodHS256 { //"I only accept tokens signed using HS256."
					return nil, jwt.ErrSignatureInvalid
				}

				return []byte(cfg.JWTSecret), nil
			},
		)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})

			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims) //It's essentially a map containing the data inside your JWT:

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token claims",
			})

			c.Abort()
			return
		}

		userID, ok := claims["user_id"]

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User ID not found in token",
			})

			c.Abort()
			return
		}

		c.Set("user_id", userID)

		c.Next() //Continue to the next handler/middleware.
	}
}
