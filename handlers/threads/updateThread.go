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
	var request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		CategoryID  int    `json:"category_id"`
	}

	id := c.Param("id")
	threadID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = c.ShouldBindJSON(&request); err != nil {
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
	/*
		Updating thread
	*/

	query :=
		`UPDATE threads 
		 SET title = $1, description = $2, category_id = $3, updated_at = NOW()
		 WHERE id = $4`

	_, err = database.Conn.Exec(context.Background(), query, request.Title, request.Description, request.CategoryID, threadID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete thread", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Updated thread successfully!."})
}
