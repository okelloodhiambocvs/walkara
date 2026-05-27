package sqlite

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite", "walkara.db")
	if err != nil {
		log.Fatal("failed to open database:", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	log.Println("SQLite database connected")

	return db
}