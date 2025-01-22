package threads

import (
	"context"
	"cvwo-winter-assignment/database"
	"cvwo-winter-assignment/handlers/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateThread(c *gin.Context) {
	id := c.Param("id")
	threadID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	/*
		Checking if user is the thread owner
	*/
	err = middleware.ValidateThreadOwnership(threadID, userID.(int))
	if err != nil {
		switch err {
		case middleware.ErrThreadNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		case middleware.ErrUnauthorized:
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
		}
	}

	var request struct {
		Title      string `json:"title"`
		Content    string `json:"content"`
		CategoryID int    `json:"category_id"`
	}

	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/*
		Check if thread is empty
	*/
	if len(request.Title) == 0 || len(request.Content) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title or Content is empty"})
		return
	}

	/*
		Updating thread
	*/

	query :=
		`UPDATE threads 
		 SET title = $1, description = $2, category_id = $3, updated_at = NOW()
		 WHERE id = $4`

	_, err = database.Conn.Exec(context.Background(), query, request.Title, request.Content, request.CategoryID, threadID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update thread", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Updated thread successfully!"})
}
