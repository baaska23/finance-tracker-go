package transactions

import "database/sql"

type transactionRepository struct {
    db *sql.DB
}


func NewTransactionRepository(db *sql.DB) *transactionRepository {
    return &transactionRepository{db: db}
}

func(r *transactionRepository) List() ([]Transaction, error) {
    rows, err := r.db.Query("SELECT * from transactions")
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    var transactions []Transaction

    for rows.Next() {
        var t Transaction
        transactions = append(transactions, t)
    }

    return transactions, nil
}