package threads

import (
	"context"
	"cvwo-winter-assignment/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DislikeThread(c *gin.Context) {
	id := c.Param("id")
	threadID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE threads 
				SET dislikes = dislikes + 1
				WHERE id = $1
				RETURNING dislikes`

	var updatedDislikes int

	err = database.Conn.QueryRow(context.Background(), query, threadID).Scan(&updatedDislikes)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to dislike thread.", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully disliked thread.", "likes": updatedDislikes})
}
