package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"time"
)

var logger *slog.Logger // Declare global logger variable
var sgclient *SafeguardClient

func init() {
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
}

const (
	redirectPort = 8400
	redirectURI  = "https://localhost:8400/callback"
)

// Returns a pointer to a SafeguardClient instance.
// New creates a new instance of SafeguardClient. It initializes the logger with the specified
// debug level and sets it as the default logger. If an existing client instance (sgclient)
// already exists, it returns that instance. Otherwise, it creates a new SafeguardClient with
// the provided appliance URL, API version, and other necessary configurations. It also starts
// a goroutine to refresh the token periodically.
//
// Parameters:
//   - applianceUrl: The URL of the appliance to connect to.
//   - apiVersion: The version of the API to use.
//   - debug: A boolean flag to enable or disable debug logging.
//
// Returns:
//
//	A pointer to the newly created or existing SafeguardClient instance.
func New(applianceUrl string, apiVersion string, debug bool) *SafeguardClient {
	var opts slog.HandlerOptions
	if debug {
		opts.Level = slog.LevelDebug
	} else {
		opts.Level = slog.LevelInfo
	}

	logger = slog.New(slog.NewTextHandler(os.Stdout, &opts))
	slog.SetDefault(logger)

	if sgclient != nil {
		return sgclient
	}

	sgclient = &SafeguardClient{
		AccessToken:   &RSTSAuthResponse{},
		ApiVersion:    apiVersion,
		HttpClient:    createTLSClient(),
		redirectPort:  redirectPort,
		redirectURI:   redirectURI,
		tokenEndpoint: applianceUrl + "/service/core/v4/Token/LoginResponse",

		// channel to signal when authentication is done
		authDone: make(chan string),
	}

	sgclient.Appliance.setUrl(applianceUrl, -1)

	ctx := context.Background()
	go sgclient.refreshToken(ctx)
	return sgclient
}

// getClusterLeaderUrl returns the URL of the cluster leader.
// This URL is used to identify the leader node in a cluster setup.
func (c *SafeguardClient) getClusterLeaderUrl() string {
	if c.ClusterLeader.getUrl() == "" {
		for {
			if c.ClusterLeader.getUrl() != "" {
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
	return c.ClusterLeader.getUrl()
}

// createTLSClient creates and returns a new HTTP client with TLS configuration.
// It loads the CA and root certificates from the specified files and adds them
// to the certificate pool. If any error occurs during this process, the function
// logs the error and exits the program.
//
// Returns:
//
//	*http.Client: A pointer to the configured HTTP client.
func createTLSClient() *http.Client {
	caCert, err := os.ReadFile("server.crt")
	if err != nil {
		logger.Error("Error loading CA certificate", "error", err)
		os.Exit(1)
	}
	rootCert, err := os.ReadFile("pam.cer")
	if err != nil {
		logger.Error("Error loading root certificate", "error", err)
		os.Exit(1)
	}

	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		logger.Error("Error adding CA certificate to pool")
		os.Exit(1)
	}
	if !caCertPool.AppendCertsFromPEM(rootCert) {
		logger.Error("Error adding root certificate to pool")
		os.Exit(1)
	}

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:    caCertPool,
				MinVersion: tls.VersionTLS12,
				MaxVersion: tls.VersionTLS13,
			},
		},
	}
}

// refreshToken refreshes the access token for the SafeguardClient.
// It waits until the initial authentication is done if necessary,
// then sets up a ticker to refresh the token periodically based on the remaining token time.
// The function supports two types of authentication providers: local and certificate-based.
// It will stop refreshing the token if the provided context is done.
//
// Parameters:
// - ctx: The context to control the lifecycle of the token refresh process.
func (c *SafeguardClient) refreshToken(ctx context.Context) {
	<-c.authDone

	if c.AccessToken.AuthProvider == "" {
		logger.Debug("token refresh skipped: no auth provider")
		return
	}

	logger.Debug("token refresh started")
	remainingTokenTime := c.RemainingTokenTime()
	ticker := time.NewTicker(remainingTokenTime - 1*time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			if c.AccessToken.AuthProvider == AuthProviderLocal {
				refreshTokenWithPassword(c)
			} else if c.AccessToken.AuthProvider == AuthProviderCertificate {
				refreshTokenWithCertificate(c)
			}
		}
	}
}

