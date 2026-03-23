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
		transactionRoutes.GET("/:id", s.TransactionHandler.GetById)
		transactionRoutes.GET("/expense", s.TransactionHandler.ListExpenses)
		transactionRoutes.GET("/income", s.TransactionHandler.ListIncomes)

		transactionRoutes.PATCH("/:id", s.TransactionHandler.Update)
		transactionRoutes.POST("/create", s.TransactionHandler.Create)
		transactionRoutes.DELETE("/:id", s.TransactionHandler.Delete)

		transactionRoutes.GET("/total-month/:month", s.TransactionHandler.GetTotalByMonth)
		transactionRoutes.GET("/summary-month/:month", s.TransactionHandler.GetSummaryByMonth)
	}

	subCategoryRoutes := r.Group("/sub-categories")
	{
		subCategoryRoutes.GET("/:id", s.SubcategoryHandler.GetById)
		subCategoryRoutes.GET("/expense", s.SubcategoryHandler.ListExpenseTypes)
		subCategoryRoutes.GET("/income", s.SubcategoryHandler.ListIncomeTypes)

		subCategoryRoutes.POST("/create", s.SubcategoryHandler.Create)
		subCategoryRoutes.PATCH("/:id", s.SubcategoryHandler.Update)
		subCategoryRoutes.DELETE("/:id", s.SubcategoryHandler.Delete)
	}

	return r
}
