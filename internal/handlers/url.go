package handlers

import (
	"encoding/json"
	"net/http"
	"url-shortener/internal/db"
	"url-shortener/internal/models"
	"url-shortener/internal/utils"
)

// ShortenURL handles the URL shortening request
func ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req models.URLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	shortID := utils.GenerateShortID()
	err := db.SaveURLMapping(shortID, req.OriginalURL)
	if err != nil {
		http.Error(w, "Failed to shorten URL", http.StatusInternalServerError)
		return
	}

	response := models.URLResponse{
		ShortenedURL: "http://localhost:8080/" + shortID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RedirectURL handles the URL redirection
func RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortID := r.URL.Path[1:]
	originalURL, err := db.GetOriginalURL(shortID)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusFound)
}
