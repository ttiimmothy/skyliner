package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"skyliner/internal/config"
	"skyliner/internal/db/models"
	httpRouter "skyliner/internal/http"
	"skyliner/internal/ws"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}
	return db
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	// Setup test database
	testDB := setupTestDB()

	// Auto-migrate test models
	err := testDB.AutoMigrate(
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
	if err != nil {
		panic("Failed to migrate test database")
	}

	// Setup config
	cfg := &config.Config{
		JWTSecret:   "test-secret-key-for-testing-purposes-only",
		CORSOrigins: []string{"http://localhost:3000"},
	}

	// Setup WebSocket hub
	hub := ws.NewHub()
	go hub.Run()

	// Setup router
	router := gin.Default()
	httpRouter.SetupRoutes(router, testDB, hub, cfg)

	return router
}

func TestSignupEndpoint(t *testing.T) {
	router := setupTestRouter()

	// Test data
	signupData := map[string]string{
		"email":      "test@example.com",
		"password":   "password123",
		"first_name": "Test",
		"last_name":  "User",
	}

	jsonData, _ := json.Marshal(signupData)
	req, _ := http.NewRequest("POST", "/api/v1/auth/signup", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code, "Expected status 201, got %d", w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Failed to unmarshal response")
	assert.Contains(t, response, "user", "Response should contain user data")
	assert.Contains(t, response, "access_token", "Response should contain access token")
	assert.Contains(t, response, "refresh_token", "Response should contain refresh token")
}

func TestLoginEndpoint(t *testing.T) {
	router := setupTestRouter()

	// First create a user
	signupData := map[string]string{
		"email":      "test@example.com",
		"password":   "password123",
		"first_name": "Test",
		"last_name":  "User",
	}

	jsonData, _ := json.Marshal(signupData)
	req, _ := http.NewRequest("POST", "/api/v1/auth/signup", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code, "Signup should succeed")

	// Now test login
	loginData := map[string]string{
		"email":    "test@example.com",
		"password": "password123",
	}

	jsonData, _ = json.Marshal(loginData)
	req, _ = http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200, got %d", w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Failed to unmarshal response")
	assert.Contains(t, response, "user", "Response should contain user data")
	assert.Contains(t, response, "access_token", "Response should contain access token")
}

func TestGetAirportsEndpoint(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/v1/airports", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200, got %d", w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Failed to unmarshal response")
	assert.Contains(t, response, "airports", "Response should contain airports data")
}

func TestSearchFlightsEndpoint(t *testing.T) {
	router := setupTestRouter()

	searchData := map[string]interface{}{
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
	}

	jsonData, _ := json.Marshal(searchData)
	req, _ := http.NewRequest("POST", "/api/v1/search", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200, got %d", w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Failed to unmarshal response")
	assert.Contains(t, response, "flights", "Response should contain flights data")
	assert.Contains(t, response, "total", "Response should contain total count")
}

func TestInvalidSignupData(t *testing.T) {
	router := setupTestRouter()

	// Test with invalid email
	invalidData := map[string]string{
		"email":      "invalid-email",
		"password":   "password123",
		"first_name": "Test",
		"last_name":  "User",
	}

	jsonData, _ := json.Marshal(invalidData)
	req, _ := http.NewRequest("POST", "/api/v1/auth/signup", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status 400 for invalid email")
}

func TestInvalidLoginData(t *testing.T) {
	router := setupTestRouter()

	// Test with non-existent user
	loginData := map[string]string{
		"email":    "nonexistent@example.com",
		"password": "password123",
	}

	jsonData, _ := json.Marshal(loginData)
	req, _ := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code, "Expected status 401 for invalid credentials")
}
