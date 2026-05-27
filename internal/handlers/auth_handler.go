package handlers

import (
	"encoding/json"
	"net/http"

	"walkara/internal/repository/sqlite"
	"walkara/internal/services"
	"walkara/internal/utils"

	"github.com/google/uuid"
)

type AuthHandler struct {
	repo *sqlite.UserRepository
	auth *services.AuthService
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

	if req.Email == "" || req.Password == "" {
		utils.JSON(w, http.StatusBadRequest, false, "email and password required", nil, "missing fields")
		return
	}

	hash, err := h.auth.HashPassword(req.Password)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "failed to hash password", nil, err.Error())
		return
	}

	id := uuid.NewString()

	err = h.repo.CreateUser(id, req.Email, hash)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "user creation failed", nil, err.Error())
		return
	}

	utils.JSON(w, http.StatusOK, true, "user registered successfully", map[string]string{
		"user_id": id,
	}, "")
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	_ = json.NewDecoder(r.Body).Decode(&req)

	if req.Email == "" || req.Password == "" {
		utils.JSON(w, http.StatusBadRequest, false, "email and password required", nil, "missing fields")
		return
	}

	id, hash, err := h.repo.GetUserByEmail(req.Email)
	if err != nil {
		utils.JSON(w, http.StatusUnauthorized, false, "invalid credentials", nil, "user not found")
		return
	}

	if !h.auth.CheckPassword(hash, req.Password) {
		utils.JSON(w, http.StatusUnauthorized, false, "invalid credentials", nil, "wrong password")
		return
	}

	token, err := h.auth.GenerateToken(id)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "failed to generate token", nil, err.Error())
		return
	}

	utils.JSON(w, http.StatusOK, true, "login successful", map[string]string{
		"token": token,
	}, "")
}