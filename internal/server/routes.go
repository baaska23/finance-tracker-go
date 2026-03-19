package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: 		[]string{"http://localhost:5173"},
		AllowMethods: 		[]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: 		[]string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials:   true,
	}))

	transactionRoutes := r.Group("/transactions")
	{
		transactionRoutes.GET("/all", s.transactionHandler.ListAll)
	}

	return r
}