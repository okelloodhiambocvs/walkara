package main

import (
	"log"
	"net/http"

	"walkara/internal/handlers"
	"walkara/internal/repository/sqlite"
)

func main() {
	db := sqlite.InitDB()
	sqlite.RunMigrations(db)

	walkRepo := sqlite.NewWalkRepository(db)
	historyRepo := sqlite.NewHistoryRepository(db)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthCheck)

	walkHandler := handlers.NewWalkHandler(walkRepo)
	mux.HandleFunc("/walk/calculate", walkHandler.CalculateWalk)

	historyHandler := handlers.NewHistoryHandler(historyRepo)
	mux.HandleFunc("/walk/history", historyHandler.GetHistory)

	summaryHandler := handlers.NewSummaryHandler(historyRepo)
	mux.HandleFunc("/walk/summary/weekly", summaryHandler.GetWeeklySummary)

	log.Println("Walkara API running on :8080")

	log.Fatal(http.ListenAndServe(":8080", mux))
}