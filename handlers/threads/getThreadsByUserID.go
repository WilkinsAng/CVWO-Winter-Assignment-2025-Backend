package threads

import (
	"context"
	"cvwo-winter-assignment/database"
	"cvwo-winter-assignment/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetThreadsByUserID(c *gin.Context) {
	id := c.Param("userID")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	threadQuery := `SELECT t.id, t.title, t.content, t.user_id, u.username, t.created_at, 
       				t.updated_at, t.likes, t.dislikes, t.category_id
					FROM threads t
					LEFT JOIN users u ON t.user_id = u.id
					WHERE t.user_id = $1
					ORDER BY t.created_at DESC`

	rows, err := database.Conn.Query(context.Background(), threadQuery, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()
	var threads []models.Thread
	for rows.Next() {
		var thread models.Thread
		err = rows.Scan(&thread.ID, &thread.Title, &thread.Content,
			&thread.UserID, &thread.Username, &thread.CreatedAt, &thread.UpdatedAt, &thread.Likes,
			&thread.Dislikes, &thread.CategoryID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		threads = append(threads, thread)
	}
	if rows.Err() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": rows.Err().Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"threads": threads})
}
