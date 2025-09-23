package main

import (
	"flag"
	"fmt"
	"log/slog"
	"miniredis/internal/core"
	"miniredis/internal/deps"
	"miniredis/internal/handlers"
	"miniredis/internal/routes"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := flag.String("addr", ":8000", "Address to open the miniredis server on")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	deps := &deps.AppDependencies{
		Logger: logger,
		Store:  core.NewStore(),
	}

	appHandlers := handlers.NewHandler(deps)

	srv := &http.Server{
		Addr:        *addr,
		Handler:     routes.NewRouter(appHandlers),
		ErrorLog:    slog.NewLogLogger(logger.Handler(), slog.LevelError),
		IdleTimeout: time.Minute,
	}

	fmt.Printf("Starting server at %v\n", *addr)

	if err := srv.ListenAndServe(); err != nil {
		logger.Error("Server failed", "error", err)
		os.Exit(1)
	}
}
