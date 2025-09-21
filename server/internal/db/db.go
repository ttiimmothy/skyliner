package db

import (
	"fmt"
	"log"

	"skyliner/internal/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto-migrate all models
	if err := db.AutoMigrate(
		&models.User{},
		&models.Airport{},
		&models.Airline{},
		&models.Flight{},
		&models.Fare{},
		&models.SeatMap{},
		&models.Seat{},
		&models.Booking{},
		&models.Itinerary{},
		&models.Segment{},
		&models.Passenger{},
		&models.Payment{},
		&models.Baggage{},
	); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database connected and migrated successfully")
	return db, nil
}
