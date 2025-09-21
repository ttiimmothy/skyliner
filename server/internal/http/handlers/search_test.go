package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"skyliner/internal/db/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupSearchTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	_ = db.AutoMigrate(
		&models.Airport{},
		&models.Airline{},
		&models.Flight{},
		&models.Fare{},
		&models.SeatMap{},
		&models.Seat{},
	)

	// Seed test data
	airports := []models.Airport{
		{Code: "JFK", Name: "John F. Kennedy International Airport", City: "New York", Country: "USA"},
		{Code: "LAX", Name: "Los Angeles International Airport", City: "Los Angeles", Country: "USA"},
		{Code: "LHR", Name: "Heathrow Airport", City: "London", Country: "UK"},
	}
	db.Create(&airports)

	airlines := []models.Airline{
		{Code: "AA", Name: "American Airlines"},
		{Code: "DL", Name: "Delta Air Lines"},
		{Code: "BA", Name: "British Airways"},
	}
	db.Create(&airlines)

	flights := []models.Flight{
		{
			Number:        "AA100",
			AirlineID:     airlines[0].ID,
			OriginID:      airports[0].ID,
			DestinationID: airports[1].ID,
			DepartureTime: time.Date(2024, 12, 25, 10, 0, 0, 0, time.UTC),
			ArrivalTime:   time.Date(2024, 12, 25, 13, 0, 0, 0, time.UTC),
			Duration:      180,
			Stops:         0,
		},
		{
			Number:        "DL200",
			AirlineID:     airlines[1].ID,
			OriginID:      airports[0].ID,
			DestinationID: airports[1].ID,
			DepartureTime: time.Date(2024, 12, 25, 14, 0, 0, 0, time.UTC),
			ArrivalTime:   time.Date(2024, 12, 25, 17, 0, 0, 0, time.UTC),
			Duration:      180,
			Stops:         0,
		},
	}
	db.Create(&flights)

	fares := []models.Fare{
		{FlightID: flights[0].ID, Class: "economy", FareType: "basic", BasePrice: 299.99, Currency: "USD", Available: 10},
		{FlightID: flights[0].ID, Class: "business", FareType: "flexible", BasePrice: 899.99, Currency: "USD", Available: 5},
		{FlightID: flights[1].ID, Class: "economy", FareType: "standard", BasePrice: 349.99, Currency: "USD", Available: 8},
	}
	db.Create(&fares)

	return db
}

func setupSearchTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	db := setupSearchTestDB()
	searchHandler := NewSearchHandler(db)

	router := gin.New()
	router.GET("/airports", searchHandler.GetAirports)
	router.GET("/airlines", searchHandler.GetAirlines)
	router.POST("/search", searchHandler.SearchFlights)
	router.GET("/flights/:id/seatmap", searchHandler.GetSeatMap)

	return router
}

func TestSearchHandler_GetAirports(t *testing.T) {
	router := setupSearchTestRouter()

	req, _ := http.NewRequest("GET", "/airports", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "airports")

	airports := response["airports"].([]interface{})
	assert.Greater(t, len(airports), 0)
}

func TestSearchHandler_GetAirlines(t *testing.T) {
	router := setupSearchTestRouter()

	req, _ := http.NewRequest("GET", "/airlines", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "airlines")

	airlines := response["airlines"].([]interface{})
	assert.Greater(t, len(airlines), 0)
}

func TestSearchHandler_SearchFlights(t *testing.T) {
	router := setupSearchTestRouter()

	tests := []struct {
		name           string
		requestBody    map[string]interface{}
		expectedStatus int
		expectError    bool
	}{
		{
			name: "valid search",
			requestBody: map[string]interface{}{
				"trip_type": "one-way",
				"legs": []map[string]string{
					{
						"origin":      "JFK",
						"destination": "LAX",
						"date":        "2024-12-25",
					},
				},
				"passengers":  1,
				"cabin":       "",
				"flexibility": 0,
			},
			expectedStatus: http.StatusOK,
			expectError:    false,
		},
		{
			name: "search with cabin filter",
			requestBody: map[string]interface{}{
				"trip_type": "one-way",
				"legs": []map[string]string{
					{
						"origin":      "JFK",
						"destination": "LAX",
						"date":        "2024-12-25",
					},
				},
				"passengers":  1,
				"cabin":       "economy",
				"flexibility": 0,
			},
			expectedStatus: http.StatusOK,
			expectError:    false,
		},
		{
			name: "search with flexibility",
			requestBody: map[string]interface{}{
				"trip_type": "one-way",
				"legs": []map[string]string{
					{
						"origin":      "JFK",
						"destination": "LAX",
						"date":        "2024-12-25",
					},
				},
				"passengers":  1,
				"cabin":       "",
				"flexibility": 1,
			},
			expectedStatus: http.StatusOK,
			expectError:    false,
		},
		{
			name: "invalid trip type",
			requestBody: map[string]interface{}{
				"trip_type": "",
				"legs": []map[string]string{
					{
						"origin":      "JFK",
						"destination": "LAX",
						"date":        "2024-12-25",
					},
				},
				"passengers":  1,
				"cabin":       "",
				"flexibility": 0,
			},
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name: "no legs",
			requestBody: map[string]interface{}{
				"trip_type":   "one-way",
				"legs":        []map[string]string{},
				"passengers":  1,
				"cabin":       "",
				"flexibility": 0,
			},
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name: "invalid date format",
			requestBody: map[string]interface{}{
				"trip_type": "one-way",
				"legs": []map[string]string{
					{
						"origin":      "JFK",
						"destination": "LAX",
						"date":        "invalid-date",
					},
				},
				"passengers":  1,
				"cabin":       "",
				"flexibility": 0,
			},
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest("POST", "/search", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if !tt.expectError {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Contains(t, response, "flights")
				assert.Contains(t, response, "total")
			}
		})
	}
}

func TestSearchHandler_GetSeatMap(t *testing.T) {
	router := setupSearchTestRouter()

	// Test with valid flight ID
	req, _ := http.NewRequest("GET", "/flights/1/seatmap", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "seat_maps")

	// Test with invalid flight ID
	req, _ = http.NewRequest("GET", "/flights/invalid/seatmap", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
