package handlers

import (
	"net/http"

	"walkara/internal/repository/sqlite"
	"walkara/internal/utils"
)

type SummaryHandler struct {
	repo *sqlite.HistoryRepository
}

func NewSummaryHandler(repo *sqlite.HistoryRepository) *SummaryHandler {
	return &SummaryHandler{repo: repo}
}

func (h *SummaryHandler) GetWeeklySummary(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")

	if userID == "" {
		utils.JSON(w, http.StatusBadRequest, false, "missing user_id", nil, "user_id is required")
		return
	}

	summary, err := h.repo.GetWeeklySummary(userID)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "failed to fetch summary", nil, err.Error())
		return
	}

	utils.JSON(w, http.StatusOK, true, "weekly summary fetched", summary, "")
}