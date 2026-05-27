package sqlite

import (
	"database/sql"
)

type WalkRepository struct {
	DB *sql.DB
}

func NewWalkRepository(db *sql.DB) *WalkRepository {
	return &WalkRepository{DB: db}
}

func (r *WalkRepository) SaveWalk(userID string, steps int, distance, calories float64) error {
	query := `
	INSERT INTO walks (user_id, steps, distance, calories)
	VALUES (?, ?, ?, ?)
	`

	_, err := r.DB.Exec(query, userID, steps, distance, calories)
	return err
}