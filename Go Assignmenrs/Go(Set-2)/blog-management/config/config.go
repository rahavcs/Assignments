package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite3", "./blog.db")
	if err != nil {
		log.Fatal("Could not connect to database: ", err)
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS blogs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		content TEXT,
		author TEXT,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatal("Could not create table: ", err)
	}
}