// refreshTokenWithCertificate refreshes the access token of the SafeguardClient
// using the provided certificate credentials.
//
// It calls the LoginWithCertificate method on the SafeguardClient instance,
// passing the certificate path and password from the client's AccessToken credentials.
//
// Parameters:
//
//	c - A pointer to the SafeguardClient instance whose token needs to be refreshed.
func refreshTokenWithCertificate(c *SafeguardClient) {
	// Refresh the token using the certificate
	c.LoginWithCertificate(c.AccessToken.getCertificate())
}

// refreshTokenWithPassword refreshes the authentication token for the SafeguardClient
// using the stored username and password credentials.
//
// Parameters:
// - c: A pointer to the SafeguardClient instance.
//
// This function calls the LoginWithPassword method on the SafeguardClient
// to obtain a new access token using the current username and password.
func refreshTokenWithPassword(c *SafeguardClient) {
	// Refresh the token using the password
	c.LoginWithPassword(c.AccessToken.getUserNamePassword())
}

// GetTokenExpirationTime returns the time when the current access token will expire
// GetTokenExpirationTime returns the expiration time of the access token.
// It calculates the expiration time by adding the token's lifespan (ExpiresIn)
// to the authentication time (AuthTime).
//
// Returns:
//
//	time.Time: The expiration time of the access token.
func (c *SafeguardClient) GetTokenExpirationTime() time.Time {
	return c.AccessToken.AuthTime.Add(time.Duration(c.AccessToken.ExpiresIn) * time.Second)
}

// IsTokenExpired checks if the current access token has expired
// IsTokenExpired checks if the current access token is expired.
// It returns true if the access token is nil, the authentication time is zero,
// or the current time is after the token's expiration time.
func (c *SafeguardClient) IsTokenExpired() bool {
	if c.AccessToken == nil || c.AccessToken.AuthTime.IsZero() {
		return true
	}
	return time.Now().After(c.GetTokenExpirationTime())
}

// RemainingTokenTime returns the duration until the token expires
// RemainingTokenTime returns the remaining time until the access token expires.
// If the access token is nil or the authentication time is zero, it returns a duration of zero.
func (c *SafeguardClient) RemainingTokenTime() time.Duration {
	if c.AccessToken == nil || c.AccessToken.AuthTime.IsZero() {
		return 0
	}
	return time.Until(c.GetTokenExpirationTime())
}

// updateClusterLeaderUrl retrieves the cluster leader host name and updates the
// cluster leader URL for the SafeguardClient. If an error occurs while getting
// the cluster leader host name, it logs the error and returns without updating
// the cluster leader URL.
func (c *SafeguardClient) updateClusterLeaderUrl() {
	clusterLeaderHostName, err := c.getClusterLeaderHostName()
	if err != nil {
		logger.Error("Failed to get cluster leader host name", "error", err)
		return
	}
	c.setClusterLeader(clusterLeaderHostName)
}

