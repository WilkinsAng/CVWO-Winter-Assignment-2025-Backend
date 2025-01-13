package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

var Conn *pgx.Conn

func ConnectToDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)
	var err error
	Conn, err = pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	/*
		Commented till i figure out how to stop it from closing immediately
	*/
	//defer Conn.Close(context.Background())

	log.Println("Connected to the database!")
}
