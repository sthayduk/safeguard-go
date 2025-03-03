package client

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveAccessTokenToEnv(t *testing.T) {
	client := &SafeguardClient{
		AccessToken: &RSTSAuthResponse{
			AccessToken: "test-token",
		},
	}

	err := client.SaveAccessTokenToEnv()
	assert.NoError(t, err)
}

func TestValidateAccessToken(t *testing.T) {
	tests := []struct {
		name      string
		token     string
		mockResp  interface{}
		mockCode  int
		wantError bool
	}{
		{
			name:      "empty token",
			token:     "",
			wantError: true,
		},
		{
			name:      "invalid token format",
			token:     "invalid-token",
			wantError: true,
		},
		{
			name:      "valid token with successful validation",
			token:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
			mockResp:  map[string]interface{}{"id": "123"},
			mockCode:  http.StatusOK,
			wantError: false,
		},
		{
			name:      "valid token format but server rejects",
			token:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
			mockCode:  http.StatusUnauthorized,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Updated path verification
				expectedPath := "/service/core/v4/me"
				if r.URL.Path != expectedPath {
					t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
					w.WriteHeader(http.StatusNotFound)
					return
				}

				// Verify fields parameter if present
				fields := r.URL.Query().Get("fields")
				if fields != "" && fields != "id" {
					t.Errorf("expected fields=id, got fields=%s", fields)
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				// Verify authorization header
				authHeader := r.Header.Get("Authorization")
				if !strings.HasPrefix(authHeader, "Bearer ") {
					t.Errorf("expected Bearer token in Authorization header, got %s", authHeader)
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.mockCode)
				if tt.mockResp != nil {
					json.NewEncoder(w).Encode(tt.mockResp)
				}
			}))
			defer ts.Close()

			client := &SafeguardClient{
				AccessToken: &RSTSAuthResponse{
					UserToken: tt.token,
				},
				Appliance:  applianceURL{},
				HttpClient: http.DefaultClient, // Add this line to fix the nil pointer
				ApiVersion: "v4",               // Add API version
			}
			client.Appliance.setUrl(ts.URL, -1)

			err := client.ValidateAccessToken()
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestExchangeRSTSTokenForSafeguard(t *testing.T) {
	tests := []struct {
		name          string
		inputToken    string
		mockResp      *RSTSAuthResponse
		mockCode      int
		wantError     bool
		expectedToken string
	}{
		{
			name:       "successful exchange",
			inputToken: "input-token",
			mockResp: &RSTSAuthResponse{
				UserToken: "new-user-token",
			},
			mockCode:      http.StatusOK,
			wantError:     false,
			expectedToken: "new-user-token",
		},
		{
			name:       "server error",
			inputToken: "input-token",
			mockCode:   http.StatusInternalServerError,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/service/core/v4/Token/LoginResponse", r.URL.Path)
				assert.Equal(t, "POST", r.Method)
				assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

				// Verify request body
				body, _ := io.ReadAll(r.Body)
				var tokenReq struct {
					StsAccessToken string `json:"StsAccessToken"`
				}
				err := json.Unmarshal(body, &tokenReq)
				assert.NoError(t, err)
				assert.Equal(t, tt.inputToken, tokenReq.StsAccessToken)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.mockCode)
				if tt.mockResp != nil {
					json.NewEncoder(w).Encode(tt.mockResp)
				}
			}))
			defer ts.Close()

			client := &SafeguardClient{
				AccessToken: &RSTSAuthResponse{
					AccessToken: tt.inputToken,
				},
				Appliance:  applianceURL{},
				HttpClient: http.DefaultClient, // Add this line
			}
			client.Appliance.setUrl(ts.URL, -1)

			err := client.exchangeRSTSTokenForSafeguard(http.DefaultClient)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedToken, client.AccessToken.getUserToken())
			}
		})
	}
}

func TestGetAuthorizationHeader(t *testing.T) {
	client := &SafeguardClient{
		AccessToken: &RSTSAuthResponse{
			UserToken: "test-user-token",
		},
	}

	req, _ := http.NewRequest("GET", "http://example.com", nil)
	req.Header = client.getAuthorizationHeader()

	assert.Equal(t, "application/json", req.Header.Get("accept"))
	assert.Equal(t, "Bearer test-user-token", req.Header.Get("Authorization"))
}

func TestHandleTokenResponse(t *testing.T) {
	tests := []struct {
		name      string
		response  *RSTSAuthResponse
		status    int
		wantError bool
	}{
		{
			name: "successful response",
			response: &RSTSAuthResponse{
				UserToken: "test-token",
			},
			status:    http.StatusOK,
			wantError: false,
		},
		{
			name:      "error response",
			status:    http.StatusUnauthorized,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a response with controlled body
			resp := &http.Response{
				StatusCode: tt.status,
				Header:     make(http.Header),
				Body:       io.NopCloser(strings.NewReader("")),
			}

			if tt.response != nil {
				body, _ := json.Marshal(tt.response)
				resp.Body = io.NopCloser(strings.NewReader(string(body)))
				resp.Header.Set("Content-Type", "application/json")
			}

			client := &SafeguardClient{}
			err := client.handleTokenResponse(resp)

			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.response.UserToken, client.AccessToken.UserToken)
			}
		})
	}
}
