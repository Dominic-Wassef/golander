package scraper

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dominic-wassef/golander/src/database"
	"github.com/gocolly/colly"
)

type Repository struct {
	Name        string
	Url         string
	Description string
	Language    string
	Stars       string
}

// Scrape function to perform the actual web scraping
func Scrape(db *database.Database, pageStart int, pageEnd int, perPage int, retryCount int, delayBetweenRetries time.Duration) {
	fmt.Println("Scraping task initiated...")

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("github.com"),
	)

	c.OnHTML(".repo-list-item", func(e *colly.HTMLElement) {
		name := e.ChildText("h3 a")
		url := e.Request.AbsoluteURL(e.ChildAttr("h3 a", "href"))
		description := e.ChildText(".repo-list-description")
		language := e.ChildText(".repo-language-color + span")
		stars := strings.TrimSpace(e.ChildText(".starring-container .social-count"))

		repo := &database.Repository{
			Name:        name,
			Url:         url,
			Description: sql.NullString{String: description, Valid: description != ""},
			Language:    sql.NullString{String: language, Valid: language != ""},
			Stars:       sql.NullString{String: stars, Valid: stars != ""},
		}

		err := db.UpsertRepo(repo)
		if err != nil {
			log.Fatalf("Failed to insert repo into database: %v", err)
		}

		fmt.Printf("Repo found and inserted into database: %+v\n", repo)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	for i := pageStart; i <= pageEnd; i++ {
		for j := 0; j < retryCount; j++ {
			err := c.Visit(fmt.Sprintf("https://github.com/search?l=Go&p=%d&q=stars%%3A%%3E0&s=stars&type=Repositories", i))
			if err != nil {
				fmt.Println("Error visiting page, retrying after delay:", err)
				time.Sleep(delayBetweenRetries)
				continue
			}
			break
		}
	}

	fmt.Println("Scraping task completed.")
}

func ScrapeWithDB(db *database.Database, pageStart int, pageEnd int, perPage int, retryCount int, delayBetweenRetries time.Duration) func() {
	return func() {
		Scrape(db, pageStart, pageEnd, perPage, retryCount, delayBetweenRetries)
	}
}
