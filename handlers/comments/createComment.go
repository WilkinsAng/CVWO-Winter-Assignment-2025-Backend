package comments

import (
	"context"
	"cvwo-winter-assignment/database"
	"cvwo-winter-assignment/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func CreateComment(c *gin.Context) {
	id := c.Param("id")

	threadID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var request struct {
		Content string `json:"content"`
	}

	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(request.Content) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comment is empty"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	query := `INSERT INTO comments (content, user_id, thread_id) VALUES ($1, $2, $3) RETURNING id`

	var commentID int

	err = database.Conn.QueryRow(context.Background(), query, request.Content, userID, threadID).Scan(&commentID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var comment models.Comment
	comment.ID = commentID
	comment.Content = request.Content
	comment.ThreadID = threadID
	comment.UserID = userID.(int)
	comment.CreatedAt = time.Now()
	c.JSON(http.StatusCreated, gin.H{"message": "Comment Created", "comment": comment})
}
