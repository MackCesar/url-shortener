package models

// URLRequest represents a request to shorten a URL
type URLRequest struct {
	OriginalURL string `json:"original_url"`
}

// URLResponse represents a response containing the shortened URL
type URLResponse struct {
	ShortenedURL string `json:"shortened_url"`
}
