package transactions

import (
	"database/sql"
	"fmt"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) List() ([]Transaction, error) {
	rows, err := r.db.Query(`
		SELECT transaction_id, date, created_at, updated_at, amount, note, category, sub_category_id
		FROM transactions
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var transactions []Transaction

	for rows.Next() {
		var t Transaction
		err := rows.Scan(
			&t.TransactionId, // transaction_id
			&t.Date,          // date
			&t.CreatedAt,     // created_at
			&t.UpdatedAt,     // updated_at
			&t.Amount,        // amount
			&t.Note,          // note
			&t.Category,      // category
			&t.SubCategoryId, // sub_category_id
		)
		if err != nil {
			return nil, err
		}
		fmt.Printf("%+v\n", t)
		transactions = append(transactions, t)
	}

	return transactions, nil
}

func (r *TransactionRepository) ListExpenses() ([]Transaction, error) {
	rows, err := r.db.Query(`
		SELECT * FROM transactions WHERE category = $1`, "expense")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var transactions []Transaction

	for rows.Next() {
		var t Transaction
		err := rows.Scan(
			&t.TransactionId, // transaction_id
			&t.Date,          // date
			&t.CreatedAt,     // created_at
			&t.UpdatedAt,     // updated_at
			&t.Amount,        // amount
			&t.Note,          // note
			&t.Category,      // category
			&t.SubCategoryId, // sub_category_id
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}

func (r *TransactionRepository) ListIncomes() ([]Transaction, error) {
	rows, err := r.db.Query(`
		SELECT * FROM transactions WHERE category = $1`, "income")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var transactions []Transaction

	for rows.Next() {
		var t Transaction
		err := rows.Scan(
			&t.TransactionId, // transaction_id
			&t.Date,          // date
			&t.CreatedAt,     // created_at
			&t.UpdatedAt,     // updated_at
			&t.Amount,        // amount
			&t.Note,          // note
			&t.Category,      // category
			&t.SubCategoryId, // sub_category_id
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}

func (r *TransactionRepository) GetById(id int) (*Transaction, error) {
	row := r.db.QueryRow(`
		SELECT transaction_id, date, created_at, updated_at, amount, note, category, sub_category_id
		FROM transactions WHERE transaction_id = $1`, id)

	var transaction Transaction

	err := row.Scan(
		&transaction.TransactionId, // transaction_id
		&transaction.Date,          // date
		&transaction.CreatedAt,     // created_at
		&transaction.UpdatedAt,     // updated_at
		&transaction.Amount,        // amount
		&transaction.Note,          // note
		&transaction.Category,      // category
		&transaction.SubCategoryId, // sub_category_id
	)

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *TransactionRepository) Update(t *Transaction) error {
	_, err := r.db.Exec(`
		UPDATE transactions
		SET amount = $1, date = $2, note = $3, category = $4, updated_at = NOW()
		WHERE transaction_id = $5
	`,
		t.Amount, t.Date, t.Note, t.Category, t.TransactionId,
	)
	return err
}

func (r *TransactionRepository) Delete(t *Transaction) error {
	_, err := r.db.Exec(`
		DELETE FROM transactions
		WHERE transaction_id = $1
	`,
		t.TransactionId,
	)
	return err
}

func (r *TransactionRepository) Create(t *Transaction) error {
	return r.db.QueryRow(
		`INSERT INTO transactions (date, created_at, updated_at, amount, note, category, sub_category_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING transaction_id`,
		t.Date, t.CreatedAt, t.UpdatedAt, t.Amount, t.Note, t.Category, t.SubCategoryId,
	).Scan(&t.TransactionId)
}
