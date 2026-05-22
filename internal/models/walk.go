package models

type WalkData struct {
	UserID     string  `json:"user_id"`
	DistanceKM  float64 `json:"distance_km"`
	Steps       int     `json:"steps"`
	Calories    float64 `json:"calories"`
	Date        string  `json:"date"`
}