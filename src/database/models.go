package database

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	Id          int64
	Name        string
	Url         string
	Description sql.NullString
	Language    sql.NullString
	Stars       sql.NullString
}

// UpsertRepo updates or inserts a new repository entry in the database.
func (db *Database) UpsertRepo(r *Repository) error {
	fmt.Printf("Upserting repository: %+v\n", r)

	query := `
	INSERT INTO repositories (name, url, description, language, stars)
	VALUES (?, ?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE description = VALUES(description), language = VALUES(language), stars = VALUES(stars)`

	res, err := db.DB.Exec(query, r.Name, r.Url, r.Description, r.Language, r.Stars)
	if err != nil {
		return fmt.Errorf("failed to upsert into database: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get inserted id: %w", err)
	}

	fmt.Printf("Upserted repository with ID: %d\n", id)

	r.Id = id
	return nil
}
