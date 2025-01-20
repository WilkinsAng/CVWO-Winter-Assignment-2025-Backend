package threads

import (
	"context"
	"cvwo-winter-assignment/database"
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

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	/*
		Check is user is the owner
	*/
	var threadOwnerID int

	userQuery := `SELECT user_id FROM threads WHERE id = $1`

	err = database.Conn.QueryRow(context.Background(), userQuery, threadID).Scan(&threadOwnerID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Thread not found."})
		return
	}
	
	if threadOwnerID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are unauthorized to update this thread."})
		return
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Updated thread successfully!."})
}
