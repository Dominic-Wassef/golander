package config_test

import (
	"os"
	"testing"

	"github.com/dominic-wassef/golander/src/config"
)

func TestLoad(t *testing.T) {
	// set up your env vars for testing
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("DB_NAME", "golander")
	os.Setenv("API_PORT", "8080")
	os.Setenv("SCRAPER_URL", "https://github.com/trending")

	// call your Load function
	c, err := config.Load()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// add your assertions
	if c.Database.Host != "localhost" {
		t.Errorf("Expected localhost, got %v", c.Database.Host)
	}
	if c.Database.Port != 3306 {
		t.Errorf("Expected 3306, got %v", c.Database.Port)
	}
	if c.Database.User != "root" {
		t.Errorf("Expected user, got %v", c.Database.User)
	}
	if c.Database.Password != "" {
		t.Errorf("Expected password, got %v", c.Database.Password)
	}
	if c.Database.DBName != "golander" {
		t.Errorf("Expected testdb, got %v", c.Database.DBName)
	}
	if c.API.Port != 8080 {
		t.Errorf("Expected 8080, got %v", c.API.Port)
	}
	if c.Scraper.URL != "https://github.com/trending" {
		t.Errorf("Expected https://github.com/trending, got %v", c.Scraper.URL)
	}
}
