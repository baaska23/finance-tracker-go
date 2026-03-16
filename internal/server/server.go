package server

import (
	"database/sql"
	"ft-service/internal/domains/transactions"
	"log"
	"net/http"
	"os"
)

func StartServer() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	repo := transactions.NewTransactionRepository(db)
	service := transactions.NewTransactionService(*repo)
	handler := transactions.NewTransactionHandler(*service)

	http.HandleFunc("/transactions", handler.Create)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}