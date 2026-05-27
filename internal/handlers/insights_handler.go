package handlers

import (
	"encoding/json"
	"net/http"

	"walkara/internal/repository/sqlite"
	"walkara/internal/services"
)

type InsightsHandler struct {
	repo    *sqlite.HistoryRepository
	service *services.InsightsService
}

func NewInsightsHandler(repo *sqlite.HistoryRepository) *InsightsHandler {
	return &InsightsHandler{
		repo:    repo,
		service: services.NewInsightsService(),
	}
}

func (h *InsightsHandler) GetWeeklyInsights(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := r.URL.Query().Get("user_id")

	if userID == "" {
		http.Error(w, "missing user_id", http.StatusBadRequest)
		return
	}

	data, err := h.repo.GetWeeklyComparison(userID)
	if err != nil {
		http.Error(w, "failed to generate insights", http.StatusInternalServerError)
		return
	}

	score := h.service.CalculateScore(
		data["this_week_steps"].(int),
		data["improvement_pct"].(float64),
	)

	message := h.service.GenerateMessage(
		score,
		data["improvement_pct"].(float64),
		data["this_week_steps"].(int),
	)

	response := map[string]interface{}{
		"user_id":         userID,
		"score":           score,
		"message":         message,
		"this_week":       data,
		"insight_version": "v1",
	}

	json.NewEncoder(w).Encode(response)
}