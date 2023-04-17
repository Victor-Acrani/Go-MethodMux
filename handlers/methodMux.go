package handlers

import (
	"net/http"
)

// methodMux structure
type methodMux struct {
	handlers    map[string]http.Handler
	middlewares []func(http.Handler) http.Handler
}

// Add a http.Handlers by http method.
func (m *methodMux) Add(pattern string, handler http.Handler) {
	m.handlers[pattern] = handler
}

// Add middleware to methodMux.
func (m *methodMux) AddMiddleware(middlewares []func(http.Handler) http.Handler) {
	m.middlewares = append(m.middlewares, middlewares...)
}

// Handler interface signature.
func (m *methodMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, ok := m.handlers[r.Method]
	if !ok {
		handler = NoRoute
	}
	m.ChainMiddlewares(m.middlewares)(handler).ServeHTTP(w, r)
}

// Call the middleware chain to run before the handler.
func (m *methodMux) ChainMiddlewares(middlewares []func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			handler = middlewares[i](handler)
		}
		return handler
	}
}

// Create a new instance of methodMux.
func NewMethodMux() *methodMux {
	var methodmux methodMux
	methodmux.handlers = make(map[string]http.Handler)
	return &methodmux
}
