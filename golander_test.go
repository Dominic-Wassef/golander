package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/dominic-wassef/golander/src/config"
	"github.com/dominic-wassef/golander/src/database"
)

type MockDBConnector struct {
	DB *sql.DB
}

func (md *MockDBConnector) Init(host string, port int, user string, password string, dbname string) (*database.Database, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)
	fmt.Println("dsn:", dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	md.DB = db

	return &database.Database{
		DB: db,
	}, nil
}

func (md *MockDBConnector) Close() error {
	if md.DB != nil {
		return md.DB.Close()
	}
	return nil
}

type MockConfigLoader struct{}

func (cl MockConfigLoader) Load() (*config.Config, error) {
	// Create a mock configuration with dummy values
	return &config.Config{
		API: config.APIConfig{
			Port: 1234,
		},
		Database: config.DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     3306,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
		Scraper: config.ScraperConfig{
			URL: "dummyURL",
		},
	}, nil
}

type MockConfigLoaderWithError struct{}

func (cl MockConfigLoaderWithError) Load() (*config.Config, error) {
	return nil, errors.New("mock error")
}

var mockDB *MockDBConnector

func TestMain(m *testing.M) {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "root")
	os.Setenv("DB_NAME", "golander")
	os.Setenv("API_PORT", "8080")
	os.Setenv("SCRAPER_URL", "https://github.com/trending")
	dbConnector := func(host string, port int, user string, password string, dbname string) (*database.Database, error) {
		if mockDB == nil {
			mockDB = &MockDBConnector{}
		}
		return mockDB.Init(host, port, user, password, dbname)
	}

	confLoader := func() (*config.Config, error) {
		mockConfig := MockConfigLoader{}
		return mockConfig.Load()
	}

	go func() {
		time.Sleep(time.Millisecond * 10)
		p, _ := os.FindProcess(os.Getpid())
		p.Signal(syscall.SIGINT)
	}()

	mainWithDependencies(dbConnector, confLoader)
}

func TestSuccessfulStartup(t *testing.T) {
	// Perform any necessary setup here
	mockDB := &MockDBConnector{}
	defer func() {
		if err := mockDB.Close(); err != nil {
			t.Errorf("Failed to close mock database: %v", err)
		}
	}()

	t.Run("successful startup", func(t *testing.T) {
		// Test logic goes here
		confLoader := func() (*config.Config, error) {
			mockConfig := MockConfigLoader{}
			return mockConfig.Load()
		}

		dbConnector := func(host string, port int, user string, password string, dbname string) (*database.Database, error) {
			fmt.Println("DB_NAME:", os.Getenv("DB_NAME"))
			db, err := mockDB.Init(host, port, user, password, dbname)
			if err != nil {
				return nil, err
			}
			return db, nil
		}

		// Wrap function call with anonymous function to capture panic
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("The code panicked: %v", r)
				}
			}()
			mainWithDependencies(dbConnector, confLoader)
		}()
		// If we reach this point, function did not panic
	})
}

func TestConfigLoadError(t *testing.T) {
	// Perform any necessary setup here
	mockDB := &MockDBConnector{}
	defer func() {
		if err := mockDB.Close(); err != nil {
			t.Errorf("Failed to close mock database: %v", err)
		}
	}()

	t.Run("config.Load error", func(t *testing.T) {
		// Test logic goes here
		confLoader := func() (*config.Config, error) {
			mockConfig := MockConfigLoaderWithError{}
			return mockConfig.Load()
		}

		dbConnector := func(host string, port int, user string, password string, dbname string) (*database.Database, error) {
			// Return an error to simulate a failed database connection
			return nil, errors.New("database connection error")
		}

		// Expect a panic from mainWithDependencies because config.Load returns an error
		defer func() {
			if r := recover(); r == nil {
				t.Error("mainWithDependencies did not panic, but it was expected to")
			}
		}()
		mainWithDependencies(dbConnector, confLoader)
	})
}
