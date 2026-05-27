package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./walkara.db")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	createTable(db)

	log.Println("SQLite database connected")

	return db
}

func createTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS walks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id TEXT,
		steps INTEGER,
		distance_km REAL,
		calories REAL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Failed to create walks table:", err)
	}
}