package main

import (
	"log"
	"net/http"

	"walkara/internal/handlers"
	"walkara/internal/middleware"
	"walkara/internal/repository/sqlite"
)

func main() {
	db := sqlite.InitDB()
	sqlite.RunMigrations(db)

	walkRepo := sqlite.NewWalkRepository(db)
	streakRepo := sqlite.NewStreakRepository(db)
	historyRepo := sqlite.NewHistoryRepository(db)
	userRepo := sqlite.NewUserRepository(db)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthCheck)

	authHandler := handlers.NewAuthHandler(userRepo)
	mux.HandleFunc("/auth/register", authHandler.Register)
	mux.HandleFunc("/auth/login", authHandler.Login)

	walkHandler := handlers.NewWalkHandler(walkRepo, streakRepo)

	historyHandler := handlers.NewHistoryHandler(historyRepo)

	summaryHandler := handlers.NewSummaryHandler(historyRepo)

	streakHandler := handlers.NewStreakHandler(streakRepo)

	insightsHandler := handlers.NewInsightsHandler(historyRepo)

	leaderboardHandler := handlers.NewLeaderboardHandler(historyRepo)

	mux.HandleFunc("/walk/calculate", middleware.AuthMiddleware(walkHandler.CalculateWalk))
	mux.HandleFunc("/walk/history", middleware.AuthMiddleware(historyHandler.GetHistory))
	mux.HandleFunc("/walk/summary/weekly", middleware.AuthMiddleware(summaryHandler.GetWeeklySummary))
	mux.HandleFunc("/walk/streak", middleware.AuthMiddleware(streakHandler.GetStreak))
	mux.HandleFunc("/walk/insights/weekly", middleware.AuthMiddleware(insightsHandler.GetWeeklyInsights))
	mux.HandleFunc("/walk/leaderboard/weekly", middleware.AuthMiddleware(leaderboardHandler.GetWeeklyLeaderboard))

	log.Println("Walkara API running on :8080")

	log.Fatal(http.ListenAndServe(":8080", mux))
}