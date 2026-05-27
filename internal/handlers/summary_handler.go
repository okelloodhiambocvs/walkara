package handlers

import (
	"encoding/json"
	"net/http"

	"walkara/internal/repository/sqlite"
)

type SummaryHandler struct {
	repo *sqlite.HistoryRepository
}

func NewSummaryHandler(repo *sqlite.HistoryRepository) *SummaryHandler {
	return &SummaryHandler{repo: repo}
}

func (h *SummaryHandler) GetWeeklySummary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := r.URL.Query().Get("user_id")

	if userID == "" {
		http.Error(w, "missing user_id", http.StatusBadRequest)
		return
	}

	summary, err := h.repo.GetWeeklySummary(userID)
	if err != nil {
		http.Error(w, "failed to fetch summary", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(summary)
}