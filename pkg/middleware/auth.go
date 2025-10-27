package middleware

import (
	"log"
	"net/http"
)

func APIKeyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey != "super-secret-key" {
			log.Printf("Authentication failed: Invalid API Key '%s'", apiKey)
			http.Error(w, "Forbidden: Invalid API Key", http.StatusForbidden)
			return
		}
		log.Println("Authentication Successful")
		next.ServeHTTP(w, r)
	})
}
