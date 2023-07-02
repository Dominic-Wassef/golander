package database

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUpsertRepo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an errr '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	database := &Database{
		DB: db,
	}

	repo := &Repository{
		Name:        "test-repo",
		Url:         "https://github.com/trending",
		Description: sql.NullString{String: "Test repo for testing", Valid: true},
		Language:    sql.NullString{String: "Go", Valid: true},
		Stars:       sql.NullString{String: "100", Valid: true},
	}

	mock.ExpectExec("INSERT INTO repositories").WithArgs(repo.Name, repo.Url, repo.Description, repo.Language, repo.Stars).WillReturnResult(sqlmock.NewResult(1, 1))

	err = database.UpsertRepo(repo)
	if err != nil {
		t.Errorf("error was not expected while upserting repo: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there werre unfulfilled expectations: %s", err)
	}
}

func TestGetAllRepos(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	database := &Database{
		DB: db,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "url", "description", "language", "stars"}).
		AddRow(1, "test-repo", "http://github.com/trending", "Test repo for testing", "Go", "100")

	mock.ExpectQuery("SELECT \\* FROM repositories").WillReturnRows(rows)

	repos, err := database.GetAllRepos()
	if err != nil {
		t.Errorf("error was not expected while getting repos: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Check that we got the correct data
	assert.Equal(t, 1, len(repos), "expected 1 repo, got %d", len(repos))
	assert.Equal(t, "test-repo", repos[0].Name)
}
