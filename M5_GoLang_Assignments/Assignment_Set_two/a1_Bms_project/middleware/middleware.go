package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs each incoming request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Method: %s, URL: %s, Time: %s", r.Method, r.URL.Path, start)
		next.ServeHTTP(w, r)
		log.Printf("Completed in %v", time.Since(start))
	})
}
