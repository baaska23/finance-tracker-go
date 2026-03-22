package transactions

import (
	"database/sql"
	"time"
)

type Transaction struct {
	TransactionId int          `json:"transaction_id"`
	Date          time.Time    `json:"date"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     sql.NullTime `json:"updated_at"`
	Amount        float64      `json:"amount"`
	Note          string       `json:"note"`
	Category      string       `json:"category"`
	SubCategoryId int          `json:"sub_category_id"`
}

type SubCategory struct {
	SubCategoryId   int    `json:"sub_category_id"`
	SubCategoryName string `json:"sub_category_name"`
}

type TransactionPatch struct {
	TransactionId *int          `json:"transaction_id"`
	Date          *time.Time    `json:"date"`
	CreatedAt     *time.Time    `json:"created_at"`
	UpdatedAt     *sql.NullTime `json:"updated_at"`
	Amount        *float64      `json:"amount"`
	Note          *string       `json:"note"`
	Category      *string       `json:"category"`
	SubCategoryId *int          `json:"sub_category_id"`
}

type Total struct {
	Expense *float64
	Income  *float64
}
