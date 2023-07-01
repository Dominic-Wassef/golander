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

1. Clone the repository to your local machine <br />

```bash
git clone https://github.com/dominic-wassef/golander.git
```

2. Navigate to the project directory <br />
```bash
cd golander
```

3. Install the necessary dependencies <br />
```bash
go mod tidy
```

4. Set up your database and provide the necessary environment variables for your application to connect.  <br />

The required variables are:  <br />
`DB_HOST`
`DB_PORT`
`DB_USER`
`DB_PASSWORD`
`DB_NAME`
`API_PORT`
`SCRAPER_URL`

Running the application <br />
To run the application, use the go run command: <br />
`go run .`

Built With <br />
`Go` - The programming language used <br />
`Gin` - The web framework used <br />
`GoColly` - The web scraping framework used <br />
`MySQL` - The database used <br />


License <br />
This project is licensed under the MIT License - see the LICENSE.md file for details