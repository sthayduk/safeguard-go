package main

import (
	"log/slog"
	"os"

	"github.com/sthayduk/safeguard-go"
)

func main() {
	// Option 1: Use the safeguard package's logger configuration with debug level
	debugLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	safeguard.SetLogger(debugLogger)

	// Option 2: Use a custom JSON logger
	// jsonLogger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
	// safeguard.SetLogger(jsonLogger)

	// Option 3: Get the configured logger and use it in your package
	logger := safeguard.GetLogger()
	logger.Info("Using safeguard logger in external package")
	logger.Debug("This is a debug message")

	// Now create a safeguard client - it will use the pre-configured logger
	client := safeguard.NewClient("https://your-appliance.com", "v4", true)

	// Your application logic here...
	_ = client
}
