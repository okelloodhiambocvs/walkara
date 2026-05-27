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

func (r *HistoryRepository) GetWeeklySummary(userID string) (map[string]interface{}, error) {
	query := `
	SELECT 
		COUNT(*) as sessions,
		COALESCE(SUM(steps), 0),
		COALESCE(SUM(distance), 0),
		COALESCE(SUM(calories), 0)
	FROM walks
	WHERE user_id = ?
	AND created_at >= datetime('now', '-7 days')
	`

	var sessions int
	var steps int
	var distance float64
	var calories float64

	err := r.DB.QueryRow(query, userID).Scan(
		&sessions,
		&steps,
		&distance,
		&calories,
	)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"user_id":     userID,
		"sessions":    sessions,
		"total_steps": steps,
		"distance_km": distance,
		"calories":    calories,
		"period":      "last_7_days",
	}, nil
}

func (r *HistoryRepository) GetWeeklyComparison(userID string) (map[string]interface{}, error) {
	query := `
	SELECT 
		COALESCE(SUM(steps), 0),
		COALESCE(SUM(distance), 0),
		COALESCE(SUM(calories), 0)
	FROM walks
	WHERE user_id = ?
	AND created_at >= datetime('now', '-7 days')
	`

	var thisWeekSteps int
	var thisWeekDistance float64
	var thisWeekCalories float64

	err := r.DB.QueryRow(query, userID).Scan(
		&thisWeekSteps,
		&thisWeekDistance,
		&thisWeekCalories,
	)

	if err != nil {
		return nil, err
	}

	// previous week
	queryPrev := `
	SELECT 
		COALESCE(SUM(steps), 0)
	FROM walks
	WHERE user_id = ?
	AND created_at BETWEEN datetime('now', '-14 days') AND datetime('now', '-7 days')
	`

	var prevWeekSteps int

	err = r.DB.QueryRow(queryPrev, userID).Scan(&prevWeekSteps)
	if err != nil {
		return nil, err
	}

	improvement := 0.0
	if prevWeekSteps > 0 {
		improvement = (float64(thisWeekSteps-prevWeekSteps) / float64(prevWeekSteps)) * 100
	}

	return map[string]interface{}{
		"user_id":            userID,
		"this_week_steps":    thisWeekSteps,
		"this_week_distance": thisWeekDistance,
		"this_week_calories": thisWeekCalories,
		"previous_week_steps": prevWeekSteps,
		"improvement_pct":    improvement,
	}, nil
}

func (r *HistoryRepository) GetWeeklyLeaderboard() ([]map[string]interface{}, error) {
	query := `
	SELECT 
		user_id,
		COALESCE(SUM(steps), 0) as total_steps,
		COALESCE(SUM(distance), 0) as total_distance,
		COALESCE(SUM(calories), 0) as total_calories
	FROM walks
	WHERE created_at >= datetime('now', '-7 days')
	GROUP BY user_id
	ORDER BY total_steps DESC
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var leaderboard []map[string]interface{}
	rank := 1

	for rows.Next() {
		var userID string
		var steps int
		var distance float64
		var calories float64

		err := rows.Scan(&userID, &steps, &distance, &calories)
		if err != nil {
			return nil, err
		}

		entry := map[string]interface{}{
			"rank":     rank,
			"user_id":  userID,
			"steps":    steps,
			"distance": distance,
			"calories": calories,
		}

		leaderboard = append(leaderboard, entry)
		rank++
	}

	return leaderboard, nil
}