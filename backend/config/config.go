package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config represents the application configuration.
type Config struct {
	Port       string
	Env        string
	JWTSecret  string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// LoadConfig loads configuration from environment variables and .env file.
func LoadConfig() *Config {
	// Attempt to load .env file. Ignore error in production if env variables are already set.
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using system environment variables instead")
	}

	return &Config{
		Port:       getEnv("PORT", "8080"),
		Env:       getEnv("ENV", "development"),
		JWTSecret:  getEnv("JWT_SECRET", "default-jwt-secret-key-change-in-production"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "secret"),
		DBName:     getEnv("DB_NAME", "learning_db"),
	}
}

// getEnv gets an environment variable or returns a default value.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
