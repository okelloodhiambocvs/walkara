package handlers

import (
	"encoding/json"
	"net/http"

	"walkara/internal/repository/sqlite"
	"walkara/internal/services"

	"github.com/google/uuid"
)

type AuthHandler struct {
	repo   *sqlite.UserRepository
	auth   *services.AuthService
}

func NewAuthHandler(repo *sqlite.UserRepository) *AuthHandler {
	return &AuthHandler{
		repo: repo,
		auth: services.NewAuthService(),
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	_ = json.NewDecoder(r.Body).Decode(&req)

	hash, _ := h.auth.HashPassword(req.Password)
	id := uuid.NewString()

	err := h.repo.CreateUser(id, req.Email, hash)
	if err != nil {
		http.Error(w, "user creation failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": id,
		"message": "user registered successfully",
	})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	_ = json.NewDecoder(r.Body).Decode(&req)

	id, hash, err := h.repo.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	if !h.auth.CheckPassword(hash, req.Password) {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	token, _ := h.auth.GenerateToken(id)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
	})
}