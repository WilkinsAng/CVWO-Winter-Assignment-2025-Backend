package threads

import (
	"context"
	"cvwo-winter-assignment/database"
	"cvwo-winter-assignment/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateThread(c *gin.Context) {
	var request struct {
		Title      string `json:"title"`
		Content    string `json:"content"`
		CategoryID int    `json:"category_id"`
		Username   string `json:"username"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(request.Title) == 0 || len(request.Content) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title or Content is empty"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var threadID int

	query := `INSERT INTO threads (title, content, user_id, category_id) 
				VALUES ($1, $2, $3, $4) RETURNING id`

	err := database.Conn.QueryRow(context.Background(), query, request.Title, request.Content, userID, request.CategoryID).Scan(&threadID)
	if err != nil {
		fmt.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var thread models.Thread
	thread.ID = threadID
	thread.UserID = userID.(int)
	thread.CategoryID = request.CategoryID
	thread.Title = request.Title
	thread.Content = request.Content
	thread.CreatedAt = time.Now()
	thread.Username = request.Username
	c.JSON(http.StatusCreated, gin.H{"message": "Thread Created", "thread": thread})
}
