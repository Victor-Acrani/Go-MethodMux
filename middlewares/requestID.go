package middlewares

import (
	"log"
	"net/http"
)

// Middleware to set a requestID.
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing request id middleware")
		next.ServeHTTP(w, r)
		log.Print("Executing request id middleware again")
	})
}
