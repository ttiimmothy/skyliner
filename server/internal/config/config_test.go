package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	// Test with default values
	cfg, err := Load()
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "postgres://postgres:postgres@localhost:5432/flightdb?sslmode=disable", cfg.DatabaseURL)
	assert.Equal(t, "redis://localhost:6379", cfg.RedisURL)
	assert.Equal(t, "change_me", cfg.JWTSecret)
	assert.Equal(t, 15*time.Minute, cfg.JWTAccessTTL)
	assert.Equal(t, 168*time.Hour, cfg.JWTRefreshTTL)
}

func TestLoadWithEnvVars(t *testing.T) {
	// Set environment variables
	os.Setenv("DATABASE_URL", "postgres://test:test@localhost:5432/testdb")
	os.Setenv("JWT_SECRET", "test-secret")
	os.Setenv("JWT_ACCESS_TTL", "30m")
	os.Setenv("JWT_REFRESH_TTL", "24h")

	// Clean up after test
	defer func() {
		os.Unsetenv("DATABASE_URL")
		os.Unsetenv("JWT_SECRET")
		os.Unsetenv("JWT_ACCESS_TTL")
		os.Unsetenv("JWT_REFRESH_TTL")
	}()

	cfg, err := Load()
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "postgres://test:test@localhost:5432/testdb", cfg.DatabaseURL)
	assert.Equal(t, "test-secret", cfg.JWTSecret)
	assert.Equal(t, 30*time.Minute, cfg.JWTAccessTTL)
	assert.Equal(t, 24*time.Hour, cfg.JWTRefreshTTL)
}

func TestParseDuration(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Duration
	}{
		{"15m", 15 * time.Minute},
		{"1h", 1 * time.Hour},
		{"168h", 168 * time.Hour},
		{"invalid", 15 * time.Minute}, // fallback
	}

	for _, test := range tests {
		result := parseDuration(test.input)
		assert.Equal(t, test.expected, result)
	}
}

func TestGetEnv(t *testing.T) {
	// Test with existing env var
	os.Setenv("TEST_VAR", "test-value")
	defer os.Unsetenv("TEST_VAR")

	assert.Equal(t, "test-value", getEnv("TEST_VAR", "default"))

	// Test with non-existing env var
	assert.Equal(t, "default-value", getEnv("NON_EXISTING_VAR", "default-value"))
}
