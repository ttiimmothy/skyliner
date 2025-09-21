package db

import (
	"testing"

	"skyliner/internal/db/models"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNew(t *testing.T) {
	// Test with SQLite for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	// Test auto-migration
	err = db.AutoMigrate(
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
	)
	assert.NoError(t, err)

	// Test that tables were created
	var count int64
	db.Raw("SELECT COUNT(*) FROM sqlite_master WHERE type='table'").Scan(&count)
	assert.Greater(t, count, int64(0))
}

func TestNewWithInvalidURL(t *testing.T) {
	// Test with invalid database URL
	_, err := New("invalid://url")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to connect to database")
}
