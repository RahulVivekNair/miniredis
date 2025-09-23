package handlers

import (
	"encoding/json"
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

type KVRequest struct {
	Value string `json:"value"`
}

// basic KV response
type KVResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// message-style response used for create/update/delete
type MessageResponse struct {
	Message string `json:"message"`
	Key     string `json:"key,omitempty"`
	Value   string `json:"value,omitempty"`
}

type AllKeysResponse struct {
	Count int      `json:"count"`
	Keys  []string `json:"keys,omitempty"`
}

// GET /
func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{"message": "Welcome to miniredis"})
}

// GET /ping
func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{"status": "OK"})
}

// GET /kv/{key}
func (h *Handler) GetKeyHandler(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	value, ok := h.deps.Store.Data.Load(key)
	if !ok {
		respondJSONError(w, http.StatusNotFound, "key not found")
		return
	}
	payload := &KVResponse{
		Key:   key,
		Value: value.(string),
	}
	respondJSON(w, http.StatusOK, payload)

}

// PUT /kv/{key}
func (h *Handler) PutKeyHandler(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]

	var body KVRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		respondJSONError(w, http.StatusNotFound, "invalid JSON")
		return
	}

	// Store/overwrite value
	_, loaded := h.deps.Store.Data.Swap(key, body.Value)
	payload := &MessageResponse{
		Key:   key,
		Value: body.Value,
	}
	if loaded {
		payload.Message = "updated"
	} else {
		payload.Message = "created"
	}
	respondJSON(w, http.StatusCreated, payload)
}

// DELETE /kv/{key}
func (h *Handler) DeleteKeyHandler(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	value, loaded := h.deps.Store.Data.LoadAndDelete(key)
	if !loaded {
		respondJSONError(w, http.StatusNotFound, "key not found")
		return
	}

	payload := &MessageResponse{
		Key:     key,
		Value:   value.(string),
		Message: "deleted",
	}
	respondJSON(w, http.StatusOK, payload)

}

func (h *Handler) ListKeysHandler(w http.ResponseWriter, r *http.Request) {
	var keys []string

	h.deps.Store.Data.Range(func(key, _ any) bool {
		keys = append(keys, key.(string))
		return true
	})

	payload := &AllKeysResponse{
		Count: len(keys),
		Keys:  keys,
	}

	respondJSON(w, http.StatusOK, payload)
}

func (h *Handler) FlushHandler(w http.ResponseWriter, r *http.Request) {
	h.deps.Store.Data.Clear()
	respondJSON(w, http.StatusOK, map[string]bool{"success": true})
}
