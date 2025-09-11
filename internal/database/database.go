package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// DB wraps the database connection
type DB struct {
	conn *sql.DB
}

// New creates a new database connection
func New(dbPath string) (*DB, error) {
	conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db := &DB{conn: conn}
	
	// Initialize the database
	if err := db.init(); err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	return db, nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.conn.Close()
}

// GetConnection returns the underlying database connection
func (db *DB) GetConnection() *sql.DB {
	return db.conn
}

// init initializes the database with tables and test data
func (db *DB) init() error {
	// Create articles table
	_, err := db.conn.Exec(`
		CREATE TABLE IF NOT EXISTS articles (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			content TEXT NOT NULL
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	// Insert test data
	_, err = db.conn.Exec(`
		INSERT OR IGNORE INTO articles (id, title, content) VALUES 
		(1, 'The Absence of Errors', 'Initial article content.'),
		(2, 'The Double Fallacy', 'Another crucial piece of the puzzle.')
	`)
	if err != nil {
		return fmt.Errorf("failed to insert test data: %w", err)
	}

	log.Println("✅ Database initialized successfully")
	return nil
}
