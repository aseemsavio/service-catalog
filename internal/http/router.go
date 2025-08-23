package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(h *Handler) http.Handler {
	r := chi.NewRouter()
	for _, mw := range Common() {
		r.Use(mw)
	}

	r.Get("/v1/healthz", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200); _, _ = w.Write([]byte("ok")) })

	r.Route("/v1", func(r chi.Router) {
		r.Get("/services", h.ListServices)
		r.Get("/services/{id}", h.GetService)
		r.Get("/services/{id}/versions", h.ListVersions)
	})

	return r
}