// setClusterLeader sets the cluster leader URL for the SafeguardClient.
// It takes the hostname of the cluster leader as a parameter and generates
// the corresponding URL. If the generated URL is the same as the current
// cluster leader URL, no changes are made. If the generated URL is the same
// as the appliance URL, the cluster leader URL is set to the appliance URL.
// Otherwise, the cluster leader URL is updated to the new URL.
//
// Parameters:
//   - clusterLeaderHostName: The hostname of the cluster leader.
//
// Logs:
//   - Debug: When setting the cluster leader, when the cluster leader is
//     unchanged, when the cluster leader is the same as the appliance URL,
//     and when updating the cluster leader URL.
//   - Error: If there is an error generating the cluster leader URL.
func (c *SafeguardClient) setClusterLeader(clusterLeaderHostName string) {
	clusterLeaderTimeout := time.Duration(10 * time.Second)

	logger.Debug("Setting cluster leader", "hostname", clusterLeaderHostName)
	clusterLeaderUrl, err := c.generateClusterLeaderURL(clusterLeaderHostName)
	if err != nil {
		logger.Error("Failed to set cluster leader", "error", err)
		return
	}

	if sgclient.ClusterLeader.getUrl() == clusterLeaderUrl {
		logger.Debug("Cluster leader unchanged", "url", clusterLeaderUrl)
	}

	if sgclient.Appliance.getUrl() == clusterLeaderUrl {
		logger.Debug("Cluster leader is same as appliance URL", "url", clusterLeaderUrl)
	}

	logger.Debug("Updating cluster leader URL",
		"old", sgclient.ClusterLeader.getUrl(),
		"new", clusterLeaderUrl)
	sgclient.ClusterLeader.setUrl(clusterLeaderUrl, clusterLeaderTimeout)
}

// generateClusterLeaderURL generates the URL for the cluster leader based on the provided
// cluster leader host name. It splits the appliance URL to extract the protocol, domain name,
// and port, and then constructs the cluster leader URL accordingly. If the domain name is empty,
// it constructs the URL without the domain name. If there is an error while splitting the appliance
// URL, it logs the error and returns an empty string along with the error.
//
// Parameters:
//   - clusterLeaderHostName: The host name of the cluster leader.
//
// Returns:
//   - string: The generated cluster leader URL.
//   - error: An error if there was an issue generating the URL.
func (c *SafeguardClient) generateClusterLeaderURL(clusterLeaderHostName string) (string, error) {
	logger.Debug("Generating cluster leader URL", "hostname", clusterLeaderHostName)
	protocol, _, domainName, port, err := splitApplianceURL(sgclient.Appliance.getUrl())
	if err != nil {
		logger.Error("Error splitting appliance URL", "error", err)
		return "", err
	}

	var clusterLeaderUrl string
	if domainName == "" {
		clusterLeaderUrl = fmt.Sprintf("%s://%s:%s", protocol, clusterLeaderHostName, port)
	} else {
		clusterLeaderUrl = fmt.Sprintf("%s://%s.%s:%s", protocol, clusterLeaderHostName, domainName, port)
	}
	logger.Debug("Generated cluster leader URL", "url", clusterLeaderUrl)
	return clusterLeaderUrl, nil
}

// getClusterLeaderHostName fetches the hostname of the cluster leader.
// It sends a request to the "Cluster/Members" endpoint with specific query parameters
// to filter for the leader and retrieve its name. The function returns the hostname
// of the cluster leader or an error if the request fails or no leader is found.
//
// Returns:
//   - string: The hostname of the cluster leader.
//   - error: An error if the request fails or no leader is found.
func (c *SafeguardClient) getClusterLeaderHostName() (string, error) {
	logger.Debug("Fetching cluster leader hostname")

	query := "Cluster/Members"
	params := url.Values{}
	params.Add("filter", "IsLeader eq true")
	params.Add("count", "false")
	params.Add("fields", "Name")

	fullPath := fmt.Sprintf("%s?%s", query, params.Encode())
	logger.Debug("Sending request for cluster leader", "path", fullPath)

	response, err := c.GetRequest(fullPath)
	if err != nil {
		logger.Error("Failed to get cluster leader response", "error", err)
		return "", err
	}

	var leaderHostName []struct {
		Name string `json:"Name"`
	}
	if err := json.Unmarshal(response, &leaderHostName); err != nil {
		logger.Error("Failed to unmarshal cluster leader response", "error", err)
		return "", err
	}

	if len(leaderHostName) == 0 {
		logger.Error("No cluster leader found in response")
		return "", fmt.Errorf("no cluster leader found")
	}

	logger.Debug("Found cluster leader", "hostname", leaderHostName[0].Name)
	return leaderHostName[0].Name, nil
}
