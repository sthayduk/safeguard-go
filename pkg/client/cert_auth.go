package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// LoginWithCertificate authenticates using a PKCS12 certificate file.
// Parameters:
// - certPath: Path to the PKCS12 certificate file
// - certPassword: Password for the certificate
// - authProvider: The authentication provider to use (e.g. "certificate")
// Returns an error if the authentication fails.
func (c *SafeguardClient) LoginWithCertificate(certPath, certPassword, authProvider string) error {
	if c.AccessToken == nil {
		c.AccessToken = &TokenResponse{}
	}

	// Read client certificate
	certData, err := os.ReadFile(certPath)
	if err != nil {
		return fmt.Errorf("read certificate file failed: %v", err)
	}

	// Create TLS config with PKCS12 cert
	tlsConfig, err := TLSConfigForPKCS12(certData, certPassword)
	if err != nil {
		return fmt.Errorf("create tls config failed: %v", err)
	}

	// Enable insecure skip verify for Safeguard appliance
	tlsConfig.InsecureSkipVerify = true
	tlsConfig.Renegotiation = tls.RenegotiateOnceAsClient

	// Create temporary client with cert config
	tempClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// Get RSTS token
	rstsToken, err := c.getRSTSTokenWithCert(tempClient, authProvider)
	if err != nil {
		return fmt.Errorf("acquire RSTS token failed: %v", err)
	}

	// Exchange for Safeguard token
	safeguardToken, err := c.exchangeRSTSToken(tempClient, rstsToken)
	if err != nil {
		return fmt.Errorf("acquire Safeguard token failed: %v", err)
	}
	c.AccessToken.AccessToken = safeguardToken
	fmt.Println("✅ Certificate authentication successful")
	return nil
}

func (c *SafeguardClient) getRSTSTokenWithCert(client *http.Client, authProvider string) (string, error) {
	// Create request body using struct
	requestBody := struct {
		GrantType string `json:"grant_type"`
		Scope     string `json:"scope"`
	}{
		GrantType: "client_credentials",
		Scope:     authProvider, // Use authProvider directly without prefix
	}

	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	logger.Printf("RSTS request body: %s", string(bodyBytes))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/RSTS/oauth2/token", c.ApplicanceURL), bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	logger.Printf("RSTS request URL: %s", req.URL.String())

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	logger.Printf("RSTS response: %s", string(body))
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("RSTS token request failed: %s", string(body))
	}

	var authResp struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(body, &authResp); err != nil {
		return "", err
	}

	if authResp.AccessToken == "" {
		return "", fmt.Errorf("no access token received")
	}

	return authResp.AccessToken, nil
}

func (c *SafeguardClient) exchangeRSTSToken(client *http.Client, rstsToken string) (string, error) {
	data := strings.NewReader(fmt.Sprintf(`{
		"StsAccessToken": "%s"
	}`, rstsToken))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/service/core/v4/Token/LoginResponse", c.ApplicanceURL), data)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("safeguard token request failed: %s", string(body))
	}

	var tokenResp struct {
		UserToken string `json:"UserToken"`
	}
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", err
	}

	return tokenResp.UserToken, nil
}
