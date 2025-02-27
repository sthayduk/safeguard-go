package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// LoginWithPassword authenticates a user using their username and password.
// It first obtains an RSTS token and then exchanges it for a Safeguard token.
// The obtained token is stored in the SafeguardClient's AccessToken field.
// Parameters:
// - username: The username of the user.
// - password: The password of the user.
// Returns an error if the login process fails, otherwise nil.
func (c *SafeguardClient) LoginWithPassword(username, password string) error {
	if c.AccessToken == nil {
		c.AccessToken = &TokenResponse{}
	}

	// Create form data with scope parameter
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", username)
	data.Set("password", password)
	data.Set("scope", "rsts:sts:primaryproviderid:local") // Correct scope parameter

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/RSTS/oauth2/token", c.ApplicanceURL), strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	logger.Info("Making RSTS login request", "url", req.URL.String())
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("RSTS login request failed: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("RSTS login failed: %s", string(body))
	}

	var rstsToken TokenResponse
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&rstsToken); err != nil {
		return fmt.Errorf("error decoding RSTS response: %v", err)
	}

	logger.Debug("RSTS Login Response", "body", string(body))

	// Exchange RSTS token for Safeguard token
	tokenReq := struct {
		StsAccessToken string `json:"StsAccessToken"`
	}{
		StsAccessToken: rstsToken.AccessToken,
	}

	tokenData, err := json.Marshal(tokenReq)
	if err != nil {
		return fmt.Errorf("error marshaling token request: %v", err)
	}

	req, err = http.NewRequest("POST", fmt.Sprintf("%s/Token/LoginResponse", c.getRootUrl()), bytes.NewBuffer(tokenData))
	if err != nil {
		return fmt.Errorf("error creating token request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err = c.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("token request failed: %v", err)
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("token exchange failed: %s", string(body))
	}

	var tokenResponse TokenResponse
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&tokenResponse); err != nil {
		return fmt.Errorf("error decoding token response: %v", err)
	}

	logger.Debug("Token Response", "body", string(body))
	c.AccessToken = &tokenResponse
	c.AccessToken.AccessToken = tokenResponse.UserToken
	fmt.Println("✅ Login successful")
	return nil
}
