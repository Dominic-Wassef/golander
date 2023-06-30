package scraper

import (
	"fmt"
	"time"
)

// New starts a new scheduler that runs the scraping at regular intervals
func New(scrapingInterval time.Duration, scrapeFunc func()) {
	ticker := time.NewTicker(scrapingInterval)

	go func() {
		for range ticker.C {
			// Call web scraping function here
			fmt.Println("Starting the scraping task...")
			scrapeFunc()
		}
	}()
}
