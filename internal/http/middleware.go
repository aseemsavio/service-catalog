package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func Common() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		middleware.RequestID,
		middleware.RealIP,
		middleware.Recoverer,
		middleware.Timeout(60 * time.Second),
	}
}
