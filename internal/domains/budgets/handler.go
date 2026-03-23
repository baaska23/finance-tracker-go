package budgets

import "github.com/gin-gonic/gin"

type BudgetHandler struct {
	service *BudgetService
}

func NewBudgetHandler(service *BudgetService) *BudgetHandler {
	return &BudgetHandler{service: service}
}

func (handler *BudgetHandler) GetById(c *gin.Context) {
	
}

func (handler *BudgetHandler) Update(c *gin.Context) {
	
}
func (handler *BudgetHandler) Delete(c *gin.Context) {
	
}
func (handler *BudgetHandler) SetMultiYear(c *gin.Context) {
	
}
