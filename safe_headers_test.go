package safeguard

import (
	"log/slog"
	"net/http"
	"os"
	"testing"
)

func TestSafeHeaders_LogValue(t *testing.T) {
	// Create a test logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	// Create test headers with sensitive authorization token
	headers := http.Header{}
	headers.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
	headers.Set("Content-Type", "application/json")
	headers.Set("Accept", "application/json")
	headers.Set("User-Agent", "SafeguardClient/1.0")

	// Test logging with SafeHeaders
	logger.Info("Request headers", "headers", NewSafeHeaders(headers))

	// The output should show:
	// - Authorization header with masked token (Bearer eyJhbG***ssw5c)
	// - Other headers unchanged
}

func TestSafeHeaders_ShortToken(t *testing.T) {
	// Test with a short token
	headers := http.Header{}
	headers.Set("Authorization", "Bearer short")
	headers.Set("Content-Type", "application/json")

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	logger.Info("Short token test", "headers", NewSafeHeaders(headers))
	// Should show "Bearer ***" for short tokens
}

func TestSafeHeaders_NonBearerAuth(t *testing.T) {
	// Test with non-Bearer authorization
	headers := http.Header{}
	headers.Set("Authorization", "Basic dXNlcjpwYXNzd29yZA==")
	headers.Set("Content-Type", "application/json")

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	logger.Info("Basic auth test", "headers", NewSafeHeaders(headers))
	// Should show "***" for non-Bearer auth
}
