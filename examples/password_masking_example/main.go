package main

import (
	"log/slog"
	"os"

	"github.com/sthayduk/safeguard-go"
)

func main() {
	// Set up debug logging to see the password masking in action
	debugLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	safeguard.SetLogger(debugLogger)

	// Create a safeguard client
	client := safeguard.NewClient("https://your-appliance.com", "v4", true)

	// Example of what would happen when checking out a password
	// This simulates the response you saw in your log
	passwordResponse := []byte(`"46j64TDDJ8wo6N45igf4"`)
	passwordPath := "AccessRequests/138-450-76-9800-1-87030a13e37a4cda845395b15ce202bb-12079/CheckOutPassword"

	// Simulate logging the response (this is what happens internally in PostRequest)
	safeResponse := safeguard.NewSafeResponseBody(passwordResponse, passwordPath)
	debugLogger.Info("Password checkout response example", "responseBody", safeResponse)

	// Example of normal response that won't be masked
	normalResponse := []byte(`{"success": true, "message": "Request processed"}`)
	normalPath := "Users/123"

	safeNormalResponse := safeguard.NewSafeResponseBody(normalResponse, normalPath)
	debugLogger.Info("Normal response example", "responseBody", safeNormalResponse)

	_ = client // Use the client variable to avoid unused variable warning
}
