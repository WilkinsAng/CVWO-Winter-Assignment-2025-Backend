package threads

import (
	"context"
	"cvwo-winter-assignment/database"
	"cvwo-winter-assignment/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetThreadByThreadID(c *gin.Context) {
	id := c.Param("id")
	threadID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var thread models.Thread
	threadQuery := `SELECT t.ID, t.Title, t.Description, t.user_id,
					t.created_at, t.updated_at, t.likes, t.dislikes,
					c.name AS category 
					FROM threads t
					LEFT JOIN categories c ON t.category_id = c.id	
                    WHERE t.id = $1`

	err = database.Conn.QueryRow(context.Background(), threadQuery, threadID).Scan(&thread.ID, &thread.Title, &thread.Description,
		&thread.UserID, &thread.CreatedAt, &thread.UpdatedAt, &thread.Likes, &thread.Dislikes, &thread.Category)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//get comments
	commentQuery := `SELECT * FROM comments 
         			WHERE thread_id = $1
         			ORDER BY created_at DESC`
	rows, err := database.Conn.Query(context.Background(), commentQuery, thread.ID)
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

	c.JSON(http.StatusOK, gin.H{"threads": thread,
		"comments": comments})
}
