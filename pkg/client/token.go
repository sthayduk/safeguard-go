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
	err := os.Setenv(envVar, c.AccessToken.getAccessToken())
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
	if c.AccessToken.getAccessToken() == "" {
		c.AccessToken.isValid = false
		return fmt.Errorf("access token is empty")
	}

	logger.Debug("Token validation",
		"length", len(c.AccessToken.getAccessToken()),
		"formatCheck", strings.HasPrefix(c.AccessToken.getAccessToken(), "ey"))

	fields := []string{"id"}
	err := c.testAccessToken(fields...)
	if err != nil {
		c.AccessToken.isValid = false
		return fmt.Errorf("invalid access token: %v", err)
	}
	c.AccessToken.isValid = true
	return nil
}

// getAuthorizationHeader sets the necessary headers for authorization and content type
// on the provided HTTP request. It adds an "accept" header with the value "application/json"
// and an "Authorization" header with the Bearer token from the SafeguardClient's AccessToken.
//
// Parameters:
//
//	req - The HTTP request to which the headers will be added.
//
// Returns:
//
//	The modified HTTP request with the added headers.
func (c *SafeguardClient) getAuthorizationHeader(req *http.Request) *http.Request {
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.AccessToken.getAccessToken())

	return req
}

// testAccessToken checks the access token by making a GET request to the "me" endpoint.
// If fields are provided, they are appended as query parameters to the request.
//
// Parameters:
//
//	fields - Optional list of fields to include in the query.
//
// Returns:
//
//	error - An error if the request fails, otherwise nil.
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

// exchangeRSTSTokenForSafeguard exchanges an RSTS token for a Safeguard token.
//
// This function sends a POST request to the Safeguard appliance to exchange the provided RSTS token
// for a Safeguard token. It constructs the request payload, sends the request, and processes the response.
//
// Parameters:
//   - client: The HTTP client used to send the request.
//   - rstsToken: The RSTS token to be exchanged.
//
// Returns:
//   - A pointer to an RSTSAuthResponse containing the Safeguard token information.
//   - An error if the request fails or the response cannot be processed.
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

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/service/core/v4/Token/LoginResponse", c.Appliance.getUrl()), bytes.NewBuffer(tokenData))
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

	safeguardResponse.setAccessToken(safeguardResponse.UserToken)
	safeguardResponse.AuthTime = time.Now()

	return &safeguardResponse, nil
}

// handleTokenResponse processes the HTTP response from a token request.
// It reads the response body and checks if the status code is OK (200).
// If the status code is not OK, it returns an error with the response body as the message.
// If the status code is OK, it decodes the response body into an RSTSAuthResponse struct
// and returns it. If decoding fails, it returns the decoding error.
//
// Parameters:
//   - resp: The HTTP response from the token request.
//
// Returns:
//   - A pointer to an RSTSAuthResponse struct if the request was successful.
//   - An error if the request failed or if decoding the response body failed.
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
