package utils

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

// GenerateShortID creates a random short ID
func GenerateShortID() string {
	b := make([]byte, 6) //6 bytes for a short ID
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
