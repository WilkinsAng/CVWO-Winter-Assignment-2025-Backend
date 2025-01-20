package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetHeader("userID")
		if userID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		userIntID, err := strconv.Atoi(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid userID format"})
			return
		}
		c.Set("userID", userIntID)
		c.Next()
	}
}
