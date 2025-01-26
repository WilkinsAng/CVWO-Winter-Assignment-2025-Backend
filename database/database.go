package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

var Conn *pgxpool.Pool

func ConnectToDB() {

	dsn := fmt.Sprintf(os.Getenv("DATABASE_URL"))
	var err error
	Conn, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to the database!")
}
