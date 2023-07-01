package config

import (
	"errors"
	"net/url"
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
	// Set defaults
	portStr := getEnv("API_PORT", "8080")
	dbportStr := getEnv("DB_PORT", "3306")
	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "root")
	password := getEnv("DB_PASSWORD", "")
	dbname := getEnv("DB_NAME", "golander")
	scraperURL := getEnv("SCRAPER_URL", "https://github.com/trending")

	// Parse integers
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}
	dbPort, err := strconv.Atoi(dbportStr)
	if err != nil {
		return nil, err
	}

	// Validate
	if port < 1 || port > 65535 {
		return nil, errors.New("invalid API port")
	}
	if dbPort < 1 || dbPort > 65535 {
		return nil, errors.New("invalid database port")
	}
	_, err = url.Parse(scraperURL)
	if err != nil {
		return nil, errors.New("invalid scraper URL")
	}

	return &Config{
		Database: DatabaseConfig{
			Host:     host,
			Port:     dbPort,
			User:     user,
			Password: password,
			DBName:   dbname,
		},
		API: APIConfig{
			Port: port,
		},
		Scraper: ScraperConfig{
			URL: scraperURL,
		},
	}, nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}
