# Golander Web Scraper

Golander is a web scraper for GitHub repositories, written in Go. It scrapes repositories from GitHub, specifically Go language repositories, and stores the data in a MySQL database. It also exposes an API to interact with the scraped data.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.17 or later recommended)
- MySQL (version 8.0 or later recommended)
- An appropriate Go IDE (like GoLand, VS Code etc.)
- Git 

### Installing

1. Clone the repository to your local machine

```bash
git clone https://github.com/dominic-wassef/golander.git
```

2. Navigate to the project directory
`cd golander`

3. Install the necessary dependencies
`go mod tidy`

4. Set up your database and provide the necessary environment variables for your application to connect to your database. The required variables are:

`DB_HOST`
`DB_PORT`
`DB_USER`
`DB_PASSWORD`
`DB_NAME`
`API_PORT`
`SCRAPER_URL`

Running the application
To run the application, use the go run command:
`go run .`

Built With
Go - The programming language used
Gin - The web framework used
GoColly - The web scraping framework used
MySQL - The database used


License
This project is licensed under the MIT License - see the LICENSE.md file for details