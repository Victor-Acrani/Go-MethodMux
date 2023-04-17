package handlers

import (
	"log"
	"net/http"
)

// GET handler to simulate a system panic
var PanicGetHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Println("Running PANIC handler")
	log.Panic("Panic simulation inside handler")
})
