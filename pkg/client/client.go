package client

import (
	"crypto/tls"
	"crypto/x509"
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

type SafeguardClient struct {
	AccessToken    *RSTSAuthResponse
	ApplicanceURL  string
	ApiVersion     string
	HttpClient     *http.Client
	tokenEndpoint  string
	redirectPort   int
	redirectURI    string
	DefaultHeaders http.Header
}

// New creates a new instance of SafeguardClient.
// Parameters:
// - applianceUrl: The URL of the Safeguard appliance.
// - apiVersion: The API version to use.
// - debug: A boolean flag to enable debug logging.
// Returns a pointer to a SafeguardClient instance.
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
	return &c
}

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

// GetTokenExpirationTime returns the time when the current access token will expire
func (c *SafeguardClient) GetTokenExpirationTime() time.Time {
	return c.AccessToken.AuthTime.Add(time.Duration(c.AccessToken.ExpiresIn) * time.Second)
}

// IsTokenExpired checks if the current access token has expired
func (c *SafeguardClient) IsTokenExpired() bool {
	if c.AccessToken == nil || c.AccessToken.AuthTime.IsZero() {
		return true
	}
	return time.Now().After(c.GetTokenExpirationTime())
}

// RemainingTokenTime returns the duration until the token expires
func (c *SafeguardClient) RemainingTokenTime() time.Duration {
	if c.AccessToken == nil || c.AccessToken.AuthTime.IsZero() {
		return 0
	}
	return time.Until(c.GetTokenExpirationTime())
}
