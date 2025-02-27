package client

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
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
		c.AccessToken = &TokenResponse{}
	}

	codeVerifier, codeChallenge := generateCodeChallenge()

	authCodeChan := make(chan string)
	errorChan := make(chan error)
	server := startHTTPSListener(authCodeChan, errorChan)
	defer server.Close()

	authURL := fmt.Sprintf("%s/RSTS/Login?response_type=code&code_challenge_method=S256&code_challenge=%s&redirect_uri=%s&port=%d",
		c.ApplicanceURL, codeChallenge, url.QueryEscape(c.redirectURI), c.redirectPort)

	openBrowser(authURL)
	fmt.Println("Please log in using your browser...")

	select {
	case authCode := <-authCodeChan:
		c.AccessToken.AuthorizationCode = authCode
		fmt.Println("\n✅ Authorization Code received")
	case err := <-errorChan:
		return fmt.Errorf("authentication failed: %v", err)
	}

	tokenResponse, err := c.exchangeToken(c.AccessToken.AuthorizationCode, codeVerifier)
	if err != nil {
		return fmt.Errorf("error retrieving token: %v", err)
	}

	c.AccessToken = tokenResponse
	fmt.Println("✅ Access Token received")
	return nil
}

func generateCodeChallenge() (string, string) {
	verifier := make([]byte, 32)
	_, err := rand.Read(verifier)
	if err != nil {
		logger.Fatalf("Error generating Code Verifier: %v", err)
	}
	codeVerifier := base64.RawURLEncoding.EncodeToString(verifier)

	hash := sha256.Sum256([]byte(codeVerifier))
	codeChallenge := base64.RawURLEncoding.EncodeToString(hash[:])

	return codeVerifier, codeChallenge
}

func startHTTPSListener(authCodeChan chan string, errorChan chan error) *http.Server {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		logger.Fatalf("Error loading certificate: %v", err)
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", redirectPort),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		logger.Println("Received callback request")
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
		logger.Printf("Starting HTTPS server on port %d", redirectPort)
		if err := server.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
			errorChan <- fmt.Errorf("HTTPS server error: %v", err)
		}
	}()

	return server
}

func (c *SafeguardClient) exchangeToken(authCode, codeVerifier string) (*TokenResponse, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", authCode)
	data.Set("redirect_uri", c.redirectURI)
	data.Set("code_verifier", codeVerifier)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/RSTS/oauth2/token", c.ApplicanceURL), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("rSTS token request failed: %s", string(body))
	}

	var rStsToken TokenResponse
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&rStsToken); err != nil {
		return nil, err
	}

	logger.Printf("rSTS Token Response: %s", string(body))
	logger.Printf("rSTS Token: %s", rStsToken.AccessToken)

	tokenReq := struct {
		StsAccessToken string `json:"StsAccessToken"`
	}{
		StsAccessToken: rStsToken.AccessToken,
	}

	tokenData, err := json.Marshal(tokenReq)
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequest("POST", fmt.Sprintf("%s/service/core/v4/Token/LoginResponse", c.ApplicanceURL), bytes.NewBuffer(tokenData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err = c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("safeguard token request failed: %s", string(body))
	}

	var tokenResponse TokenResponse
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&tokenResponse); err != nil {
		return nil, err
	}

	logger.Printf("Safeguard Token Response: %s", string(body))
	logger.Printf("Safeguard Token: %s", tokenResponse.UserToken)

	tokenResponse.AccessToken = tokenResponse.UserToken

	return &tokenResponse, nil
}
