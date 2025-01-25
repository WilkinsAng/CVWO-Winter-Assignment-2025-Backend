package auth

import (
	"context"
	"cvwo-winter-assignment/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Signup(c *gin.Context) {
	var request struct {
		Username string `json:"username" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := database.Conn.Exec(context.Background(), "INSERT INTO users (username) VALUES ($1)", request.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Singup Successful", "username": request.Username})
}
