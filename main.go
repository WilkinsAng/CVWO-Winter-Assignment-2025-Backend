package main

import (
	"cvwo-winter-assignment/database"
	"cvwo-winter-assignment/handlers/auth"
	"cvwo-winter-assignment/handlers/auth/middleware"
	"cvwo-winter-assignment/handlers/threads"
	"cvwo-winter-assignment/initialize"
	"fmt"
	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, //Change * to backend when deploying
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Authentication Routes
	router.POST("/signup", auth.Signup)
	router.POST("/login", auth.Login)

	// Thread CRUD Routes
	router.POST("/threads", middleware.AuthMiddleware(), threads.CreateThread)
	router.GET("/threads", threads.GetAllThreads)
	router.GET("/threads/:id", threads.GetThreadByID)
	//router.PUT("/threads/:id", threads.UpdateThread)
	//router.DELETE("/threads/:id", threads.DeleteThread)

	fmt.Printf("Server running on http://localhost:%v", os.Getenv("PORT"))

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
