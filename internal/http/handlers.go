package http

import (
	"encoding/json"
	"net/http"
	"services-catalog/internal/service"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct{ svc *service.Svc }

func NewHandler(s *service.Svc) *Handler { return &Handler{svc: s} }

func (h *Handler) ListServices(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := service.ListOpts{
		Query:    q.Get("query"),
		SortBy:   q.Get("sort"),
		Order:    q.Get("order"),
		Page:     atoiOr(q.Get("page"), 1),
		PageSize: atoiOr(q.Get("page_size"), 20),
	}
	items, total, err := h.svc.List(r.Context(), opts)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	resp := ListResponse[service.ServiceDTO]{Data: items}
	resp.Meta.Page = opts.Page
	resp.Meta.PageSize = opts.PageSize
	resp.Meta.Total = total
	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) GetService(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := service.ParseUUID(idStr)
	if err != nil {
		writeErr(w, http.StatusBadRequest, err.Error())
		return
	}
	it, err := h.svc.Get(r.Context(), id)
	if err != nil {
		writeErr(w, http.StatusNotFound, "service not found")
		return
	}
	writeJSON(w, http.StatusOK, it)
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func writeErr(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, code, map[string]string{"error": msg})
}

func atoiOr(s string, d int) int {
	if v, err := strconv.Atoi(s); err == nil && v > 0 {
		return v
	}
	return d
}
