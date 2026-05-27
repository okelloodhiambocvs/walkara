package handlers

import (
	"encoding/json"
	"net/http"

	"walkara/internal/repository/sqlite"
)

type LeaderboardHandler struct {
	repo *sqlite.HistoryRepository
}

func NewLeaderboardHandler(repo *sqlite.HistoryRepository) *LeaderboardHandler {
	return &LeaderboardHandler{repo: repo}
}

func (h *LeaderboardHandler) GetWeeklyLeaderboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := h.repo.GetWeeklyLeaderboard()
	if err != nil {
		http.Error(w, "failed to fetch leaderboard", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"period":      "last_7_days",
		"leaderboard": data,
	})
}