package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// LoginWithCertificate authenticates using a PKCS12 certificate file.
// Parameters:
// - certPath: Path to the PKCS12 certificate file
// - certPassword: Password for the certificate
// - authProvider: The authentication provider to use (e.g. "certificate")
// Returns an error if the authentication fails.
func (c *SafeguardClient) LoginWithCertificate(certPath, certPassword string) error {
	if c.AccessToken == nil {
		c.AccessToken = &RSTSAuthResponse{
			AuthProvider: AuthProviderCertificate,
		}
	} else {
		c.AccessToken.AuthProvider = AuthProviderCertificate
	}

	// Read client certificate
	certData, err := os.ReadFile(certPath)
	if err != nil {
		return fmt.Errorf("read certificate file failed: %v", err)
	}

	// Create TLS config with PKCS12 cert
	tlsConfig, err := tLSConfigForPKCS12(certData, certPassword)
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
	rstsToken, err := c.getRSTSTokenWithCert(tempClient, AuthProviderCertificate)
	if err != nil {
		return fmt.Errorf("acquire RSTS token failed: %v", err)
	}

	// Exchange for Safeguard token
	safeguardToken, err := c.exchangeRSTSToken(tempClient, rstsToken)
	if err != nil {
		return fmt.Errorf("acquire Safeguard token failed: %v", err)
	}
	c.AccessToken.UserToken = safeguardToken
	c.AccessToken.AccessToken = safeguardToken
	c.AccessToken.credentials.certPath = certPath
	c.AccessToken.credentials.certPassword = certPassword
	fmt.Println("✅ Certificate authentication successful")

	return nil
}

func (c *SafeguardClient) getRSTSTokenWithCert(client *http.Client, authProvider AuthProvider) (string, error) {
	requestBody := struct {
		GrantType string `json:"grant_type"`
		Scope     string `json:"scope"`
	}{
		GrantType: "client_credentials",
		Scope:     authProvider.String(),
	}

	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/RSTS/oauth2/token", c.ApplicanceURL), bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	rstsResp, err := handleTokenResponse(resp)
	if err != nil {
		return "", fmt.Errorf("RSTS token request failed: %v", err)
	}

	// Store RSTS response
	c.AccessToken = rstsResp
	return rstsResp.AccessToken, nil
}

func (c *SafeguardClient) exchangeRSTSToken(client *http.Client, rstsToken string) (string, error) {
	safeguardResponse, err := c.exchangeRSTSTokenForSafeguard(client, rstsToken)
	if err != nil {
		return "", err
	}

	c.AccessToken = safeguardResponse
	return safeguardResponse.UserToken, nil
}
