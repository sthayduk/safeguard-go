package client

import (
	"testing"
	"time"
)

func TestSetClusterLeader(t *testing.T) {
	tests := []struct {
		name                string
		applianceURL        string
		clusterLeaderHost   string
		expectedClusterURL  string
		expectedLogContains string
	}{
		{
			name:                "Only a Hostname is provided",
			applianceURL:        "https://leader",
			clusterLeaderHost:   "leader",
			expectedClusterURL:  "https://leader:443",
			expectedLogContains: "Cluster leader is the same as appliance URL",
		},
		{
			name:                "Cluster leader is the same as appliance URL",
			applianceURL:        "https://leader.example.com",
			clusterLeaderHost:   "leader",
			expectedClusterURL:  "https://leader.example.com:443",
			expectedLogContains: "Cluster leader is the same as appliance URL",
		},
		{
			name:                "Cluster leader is different from appliance URL (Hostname-only)",
			applianceURL:        "https://appliance",
			clusterLeaderHost:   "leader",
			expectedClusterURL:  "https://leader:443",
			expectedLogContains: "Cluster leader set to:",
		},
		{
			name:                "Cluster leader is different from appliance URL (Full domain)",
			applianceURL:        "https://appliance.example.com",
			clusterLeaderHost:   "leader",
			expectedClusterURL:  "https://leader.example.com:443",
			expectedLogContains: "Cluster leader set to:",
		},
		{
			name:                "Cluster leader with subdomain in the middle",
			applianceURL:        "https://appliance.sub.example.com",
			clusterLeaderHost:   "leader",
			expectedClusterURL:  "https://leader.sub.example.com:443",
			expectedLogContains: "Cluster leader set to:",
		},
		{
			name:                "Cluster leader with multiple subdomains in the middle",
			applianceURL:        "https://appliance.sub1.sub2.example.com",
			clusterLeaderHost:   "leader",
			expectedClusterURL:  "https://leader.sub1.sub2.example.com:443",
			expectedLogContains: "Cluster leader set to:",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sgclient = &SafeguardClient{
				ApplicanceURL: tt.applianceURL,
			}

			sgclient.setClusterLeader(tt.clusterLeaderHost)

			if sgclient.ClusterLeaderUrl != tt.expectedClusterURL {
				t.Errorf("expected %s, got %s", tt.expectedClusterURL, sgclient.ClusterLeaderUrl)
			}
		})
	}
}
func TestGetTokenExpirationTime(t *testing.T) {
	tests := []struct {
		name           string
		authTime       time.Time
		expiresIn      int
		expectedExpiry time.Time
	}{
		{
			name:           "Token expires in 1 hour",
			authTime:       time.Now(),
			expiresIn:      3600,
			expectedExpiry: time.Now().Add(1 * time.Hour),
		},
		{
			name:           "Token expires in 30 minutes",
			authTime:       time.Now(),
			expiresIn:      1800,
			expectedExpiry: time.Now().Add(30 * time.Minute),
		},
		{
			name:           "Token expires in 0 seconds",
			authTime:       time.Now(),
			expiresIn:      0,
			expectedExpiry: time.Now(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sgclient = &SafeguardClient{
				AccessToken: &RSTSAuthResponse{
					AuthTime:  tt.authTime,
					ExpiresIn: tt.expiresIn,
				},
			}

			expiryTime := sgclient.GetTokenExpirationTime()

			if !expiryTime.Truncate(time.Second).Equal(tt.expectedExpiry.Truncate(time.Second)) {
				t.Errorf("expected %v, got %v", tt.expectedExpiry.Truncate(time.Second), expiryTime.Truncate(time.Second))
			}
		})
	}
}
func TestIsTokenExpired(t *testing.T) {
	tests := []struct {
		name           string
		authTime       time.Time
		expiresIn      int
		expectedResult bool
	}{
		{
			name:           "Token is expired",
			authTime:       time.Now().Add(-2 * time.Hour),
			expiresIn:      3600,
			expectedResult: true,
		},
		{
			name:           "Token is not expired",
			authTime:       time.Now(),
			expiresIn:      3600,
			expectedResult: false,
		},
		{
			name:           "Token is nil",
			authTime:       time.Time{},
			expiresIn:      0,
			expectedResult: true,
		},
		{
			name:           "AuthTime is zero",
			authTime:       time.Time{},
			expiresIn:      3600,
			expectedResult: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sgclient = &SafeguardClient{
				AccessToken: &RSTSAuthResponse{
					AuthTime:  tt.authTime,
					ExpiresIn: tt.expiresIn,
				},
			}

			result := sgclient.IsTokenExpired()

			if result != tt.expectedResult {
				t.Errorf("expected %v, got %v", tt.expectedResult, result)
			}
		})
	}
}
func TestRemainingTokenTime(t *testing.T) {
	tests := []struct {
		name             string
		authTime         time.Time
		expiresIn        int
		expectedMinRange time.Duration
		expectedMaxRange time.Duration
		expectZeroOrLess bool
	}{
		{
			name:             "Token has 1 hour remaining",
			authTime:         time.Now(),
			expiresIn:        3600,
			expectedMinRange: 59 * time.Minute,
			expectedMaxRange: 60 * time.Minute,
			expectZeroOrLess: false,
		},
		{
			name:             "Token has 30 minutes remaining",
			authTime:         time.Now(),
			expiresIn:        1800,
			expectedMinRange: 29 * time.Minute,
			expectedMaxRange: 30 * time.Minute,
			expectZeroOrLess: false,
		},
		{
			name:             "Token has expired",
			authTime:         time.Now().Add(-2 * time.Hour),
			expiresIn:        3600,
			expectZeroOrLess: true,
		},
		{
			name:             "Token is nil",
			authTime:         time.Time{},
			expiresIn:        0,
			expectZeroOrLess: true,
		},
		{
			name:             "AuthTime is zero",
			authTime:         time.Time{},
			expiresIn:        3600,
			expectZeroOrLess: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sgclient = &SafeguardClient{
				AccessToken: &RSTSAuthResponse{
					AuthTime:  tt.authTime,
					ExpiresIn: tt.expiresIn,
				},
			}

			result := sgclient.RemainingTokenTime()

			if tt.expectZeroOrLess {
				if result > 0 {
					t.Errorf("expected <= 0, got %v", result)
				}
			} else {
				if result < tt.expectedMinRange || result > tt.expectedMaxRange {
					t.Errorf("expected between %v and %v, got %v",
						tt.expectedMinRange,
						tt.expectedMaxRange,
						result)
				}
			}
		})
	}
}
