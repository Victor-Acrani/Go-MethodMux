package middlewares

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
)

// key type for passing value trough context
type key string

const (
	KeyRequestID key = "request_id"
)

// Middleware to set a requestID.
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing request id middleware")

		// create uuid for request
		uuid, err := uuid.NewUUID()
		if err != nil {
			log.Printf("RequestID middleware error: %v", err)
			http.Error(w, "Something went wrong!", http.StatusInternalServerError)
		}

		// pass uuid trough context
		ctx := context.WithValue(r.Context(), KeyRequestID, uuid.String())
		next.ServeHTTP(w, r.WithContext(ctx))
		log.Print("Executing request id middleware again")
	})
}
