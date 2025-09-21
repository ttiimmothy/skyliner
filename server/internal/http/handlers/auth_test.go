package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"skyliner/internal/config"
	"skyliner/internal/db/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupAuthTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	_ = db.AutoMigrate(&models.User{})
	return db
}

func setupAuthTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	db := setupAuthTestDB()
	cfg := &config.Config{
		JWTSecret:     "test-secret",
		JWTAccessTTL:  15 * 60 * 1000000000,       // 15 minutes in nanoseconds
		JWTRefreshTTL: 168 * 60 * 60 * 1000000000, // 168 hours in nanoseconds
	}

	router := gin.New()
	authHandler := NewAuthHandler(db, cfg)

	router.POST("/signup", authHandler.Signup)
	router.POST("/login", authHandler.Login)
	router.POST("/google", authHandler.GoogleAuth)
	router.POST("/refresh", authHandler.Refresh)
	router.POST("/logout", authHandler.Logout)

	return router
}

func TestAuthHandler_Signup(t *testing.T) {
	router := setupAuthTestRouter()

	tests := []struct {
		name           string
		requestBody    map[string]string
		expectedStatus int
		expectError    bool
	}{
		{
			name: "valid signup",
			requestBody: map[string]string{
				"email":      "test@example.com",
				"password":   "password123",
				"first_name": "Test",
				"last_name":  "User",
			},
			expectedStatus: http.StatusCreated,
			expectError:    false,
		},
		{
			name: "invalid email",
			requestBody: map[string]string{
				"email":      "invalid-email",
				"password":   "password123",
				"first_name": "Test",
				"last_name":  "User",
			},
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name: "short password",
			requestBody: map[string]string{
				"email":      "test@example.com",
				"password":   "123",
				"first_name": "Test",
				"last_name":  "User",
			},
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name: "missing first name",
			requestBody: map[string]string{
				"email":     "test@example.com",
				"password":  "password123",
				"last_name": "User",
			},
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if !tt.expectError {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Contains(t, response, "user")
				assert.Contains(t, response, "access_token")
				assert.Contains(t, response, "refresh_token")
			}
		})
	}
}

func TestAuthHandler_Login(t *testing.T) {
	router := setupAuthTestRouter()

	// First create a user
	signupData := map[string]string{
		"email":      "test@example.com",
		"password":   "password123",
		"first_name": "Test",
		"last_name":  "User",
	}

	jsonData, _ := json.Marshal(signupData)
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	tests := []struct {
		name           string
		requestBody    map[string]string
		expectedStatus int
		expectError    bool
	}{
		{
			name: "valid login",
			requestBody: map[string]string{
				"email":    "test@example.com",
				"password": "password123",
			},
			expectedStatus: http.StatusOK,
			expectError:    false,
		},
		{
			name: "invalid email",
			requestBody: map[string]string{
				"email":    "nonexistent@example.com",
				"password": "password123",
			},
			expectedStatus: http.StatusUnauthorized,
			expectError:    true,
		},
		{
			name: "invalid password",
			requestBody: map[string]string{
				"email":    "test@example.com",
				"password": "wrongpassword",
			},
			expectedStatus: http.StatusUnauthorized,
			expectError:    true,
		},
		{
			name: "invalid email format",
			requestBody: map[string]string{
				"email":    "invalid-email",
				"password": "password123",
			},
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if !tt.expectError {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Contains(t, response, "user")
				assert.Contains(t, response, "access_token")
				assert.Contains(t, response, "refresh_token")
			}
		})
	}
}

func TestAuthHandler_GoogleAuth(t *testing.T) {
	router := setupAuthTestRouter()

	req, _ := http.NewRequest("POST", "/google", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotImplemented, w.Code)
}

func TestAuthHandler_Refresh(t *testing.T) {
	router := setupAuthTestRouter()

	req, _ := http.NewRequest("POST", "/refresh", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotImplemented, w.Code)
}

func TestAuthHandler_Logout(t *testing.T) {
	router := setupAuthTestRouter()

	req, _ := http.NewRequest("POST", "/logout", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "message")
	assert.Equal(t, "Logged out successfully", response["message"])
}
