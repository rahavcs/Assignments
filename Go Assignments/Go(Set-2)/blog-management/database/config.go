package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "blogs.db")
	if err != nil {
		log.Fatalf("Could not open database: %v", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS blogs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		author TEXT NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	if _, err := DB.Exec(createTable); err != nil {
		log.Fatalf("Could not create table: %v", err)
	}
	log.Println("Database initialized successfully.")
}
