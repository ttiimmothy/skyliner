package db

import (
	"log"
	"time"

	"skyliner/internal/db/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	ClassBusiness = "business"
	ClassFirst    = "first"
	ClassEconomy  = "economy"
)

func Seed(db *gorm.DB) error {
	log.Println("Seeding database...")

	// Create users
	if err := seedUsers(db); err != nil {
		return err
	}

	// Create airports
	if err := seedAirports(db); err != nil {
		return err
	}

	// Create airlines
	if err := seedAirlines(db); err != nil {
		return err
	}

	// Create flights
	if err := seedFlights(db); err != nil {
		return err
	}

	// Create seat maps
	if err := seedSeatMaps(db); err != nil {
		return err
	}

	log.Println("Database seeded successfully")
	return nil
}

func seedUsers(db *gorm.DB) error {
	// Check if users already exist
	var count int64
	db.Model(&models.User{}).Count(&count)
	if count > 0 {
		return nil
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	users := []models.User{
		{
			Email:        "traveler@example.com",
			PasswordHash: string(hashedPassword),
			FirstName:    "John",
			LastName:     "Traveler",
			Role:         models.RoleTraveler,
			IsActive:     true,
		},
		{
			Email:        "agent@example.com",
			PasswordHash: string(hashedPassword),
			FirstName:    "Jane",
			LastName:     "Agent",
			Role:         models.RoleAgent,
			IsActive:     true,
		},
		{
			Email:        "admin@example.com",
			PasswordHash: string(hashedPassword),
			FirstName:    "Admin",
			LastName:     "User",
			Role:         models.RoleAdmin,
			IsActive:     true,
		},
	}

	return db.Create(&users).Error
}

func seedAirports(db *gorm.DB) error {
	var count int64
	db.Model(&models.Airport{}).Count(&count)
	if count > 0 {
		return nil
	}

	airports := []models.Airport{
		{Code: "JFK", Name: "John F. Kennedy International Airport", City: "New York", Country: "USA", Latitude: 40.6413, Longitude: -73.7781},
		{Code: "LAX", Name: "Los Angeles International Airport", City: "Los Angeles", Country: "USA", Latitude: 33.9416, Longitude: -118.4085},
		{Code: "LHR", Name: "London Heathrow Airport", City: "London", Country: "UK", Latitude: 51.4700, Longitude: -0.4543},
		{Code: "CDG", Name: "Charles de Gaulle Airport", City: "Paris", Country: "France", Latitude: 49.0097, Longitude: 2.5479},
		{Code: "NRT", Name: "Narita International Airport", City: "Tokyo", Country: "Japan", Latitude: 35.7720, Longitude: 140.3928},
		{Code: "SFO", Name: "San Francisco International Airport", City: "San Francisco", Country: "USA", Latitude: 37.6213, Longitude: -122.3790},
	}

	return db.Create(&airports).Error
}

func seedAirlines(db *gorm.DB) error {
	var count int64
	db.Model(&models.Airline{}).Count(&count)
	if count > 0 {
		return nil
	}

	airlines := []models.Airline{
		{Code: "AA", Name: "American Airlines"},
		{Code: "DL", Name: "Delta Air Lines"},
		{Code: "UA", Name: "United Airlines"},
		{Code: "BA", Name: "British Airways"},
		{Code: "AF", Name: "Air France"},
		{Code: "JL", Name: "Japan Airlines"},
	}

	return db.Create(&airlines).Error
}

func seedFlights(db *gorm.DB) error {
	var count int64
	db.Model(&models.Flight{}).Count(&count)
	if count > 0 {
		return nil
	}

	// Get airports and airlines for foreign keys
	var airports []models.Airport
	var airlines []models.Airline
	db.Find(&airports)
	db.Find(&airlines)

	if len(airports) < 2 || len(airlines) < 1 {
		return nil
	}

	now := time.Now()
	flights := []models.Flight{
		{
			Number:        "AA100",
			AirlineID:     airlines[0].ID,
			OriginID:      airports[0].ID, // JFK
			DestinationID: airports[1].ID, // LAX
			DepartureTime: now.Add(24 * time.Hour),
			ArrivalTime:   now.Add(24*time.Hour + 5*time.Hour + 30*time.Minute),
			Duration:      330, // 5h 30m
			Stops:         0,
		},
		{
			Number:        "BA200",
			AirlineID:     airlines[3].ID,
			OriginID:      airports[0].ID, // JFK
			DestinationID: airports[2].ID, // LHR
			DepartureTime: now.Add(48 * time.Hour),
			ArrivalTime:   now.Add(48*time.Hour + 7*time.Hour),
			Duration:      420, // 7h
			Stops:         0,
		},
		{
			Number:        "AF300",
			AirlineID:     airlines[4].ID,
			OriginID:      airports[2].ID, // LHR
			DestinationID: airports[3].ID, // CDG
			DepartureTime: now.Add(72 * time.Hour),
			ArrivalTime:   now.Add(72*time.Hour + 1*time.Hour + 15*time.Minute),
			Duration:      75, // 1h 15m
			Stops:         0,
		},
	}

	if err := db.Create(&flights).Error; err != nil {
		return err
	}

	// Create fares for each flight
	for _, flight := range flights {
		fares := []models.Fare{
			{
				FlightID:  flight.ID,
				Class:     "economy",
				FareType:  "basic",
				BasePrice: 299.99,
				Currency:  "USD",
				Available: 50,
			},
			{
				FlightID:  flight.ID,
				Class:     "economy",
				FareType:  "standard",
				BasePrice: 399.99,
				Currency:  "USD",
				Available: 30,
			},
			{
				FlightID:  flight.ID,
				Class:     "business",
				FareType:  "flexible",
				BasePrice: 899.99,
				Currency:  "USD",
				Available: 12,
			},
		}
		db.Create(&fares)
	}

	return nil
}

func seedSeatMaps(db *gorm.DB) error {
	var count int64
	db.Model(&models.SeatMap{}).Count(&count)
	if count > 0 {
		return nil
	}

	var flights []models.Flight
	db.Find(&flights)

	for _, flight := range flights {
		// Create seat maps for each class
		classes := []string{"economy", "business", "first"}
		for _, class := range classes {
			seatMap := models.SeatMap{
				FlightID: flight.ID,
				Class:    class,
			}
			db.Create(&seatMap)

			// Create seats for this map
			rows := 20
			if class == ClassBusiness {
				rows = 8
			} else if class == ClassFirst {
				rows = 4
			}

			columns := []string{"A", "B", "C", "D", "E", "F"}
			if class == ClassBusiness {
				columns = []string{"A", "B", "C", "D"}
			} else if class == ClassFirst {
				columns = []string{"A", "B"}
			}

			for row := 1; row <= rows; row++ {
				for _, col := range columns {
					seat := models.Seat{
						SeatMapID: seatMap.ID,
						Row:       row,
						Column:    col,
						Class:     class,
						Status:    models.SeatAvailable,
					}
					if class == "business" {
						price := 50.0
						seat.Price = &price
					} else if class == "first" {
						price := 100.0
						seat.Price = &price
					}
					db.Create(&seat)
				}
			}
		}
	}

	return nil
}
