package domain

type Walk struct {
	ID         string  `json:"id"`
	UserID     string  `json:"user_id"`
	DistanceKM float64 `json:"distance_km"`
	Steps      int     `json:"steps"`
	Calories   float64 `json:"calories"`
	Pace       float64 `json:"pace"` // km/h
	Date       string  `json:"date"`
}
