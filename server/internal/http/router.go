package http

import (
	"skyliner/internal/config"
	"skyliner/internal/http/handlers"
	"skyliner/internal/http/middleware"
	"skyliner/internal/ws"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB, hub *ws.Hub, cfg *config.Config) {
	// Middleware
	router.Use(middleware.CORS(cfg.CORSOrigins))
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": "skyliner-api",
			"version": "1.0.0",
		})
	})

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(db, cfg)
	searchHandler := handlers.NewSearchHandler(db)
	bookingHandler := handlers.NewBookingHandler(db, cfg)
	paymentHandler := handlers.NewPaymentHandler(db, cfg)

	// API routes
	api := router.Group("/api/v1")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/signup", authHandler.Signup)
			auth.POST("/login", authHandler.Login)
			auth.POST("/google", authHandler.GoogleAuth)
			auth.POST("/refresh", authHandler.Refresh)
			auth.POST("/logout", authHandler.Logout)
		}

		// Public routes
		api.GET("/airports", searchHandler.GetAirports)
		api.GET("/airlines", searchHandler.GetAirlines)
		api.POST("/search", searchHandler.SearchFlights)
		api.GET("/flights/:id/seatmap", searchHandler.GetSeatMap)

		// Protected routes
		protected := api.Group("")
		protected.Use(middleware.AuthRequired(cfg.JWTSecret))
		{
			// Booking routes
			bookings := protected.Group("/bookings")
			{
				bookings.POST("", bookingHandler.CreateBooking)
				bookings.GET("/:id", bookingHandler.GetBooking)
				bookings.POST("/:id/issue", bookingHandler.IssueBooking)
				bookings.POST("/:id/cancel", bookingHandler.CancelBooking)
			}

			// Payment routes
			payments := protected.Group("/payments")
			{
				payments.POST("/checkout-session", paymentHandler.CreateCheckoutSession)
				payments.POST("/billing-portal", paymentHandler.CreateBillingPortal)
			}
		}

		// Admin/Agent routes
		admin := api.Group("/admin")
		admin.Use(middleware.AuthRequired(cfg.JWTSecret))
		admin.Use(middleware.AdminRequired())
		{
			admin.GET("/bookings", bookingHandler.GetAllBookings)
			admin.POST("/bookings/:id/waive", bookingHandler.WaiveBooking)
			admin.POST("/reprice", bookingHandler.RepriceBookings)
		}
	}

	// WebSocket endpoint
	router.GET("/ws", func(c *gin.Context) {
		ws.HandleWebSocket(c, hub)
	})

	// Stripe webhook
	router.POST("/webhooks/stripe", paymentHandler.HandleStripeWebhook)
}
