package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	DB *sql.DB
}

func Init(host string, port int, user string, password string, dbname string) (*Database, error) {
	portStr := os.Getenv("DB_PORT")
	if portStr == "" {
		portStr = "3306" // default value
	}
	dbPort, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, dbPort, dbname)

	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil
}

func (database *Database) Close() error {
	return database.DB.Close()
}

func (db *Database) CreateRepoTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS repositories (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		url VARCHAR(255) NOT NULL UNIQUE,
		description VARCHAR(1000),
		language VARCHAR(100),
		stars VARCHAR(100)
	)`

	_, err := db.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}
