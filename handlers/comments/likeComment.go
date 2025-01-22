package comments

import (
	"context"
	"cvwo-winter-assignment/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func LikeComment(c *gin.Context) {
	id := c.Param("id")
	commentID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE comments 
				SET likes = likes + 1
				WHERE id = $1
				RETURNING likes`

	var updatedLikes int

	err = database.Conn.QueryRow(context.Background(), query, commentID).Scan(&updatedLikes)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to like comment.", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully liked comment.", "likes": updatedLikes})
}
