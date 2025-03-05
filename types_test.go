package safeguard

import (
	"testing"
	"time"
)

func TestApplianceURLIsExpired(t *testing.T) {
	tests := []struct {
		name       string
		lastUpdate time.Time
		cacheTime  time.Duration
		expected   bool
	}{
		{
			name:       "Not expired",
			lastUpdate: time.Now().Add(-5 * time.Minute),
			cacheTime:  10 * time.Minute,
			expected:   false,
		},
		{
			name:       "Expired",
			lastUpdate: time.Now().Add(-15 * time.Minute),
			cacheTime:  10 * time.Minute,
			expected:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &applianceURL{
				lastUpdate: tt.lastUpdate,
				cacheTime:  tt.cacheTime,
			}
			if got := a.isExpired(); got != tt.expected {
				t.Errorf("isExpired() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestApplianceURLGetSetUrl(t *testing.T) {
	tests := []struct {
		name      string
		url       string
		cacheTime time.Duration
	}{
		{
			name:      "Basic URL",
			url:       "https://example.com",
			cacheTime: 10 * time.Minute,
		},
		{
			name:      "URL with port",
			url:       "https://example.com:8443",
			cacheTime: 5 * time.Minute,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &applianceURL{}
			a.setUrl(tt.url, tt.cacheTime)

			if got := a.getUrl(); got != tt.url {
				t.Errorf("getUrl() = %v, want %v", got, tt.url)
			}
		})
	}
}

func TestRSTSAuthResponseTokenMethods(t *testing.T) {
	tests := []struct {
		name      string
		token     string
		username  string
		password  string
		certPath  string
		certPass  string
		userToken string
	}{
		{
			name:      "Basic credentials",
			token:     "test-token",
			username:  "testuser",
			password:  "testpass",
			certPath:  "/path/to/cert",
			certPass:  "certpass",
			userToken: "user-token",
		},
		{
			name:      "Empty credentials",
			token:     "",
			username:  "",
			password:  "",
			certPath:  "",
			certPass:  "",
			userToken: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RSTSAuthResponse{}

			// Test access token
			r.setAccessToken(tt.token)
			if got := r.getAccessToken(); got != tt.token {
				t.Errorf("getAccessToken() = %v, want %v", got, tt.token)
			}

			// Test username/password
			r.setUserNamePassword(tt.username, tt.password)
			username, password := r.getUserNamePassword()
			if username != tt.username || password != tt.password {
				t.Errorf("getUserNamePassword() = (%v, %v), want (%v, %v)",
					username, password, tt.username, tt.password)
			}

			// Test certificate
			r.setCertificate(tt.certPath, tt.certPass)
			certPath, certPass := r.getCertificate()
			if certPath != tt.certPath || certPass != tt.certPass {
				t.Errorf("getCertificate() = (%v, %v), want (%v, %v)",
					certPath, certPass, tt.certPath, tt.certPass)
			}

			// Test user token
			r.setUserToken(tt.userToken)
			if got := r.UserToken; got != tt.userToken {
				t.Errorf("UserToken = %v, want %v", got, tt.userToken)
			}
		})
	}
}

func TestAuthProviderString(t *testing.T) {
	tests := []struct {
		name     string
		provider AuthProvider
		want     string
	}{
		{
			name:     "Certificate provider",
			provider: AuthProviderCertificate,
			want:     "rsts:sts:primaryproviderid:certificate",
		},
		{
			name:     "Local provider",
			provider: AuthProviderLocal,
			want:     "rsts:sts:primaryproviderid:local",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.provider.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
