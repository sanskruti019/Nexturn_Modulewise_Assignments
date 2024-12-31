package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Initialize the SQLite database
func InitDatabase() error {
	var err error
	db, err = sql.Open("sqlite3", "inventory.db")
	if err != nil {
		return err
	}

	// Create products table
	query := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		price REAL NOT NULL,
		stock INTEGER NOT NULL,
		category_id INTEGER
	);
	`
	_, err = db.Exec(query)
	return err
}
