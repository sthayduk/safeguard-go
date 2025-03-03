package client

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// LoginWithOauth initiates the OAuth2.0 authorization code flow to obtain an access token.
// It generates a code challenge and starts an HTTPS listener to receive the authorization code.
// The user is prompted to log in using their browser, and upon successful login, the authorization code
// is exchanged for an access token.
// Returns an error if the authentication or token exchange process fails.
func (c *SafeguardClient) LoginWithOauth() error {
	if c.AccessToken == nil {
		c.AccessToken = &RSTSAuthResponse{}
	}

	codeVerifier, codeChallenge := generateCodeChallenge()

	authCodeChan := make(chan string)
	errorChan := make(chan error)
	server := startHTTPSListener(authCodeChan, errorChan)
	defer server.Close()

	authURL := fmt.Sprintf("%s/RSTS/Login?response_type=code&code_challenge_method=S256&code_challenge=%s&redirect_uri=%s&port=%d",
		c.Appliance.getUrl(), codeChallenge, url.QueryEscape(c.redirectURI), c.redirectPort)

	openBrowser(authURL)
	fmt.Println("Please log in using your browser...")

	select {
	case authCode := <-authCodeChan:
		c.AccessToken.AuthorizationCode = authCode
		fmt.Println("\n✅ Authorization Code received")
	case err := <-errorChan:
		return fmt.Errorf("authentication failed: %v", err)
	}

	err := c.getRSTSTokenWithOauth(c.AccessToken.AuthorizationCode, codeVerifier)
	if err != nil {
		return fmt.Errorf("error retrieving token: %v", err)
	}

	// Exchange for Safeguard token
	err = c.exchangeRSTSTokenForSafeguard(c.HttpClient)
	if err != nil {
		return fmt.Errorf("acquire Safeguard token failed: %v", err)
	}

	fmt.Println("✅ Access Token received")
	c.authDone <- "Done"

	return nil
}

func generateCodeChallenge() (string, string) {
	verifier := make([]byte, 32)
	_, err := rand.Read(verifier)
	if err != nil {
		logger.Error("Error generating Code Verifier", "error", err)
		panic(fmt.Sprintf("Error generating Code Verifier: %v", err))
	}
	codeVerifier := base64.RawURLEncoding.EncodeToString(verifier)

	hash := sha256.Sum256([]byte(codeVerifier))
	codeChallenge := base64.RawURLEncoding.EncodeToString(hash[:])

	return codeVerifier, codeChallenge
}

func startHTTPSListener(authCodeChan chan string, errorChan chan error) *http.Server {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		logger.Error("Error loading certificate", "error", err)
		panic(fmt.Sprintf("Error loading certificate: %v", err))
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", redirectPort),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Received callback request")
		if err := r.URL.Query().Get("error"); err != "" {
			errorDesc := r.URL.Query().Get("error_description")
			errorChan <- fmt.Errorf("%s: %s", err, errorDesc)
			w.Write([]byte("Authentication failed. You can close this window."))
			return
		}

		authCode := r.URL.Query().Get("code")
		if authCode != "" {
			authCodeChan <- authCode
			w.Write([]byte("Authentication successful! You can close this window."))
		} else {
			errorChan <- fmt.Errorf("no authorization code received")
			w.Write([]byte("Authentication failed. No authorization code received."))
		}
	})

	go func() {
		logger.Info("Starting HTTPS server", "port", redirectPort)
		if err := server.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
			errorChan <- fmt.Errorf("HTTPS server error: %v", err)
		}
	}()

	return server
}

// getRSTSTokenWithOauth exchanges an authorization code for an RSTS token using OAuth2.
// It sends a POST request to the RSTS token endpoint with the necessary parameters.
//
// Parameters:
//
//	authCode - The authorization code received from the OAuth2 authorization server.
//	codeVerifier - The code verifier used in the PKCE (Proof Key for Code Exchange) flow.
//
// Returns:
//
//	error - An error if the token request fails or if there is an issue handling the token response.
func (c *SafeguardClient) getRSTSTokenWithOauth(authCode, codeVerifier string) error {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", authCode)
	data.Set("redirect_uri", c.redirectURI)
	data.Set("code_verifier", codeVerifier)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/RSTS/oauth2/token", c.Appliance.getUrl()), strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.HttpClient.Do(req)
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
