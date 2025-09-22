package handlers

import (
	"miniredis/internal/deps"
	"net/http"
)

type Handler struct {
	deps *deps.AppDependencies
}

func NewHandler(deps *deps.AppDependencies) *Handler {
	return &Handler{
		deps: deps,
	}
}

// GET /
func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

// GET /ping
func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

// GET /kv/{key}
func (h *Handler) GetKeyHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

// PUT /kv/{key}
func (h *Handler) PutKeyHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

// DELETE /kv/{key}
func (h *Handler) DeleteKeyHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func (h *Handler) ListKeysHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func (h *Handler) FlushHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}
