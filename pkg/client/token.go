package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// SaveAccessTokenToEnv saves the access token to an environment variable.
// Returns an error if the operation fails.
func (c *SafeguardClient) SaveAccessTokenToEnv() error {
	envVar := "SAFEGUARD_ACCESS_TOKEN"
	err := os.Setenv(envVar, c.AccessToken.AccessToken)
	if err != nil {
		logger.Error("Error saving access token to environment variable", "error", err)
		return err
	}

	fmt.Printf("Access token saved to environment variable: %s\n", envVar)
	return nil
}

// ValidateAccessToken checks if the access token in the SafeguardClient is valid.
// It returns an error if the access token is empty or invalid.
// Returns nil if the access token is valid, otherwise an error.
func (c *SafeguardClient) ValidateAccessToken() error {
	if c.AccessToken.AccessToken == "" {
		return fmt.Errorf("access token is empty")
	}

	logger.Debug("Token validation",
		"length", len(c.AccessToken.AccessToken),
		"formatCheck", strings.HasPrefix(c.AccessToken.AccessToken, "ey"))

	err := c.testAccessToken()
	if err != nil {
		return fmt.Errorf("invalid access token: %v", err)
	}
	return nil
}

func (c *SafeguardClient) getAuthorizationHeader(req *http.Request) *http.Request {
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.AccessToken.AccessToken)

	return req
}

func (c *SafeguardClient) testAccessToken(fields ...string) error {
	query := "me"
	if len(fields) > 0 {
		query += "?fields=" + fields[0]
		for _, field := range fields[1:] {
			query += "," + field
		}
	}
	_, err := c.GetRequest(query)
	if err != nil {
		return err
	}

	return nil
}

// exchangeRSTSTokenForSafeguard exchanges an RSTS token for a Safeguard token
func (c *SafeguardClient) exchangeRSTSTokenForSafeguard(client *http.Client, rstsToken string) (*RSTSAuthResponse, error) {
	tokenReq := struct {
		StsAccessToken string `json:"StsAccessToken"`
	}{
		StsAccessToken: rstsToken,
	}

	tokenData, err := json.Marshal(tokenReq)
	if err != nil {
		return nil, fmt.Errorf("error marshaling token request: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/service/core/v4/Token/LoginResponse", c.ApplicanceURL), bytes.NewBuffer(tokenData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("safeguard token request failed: %s", string(body))
	}

	var safeguardResponse RSTSAuthResponse
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&safeguardResponse); err != nil {
		return nil, err
	}

	// If we have existing RSTS token info, merge it
	if c.AccessToken != nil {
		safeguardResponse.RefreshToken = c.AccessToken.RefreshToken
		safeguardResponse.TokenType = c.AccessToken.TokenType
		safeguardResponse.ExpiresIn = c.AccessToken.ExpiresIn
		safeguardResponse.Scope = c.AccessToken.Scope
	}

	safeguardResponse.AccessToken = safeguardResponse.UserToken
	safeguardResponse.AuthTime = time.Now()

	return &safeguardResponse, nil
}

// handleTokenResponse processes an HTTP response containing a token
func handleTokenResponse(resp *http.Response) (*RSTSAuthResponse, error) {
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("token request failed: %s", string(body))
	}

	var tokenResp RSTSAuthResponse
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&tokenResp); err != nil {
		return nil, err
	}

	return &tokenResp, nil
}
