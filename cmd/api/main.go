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

	// Repository
	walkRepo := sqlite.NewWalkRepository(db)

	// Handlers
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthCheck)

	walkHandler := handlers.NewWalkHandler(walkRepo)
	mux.HandleFunc("/walk/calculate", walkHandler.CalculateWalk)

	log.Println("Walkara API running on :8080")

	log.Fatal(http.ListenAndServe(":8080", mux))
}