package routes

import (
	"miniredis/internal/handlers"
	"net/http"
	"os"

	ghandler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func NewRouter(h *handlers.Handler) http.Handler {
	r := mux.NewRouter()

	// Home + health
	r.HandleFunc("/", h.HomeHandler).Methods("GET")
	r.HandleFunc("/ping", h.HealthHandler).Methods("GET")

	// KV operations (single key)
	r.HandleFunc("/kv/{key}", h.GetKeyHandler).Methods("GET")
	r.HandleFunc("/kv/{key}", h.PutKeyHandler).Methods("PUT")
	r.HandleFunc("/kv/{key}", h.DeleteKeyHandler).Methods("DELETE")

	// KV operations (whole store)
	r.HandleFunc("/kv", h.ListKeysHandler).Methods("GET")
	r.HandleFunc("/kv", h.FlushHandler).Methods("DELETE")

	middlewareChain := alice.New(
		func(next http.Handler) http.Handler {
			return ghandler.RecoveryHandler()(next) // panic-safe
		},
		ghandler.CompressHandler, // gzip responses
		func(next http.Handler) http.Handler {
			return ghandler.LoggingHandler(os.Stdout, next) // log requests
		},
	)

	return middlewareChain.Then(r)

}
