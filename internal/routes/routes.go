package routes

import (
	"miniredis/internal/handlers"

	"github.com/gorilla/mux"
)

func NewRouter(h *handlers.Handler) *mux.Router {
	r := mux.NewRouter()

	// Home + health
	r.HandleFunc("/", h.HomeHandler).Methods("GET")
	r.HandleFunc("/ping", h.HealthHandler).Methods("GET")

	// KV operations (single key)
	r.HandleFunc("/kv{key}", h.GetKeyHandler).Methods("GET")
	r.HandleFunc("/kv/{key}", h.PutKeyHandler).Methods("PUT")
	r.HandleFunc("/kv/{key}", h.DeleteKeyHandler).Methods("DELETE")

	// KV operations (whole store)
	r.HandleFunc("/kv", h.ListKeysHandler).Methods("GET")
	r.HandleFunc("/kv", h.FlushHandler).Methods("DELETE")

	return r
}
