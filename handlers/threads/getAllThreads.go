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

	/*
		Setting number of threads per page to be 10
	*/
	const threadPerPage = 10

	/*
		Getting query data from request
	*/
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	categoryStr := c.DefaultQuery("category", "")

	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	/*
		Calculating offset for page 2 onwards
	*/
	offset := (page - 1) * threadPerPage

	/*
		Unfortunately can't use * as it gives unnecessary info
	*/
	query := `SELECT t.ID, t.Title, t.Content, t.user_id,
				t.created_at, t.updated_at, t.likes, t.dislikes, t.category_ID
       			FROM threads t`

	args := []interface{}{threadPerPage, offset}

	/*
		If there is a string in the category, select threads in that category
		else, show all threads
	*/
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

	/*
		Putting threads into a slice
	*/
	var threads []models.Thread
	for rows.Next() {
		var thread models.Thread
		if err := rows.Scan(&thread.ID, &thread.Title, &thread.Content, &thread.UserID,
			&thread.CreatedAt, &thread.UpdatedAt, &thread.Likes, &thread.Dislikes, &thread.CategoryID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		threads = append(threads, thread)
	}

	totalThreads, err := GetNumberOfThreads(categoryStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count threads"})
		return
	}

	totalPages := (totalThreads + threadPerPage - 1) / threadPerPage

	c.JSON(http.StatusOK, gin.H{"message": "Threads GET",
		"threads":      threads,
		"page":         page,
		"totalPages":   totalPages,
		"totalThreads": totalThreads})
}
