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

	threadQuery := `SELECT t.id, t.title, t.content, t.user_id, u.username, t.created_at,
					t.updated_at, t.likes, t.dislikes, t.category_id
					FROM threads t
					LEFT JOIN users u ON t.user_id = u.id
                    WHERE t.id = $1`

	err = database.Conn.QueryRow(context.Background(), threadQuery, threadID).Scan(&thread.ID, &thread.Title, &thread.Content,
		&thread.UserID, &thread.Username, &thread.CreatedAt, &thread.UpdatedAt, &thread.Likes, &thread.Dislikes, &thread.CategoryID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"thread": thread})
}
