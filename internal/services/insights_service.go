package services

import "math"

type InsightsService struct{}

func NewInsightsService() *InsightsService {
	return &InsightsService{}
}

func (s *InsightsService) CalculateScore(steps int, improvement float64) int {
	score := float64(steps)/100.0 + improvement*0.5

	if score > 100 {
		score = 100
	}
	if score < 0 {
		score = 0
	}

	return int(math.Round(score))
}

func (s *InsightsService) GenerateMessage(score int, improvement float64, steps int) string {
	if steps == 0 {
		return "You were inactive this week. Try starting with a short walk tomorrow."
	}

	if score >= 80 {
		return "Excellent activity! You are in a strong fitness zone this week."
	}

	if score >= 50 {
		return "Good progress. You are improving your activity levels."
	}

	if improvement > 0 {
		return "You are improving, keep going to build consistency."
	}

	return "Try to increase your walking consistency for better health benefits."
}