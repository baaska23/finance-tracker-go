package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading env file")
    }
    connStr := os.Getenv("DATABASE_URL")
    conn, connErr := pgx.Connect(context.Background(), connStr)
    if connErr != nil {
        log.Fatal("Unable to connect to database:", connErr)
    }
    defer conn.Close(context.Background())
    log.Println("Successfully connected")
}