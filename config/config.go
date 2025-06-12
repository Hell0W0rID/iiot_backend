package config

import (
	"os"
	"strconv"
)

// Config holds all configuration for the application
type Config struct {
	// Server configuration
	Port int
	Host string

	// Database configuration
	DatabaseURL  string
	DatabaseHost string
	DatabasePort int
	DatabaseUser string
	DatabasePass string
	DatabaseName string

	// Redis configuration
	RedisURL  string
	RedisHost string
	RedisPort int

	// Security configuration
	JWTSecret     string
	APIKeyEnabled bool

	// Service configuration
	ServiceName        string
	ServiceVersion     string
	ServiceDescription string

	// Registry configuration
	RegistryHost string
	RegistryPort int

	// Logging configuration
	LogLevel  string
	LogFormat string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	config := &Config{
		// Server defaults
		Port: getEnvAsInt("PORT", 8000),
		Host: getEnv("HOST", "0.0.0.0"),

		// Database configuration
		DatabaseURL:  getEnv("DATABASE_URL", ""),
		DatabaseHost: getEnv("PGHOST", "localhost"),
		DatabasePort: getEnvAsInt("PGPORT", 5432),
		DatabaseUser: getEnv("PGUSER", "postgres"),
		DatabasePass: getEnv("PGPASSWORD", "postgres"),
		DatabaseName: getEnv("PGDATABASE", "iiot_backend"),

		// Redis configuration
		RedisURL:  getEnv("REDIS_URL", ""),
		RedisHost: getEnv("REDIS_HOST", "localhost"),
		RedisPort: getEnvAsInt("REDIS_PORT", 6379),

		// Security configuration
		JWTSecret:     getEnv("JWT_SECRET", "iiot-backend-secret-key"),
		APIKeyEnabled: getEnvAsBool("API_KEY_ENABLED", false),

		// Service configuration
		ServiceName:        "iiot-backend",
		ServiceVersion:     "1.0.0",
		ServiceDescription: "Industrial IoT Backend System",

		// Registry configuration
		RegistryHost: getEnv("REGISTRY_HOST", "localhost"),
		RegistryPort: getEnvAsInt("REGISTRY_PORT", 8500),

		// Logging configuration
		LogLevel:  getEnv("LOG_LEVEL", "info"),
		LogFormat: getEnv("LOG_FORMAT", "json"),
	}

	return config, nil
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}
