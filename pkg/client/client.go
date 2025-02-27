package client

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
)

var logger *log.Logger // Declare global logger variable

func init() {
	logger = log.New(os.Stdout, "", log.LstdFlags)
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
	if debug {
		logger.SetFlags(log.LstdFlags | log.Lshortfile)
	} else {
		logger.SetFlags(log.LstdFlags)
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

	return &c
}

func createTLSClient() *http.Client {
	caCert, err := os.ReadFile("server.crt")
	if err != nil {
		logger.Fatalf("Error loading CA certificate: %v", err)
	}
	rootCert, err := os.ReadFile("pam.cer")
	if err != nil {
		logger.Fatalf("Error loading root certificate: %v", err)
	}

	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		logger.Fatalf("Error adding CA certificate to pool")
	}
	if !caCertPool.AppendCertsFromPEM(rootCert) {
		logger.Fatalf("Error adding root certificate to pool")
	}

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}
}
