package threads

import (
	"context"
	"cvwo-winter-assignment/database"
	"cvwo-winter-assignment/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllThreads(c *gin.Context) {

	const pageSize = 10

	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)

	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	offset := (page - 1) * pageSize

	query := `SELECT * FROM threads
				ORDER BY created_at DESC
				LIMIT $1 OFFSET $2`

	rows, err := database.Conn.Query(context.Background(), query, pageSize, offset)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	var threads []models.Thread
	for rows.Next() {
		var thread models.Thread
		if err := rows.Scan(&thread.ID, &thread.Title, &thread.Description, &thread.UserID, &thread.CreatedAt, &thread.UpdatedAt,
			&thread.Likes, &thread.Dislikes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		threads = append(threads, thread)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Threads GET",
		"threads": threads,
		"page":    page})
}
