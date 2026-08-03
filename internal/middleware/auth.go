package middleware

import (
	"net/http"
	"strings"

	"go-tweets/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifies the Bearer access token and puts the caller's
// identity on the gin context as "userID" / "username".
func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header is required",
			})
			return
		}

		parts := strings.Fields(header)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header must be in the format: Bearer <token>",
			})
			return
		}

		userID, username, err := jwt.ValidateToken(parts[1], secretKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Set("userID", userID)
		c.Set("username", username)

		c.Next()
	}
}
