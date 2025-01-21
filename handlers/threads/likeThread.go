package threads

import (
	"context"
	"cvwo-winter-assignment/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func LikeThread(c *gin.Context) {
	id := c.Param("id")
	threadID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE threads 
				SET likes = likes + 1
				WHERE id = $1
				RETURNING likes`

	var updatedLikes int

	err = database.Conn.QueryRow(context.Background(), query, threadID).Scan(&updatedLikes)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to like thread.", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully liked thread.", "likes": updatedLikes})
}
