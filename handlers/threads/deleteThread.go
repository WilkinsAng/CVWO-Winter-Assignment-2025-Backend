package threads

import (
	"context"
	"cvwo-winter-assignment/database"
	"cvwo-winter-assignment/handlers/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteThread(c *gin.Context) {
	id := c.Param("id")
	threadID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err = middleware.ValidateThreadOwnership(threadID, userID.(int))
	if err != nil {
		switch err {
		case middleware.ErrThreadNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		case middleware.ErrUnauthorized:
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
		}
	}

	query := `DELETE FROM threads WHERE id = $1`
	_, err = database.Conn.Exec(context.Background(), query, threadID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete thread", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted thread"})
}
