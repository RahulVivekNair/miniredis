package deps

import (
	"log/slog"
	"miniredis/internal/core"
)

type AppDependencies struct {
	Logger *slog.Logger
	Store  *core.Store
}
