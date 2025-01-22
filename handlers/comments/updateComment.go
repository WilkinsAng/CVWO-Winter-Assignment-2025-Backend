package comments

import (
	"context"
	"cvwo-winter-assignment/database"
	"cvwo-winter-assignment/handlers/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateComment(c *gin.Context) {
	id := c.Param("id")
	commentID, err := strconv.Atoi(id)
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
		Checking if user is the comment owner
	*/
	err = middleware.ValidateCommentOwnership(commentID, userID.(int))
	if err != nil {
		switch err {
		case middleware.ErrCommentNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		case middleware.ErrUnauthorizedComment:
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
		}
	}

	var request struct {
		Content string `json:"content"`
	}

	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/*
		Check if comment is empty
	*/
	if len(request.Content) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comment is empty"})
		return
	}

	query := `UPDATE comments 
			  SET content = $1 
			  WHERE id = $2`

	_, err = database.Conn.Exec(context.Background(), query, request.Content, commentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update comment", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully!"})
}
