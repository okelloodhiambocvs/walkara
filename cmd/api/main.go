package main

import (
	"log"
	"net/http"

	"walkara/config"
	"walkara/internal/handlers"
	"walkara/internal/repository/sqlite"
	"walkara/internal/services"
)

func main() {
	cfg := config.LoadConfig()

	// production JWT secret injection
	services.JWTSecret = []byte(cfg.JWTSecret)

	// database init
	db := sqlite.InitDB(cfg.DBPath)
	sqlite.RunMigrations(db)

	// repositories
	walkRepo := sqlite.NewWalkRepository(db)
	streakRepo := sqlite.NewStreakRepository(db)
	historyRepo := sqlite.NewHistoryRepository(db)
	userRepo := sqlite.NewUserRepository(db)

	// handlers
	authHandler := handlers.NewAuthHandler(userRepo)
	walkHandler := handlers.NewWalkHandler(walkRepo, streakRepo)
	historyHandler := handlers.NewHistoryHandler(historyRepo)
	summaryHandler := handlers.NewSummaryHandler(historyRepo)
	streakHandler := handlers.NewStreakHandler(streakRepo)
	insightsHandler := handlers.NewInsightsHandler(historyRepo)
	leaderboardHandler := handlers.NewLeaderboardHandler(historyRepo)

	mux := http.NewServeMux()

	// health
	mux.HandleFunc("/health", handlers.HealthCheck)

	// auth
	mux.HandleFunc("/auth/register", authHandler.Register)
	mux.HandleFunc("/auth/login", authHandler.Login)

	// walk system
	mux.HandleFunc("/walk/calculate", walkHandler.CalculateWalk)
	mux.HandleFunc("/walk/history", historyHandler.GetHistory)
	mux.HandleFunc("/walk/summary/weekly", summaryHandler.GetWeeklySummary)
	mux.HandleFunc("/walk/streak", streakHandler.GetStreak)
	mux.HandleFunc("/walk/insights/weekly", insightsHandler.GetWeeklyInsights)
	mux.HandleFunc("/walk/leaderboard/weekly", leaderboardHandler.GetWeeklyLeaderboard)

	log.Println("Walkara running on port:", cfg.Port)

	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}