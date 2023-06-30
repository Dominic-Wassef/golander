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
	fmt.Println("Starting the Golander Web Scraper...")

	// load the application configuration
	conf, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	var db *database.Database
	// Retry database connection
	for i := 0; i < 10; i++ {
		// init the db
		db, err = database.Init(conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Password, conf.Database.DBName)
		if err == nil {
			break
		}
		log.Println("Could not connect to database, retrying...")
		time.Sleep(time.Second * 5)
	}

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	} else {
		fmt.Println("Successfully connected to the database")
	}

	defer db.Close()

	err = db.CreateRepoTable()
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	fmt.Println("Tables created!")

	apiServer := api.NewServer(db)
	go apiServer.Start(fmt.Sprintf(":%d", conf.API.Port)) // Start API server in a separate goroutine

	// Create a new scheduler with a 24-hour interval
	scraper.New(4*time.Second, scraper.ScrapeWithDB(db)) // change interval to 24 hours or your preferred timing

	// Wait for a termination signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	fmt.Println("Shutting down...")
}
