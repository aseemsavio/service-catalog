package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

// Common returns a slice of common middleware functions for HTTP handlers
func Common() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		middleware.RequestID,
		middleware.RealIP,
		middleware.Recoverer,
		middleware.Timeout(60 * time.Second),
	}
}
