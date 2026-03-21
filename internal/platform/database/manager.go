package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type DBManager struct {
	NeonDB *sql.DB
}

func NewDBManager() (*DBManager, error) {
	manager := &DBManager{}

	db, err := initNeon()
	if err != nil {
		return nil, fmt.Errorf("failed to init neon db: %w", &err)
	}

	manager.NeonDB = db

	return manager, nil
}

func initNeon() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("Error loading env file")
	}
	connStr := os.Getenv("DATABASE_URL")
	db, connErr := sql.Open("pgx", connStr)
	if connErr != nil {
		log.Fatal("Unable to connect to database:", connErr)
	}

	log.Println("Successfully connected")
	return db, nil
}
