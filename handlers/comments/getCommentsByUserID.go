package comments

import (
	"context"
	"cvwo-winter-assignment/database"
	"cvwo-winter-assignment/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCommentsByUserID(c *gin.Context) {
	id := c.Param("userID")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	commentQuery := `SELECT * FROM comments 
         			WHERE user_id = $1
         			ORDER BY created_at DESC`
	rows, err := database.Conn.Query(context.Background(), commentQuery, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err = rows.Scan(&comment.ID, &comment.ThreadID, &comment.UserID, &comment.Content,
			&comment.Likes, &comment.Dislikes, &comment.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		comments = append(comments, comment)
	}
	if rows.Err() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": rows.Err().Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comments": comments})
}
