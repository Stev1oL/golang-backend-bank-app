package config

import "os"

type Config struct {
	DatabaseURL string
	ServerPort  string
}

func LoadConfig() *Config {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:level1ngs0lo@localhost:5432/bankdb?sslmode=disable"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		DatabaseURL: dbURL,
		ServerPort:  port,
	}
}
