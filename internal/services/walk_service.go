package services

import (
	"math"
	"walkara/internal/repository/sqlite"
)

type WalkService struct {
	repo *sqlite.WalkRepository
}

func NewWalkService(repo *sqlite.WalkRepository) *WalkService {
	return &WalkService{repo: repo}
}

func (s *WalkService) StepsToKM(steps int) float64 {
	return math.Round(float64(steps)*0.0008*100) / 100
}

func (s *WalkService) EstimateCalories(steps int) float64 {
	return math.Round(float64(steps)*0.04*100) / 100
}

func (s *WalkService) SaveWalk(userID string, steps int) (float64, float64, error) {
	distance := s.StepsToKM(steps)
	calories := s.EstimateCalories(steps)

	err := s.repo.SaveWalk(userID, steps, distance, calories)
	return distance, calories, err
}