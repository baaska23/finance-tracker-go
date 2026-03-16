package transactions

import (
	"errors"
	"time"
)

type transactionService struct {
	repo transactionRepository
}

func NewTransactionService(repo transactionRepository) *transactionService {
	return &transactionService{repo: repo}
}

func (t *transactionService) CreateTransaction(req Transaction) (*Transaction, error) {
	if req.Amount <= 0 {
		return nil, errors.New("amount should be greater than 0")
	}

	neu := &Transaction{
		TransactionId: req.TransactionId,
		Amount: req.Amount,
		Date: time.Now(),
		CreatedAt: time.Now(),
		Note: req.Note,
		Category: req.Category,
	}

	return neu, nil
}