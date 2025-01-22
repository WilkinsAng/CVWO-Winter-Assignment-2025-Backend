package comments

import (
	"context"
	"cvwo-winter-assignment/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DislikeComment(c *gin.Context) {
	id := c.Param("id")
	commentID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE comments 
				SET dislikes = dislikes + 1
				WHERE id = $1
				RETURNING dislikes`

	var updatedDislikes int

	err = database.Conn.QueryRow(context.Background(), query, commentID).Scan(&updatedDislikes)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to dislike comment.", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully disliked comment.", "likes": updatedDislikes})
}
