package main

import (
	"log"

	"skyliner/internal/config"
	"skyliner/internal/db"
	"skyliner/internal/http"
	"skyliner/internal/ws"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Initialize database
	database, err := db.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations
	if err := db.Migrate(database); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Initialize WebSocket hub
	hub := ws.NewHub()
	go hub.Run()

	// Initialize HTTP server
	router := gin.Default()

	// Setup routes
	http.SetupRoutes(router, database, hub, cfg)

	log.Printf("Server starting on port %s", "8080")
	if err := router.Run(":" + "8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
