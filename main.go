package main

import (
	"cvwo-winter-assignment/database"
	"cvwo-winter-assignment/initialize"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func init() {
	initialize.LoadEnvVar()
	database.ConnectToDB()
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to the Forum API!")
	})

	fmt.Printf("Server running on http://localhost:%v", os.Getenv("PORT"))

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
