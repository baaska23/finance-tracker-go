// Package server
package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"secret_room_backend_v2/internal/domains/auth"
	"secret_room_backend_v2/internal/domains/billing"
	"secret_room_backend_v2/internal/domains/content"
	"secret_room_backend_v2/internal/domains/user"
	"secret_room_backend_v2/internal/platform/database"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port           int
	dbManager      *database.DBManager
	userHandler    *user.Handler
	billingHandler *billing.Handler
	authHandler    *auth.Handler
	contentHandler *content.Handler
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	dbManager, err := database.NewDBManager()
	if err != nil {
		log.Fatalf("Failed to initialize database manager: %v", err)
	}

	contentRepo := content.NewRepository(dbManager.RTV)
	contentService := content.NewService(contentRepo)
	contentHandler := content.NewHandler(contentService)

	userRepo := user.NewRepository(dbManager.RTV) // or use a separate auth/user DB when available
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	billingRepo := billing.NewRepository() // Update this when billing DB is added
	billingService := billing.NewService(billingRepo)
	billingHandler := billing.NewHandler(billingService)

	authRepo := auth.NewRepository() // Update this when auth DB is added
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	srv := &Server{
		port:           port,
		dbManager:      dbManager,
		userHandler:    userHandler,
		billingHandler: billingHandler,
		authHandler:    authHandler,
		contentHandler: contentHandler,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", srv.port),
		Handler:      srv.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func (s *Server) Close() error {
	if s.dbManager != nil {
		return s.dbManager.Close()
	}
	return nil
}
