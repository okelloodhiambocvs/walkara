package handlers

import (
	"encoding/json"
	"net/http"
	"walkara/internal/services"
)

type WalkHandler struct {
	service *services.WalkService
}

func NewWalkHandler() *WalkHandler {
	return &WalkHandler{
		service: services.NewWalkService(),
	}
}

func (h *WalkHandler) CalculateWalk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req struct {
		Steps int `json:"steps"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	distance := h.service.StepsToKM(req.Steps)
	calories := h.service.EstimateCalories(req.Steps)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"steps":    req.Steps,
		"distance": distance,
		"calories": calories,
		"message":  "Walkara activity processed",
	})
}