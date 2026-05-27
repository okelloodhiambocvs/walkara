package main

import (
	"log"
	"net/http"

	"walkara/internal/database"
	"walkara/internal/handlers"
)

func main() {
	db := database.InitDB()

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthCheck)

	walkHandler := handlers.NewWalkHandler(db)
	mux.HandleFunc("/walk/calculate", walkHandler.CalculateWalk)

	log.Println("Walkara API running on :8080")

	log.Fatal(http.ListenAndServe(":8080", mux))
}