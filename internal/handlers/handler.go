package handlers

import (
	"encoding/json"
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

type payload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
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
	value, ok := h.deps.Store.Data.Load(key)
	if !ok {
		respondJSON(w, http.StatusNotFound, map[string]string{"error": "key not found"})
		return
	}
	payload := &payload{
		Key:   key,
		Value: value.(string),
	}
	respondJSON(w, http.StatusOK, payload)

}

// PUT /kv/{key}
func (h *Handler) PutKeyHandler(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]

	var body payload
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
		return
	}

	// Check if the key already exists
	_, existed := h.deps.Store.Data.Load(key)

	// Store/overwrite value
	h.deps.Store.Data.Store(key, body.Value)

	if existed {
		respondJSON(w, http.StatusOK, map[string]string{
			"message": "updated",
			"key":     key,
			"value":   body.Value,
		})
	} else {
		respondJSON(w, http.StatusCreated, map[string]string{
			"message": "created",
			"key":     key,
			"value":   body.Value,
		})
	}
}

// DELETE /kv/{key}
func (h *Handler) DeleteKeyHandler(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	fmt.Fprintf(w, "You delete key: %v", key)
}

func (h *Handler) ListKeysHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You got all")
}

func (h *Handler) FlushHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You flush all")
}
