package client

import (
	"net/http"
	"sync"
	"time"
)

type SafeguardClient struct {
	AccessToken    *RSTSAuthResponse
	Appliance      applianceURL
	ClusterLeader  applianceURL
	ApiVersion     string
	HttpClient     *http.Client
	tokenEndpoint  string
	redirectPort   int
	redirectURI    string
	DefaultHeaders http.Header
}

type applianceURL struct {
	sync.RWMutex

	Protocol   string
	Hostname   string
	DomainName string
	Port       string
	Url        string
}

func (a *applianceURL) getUrl() string {
	a.RWMutex.RLock()
	defer a.RWMutex.RUnlock()
	return a.Url
}

func (a *applianceURL) setUrl(url string) {
	a.RWMutex.Lock()
	defer a.RWMutex.Unlock()
	a.Url = url
}

// RSTSAuthResponse represents the complete authentication response from both RSTS and Safeguard
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

func (a *RSTSAuthResponse) getAccessToken() string {
	a.RWMutex.RLock()
	defer a.RWMutex.RUnlock()
	return a.AccessToken
}

func (a *RSTSAuthResponse) setAccessToken(accessToken string) {
	a.RWMutex.Lock()
	defer a.RWMutex.Unlock()
	a.AccessToken = accessToken
}

func (a *RSTSAuthResponse) setUserNamePassword(username, password string) {
	a.RWMutex.Lock()
	defer a.RWMutex.Unlock()
	a.credentials.username = username
	a.credentials.password = password
}

func (a *RSTSAuthResponse) getUserNamePassword() (string, string) {
	a.RWMutex.RLock()
	defer a.RWMutex.RUnlock()
	return a.credentials.username, a.credentials.password
}

func (a *RSTSAuthResponse) setCertificate(certPath, certPassword string) {
	a.RWMutex.Lock()
	defer a.RWMutex.Unlock()
	a.credentials.certPath = certPath
	a.credentials.certPassword = certPassword
}

func (a *RSTSAuthResponse) getCertificate() (string, string) {
	a.RWMutex.RLock()
	defer a.RWMutex.RUnlock()
	return a.credentials.certPath, a.credentials.certPassword
}

func (a *RSTSAuthResponse) setUserToken(userToken string) {
	a.RWMutex.Lock()
	defer a.RWMutex.Unlock()
	a.UserToken = userToken
}

type Credentials struct {
	username     string
	password     string
	certPath     string
	certPassword string
}

// AuthProvider represents the type of authentication provider
type AuthProvider string

const (
	// AuthProviderCertificate represents certificate-based authentication
	AuthProviderCertificate AuthProvider = "rsts:sts:primaryproviderid:certificate"
	// AuthProviderLocal represents local username/password authentication
	AuthProviderLocal AuthProvider = "rsts:sts:primaryproviderid:local"
)

// String returns the string representation of the AuthProvider
func (a AuthProvider) String() string {
	return string(a)
}
