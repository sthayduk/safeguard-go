package main

import (
	"log/slog"
	"os"

	"github.com/sthayduk/safeguard-go"
)

func main() {
	// Example 1: Simple debug logging setup
	debugLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	safeguard.SetLogger(debugLogger)
	logger1 := safeguard.GetLogger()
	logger1.Info("Example 1: Debug logging enabled")

	// Example 2: JSON logging to stderr
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelWarn}))
	safeguard.SetLogger(jsonLogger)
	logger2 := safeguard.GetLogger()
	logger2.Warn("Example 2: JSON logging to stderr")

	// Example 3: Custom structured logging with context
	baseLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true, // Add source file and line information
	}))
	contextLogger := baseLogger.With("service", "my-app", "version", "1.0.0")
	safeguard.SetLogger(contextLogger)
	logger3 := safeguard.GetLogger()
	logger3.Debug("Example 3: Custom structured logging with context")

	// Example 4: Using the default logger
	safeguard.SetLogger(slog.Default())
	logger4 := safeguard.GetLogger()
	logger4.Info("Example 4: Using default slog logger")

	// Create safeguard client - it will use the configured logger
	client := safeguard.NewClient("https://your-appliance.com", "v4", true)
	_ = client
}
