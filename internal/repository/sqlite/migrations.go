package sqlite

import (
	"database/sql"
	"log"
)

func RunMigrations(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS walks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id TEXT,
		steps INTEGER,
		distance REAL,
		calories REAL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("failed to run migrations:", err)
	}

	log.Println("Database migrations completed")
}