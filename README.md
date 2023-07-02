# Golander Web Scraper

Golander is a web scraper for GitHub repositories, written in Go. It scrapes repositories from GitHub, specifically Go language repositories, and stores the data in a MySQL database. It also exposes an API to interact with the scraped data.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

## Prerequisites

- Go (version 1.17 or later recommended)
- MySQL (version 8.0 or later recommended)
- An appropriate Go IDE (like GoLand, VS Code etc.)
- Git 

## Installing

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
```bash
go run .
```

## Tests
These tests cover the main functions and components of the project, including the API, database operations, configuration loading, models, scraper's scheduler, and the main application behavior.

### Running the tests
To run all tests in the project, navigate to the project root directory and use the following command: <br />
```bash
go test ./...
```

This command will recursively run all tests in the project. <br />

To run a specific test file, navigate to the file's directory and use the following command: <br />
```bash
go test <file_name>
```
Replace <filename> with the name of the test file you want to run. <br />

`api_test.go `<br />
This file tests the pingHandler function of the API. It checks if the API returns a pong response when pinged.

`config_test.go` <br />
This file tests the Load function of the config package. It checks if the environment variables are loaded correctly into the configuration.

`db_test.go `<br />
This file tests the Init function of the database package and the UpsertRepo and GetAllRepos functions. Mocking is used to simulate the MySQL database and interactions with it.

`models_test.go` <br />
This file tests the UpsertRepo and GetAllRepos functions in the database package. It checks if the functions correctly upsert and retrieve repositories from the database.

`scheduler_test.go`<br />
This file tests the NewTask and NewScheduler functions in the scraper package. It checks if tasks are correctly scheduled and whether they can be started and stopped as expected.

`golander_test.go` <br />
This file tests the main application functions, including database connection and configuration loading during application startup. Two main scenarios are covered: successful startup and startup with a configuration loading error.

Built With <br />
`Go` - The programming language used <br />
`Gin` - The web framework used <br />
`GoColly` - The web scraping framework used <br />
`MySQL` - The database used <br />


License <br />
This project is licensed under the MIT License - see the LICENSE.md file for details