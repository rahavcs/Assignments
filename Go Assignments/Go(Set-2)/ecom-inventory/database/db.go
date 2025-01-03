package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	// Open the database
	DB, err = sql.Open("sqlite3", "./inventory.db")
	if err != nil {
		log.Fatal("Could not connect to the database: ", err)
	}

	// Create the products table if it doesn't exist
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		description TEXT,
		price REAL,
		stock INTEGER,
		category_id INTEGER
	)`)
	if err != nil {
		log.Fatal("Could not create table: ", err)
	}
}
