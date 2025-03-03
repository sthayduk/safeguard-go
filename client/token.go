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

// SaveAccessTokenToEnv saves the current access token to an environment variable.
// This allows persistence of the token across sessions.
//
// Returns:
//   - error: An error if saving the token fails.
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

// ValidateAccessToken checks if the current access token is valid by testing it
// against the Safeguard API. It verifies both token format and server acceptance.
//
// Returns:
//   - error: An error if the token is invalid or validation fails.
func (c *SafeguardClient) ValidateAccessToken() error {
	if c.AccessToken.getUserToken() == "" {
		c.AccessToken.isValid = false
		return fmt.Errorf("access token is empty")
	}

	logger.Debug("Token validation",
		"length", len(c.AccessToken.getUserToken()),
		"formatCheck", strings.HasPrefix(c.AccessToken.getUserToken(), "ey"))

	fields := []string{"id"}
	err := c.testAccessToken(fields...)
	if err != nil {
		c.AccessToken.isValid = false
		return fmt.Errorf("invalid access token: %v", err)
	}
	c.AccessToken.isValid = true
	return nil
}

// getAuthorizationHeader prepares authorization headers for API requests.
// It formats the access token according to the required Bearer scheme.
//
// Parameters:
//   - req: The HTTP request to modify with authorization headers.
//
// Returns:
//   - *http.Request: The modified request with authorization headers.
func (c *SafeguardClient) getAuthorizationHeader() http.Header {
	headers := http.Header{}
	headers.Set("accept", "application/json")
	headers.Set("Authorization", "Bearer "+c.AccessToken.getUserToken())

	return headers
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
//
// Returns:
//   - A pointer to an RSTSAuthResponse containing the Safeguard token information.
//   - An error if the request fails or the response cannot be processed.
func (c *SafeguardClient) exchangeRSTSTokenForSafeguard(client *http.Client) error {
	// tokenReq required because LoginResponse need AccessToken as StsAccessToken
	tokenReq := struct {
		StsAccessToken string `json:"StsAccessToken"`
	}{
		StsAccessToken: c.AccessToken.getAccessToken(),
	}

	tokenData, err := json.Marshal(tokenReq)
	if err != nil {
		return fmt.Errorf("error marshaling token request: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/service/core/v4/Token/LoginResponse", c.Appliance.getUrl()), bytes.NewBuffer(tokenData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("safeguard token request failed: %s", string(body))
	}

	var safeguardResponse RSTSAuthResponse
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&safeguardResponse); err != nil {
		return err
	}

	c.AccessToken.setUserToken(safeguardResponse.UserToken)
	c.AccessToken.AuthTime = time.Now()

	return nil
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
func (c *SafeguardClient) handleTokenResponse(resp *http.Response) error {
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("token request failed: %s", string(body))
	}

	var tokenResp *RSTSAuthResponse
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&tokenResp); err != nil {
		return err
	}

	c.AccessToken = tokenResp
	c.AccessToken.AuthProvider = AuthProvider(tokenResp.Scope)
	return nil
}
