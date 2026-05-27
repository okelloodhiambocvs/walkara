package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"walkara/internal/services"
)

type WalkHandler struct {
	service *services.WalkService
	db      *sql.DB
}

func NewWalkHandler(db *sql.DB) *WalkHandler {
	return &WalkHandler{
		service: services.NewWalkService(),
		db:      db,
	}
}