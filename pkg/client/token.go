package client

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// SaveAccessTokenToEnv saves the access token to an environment variable.
// Returns an error if the operation fails.
func (c *SafeguardClient) SaveAccessTokenToEnv() error {
	envVar := "SAFEGUARD_ACCESS_TOKEN"
	err := os.Setenv(envVar, c.AccessToken.AccessToken)
	if err != nil {
		logger.Fatalf("Error saving access token to environment variable: %v", err)
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

	logger.Printf("Token length: %d", len(c.AccessToken.AccessToken))
	logger.Printf("Token format check - starts with 'ey': %v", strings.HasPrefix(c.AccessToken.AccessToken, "ey"))

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
