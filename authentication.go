package safeguard

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

// RSTSAuthResponse encapsulates authentication data from both RSTS and Safeguard systems.
// It includes tokens, authentication status, and credentials with thread-safe access.
type RSTSAuthResponse struct {
	sync.RWMutex

	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`

	// Safeguard specific fields
	UserToken         string       `json:"UserToken"`
	Status            string       `json:"Status"`
	IdentityProvider  string       `json:"IdentityProvider"`
	AuthorizationCode string       `json:"-"` // Used internally for OAuth flow
	AuthTime          time.Time    `json:"-"` // Time when the token was received
	AuthProvider      AuthProvider `json:"-"` // Type of authentication provider
	credentials       Credentials  `json:"-"`
	isValid           bool         `json:"-"`
}

// getAccessToken safely retrieves the current access token.
func (a *RSTSAuthResponse) getAccessToken() string {
	a.RWMutex.RLock()
	defer a.RWMutex.RUnlock()
	return a.AccessToken
}

// setAccessToken safely updates the access token.
//
// Parameters:
//   - accessToken: The new access token to store
func (a *RSTSAuthResponse) setAccessToken(accessToken string) {
	a.RWMutex.Lock()
	defer a.RWMutex.Unlock()
	a.AccessToken = accessToken
}

// getUserToken safely retrieves the current user token.
func (a *RSTSAuthResponse) getUserToken() string {
	a.RWMutex.RLock()
	defer a.RWMutex.RUnlock()
	return a.UserToken
}

// setUserToken safely updates the user token.
//
// Parameters:
//   - userToken: The new user token to store
func (a *RSTSAuthResponse) setUserToken(userToken string) {
	a.RWMutex.Lock()
	defer a.RWMutex.Unlock()
	a.UserToken = userToken
}

// setUserNamePassword safely stores username and password credentials.
//
// Parameters:
//   - username: The username to store
//   - password: The password to store
func (a *RSTSAuthResponse) setUserNamePassword(username, password string) {
	a.RWMutex.Lock()
	defer a.RWMutex.Unlock()
	a.credentials.username = username
	a.credentials.password = password
}

// getUserNamePassword safely retrieves the stored username and password.
//
// Returns:
//   - string: The stored username
//   - string: The stored password
func (a *RSTSAuthResponse) getUserNamePassword() (string, string) {
	a.RWMutex.RLock()
	defer a.RWMutex.RUnlock()
	return a.credentials.username, a.credentials.password
}

// setCertificate safely stores certificate credentials.
//
// Parameters:
//   - certPath: Path to the certificate file
//   - certPassword: Password for the certificate
func (a *RSTSAuthResponse) setCertificate(certPath, certPassword string) {
	a.RWMutex.Lock()
	defer a.RWMutex.Unlock()
	a.credentials.certPath = certPath
	a.credentials.certPassword = certPassword
}

// getCertificate safely retrieves the stored certificate credentials.
//
// Returns:
//   - string: Path to the certificate file
//   - string: Password for the certificate
func (a *RSTSAuthResponse) getCertificate() (string, string) {
	a.RWMutex.RLock()
	defer a.RWMutex.RUnlock()
	return a.credentials.certPath, a.credentials.certPassword
}

// Credentials stores various authentication credentials securely.
type Credentials struct {
	username     string
	password     string
	certPath     string
	certPassword string
}

// LoginWithOauth initiates the OAuth2.0 authorization code flow to obtain an access token.
// It generates a code challenge and starts a TCP listener to receive the authorization code.
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
	listener := startTCPListener(authCodeChan, errorChan)
	defer listener.Close()

	redirectURI := "urn:InstalledApplicationTcpListener"
	authURL := fmt.Sprintf("%s/RSTS/Login?response_type=code&code_challenge_method=S256&code_challenge=%s&redirect_uri=%s&port=%d",
		c.Appliance.getUrl(), codeChallenge, url.QueryEscape(redirectURI), c.redirectPort)

	openBrowser(authURL)
	fmt.Println("Please log in using your browser...")

	select {
	case authCode := <-authCodeChan:
		c.AccessToken.AuthorizationCode = authCode
		fmt.Println("\n✅ Authorization Code received")
	case err := <-errorChan:
		return fmt.Errorf("authentication failed: %v", err)
	}

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", c.AccessToken.AuthorizationCode)
	data.Set("redirect_uri", redirectURI)
	data.Set("code_verifier", codeVerifier)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/RSTS/oauth2/token", c.Appliance.getUrl()), strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("RSTS token request failed: %v", err)
	}
	defer resp.Body.Close()

	err = c.handleTokenResponse(resp)
	if err != nil {
		return fmt.Errorf("error retrieving token: %v", err)
	}

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

func startTCPListener(authCodeChan chan string, errorChan chan error) net.Listener {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", redirectPort))
	if err != nil {
		logger.Error("Error starting TCP listener", "error", err)
		errorChan <- err
		return nil
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			errorChan <- err
			return
		}
		defer conn.Close()

		buffer := make([]byte, 4096)
		n, err := conn.Read(buffer)
		if err != nil {
			errorChan <- err
			return
		}

		request := string(buffer[:n])
		logger.Debug("Received request", "request", request)

		authCode := extractAuthCode(request)
		if authCode != "" {
			// Simple HTTP success response
			response := "HTTP/1.1 200 OK\r\nConnection: close\r\n\r\nAuthentication successful"
			conn.Write([]byte(response))
			authCodeChan <- authCode
		} else {
			errorChan <- fmt.Errorf("no authorization code received in request: %s", request)
		}
	}()

	return listener
}

func extractAuthCode(request string) string {
	re := regexp.MustCompile(`GET /\?(.+) HTTP`)
	match := re.FindStringSubmatch(request)
	if len(match) < 2 {
		logger.Error("No URL parameters found in request")
		return ""
	}

	params, err := url.ParseQuery(match[1])
	if err != nil {
		logger.Error("Failed to parse query parameters", "error", err)
		return ""
	}

	// Look specifically for the 'oauth' parameter
	if code := params.Get("oauth"); code != "" {
		logger.Debug("Found oauth code", "code", code[:30]+"...") // Log first 30 chars
		return code
	}

	logger.Error("No oauth parameter found in query", "params", params)
	return ""
}

// LoginWithPassword authenticates a user using their username and password.
// It first obtains an RSTS token and then exchanges it for a Safeguard token.
// The obtained token is stored in the SafeguardClient's AccessToken field.
//
// Parameters:
//   - username: The username of the user.
//   - password: The password of the user.
//
// Returns:
//   - error: An error if the login process fails, otherwise nil.
//
// The function automatically handles:
//   - RSTS token acquisition
//   - Token exchange for Safeguard access
//   - Token storage and management
//   - Error handling and logging
func (c *SafeguardClient) LoginWithPassword(username, password string) error {
	if c.AccessToken == nil {
		c.AccessToken = &RSTSAuthResponse{
			AuthProvider: AuthProviderLocal,
		}
	} else {
		c.AccessToken.AuthProvider = AuthProviderLocal
	}

	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", username)
	data.Set("password", password)
	data.Set("scope", AuthProviderLocal.String())

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/RSTS/oauth2/token", c.Appliance.getUrl()), strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("RSTS login request failed: %v", err)
	}
	defer resp.Body.Close()

	err = c.handleTokenResponse(resp)
	if err != nil {
		return fmt.Errorf("RSTS login failed: %v", err)
	}

	err = c.exchangeRSTSTokenForSafeguard(c.HttpClient)
	if err != nil {
		return fmt.Errorf("token exchange failed: %v", err)
	}

	c.AccessToken.setUserNamePassword(username, password)
	fmt.Println("✅ Login successful")
	c.authDone <- "Done"
	return nil
}

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
