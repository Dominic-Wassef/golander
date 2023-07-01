package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dominic-wassef/golander/src/api"
	"github.com/dominic-wassef/golander/src/config"
	"github.com/dominic-wassef/golander/src/database"
	"github.com/dominic-wassef/golander/src/scraper"
)

func main() {
	log.Println("Starting the Golander Web Scraper...")

	// Load the application configuration
	conf, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Init the DB with retries
	db, err := initDatabaseWithRetry(conf, 10, 5*time.Second)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	err = db.CreateRepoTable()
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	log.Println("Tables created!")

	apiServer := api.NewServer(db)
	go apiServer.Start(fmt.Sprintf(":%d", conf.API.Port)) // Start API server in a separate goroutine

	// Create a new scraping task with a 24-hour interval
	scrapeFunc := scraper.ScrapeWithDB(db, 1, 10, 30, 5, time.Second*2)
	go func() {
		for {
			scrapeFunc()
			time.Sleep(24 * time.Hour)
		}
	}()

	// Wait for a termination signal
	waitForTermination()
	log.Println("Shutting down...")
}

func initDatabaseWithRetry(conf *config.Config, maxRetries int, delayBetweenRetries time.Duration) (*database.Database, error) {
	var db *database.Database
	var err error

	for i := 0; i < maxRetries; i++ {
		db, err = database.Init(conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Password, conf.Database.DBName)
		if err == nil {
			log.Println("Successfully connected to the database.")
			return db, nil
		}
		log.Printf("Failed to connect to the database (attempt %d/%d): %v. Retrying after delay...", i+1, maxRetries, err)
		time.Sleep(delayBetweenRetries)
	}

	return nil, fmt.Errorf("failed to connect to the database after %d attempts: %w", maxRetries, err)
}

func waitForTermination() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
