package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"skyliner/internal/config"
	"skyliner/internal/db/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	_ = db.AutoMigrate(&models.User{})
	return db
}

func TestCORS(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	origins := []string{"http://localhost:3000", "https://example.com"}
	router.Use(CORS(origins))
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	// Test allowed origin
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "http://localhost:3000", w.Header().Get("Access-Control-Allow-Origin"))

	// Test disallowed origin
	req, _ = http.NewRequest("GET", "/test", nil)
	req.Header.Set("Origin", "http://malicious.com")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "", w.Header().Get("Access-Control-Allow-Origin"))

	// Test OPTIONS request
	req, _ = http.NewRequest("OPTIONS", "/test", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestAuthRequired(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	cfg := &config.Config{JWTSecret: "test-secret"}
	router.Use(AuthRequired(cfg.JWTSecret))
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	// Test without token
	req, _ := http.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Test with invalid token
	req, _ = http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Test with valid token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": float64(1),
		"exp":     time.Now().Add(time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(cfg.JWTSecret))

	req, _ = http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAdminRequired(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := setupTestDB()

	// Test without user_id
	router1 := gin.New()
	router1.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	router1.Use(AdminRequired())
	router1.GET("/admin", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	req, _ := http.NewRequest("GET", "/admin", nil)
	w := httptest.NewRecorder()
	router1.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Test with non-admin user
	user := models.User{
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
		Role:      models.RoleTraveler,
	}
	db.Create(&user)

	router2 := gin.New()
	router2.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("user_id", user.ID)
		c.Next()
	})
	router2.Use(AdminRequired())
	router2.GET("/admin", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	req, _ = http.NewRequest("GET", "/admin", nil)
	w = httptest.NewRecorder()
	router2.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)

	// Test with admin user
	admin := models.User{
		Email:     "admin@example.com",
		FirstName: "Admin",
		LastName:  "User",
		Role:      models.RoleAdmin,
	}
	db.Create(&admin)

	router3 := gin.New()
	router3.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("user_id", admin.ID)
		c.Next()
	})
	router3.Use(AdminRequired())
	router3.GET("/admin", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	req, _ = http.NewRequest("GET", "/admin", nil)
	w = httptest.NewRecorder()
	router3.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
