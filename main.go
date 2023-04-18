package main

import (
	"log"
	"net/http"

	"github.com/Victor-Acrani/Go-MethodMux/handlers"
	"github.com/Victor-Acrani/Go-MethodMux/middlewares"
)

// port gives the http port
const port = ":8080"

func main() {
	// create net/http mux
	mux := http.NewServeMux()

	//set chain of middlewares (the order of insertion changes the behavior of the chain)
	chain := []func(http.Handler) http.Handler{
		middlewares.RequestID,
		middlewares.Logger,
		middlewares.Recover,
	}

	// create custom method mux
	methodmux := handlers.NewMethodMux()
	methodmux.AddMiddleware(chain)
	methodmux.Add(http.MethodGet, handlers.DefaultGetHandler)
	methodmux.Add(http.MethodPost, handlers.DefaultPostHandler)
	methodmux.Add(http.MethodPut, handlers.DefaultPostHandler)
	methodmux.Add(http.MethodDelete, handlers.DefaultPostHandler)
	mux.Handle("/handlers", methodmux)

	// create handlers to simulate a panic
	methodmuxPanic := handlers.NewMethodMux()
	methodmuxPanic.AddMiddleware(chain)
	methodmuxPanic.Add(http.MethodGet, handlers.PanicGetHandler)
	mux.Handle("/panic", methodmuxPanic)

	/*
		define a default handler for non defined paths
		***This is necessary because the string "/" in method mux.Handle("/", methodmuxNotFound)
		is a regular expression that matches any path.***
	*/
	methodmuxNotFound := handlers.NewMethodMux()
	methodmuxNotFound.AddMiddleware(chain)
	methodmuxNotFound.Add(http.MethodGet, handlers.NoRoute)
	mux.Handle("/", methodmuxNotFound)

	// listen and serve default server
	log.Printf("Listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
