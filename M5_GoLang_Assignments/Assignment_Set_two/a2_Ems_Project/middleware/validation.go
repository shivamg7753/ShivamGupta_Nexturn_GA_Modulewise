package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// ValidationMiddleware validates that the request body contains required fields
func ValidationMiddleware(requiredFields []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Ensure request body is not empty
			if r.ContentLength == 0 {
				http.Error(w, `{"error": "Request body cannot be empty"}`, http.StatusBadRequest)
				return
			}

			// Parse the JSON payload
			var payload map[string]interface{}
			bodyBytes, _ := io.ReadAll(r.Body) // Read the original body
			r.Body.Close()
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Reset r.Body for downstream handlers

			if err := json.Unmarshal(bodyBytes, &payload); err != nil {
				http.Error(w, `{"error": "Invalid JSON payload: `+err.Error()+`"}`, http.StatusBadRequest)
				return
			}

			// Check for required fields
			for _, field := range requiredFields {
				if _, exists := payload[field]; !exists {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(map[string]string{"error": "Missing required field: " + field})
					return
				}
			}

			// Reset request body for the next handler
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			next.ServeHTTP(w, r)
		})
	}
}
