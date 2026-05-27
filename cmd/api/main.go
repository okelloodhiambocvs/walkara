package main

import (
	"log"
	"net/http"
	"walkara/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	healthHandler := handlers.HealthCheck
	walkHandler := handlers.NewWalkHandler()

	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/walk/calculate", walkHandler.CalculateWalk)

	log.Println("Walkara API running on :8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}