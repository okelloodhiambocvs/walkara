package sqlite

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitDB(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	return db
}