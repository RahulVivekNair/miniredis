package handlers

import (
	"fmt"
	"miniredis/internal/deps"
	"net/http"

	"github.com/gorilla/mux"
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
	fmt.Fprintln(w, "Home")
}

// GET /ping
func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

// GET /kv/{key}
func (h *Handler) GetKeyHandler(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	fmt.Fprintf(w, "You got key: %v", key)
}

// PUT /kv/{key}
func (h *Handler) PutKeyHandler(w http.ResponseWriter, r *http.Request) {

}

// DELETE /kv/{key}
func (h *Handler) DeleteKeyHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func (h *Handler) ListKeysHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You got all")
}

func (h *Handler) FlushHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}
