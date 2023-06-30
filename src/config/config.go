package config

import (
	"os"
	"strconv"
)

type Config struct {
	Database DatabaseConfig
	API      APIConfig
	Scraper  ScraperConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type APIConfig struct {
	Port int
}

type ScraperConfig struct {
	URL string
}

func Load() (*Config, error) {
	portStr := os.Getenv("API_PORT")
	if portStr == "" {
		portStr = "8080" // default value
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	dbportStr := os.Getenv("DB_PORT") // Use "DB_PORT" not "API_PORT"
	if dbportStr == "" {
		dbportStr = "3306" // default value
	}
	dbPort, err := strconv.Atoi(dbportStr)
	if err != nil {
		return nil, err
	}

	return &Config{
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     dbPort,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
		API: APIConfig{
			Port: port,
		},
		Scraper: ScraperConfig{
			URL: os.Getenv("SCRAPER_URL"),
		},
	}, nil
}
