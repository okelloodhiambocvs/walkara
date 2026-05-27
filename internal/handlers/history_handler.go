package handlers

import (
	"encoding/json"
	"net/http"

	"walkara/internal/repository/sqlite"
)

type HistoryHandler struct {
	repo *sqlite.HistoryRepository
}

func NewHistoryHandler(repo *sqlite.HistoryRepository) *HistoryHandler {
	return &HistoryHandler{repo: repo}
}

func (h *HistoryHandler) GetHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := r.URL.Query().Get("user_id")

	if userID == "" {
		http.Error(w, "missing user_id", http.StatusBadRequest)
		return
	}

	history, err := h.repo.GetWalkHistory(userID)
	if err != nil {
		http.Error(w, "failed to fetch history", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": userID,
		"data":    history,
	})
}