package budgets

type BudgetService struct {
	repo *BudgetRepository
}

func NewBudgetService(repo *BudgetRepository) *BudgetService {
	return &BudgetService{repo: repo}
}
