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

	if c.HttpClient.Transport == nil {
		c.HttpClient.Transport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
	} else {
		if transport, ok := c.HttpClient.Transport.(*http.Transport); ok {
			transport.TLSClientConfig.Certificates = tlsConfig.Certificates
			transport.TLSClientConfig.Renegotiation = tls.RenegotiateFreelyAsClient
		} else {
			return fmt.Errorf("existing transport is not an *http.Transport")
		}
	}

	// Get RSTS token
	err = c.getRSTSTokenWithCert(c.HttpClient, AuthProviderCertificate)
	if err != nil {
		return fmt.Errorf("acquire RSTS token failed: %v", err)
	}

	// Exchange for Safeguard token
	err = c.exchangeRSTSTokenForSafeguard(c.HttpClient)
	if err != nil {
		return fmt.Errorf("acquire Safeguard token failed: %v", err)
	}

	c.AccessToken.setCertificate(certPath, certPassword)
	fmt.Println("✅ Certificate authentication successful")
	c.authDone <- "Done"
	return nil
}

// getRSTSTokenWithCert retrieves an RSTS token using client certificate authentication.
// It sends a POST request to the RSTS endpoint with the required grant type and scope.
// The function returns the access token as a string or an error if the request fails.
//
// Parameters:
//   - client: An HTTP client to send the request.
//   - authProvider: An AuthProvider instance that provides the scope for the request.
//
// Returns:
//   - A string containing the access token.
//   - An error if the request fails or if there is an issue with the response.
func (c *SafeguardClient) getRSTSTokenWithCert(client *http.Client, authProvider AuthProvider) error {
	requestBody := struct {
		GrantType string `json:"grant_type"`
		Scope     string `json:"scope"`
	}{
		GrantType: "client_credentials",
		Scope:     authProvider.String(),
	}

	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/RSTS/oauth2/token", c.Appliance.getUrl()), bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = c.handleTokenResponse(resp)
	if err != nil {
		return fmt.Errorf("RSTS token request failed: %v", err)
	}

	return nil
}
