package threads

import (
	"context"
	"cvwo-winter-assignment/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateThread(c *gin.Context) {
	var request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		CategoryID  int    `json:"category_id"`
	}

	/*
		Parsing the Request
	*/
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/*
		Auth check
	*/
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	/*
		Inserting thread
	*/
	var threadID int

	query := `INSERT INTO threads (title, description, user_id, category_id) 
				VALUES ($1, $2, $3, $4) RETURNING id`

	err := database.Conn.QueryRow(context.Background(), query, request.Title, request.Description, userID, request.CategoryID).Scan(&threadID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Thread Created", "id": threadID})
}
