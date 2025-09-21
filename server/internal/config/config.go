package config

import (
	"os"
	"time"
)

type Config struct {
	DatabaseURL         string
	RedisURL            string
	JWTSecret           string
	JWTAccessTTL        time.Duration
	JWTRefreshTTL       time.Duration
	GoogleClientID      string
	StripeSecretKey     string
	StripeWebhookSecret string
	CORSOrigins         []string
	Port                string
}

func Load() (*Config, error) {
	config := &Config{
		DatabaseURL:         getEnv("DATABASE_URL", "postgres://myuser:mysecretpassword@localhost:5432/skyliner?sslmode=disable"),
		RedisURL:            getEnv("REDIS_URL", "redis://localhost:6379"),
		JWTSecret:           getEnv("JWT_SECRET", "change_me"),
		JWTAccessTTL:        parseDuration(getEnv("JWT_ACCESS_TTL", "15m")),
		JWTRefreshTTL:       parseDuration(getEnv("JWT_REFRESH_TTL", "168h")),
		GoogleClientID:      getEnv("GOOGLE_CLIENT_ID", ""),
		StripeSecretKey:     getEnv("STRIPE_SECRET_KEY", ""),
		StripeWebhookSecret: getEnv("STRIPE_WEBHOOK_SECRET", ""),
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func parseDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		// Default fallback
		if s == "15m" {
			return 15 * time.Minute
		}
		if s == "168h" {
			return 168 * time.Hour
		}
		return 15 * time.Minute
	}
	return d
}
