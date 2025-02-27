package client

import (
	"crypto/tls"
	"crypto/x509"
	"log/slog"
	"net/http"
	"os"
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
	AccessToken    *TokenResponse
	ApplicanceURL  string
	ApiVersion     string
	HttpClient     *http.Client
	tokenEndpoint  string
	redirectPort   int
	redirectURI    string
	DefaultHeaders http.Header
}

type TokenResponse struct {
	AccessToken       string `json:"access_token"`
	AuthorizationCode string `json:"authorization_code"`
	TokenType         string `json:"token_type,omitempty"`
	ExpiresIn         int    `json:"expires_in,omitempty"`
	UserToken         string `json:"UserToken,omitempty"`
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
		AccessToken:   &TokenResponse{},
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
