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

	const threadPerPage = 10

	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	categoryStr := c.DefaultQuery("categoryID", "")

	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	offset := (page - 1) * threadPerPage

	query := `SELECT t.id, t.title, t.content, t.user_id, u.username, t.created_at,
			t.updated_at, t.likes, t.dislikes, t.category_id
			FROM threads t
			LEFT JOIN users u ON t.user_id = u.id`

	args := []interface{}{threadPerPage, offset}

	if categoryStr != "" {
		categoryID, err := strconv.Atoi(categoryStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category number."})
			return
		}
		query += " WHERE t.category_ID = $3 ORDER BY t.created_at DESC LIMIT $1 OFFSET $2"
		args = append(args, categoryID)
	} else {
		query += " ORDER BY t.created_at DESC LIMIT $1 OFFSET $2"
	}
	rows, err := database.Conn.Query(context.Background(), query, args...)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	var threads []models.Thread
	for rows.Next() {
		var thread models.Thread
		if err := rows.Scan(&thread.ID, &thread.Title, &thread.Content, &thread.UserID, &thread.Username,
			&thread.CreatedAt, &thread.UpdatedAt, &thread.Likes, &thread.Dislikes, &thread.CategoryID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		threads = append(threads, thread)
	}

	totalThreads, err := GetNumberOfThreads(categoryStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := (totalThreads + threadPerPage - 1) / threadPerPage

	c.JSON(http.StatusOK, gin.H{"message": "Threads GET",
		"threads":      threads,
		"page":         page,
		"totalPages":   totalPages,
		"totalThreads": totalThreads})
}
