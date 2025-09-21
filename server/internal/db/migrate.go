package db

import (
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	// GORM AutoMigrate is already called in db.New()
	// This function is kept for future manual migrations if needed
	log.Println("Database migrations completed")
	return nil
}
