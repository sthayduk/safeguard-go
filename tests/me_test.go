package tests

import (
	"testing"
	"time"

	"github.com/sthayduk/safeguard-go/models"
)

func TestGetMaximumReleaseDuration(t *testing.T) {
	tests := []struct {
		name     string
		props    models.RequesterProperties
		expected time.Duration
	}{
		{
			name: "Zero duration",
			props: models.RequesterProperties{
				MaximumReleaseDurationDays:    0,
				MaximumReleaseDurationHours:   0,
				MaximumReleaseDurationMinutes: 0,
			},
			expected: 0,
		},
		{
			name: "Only days",
			props: models.RequesterProperties{
				MaximumReleaseDurationDays:    2,
				MaximumReleaseDurationHours:   0,
				MaximumReleaseDurationMinutes: 0,
			},
			expected: 48 * time.Hour,
		},
		{
			name: "Only hours",
			props: models.RequesterProperties{
				MaximumReleaseDurationDays:    0,
				MaximumReleaseDurationHours:   5,
				MaximumReleaseDurationMinutes: 0,
			},
			expected: 5 * time.Hour,
		},
		{
			name: "Only minutes",
			props: models.RequesterProperties{
				MaximumReleaseDurationDays:    0,
				MaximumReleaseDurationHours:   0,
				MaximumReleaseDurationMinutes: 30,
			},
			expected: 30 * time.Minute,
		},
		{
			name: "Days, hours, and minutes",
			props: models.RequesterProperties{
				MaximumReleaseDurationDays:    1,
				MaximumReleaseDurationHours:   2,
				MaximumReleaseDurationMinutes: 30,
			},
			expected: 26*time.Hour + 30*time.Minute,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.props.GetMaximumReleaseDuration(); got != tt.expected {
				t.Errorf("GetMaximumReleaseDuration() = %v, want %v", got, tt.expected)
			}
		})
	}
}
func TestGetDefaultReleaseDuration(t *testing.T) {
	tests := []struct {
		name     string
		props    models.RequesterProperties
		expected time.Duration
	}{
		{
			name: "Zero duration",
			props: models.RequesterProperties{
				DefaultReleaseDurationDays:    0,
				DefaultReleaseDurationHours:   0,
				DefaultReleaseDurationMinutes: 0,
			},
			expected: 0,
		},
		{
			name: "Only days",
			props: models.RequesterProperties{
				DefaultReleaseDurationDays:    2,
				DefaultReleaseDurationHours:   0,
				DefaultReleaseDurationMinutes: 0,
			},
			expected: 48 * time.Hour,
		},
		{
			name: "Only hours",
			props: models.RequesterProperties{
				DefaultReleaseDurationDays:    0,
				DefaultReleaseDurationHours:   5,
				DefaultReleaseDurationMinutes: 0,
			},
			expected: 5 * time.Hour,
		},
		{
			name: "Only minutes",
			props: models.RequesterProperties{
				DefaultReleaseDurationDays:    0,
				DefaultReleaseDurationHours:   0,
				DefaultReleaseDurationMinutes: 30,
			},
			expected: 30 * time.Minute,
		},
		{
			name: "Days, hours, and minutes",
			props: models.RequesterProperties{
				DefaultReleaseDurationDays:    1,
				DefaultReleaseDurationHours:   2,
				DefaultReleaseDurationMinutes: 30,
			},
			expected: 26*time.Hour + 30*time.Minute,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.props.GetDefaultReleaseDuration(); got != tt.expected {
				t.Errorf("GetDefaultReleaseDuration() = %v, want %v", got, tt.expected)
			}
		})
	}
}
