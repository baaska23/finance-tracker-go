package server

import (
	"ft-service/internal/platform/middleware"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	transactionRoutes := r.Group("/transactions")
	{
		transactionRoutes.GET("/all", middleware.BasicAuth(), s.TransactionHandler.ListAll)
		transactionRoutes.GET("/:id", middleware.BasicAuth(), s.TransactionHandler.GetById)
		transactionRoutes.GET("/expense", middleware.BasicAuth(), s.TransactionHandler.ListExpenses)
		transactionRoutes.GET("/income", middleware.BasicAuth(), s.TransactionHandler.ListIncomes)

		transactionRoutes.PATCH("/:id", middleware.BasicAuth(), s.TransactionHandler.Update)
		transactionRoutes.POST("/create", middleware.BasicAuth(), s.TransactionHandler.Create)
		transactionRoutes.DELETE("/:id", middleware.BasicAuth(), s.TransactionHandler.Delete)

		transactionRoutes.GET("/total-month/:month", middleware.BasicAuth(), s.TransactionHandler.GetTotalByMonth)
		transactionRoutes.GET("/summary-month/:month", middleware.BasicAuth(), s.TransactionHandler.GetSummaryByMonth)
	}

	subCategoryRoutes := r.Group("/sub-categories")
	{
		subCategoryRoutes.GET("/:id", middleware.BasicAuth(), s.SubcategoryHandler.GetById)
		subCategoryRoutes.GET("/expense", middleware.BasicAuth(), s.SubcategoryHandler.ListExpenseTypes)
		subCategoryRoutes.GET("/income", middleware.BasicAuth(), s.SubcategoryHandler.ListIncomeTypes)

		subCategoryRoutes.POST("/create", middleware.BasicAuth(), s.SubcategoryHandler.Create)
		subCategoryRoutes.PATCH("/:id", middleware.BasicAuth(), s.SubcategoryHandler.Update)
		subCategoryRoutes.DELETE("/:id", middleware.BasicAuth(), s.SubcategoryHandler.Delete)
	}

	budgetRoutes := r.Group("/budgets")
	{
		budgetRoutes.GET("/:category_id/:month", middleware.BasicAuth(), s.BudgetHandler.GetById)
		budgetRoutes.PUT("/:category_id/:month", middleware.BasicAuth(), s.BudgetHandler.Update)
		budgetRoutes.DELETE("/:category_id/:month", middleware.BasicAuth(), s.BudgetHandler.Delete)
		budgetRoutes.POST("/set-multi-year", middleware.BasicAuth(), s.BudgetHandler.SetMultiYear)
	}

	return r
}
