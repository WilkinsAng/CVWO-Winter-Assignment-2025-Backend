package handlers

import (
	"context"
	"cvwo-winter-assignment/database"
	"cvwo-winter-assignment/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Signup(c *gin.Context) {
	var user models.User

	/*
		if err := c.BindJSON(&user); err != nil
				return
			}
	*/
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}
	_, err := database.Conn.Exec(context.Background(), "INSERT INTO users (username) VALUES ($1)", user.Username)
	if err != nil {
		//if err.Error() == `ERROR: duplicate key value violates unique constraint "users_username_key" (SQLSTATE 23505)` {
		//	c.JSON(http.StatusBadRequest, gin.H{"error": "Username already taken"})
		//	return
		//}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User Created", "username": user.Username})
}

//func Login(c *gin.Context) {}
