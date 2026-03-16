package transactions

import "database/sql"

type transactionRepository struct {
    db *sql.DB
}


func NewTransactionRepository(db *sql.DB) *transactionRepository {
    return &transactionRepository{db: db}
}