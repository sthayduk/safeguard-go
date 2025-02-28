package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log/slog"
	"net/http"
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

	c := SafeguardClient{
		AccessToken:   &RSTSAuthResponse{},
		ApiVersion:    apiVersion,
		ApplicanceURL: applianceUrl,
		HttpClient:    createTLSClient(),
		redirectPort:  redirectPort,
		redirectURI:   redirectURI,
		tokenEndpoint: applianceUrl + "/service/core/v4/Token/LoginResponse",
	}

	sgclient = &c

	ctx := context.Background()
	go c.refreshToken(ctx)

	return &c
}

// GetClusterLeaderUrl returns the URL of the cluster leader.
// This URL is used to identify the leader node in a cluster setup.
func (c *SafeguardClient) GetClusterLeaderUrl() string {
	return c.ClusterLeaderUrl
}

// SetClusterLeader sets the cluster leader URL for the SafeguardClient.
// It constructs the URL based on the provided cluster leader host name and the
// existing appliance URL components (protocol, domain name, and port).
//
// Parameters:
//   - clusterLeaderHostName: The host name of the cluster leader.
//
// If the constructed cluster leader URL is the same as the current appliance URL,
// it sets the ClusterLeaderUrl to the appliance URL and logs a debug message.
// Otherwise, it updates the ClusterLeaderUrl with the new cluster leader URL and logs
// the change.
//
// Logs:
//   - Error: If there is an error splitting the appliance URL.
//   - Debug: If the cluster leader is the same as the appliance URL or if the cluster leader is set successfully.
func (c *SafeguardClient) SetClusterLeader(clusterLeaderHostName string) {
	protocol, _, domainName, port, err := sgclient.splitApplianceURL()
	if err != nil {
		logger.Error("Error splitting appliance URL", "error", err)
		return
	}

	var clusterLeaderUrl string
	if domainName == "" {
		clusterLeaderUrl = fmt.Sprintf("%s://%s:%s", protocol, clusterLeaderHostName, port)
	} else {
		clusterLeaderUrl = fmt.Sprintf("%s://%s.%s:%s", protocol, clusterLeaderHostName, domainName, port)
	}

	if sgclient.ApplicanceURL == clusterLeaderUrl {
		logger.Debug("Cluster leader is the same as appliance URL")
		sgclient.ClusterLeaderUrl = sgclient.ApplicanceURL
		return
	}

	sgclient.ClusterLeaderUrl = clusterLeaderUrl
	logger.Debug("Cluster leader set to:", "url", clusterLeaderUrl)
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
				RootCAs: caCertPool,
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
func (c SafeguardClient) refreshToken(ctx context.Context) {
	if c.AccessToken.AuthProvider == "" {
		logger.Debug("token refresh skipped: no auth provider")
		return
	}

	// Wait until Authentication is done
	if c.AccessToken.AuthTime.IsZero() {
		logger.Debug("wait until authentication is done")

		for {
			if !c.AccessToken.AuthTime.IsZero() {
				break
			}
			time.Sleep(1 * time.Second)
		}
	}

	remainingTokenTime := c.RemainingTokenTime()
	ticker := time.NewTicker(remainingTokenTime - 1*time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			if c.AccessToken.AuthProvider == AuthProviderLocal {
				c.LoginWithPassword(c.AccessToken.credentials.username, c.AccessToken.credentials.password)
			} else if c.AccessToken.AuthProvider == AuthProviderCertificate {
				c.LoginWithCertificate(c.AccessToken.credentials.certPath, c.AccessToken.credentials.certPassword)
			}
		}
	}
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
