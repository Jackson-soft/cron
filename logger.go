package cron

import (
	"io"
	"log/slog"
)

var (
	// DefaultLogger is used by Cron if none is specified.
	DefaultLogger Logger = slog.Default()

	// DiscardLogger can be used by callers to discard all log messages.
	DiscardLogger Logger = slog.New(slog.NewTextHandler(io.Discard, nil))
)

// Logger is the interface used in this package for logging, so that any backend can be plugged in.
type Logger interface {
	Error(msg string, args ...any)
	Info(msg string, args ...any)
}
