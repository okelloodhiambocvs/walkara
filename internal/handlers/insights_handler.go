package handlers

import (
	"net/http"

	"walkara/internal/repository/sqlite"
	"walkara/internal/services"
	"walkara/internal/utils"
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
	userID := r.URL.Query().Get("user_id")

	if userID == "" {
		utils.JSON(w, http.StatusBadRequest, false, "missing user_id", nil, "user_id is required")
		return
	}

	data, err := h.repo.GetWeeklyComparison(userID)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "failed to generate insights", nil, err.Error())
		return
	}

	steps := data["this_week_steps"].(int)
	improvement := data["improvement_pct"].(float64)

	score := h.service.CalculateScore(steps, improvement)
	message := h.service.GenerateMessage(score, improvement, steps)

	response := map[string]interface{}{
		"user_id":         userID,
		"score":           score,
		"message":         message,
		"this_week":       data,
		"insight_version": "v1",
	}

	utils.JSON(w, http.StatusOK, true, "weekly insights generated", response, "")
}