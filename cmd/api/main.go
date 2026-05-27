package main

import (
	"log"
	"net/http"

	"walkara/internal/handlers"
	"walkara/internal/repository/sqlite"
)

func main() {
	// DB init
	db := sqlite.InitDB()
	sqlite.RunMigrations(db)

	// Repositories
	walkRepo := sqlite.NewWalkRepository(db)
	historyRepo := sqlite.NewHistoryRepository(db)

	// Handlers
	mux := http.NewServeMux()

	// Health endpoint
	mux.HandleFunc("/health", handlers.HealthCheck)

	// Walk calculation (save activity)
	walkHandler := handlers.NewWalkHandler(walkRepo)
	mux.HandleFunc("/walk/calculate", walkHandler.CalculateWalk)

	// Walk history (Strava-like feed)
	historyHandler := handlers.NewHistoryHandler(historyRepo)
	mux.HandleFunc("/walk/history", historyHandler.GetHistory)

	log.Println("Walkara API running on :8080")

	log.Fatal(http.ListenAndServe(":8080", mux))
}