package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/tickets", s.userHandler.UserOrdersHandler)
		userRoutes.GET("/phone", s.userHandler.UserPhoneHandler)
		userRoutes.GET("/subscriptions", s.userHandler.UserSubscriptionsHandler)
		userRoutes.POST("/subscriptions/cancel", s.userHandler.UserSubscriptionCancelHandler)
		// userRoutes.GET("/parentalPin", s.userHandler.UserParentalPinHandler)
	}

	billinRoutes := r.Group("/billing")
	// i think this part is done
	{
		billinRoutes.POST("/purchase", s.billingHandler.BillingPurchaseHandler)
	}

	authRoutes := r.Group("/auth")
	// i think this part is done
	{
		authRoutes.POST("/otp/send", s.authHandler.SendOtpHandler)
		authRoutes.POST("/otp/verify", s.authHandler.VerifyOtpHandler)
		authRoutes.POST("/otp/resend", s.authHandler.ResendOtpHandler)
	}
	contentRoutes := r.Group("/contents")
	{
		contentRoutes.GET("/recommended", s.contentHandler.ContentRecommendedHandler)
		contentRoutes.GET("/search", s.contentHandler.ContentSearchHandler)
		contentRoutes.GET("/get-video-source", s.contentHandler.ContentVideoSourceHandler)
		contentRoutes.GET("/billing-info", s.contentHandler.ContentBillingInfoHandler)
		contentRoutes.GET("/faq", s.contentHandler.FaqHandler)
		// ene 2 iig yaydo baiz neg arga olno baiga
		// contentRoutes.GET("/tempCarousel", handlers.TempCarouselHandler)
		// contentRoutes.GET("/tempCarouselSvod", handlers.TempCarouselSvodHandler)
	}

	return r
}

func (s *http.Server) RegisterRoutes() http.Handler{
	r := gin.Default()
	
}
