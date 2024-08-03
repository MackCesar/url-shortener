package main

import (
	"log"
	"net/http"
	"url-shortener/internal/db"
	"url-shortener/internal/handlers"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Set uo HTTP handlers
	http.HandleFunc("/shorten", handlers.ShortenURL)
	http.HandleFunc("/", handlers.RedirectURL)

	// Start the HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
