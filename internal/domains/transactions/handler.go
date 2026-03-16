package transactions

import (
	"encoding/json"
	"net/http"
)

type TransactionHandler struct {
	service transactionService
}

func NewTransactionHandler(service transactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (handler *TransactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req Transaction

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	transaction, err := handler.service.CreateTransaction(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}