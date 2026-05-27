package sqlite

import (
	"database/sql"
)

type HistoryRepository struct {
	DB *sql.DB
}

func NewHistoryRepository(db *sql.DB) *HistoryRepository {
	return &HistoryRepository{DB: db}
}

type WalkRecord struct {
	ID        int     `json:"id"`
	UserID    string  `json:"user_id"`
	Steps     int     `json:"steps"`
	Distance  float64 `json:"distance"`
	Calories  float64 `json:"calories"`
	CreatedAt string  `json:"created_at"`
}

func (r *HistoryRepository) GetWalkHistory(userID string) ([]WalkRecord, error) {
	query := `
	SELECT id, user_id, steps, distance, calories, created_at
	FROM walks
	WHERE user_id = ?
	ORDER BY created_at DESC
	`

	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []WalkRecord

	for rows.Next() {
		var w WalkRecord
		err := rows.Scan(
			&w.ID,
			&w.UserID,
			&w.Steps,
			&w.Distance,
			&w.Calories,
			&w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		history = append(history, w)
	}

	return history, nil
}