package services

import "math"

type WalkService struct{}

func NewWalkService() *WalkService {
	return &WalkService{}
}

// Convert steps to km (basic MVP formula)
func (s *WalkService) StepsToKM(steps int) float64 {
	return math.Round(float64(steps)*0.0008*100) / 100
}

// Calories estimation (basic average formula)
func (s *WalkService) EstimateCalories(steps int) float64 {
	return math.Round(float64(steps)*0.04*100) / 100
}