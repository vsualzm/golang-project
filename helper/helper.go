package helper

import (
	"log"
	"net/http"
	"time"
)

// ANSI color codes
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the request start with Cyan color
		log.Printf("%sStarted %s %s%s", Cyan, r.Method, r.URL.Path, Reset)

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log the request completion with Green for success
		log.Printf("%sCompleted %s %s in %s%v%s", Green, r.Method, r.URL.Path, Yellow, time.Since(start), Reset)
	})
}
