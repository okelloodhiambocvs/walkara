package services

import (
	"math"
	"walkara/internal/repository/sqlite"
)

type WalkService struct {
	walkRepo   *sqlite.WalkRepository
	streakRepo *sqlite.StreakRepository
}

func NewWalkService(w *sqlite.WalkRepository, s *sqlite.StreakRepository) *WalkService {
	return &WalkService{
		walkRepo:   w,
		streakRepo: s,
	}
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

	// 1. Save walk (critical path)
	err := s.walkRepo.SaveWalk(userID, steps, distance, calories)
	if err != nil {
		return 0, 0, err
	}

	// 2. Update streak (non-critical path)
	// We intentionally do NOT fail walk saving if streak fails
	if s.streakRepo != nil {
		_ = s.streakRepo.UpdateStreak(userID)
	}

	return distance, calories, nil
}