package handlers

import (
	"encoding/json"
	"net/http"

	"walkara/internal/repository/sqlite"
)

type StreakHandler struct {
	repo *sqlite.StreakRepository
}

func NewStreakHandler(repo *sqlite.StreakRepository) *StreakHandler {
	return &StreakHandler{repo: repo}
}

func (h *StreakHandler) GetStreak(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := r.URL.Query().Get("user_id")

	if userID == "" {
		http.Error(w, "missing user_id", http.StatusBadRequest)
		return
	}

	streak, err := h.repo.GetStreak(userID)
	if err != nil {
		http.Error(w, "streak not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(streak)
}