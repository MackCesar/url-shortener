package db

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitDB initializes the database
func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./urlshortener.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	createTableQuery := `
    CREATE TABLE IF NOT EXISTS url_mapping (
        short_id TEXT PRIMARY KEY,
        original_url TEXT
    );`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

// SaveURLMapping stores the mapping between short ID and original URL
func SaveURLMapping(shortID, originalURL string) error {
	_, err := db.Exec("INSERT INTO url_mapping(short_id, original_url) VALUES (?, ?)", shortID, originalURL)
	return err
}

// GetOriginalURL retrieves the original URL from the database using the short ID
func GetOriginalURL(shortID string) (string, error) {
	var originalURL string
	err := db.QueryRow("SELECT original_url FROM url_mapping WHERE short_id = ?", shortID).Scan(&originalURL)
	if err == sql.ErrNoRows {
		return "", errors.New("no matching URL found")
	}
	return originalURL, err
}
