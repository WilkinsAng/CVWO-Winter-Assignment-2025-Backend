package comments

import (
	"context"
	"cvwo-winter-assignment/database"
	"cvwo-winter-assignment/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCommentByThreadID(c *gin.Context) {
	id := c.Param("id")
	threadID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	commentQuery := `SELECT c.id, c.thread_id, c.user_id, u.username,
       				c.content, c.likes, c.dislikes, c.created_at
         			FROM comments c
         			INNER JOIN users u ON c.user_id = u.id
         			WHERE c.thread_id = $1
         			ORDER BY c.created_at DESC`

	rows, err := database.Conn.Query(context.Background(), commentQuery, threadID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err = rows.Scan(&comment.ID, &comment.ThreadID, &comment.UserID, &comment.Username,
			&comment.Content, &comment.Likes, &comment.Dislikes, &comment.CreatedAt)
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
