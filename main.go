package main

import (
	"fmt"
	"net/http"

	"github.com/rednexx46/cipher-api/handlers"
)

func main() {
	// Set up routes for the API
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/encrypt", handlers.EncryptHandler)
	http.HandleFunc("/decrypt", handlers.DecryptHandler)

	fmt.Println("🔐 Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
