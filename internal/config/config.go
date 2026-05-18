package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	DBUrlMigration string
	SecretJWT      string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Port:           getEnv("PORT", "8080"),
		DBUrlMigration: getEnv("DATABASE_URL", ""),
		SecretJWT:      getEnv("SECRET_JWT", ""),

		DBHost:     getEnv("DB_HOST", ""),
		DBPort:     getEnv("DB_PORT", ""),
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", ""),
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}