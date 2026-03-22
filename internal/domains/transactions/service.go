package transactions

import (
	"database/sql"
	"errors"
	"time"
)

type TransactionService struct {
	repo *TransactionRepository
}

func NewTransactionService(repo *TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (t *TransactionService) CreateTransaction(req Transaction) (*Transaction, error) {
	if req.Amount <= 0 {
		return nil, errors.New("amount should be greater than 0")
	}

	newTransaction := &Transaction{
		// TransactionId: int.serial
		Amount:        req.Amount,
		Date:          time.Now(),
		CreatedAt:     time.Now(),
		UpdatedAt:     sql.NullTime{},
		Note:          req.Note,
		Category:      req.Category,
		SubCategoryId: req.SubCategoryId,
	}

	err := t.repo.Create(newTransaction)
	if err != nil {
		return nil, err
	}
	return newTransaction, nil
}

func (t *TransactionService) UpdateTransaction(id int, req TransactionPatch) (*Transaction, error) {
	selectedTransaction, err := t.GetById(*req.TransactionId)
	if err != nil {
		return nil, err
	}

	if req.Amount != nil {
		selectedTransaction.Amount = *req.Amount
	}
	if req.Date != nil {
		selectedTransaction.Date = *req.Date
	}
	if req.Note != nil {
		selectedTransaction.Note = *req.Note
	}
	if req.Category != nil {
		selectedTransaction.Category = *req.Category
	}

	err = t.repo.Update(selectedTransaction)
	if err != nil {
		return nil, err
	}

	return selectedTransaction, nil
}

func (t *TransactionService) DeleteTransaction(id int) (*Transaction, error) {
	selectedTransaction, err := t.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	err = t.repo.Delete(selectedTransaction)
	if err != nil {
		return nil, err
	}

	return selectedTransaction, nil
}

func (t *TransactionService) ListTransaction() ([]Transaction, error) {
	transactions, err := t.repo.List()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (t *TransactionService) ListExpenses() ([]Transaction, error) {
	transactions, err := t.repo.ListExpenses()
	if err != nil {
		return nil, err
	}
	return transactions, err
}

func (t *TransactionService) ListIncomes() ([]Transaction, error) {
	transactions, err := t.repo.ListIncomes()
	if err != nil {
		return nil, err
	}
	return transactions, err
}

func (t *TransactionService) GetById(id int) (*Transaction, error) {
	transaction, err := t.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t *TransactionService) GetTotalByMonth(month string) (Total, error) {
	totals, err := t.repo.GetTotalByMonth(month)
	if err != nil {
		return Total{}, err
	}
	return totals, err
}

func (t *TransactionService) GetSummaryByMonth(month string) ([]Transaction, error) {
	transactions, err := t.repo.GetSummaryByMonth(month)
	if err != nil {
		return nil, err
	}
	return transactions, err
}
