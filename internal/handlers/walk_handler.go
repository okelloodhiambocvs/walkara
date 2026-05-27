package handlers

import (
	"encoding/json"
	"net/http"

	"walkara/internal/repository/sqlite"
	"walkara/internal/services"
)

type WalkHandler struct {
	service *services.WalkService
}

func NewWalkHandler(wRepo *sqlite.WalkRepository, sRepo *sqlite.StreakRepository) *WalkHandler {
	return &WalkHandler{
		service: services.NewWalkService(wRepo, sRepo),
	}
}

func (h *WalkHandler) CalculateWalk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req struct {
		UserID string `json:"user_id"`
		Steps  int    `json:"steps"`
	}

	_ = json.NewDecoder(r.Body).Decode(&req)

	distance, calories, err := h.service.SaveWalk(req.UserID, req.Steps)
	if err != nil {
		http.Error(w, "failed to save walk", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id":  req.UserID,
		"steps":    req.Steps,
		"distance": distance,
		"calories": calories,
		"message":  "Walk saved successfully",
	})
}