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

	CREATE TABLE IF NOT EXISTS user_streaks (
		user_id TEXT PRIMARY KEY,
		current_streak INTEGER DEFAULT 0,
		longest_streak INTEGER DEFAULT 0,
		last_active_date DATE
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("failed to run migrations:", err)
	}

	log.Println("Database migrations completed")
}