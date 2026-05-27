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

	// repositories
	walkRepo := sqlite.NewWalkRepository(db)
	streakRepo := sqlite.NewStreakRepository(db)
	historyRepo := sqlite.NewHistoryRepository(db)
	userRepo := sqlite.NewUserRepository(db)

	mux := http.NewServeMux()

	// health
	mux.HandleFunc("/health", handlers.HealthCheck)

	// auth
	authHandler := handlers.NewAuthHandler(userRepo)
	mux.HandleFunc("/auth/register", authHandler.Register)
	mux.HandleFunc("/auth/login", authHandler.Login)

	// walk system
	walkHandler := handlers.NewWalkHandler(walkRepo, streakRepo)
	mux.HandleFunc("/walk/calculate", walkHandler.CalculateWalk)

	historyHandler := handlers.NewHistoryHandler(historyRepo)
	mux.HandleFunc("/walk/history", historyHandler.GetHistory)

	summaryHandler := handlers.NewSummaryHandler(historyRepo)
	mux.HandleFunc("/walk/summary/weekly", summaryHandler.GetWeeklySummary)

	streakHandler := handlers.NewStreakHandler(streakRepo)
	mux.HandleFunc("/walk/streak", streakHandler.GetStreak)

	insightsHandler := handlers.NewInsightsHandler(historyRepo)
	mux.HandleFunc("/walk/insights/weekly", insightsHandler.GetWeeklyInsights)

	leaderboardHandler := handlers.NewLeaderboardHandler(historyRepo)
	mux.HandleFunc("/walk/leaderboard/weekly", leaderboardHandler.GetWeeklyLeaderboard)

	log.Println("Walkara API running on :8080")

	log.Fatal(http.ListenAndServe(":8080", mux))
}