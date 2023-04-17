package handlers

import (
	"log"
	"net/http"
)

// --- Dummy handlers to test the API ---

// GET handler
var DefaultGetHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Println("Running GET handler")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running GET handler"))
})

// POST handler
var DefaultPostHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Println("Running POST handler")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Running POST handler"))
})

// PUT handler
var DefaultPutHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Println("Running PUT handler")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running PUT handler"))
})

// DELETE handler
var DefaultDeleteHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Println("Running DELETE handler")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running DELETE handler"))
})
