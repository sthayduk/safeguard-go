package safeguard

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"testing"
	"time"
)

func setupTestClient() *SafeguardClient {
	return &SafeguardClient{
		AccessToken: &RSTSAuthResponse{
			AccessToken: "initial-token",
			AuthTime:    time.Now(),
			ExpiresIn:   3600,
		},
		Appliance: applianceURL{
			Url: "https://appliance.example.com:443",
		},
		ClusterLeader: applianceURL{
			Url: "https://leader.example.com:443",
		},
		HttpClient: &http.Client{},
	}
}

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
			client := setupTestClient()
			client.Appliance.setUrl(tt.applianceURL, -1)
			client.setClusterLeader(tt.clusterLeaderHost)

			if client.ClusterLeader.getUrl() != tt.expectedClusterURL {
				t.Errorf("expected %s, got %s", tt.expectedClusterURL, client.ClusterLeader.getUrl())
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
			client := setupTestClient()
			client.AccessToken.AuthTime = tt.authTime
			client.AccessToken.ExpiresIn = tt.expiresIn

			expiryTime := client.GetTokenExpirationTime()

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
			client := setupTestClient()
			client.AccessToken.AuthTime = tt.authTime
			client.AccessToken.ExpiresIn = tt.expiresIn

			result := client.IsTokenExpired()

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
			client := setupTestClient()
			client.AccessToken.AuthTime = tt.authTime
			client.AccessToken.ExpiresIn = tt.expiresIn

			result := client.RemainingTokenTime()

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

func TestConcurrentAccess(t *testing.T) {
	const (
		numReaders    = 100
		numWriters    = 10
		numOperations = 1000
	)

	client := setupTestClient()
	var wg sync.WaitGroup
	errorChan := make(chan error, numReaders+numWriters)

	// Start readers
	for i := 0; i < numReaders; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := 0; j < numOperations; j++ {
				url := client.ClusterLeader.getUrl()
				if url == "" {
					errorChan <- fmt.Errorf("reader %d: empty URL at operation %d", id, j)
					return
				}
			}
		}(i)
	}

	// Start writers
	for i := 0; i < numWriters; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := 0; j < numOperations; j++ {
				newLeader := fmt.Sprintf("leader%d-%d", id, j)
				client.ClusterLeader.setUrl(fmt.Sprintf("https://%s:443", newLeader), 10*time.Minute)
			}
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	close(errorChan)

	// Check for any errors
	var errors []error
	for err := range errorChan {
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		t.Errorf("Found %d concurrent access errors:", len(errors))
		for _, err := range errors {
			t.Error(err)
		}
	}
}

