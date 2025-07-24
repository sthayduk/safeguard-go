package safeguard

import (
	"log/slog"
	"os"
	"testing"
)

func TestSafeResponseBody_LogValue(t *testing.T) {
	// Create a test logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	// Test password response body
	passwordBody := []byte(`"46j64TDDJ8wo6N45igf4"`)
	passwordPath := "AccessRequests/138-450-76-9800-1-87030a13e37a4cda845395b15ce202bb-12079/CheckOutPassword"

	// Test logging with SafeResponseBody for password
	logger.Info("Password checkout response", "responseBody", NewSafeResponseBody(passwordBody, passwordPath))

	// Test normal response body
	normalBody := []byte(`{"id": 123, "name": "test"}`)
	normalPath := "Users/123"

	// Test logging with SafeResponseBody for normal response
	logger.Info("Normal response", "responseBody", NewSafeResponseBody(normalBody, normalPath))

	// The output should show:
	// - Password response with masked content (e.g., "46***f4")
	// - Normal response unchanged
}

func TestSafeResponseBody_ShortPassword(t *testing.T) {
	// Test with a short password
	passwordBody := []byte(`"short"`)
	passwordPath := "AccessRequests/123/CheckOutPassword"

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	logger.Info("Short password test", "responseBody", NewSafeResponseBody(passwordBody, passwordPath))
	// Should show "***" for short passwords
}

func TestSafeResponseBody_NonPasswordResponse(t *testing.T) {
	// Test with quoted string that's not a password (contains spaces)
	normalBody := []byte(`"This is a normal response"`)
	normalPath := "AccessRequests/123/SomeEndpoint"

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	logger.Info("Normal quoted response test", "responseBody", NewSafeResponseBody(normalBody, normalPath))
	// Should show the full content since it contains spaces (not a password)
}

func TestSafeResponseBody_EmptyResponse(t *testing.T) {
	// Test with empty response
	emptyBody := []byte(``)
	passwordPath := "AccessRequests/123/CheckOutPassword"

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	logger.Info("Empty response test", "responseBody", NewSafeResponseBody(emptyBody, passwordPath))
	// Should handle empty response gracefully
}
