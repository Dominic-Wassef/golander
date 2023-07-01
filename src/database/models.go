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

func (db *Database) GetAllRepos() ([]*Repository, error) {
	rows, err := db.DB.Query("SELECT * FROM repositories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	repos := make([]*Repository, 0)
	for rows.Next() {
		repo := new(Repository)
		err := rows.Scan(&repo.Id, &repo.Name, &repo.Url, &repo.Description, &repo.Language, &repo.Stars)
		if err != nil {
			return nil, err
		}
		repos = append(repos, repo)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return repos, nil
}

func (db *Database) GetRepoByID(id int64) (*Repository, error) {
	repo := new(Repository)
	err := db.DB.QueryRow("SELECT * FROM repositories WHERE id=?", id).Scan(&repo.Id, &repo.Name, &repo.Url, &repo.Description, &repo.Language, &repo.Stars)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (db *Database) DeleteRepoByID(id int64) error {
	_, err := db.DB.Exec("DELETE FROM repositories WHERE id=?", id)
	return err
}
