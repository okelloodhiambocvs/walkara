package main

import (
	"log"
	"net/http"
	"walkara/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthCheck)

	walkHandler := handlers.NewWalkHandler()
	mux.HandleFunc("/walk/calculate", walkHandler.CalculateWalk)

	log.Println("Walkara API running on :8080")

	log.Fatal(http.ListenAndServe(":8080", mux))
}