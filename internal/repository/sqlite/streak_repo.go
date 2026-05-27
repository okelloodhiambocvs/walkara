package sqlite

import (
	"database/sql"
	"time"
)

type StreakRepository struct {
	DB *sql.DB
}

func NewStreakRepository(db *sql.DB) *StreakRepository {
	return &StreakRepository{DB: db}
}

type Streak struct {
	UserID         string `json:"user_id"`
	CurrentStreak  int    `json:"current_streak"`
	LongestStreak  int    `json:"longest_streak"`
	LastActiveDate string `json:"last_active_date"`
}

func (r *StreakRepository) UpdateStreak(userID string) error {
	today := time.Now().Format("2006-01-02")

	var lastDate string
	var current, longest int

	err := r.DB.QueryRow(`
		SELECT last_active_date, current_streak, longest_streak
		FROM user_streaks
		WHERE user_id = ?
	`, userID).Scan(&lastDate, &current, &longest)

	if err != nil {
		if err == sql.ErrNoRows {
			_, err = r.DB.Exec(`
				INSERT INTO user_streaks (user_id, current_streak, longest_streak, last_active_date)
				VALUES (?, 1, 1, ?)
			`, userID, today)
			return err
		}
		return err
	}

	// If already updated today → do nothing
	if lastDate == today {
		return nil
	}

	// Check if yesterday (continue streak)
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	if lastDate == yesterday {
		current++
	} else {
		current = 1
	}

	if current > longest {
		longest = current
	}

	_, err = r.DB.Exec(`
		UPDATE user_streaks
		SET current_streak = ?, longest_streak = ?, last_active_date = ?
		WHERE user_id = ?
	`, current, longest, today, userID)

	return err
}

func (r *StreakRepository) GetStreak(userID string) (*Streak, error) {
	var s Streak

	err := r.DB.QueryRow(`
		SELECT user_id, current_streak, longest_streak, last_active_date
		FROM user_streaks
		WHERE user_id = ?
	`, userID).Scan(&s.UserID, &s.CurrentStreak, &s.LongestStreak, &s.LastActiveDate)

	if err != nil {
		return nil, err
	}

	return &s, nil
}