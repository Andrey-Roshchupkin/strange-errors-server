package config

import (
	"os"
)

// Config holds the application configuration
type Config struct {
	Port     string
	DBPath   string
	LogLevel string
}

// LoadConfig loads configuration from environment variables with defaults
func LoadConfig() *Config {
	return &Config{
		Port:     getEnv("PORT", ":3000"),
		DBPath:   getEnv("DB_PATH", "./database.db"),
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}
}

// getEnv gets an environment variable with a fallback default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
