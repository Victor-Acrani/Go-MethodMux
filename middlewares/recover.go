package middlewares

import (
	"log"
	"net/http"
)

// Middleware for recovering.
func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing recover middleware")
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				http.Error(w, "Something went wrong!", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
		log.Print("Executing recover middleware again")
	})
}
