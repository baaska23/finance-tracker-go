package server

import (
	"fmt"
	"ft-service/internal/domains/transactions"
	"ft-service/internal/platform/database"
	"net/http"
	"time"
)

type Server struct {
	DBManager          *database.DBManager
	TransactionHandler *transactions.TransactionHandler
}

func NewServer() *http.Server {
	dbManager, err := database.NewDBManager()
	if err != nil {
		panic(fmt.Errorf("Failed to init database manager: %w", err))
	}

	transactionRepo := transactions.NewTransactionRepository(dbManager.NeonDB)
	transactionService := transactions.NewTransactionService(transactionRepo)
	transactionHandler := transactions.NewTransactionHandler(transactionService)

	srv := Server{
		DBManager:          dbManager,
		TransactionHandler: transactionHandler,
	}

	server := &http.Server{
		Addr:         ":8080", // Make sure to set the port!
		Handler:      srv.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