func TestAccessTokenConcurrency(t *testing.T) {
	const (
		numReaders    = 50
		numWriters    = 50
		numOperations = 500
	)

	client := &SafeguardClient{
		AccessToken: &RSTSAuthResponse{
			AccessToken: "initial-token",
			AuthTime:    time.Now(),
			ExpiresIn:   3600,
		},
	}

	errorChan := make(chan error, numReaders+numWriters)
	doneChan := make(chan bool, numReaders+numWriters)

	// Start readers
	for i := 0; i < numReaders; i++ {
		go func(id int) {
			defer func() { doneChan <- true }()

			for j := 0; j < numOperations; j++ {

				token := client.AccessToken.getAccessToken()

				if token == "" {
					errorChan <- fmt.Errorf("reader %d: invalid token at operation %d", id, j)
					return
				}
			}
		}(i)
	}

	// Start writers
	for i := 0; i < numWriters; i++ {
		go func(id int) {
			defer func() { doneChan <- true }()

			for j := 0; j < numOperations; j++ {

				newToken := fmt.Sprintf("token-%d-%d", id, j)
				client.AccessToken.setAccessToken(newToken)

			}
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < numReaders+numWriters; i++ {
		<-doneChan
	}

	// Check for any errors
	close(errorChan)
	var errors []error
	for err := range errorChan {
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		t.Errorf("Found %d concurrent access errors:", len(errors))
		for _, err := range errors {
			t.Error(err)
		}
	}
}

func TestCredentialsConcurrency(t *testing.T) {
	const (
		numReaders    = 50
		numWriters    = 10
		numOperations = 1000
	)

	client := &SafeguardClient{
		AccessToken: &RSTSAuthResponse{
			credentials: Credentials{
				username:     "initial-user",
				password:     "initial-pass",
				certPath:     "initial-cert",
				certPassword: "initial-cert-pass",
			},
		},
	}

	errorChan := make(chan error, numReaders+numWriters)
	doneChan := make(chan bool, numReaders+numWriters)

	// Test concurrent reads of credentials
	for i := 0; i < numReaders; i++ {
		go func(id int) {
			defer func() { doneChan <- true }()

			for j := 0; j < numOperations; j++ {

				username, password := client.AccessToken.getUserNamePassword()
				certPath, certPass := client.AccessToken.getCertificate()

				if username == "" || password == "" || certPath == "" || certPass == "" {
					errorChan <- fmt.Errorf("reader %d: empty credentials at operation %d", id, j)
					return
				}
			}
		}(i)
	}

	// Test concurrent writes to credentials
	for i := 0; i < numWriters; i++ {
		go func(id int) {
			defer func() { doneChan <- true }()

			for j := 0; j < numOperations; j++ {

				username := fmt.Sprintf("user-%d-%d", id, j)
				password := fmt.Sprintf("pass-%d-%d", id, j)
				client.AccessToken.setUserNamePassword(username, password)

				certPath := fmt.Sprintf("cert-%d-%d", id, j)
				certPass := fmt.Sprintf("cert-pass-%d-%d", id, j)
				client.AccessToken.setCertificate(certPath, certPass)

			}
		}(i)
	}

	// Wait for completion
	for i := 0; i < numReaders+numWriters; i++ {
		<-doneChan
	}

	// Check for errors
	close(errorChan)
	var errors []error
	for err := range errorChan {
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		t.Errorf("Found %d credential concurrency errors:", len(errors))
		for _, err := range errors {
			t.Error(err)
		}
	}
}

func TestMutexDeadlockPrevention(t *testing.T) {
	client := setupTestClient()

	done := make(chan bool, 2)
	const iterations = 1000

	// First goroutine: simplified token operations
	go func() {
		defer func() { done <- true }()

		for i := 0; i < iterations; i++ {

			token := client.AccessToken.getAccessToken()
			if token == "" {
				t.Error("Empty token in first goroutine")

				return
			}
			client.AccessToken.setAccessToken(fmt.Sprintf("token-1-%d", i))
			client.ClusterLeader.setUrl(fmt.Sprintf("https://leader-1-%d.example.com:443", i), 10*time.Minute)
		}
	}()

	// Second goroutine: simplified operations in reverse order
	go func() {
		defer func() { done <- true }()

		for i := 0; i < iterations; i++ {

			url := client.ClusterLeader.getUrl()
			if url == "" {
				t.Error("Empty URL in second goroutine")

				return
			}
			token := client.AccessToken.getAccessToken()
			if token == "" {
				t.Error("Empty token in second goroutine")

				return
			}
			client.AccessToken.setAccessToken(fmt.Sprintf("token-2-%d", i))
		}
	}()

	// Wait with timeout for both goroutines
	for i := 0; i < 2; i++ {
		select {
		case <-done:
			continue
		case <-time.After(5 * time.Second):
			t.Fatal("Deadlock detected - test timed out")
		}
	}
}

func TestRWMutexStress(t *testing.T) {
	const (
		numGoroutines = 1000
		numOperations = 1000
	)

	client := &SafeguardClient{
		AccessToken: &RSTSAuthResponse{
			AccessToken: "test-token",
		},
		ClusterLeader: applianceURL{
			Url: "https://leader.example.com",
		},
	}

	var wg sync.WaitGroup
	errors := make(chan error, numGoroutines*2)

	// Launch many goroutines that randomly read or write
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := 0; j < numOperations; j++ {
				if rand.Float32() < 0.2 { // 20% chance of writing
					// Write operation

					client.AccessToken.setAccessToken(fmt.Sprintf("token-%d-%d", id, j))
					client.ClusterLeader.setUrl(fmt.Sprintf("https://leader-%d-%d.example.com", id, j), 10*time.Minute)

				} else {
					// Read operation

					token := client.AccessToken.getAccessToken()
					url := client.ClusterLeader.getUrl()

					if token == "" || url == "" {
						errors <- fmt.Errorf("goroutine %d: empty data at operation %d", id, j)
					}
				}
			}
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	close(errors)

	// Check for any errors
	var errs []error
	for err := range errors {
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		t.Errorf("Found %d stress test errors:", len(errs))
		for _, err := range errs {
			t.Error(err)
		}
	}
}
