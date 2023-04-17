package handlers

import (
	"log"
	"net/http"
)

// Handler for paths that not exists
var NoRoute = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Println("404 page not found")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 page not found"))
})
