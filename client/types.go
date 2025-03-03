package client

import (
	"log/slog"
	"net/http"
	"sync"
	"time"
)

// SafeguardClient represents the main client for interacting with the Safeguard API.
// It handles authentication, request routing, and session management.
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
	authDone       chan string
	Logger         *slog.Logger
}

// applianceURL represents a Safeguard appliance URL with thread-safe access
// and caching capabilities. It maintains the components of the URL and handles
// cache expiration for URL refreshing.
type applianceURL struct {
	sync.RWMutex

	Protocol   string
	Hostname   string
	DomainName string
	Port       string
	Url        string

	lastUpdate time.Time
	cacheTime  time.Duration
}

// isExpired checks if the cached URL has exceeded its cache duration.
// Returns true if the cache has expired or if cacheTime is 0.
// Returns false if cacheTime is -1 (infinite cache).
func (a *applianceURL) isExpired() bool {
	a.RWMutex.RLock()
	defer a.RWMutex.RUnlock()

	if a.cacheTime == 0 {
		return true
	}

	if a.cacheTime == -1 {
		return false
	}

	return time.Since(a.lastUpdate) > a.cacheTime
}

// getUrl returns the current appliance URL in a thread-safe manner.
func (a *applianceURL) getUrl() string {
	a.RWMutex.RLock()
	defer a.RWMutex.RUnlock()
	return a.Url
}

// setUrl updates the appliance URL and its components with thread safety.
// It parses the URL into its components and updates the cache timestamp.
//
// Parameters:
//   - url: The complete URL string to set
//   - cacheTime: Duration for which the URL should be cached. Use -1 for infinite cache.
func (a *applianceURL) setUrl(url string, cacheTime time.Duration) {
	a.RWMutex.Lock()
	defer a.RWMutex.Unlock()

	var err error
	a.Protocol, a.Hostname, a.DomainName, a.Port, err = splitApplianceURL(url)
	if err != nil {
		logger.Error("Failed to split appliance URL", "error", err)
		return
	}

	a.Url = url
	a.lastUpdate = time.Now()
	a.cacheTime = cacheTime
}

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

// AuthProvider represents the supported authentication provider types.
type AuthProvider string

// AuthProvider constants define the supported authentication methods.
const (
	// AuthProviderCertificate represents certificate-based authentication
	AuthProviderCertificate AuthProvider = "rsts:sts:primaryproviderid:certificate"
	// AuthProviderLocal represents local username/password authentication
	AuthProviderLocal AuthProvider = "rsts:sts:primaryproviderid:local"
)

// String returns the string representation of the AuthProvider.
//
// Returns:
//   - string: The provider identifier string used in authentication requests.
func (a AuthProvider) String() string {
	return string(a)
}

// AccessRequestType represents the type of access request
type AccessRequestType string

// AccessRequestType constants define the supported types of access requests
const (
	AccessRequestTypePassword              AccessRequestType = "Password"
	AccessRequestTypeSsh                   AccessRequestType = "Ssh"
	AccessRequestTypeRemoteDesktop         AccessRequestType = "RemoteDesktop"
	AccessRequestTypeRemoteDesktopManager  AccessRequestType = "RemoteDesktopManager"
	AccessRequestTypeRemoteSsh             AccessRequestType = "RemoteSsh"
	AccessRequestTypeRemoteSshPrivateKey   AccessRequestType = "RemoteSshPrivateKey"
	AccessRequestTypeRemoteDesktopAccounts AccessRequestType = "RemoteDesktopAccounts"
	AccessRequestTypeRemoteDesktopServices AccessRequestType = "RemoteDesktopServices"
	AccessRequestTypeApi                   AccessRequestType = "Api"
	AccessRequestTypeApiExternal           AccessRequestType = "ApiExternal"
	AccessRequestTypeRemoteProcess         AccessRequestType = "RemoteProcess"
	AccessRequestTypeResourceCertificate   AccessRequestType = "ResourceCertificate"
	AccessRequestTypeFile                  AccessRequestType = "File"
	AccessRequestTypeSshKey                AccessRequestType = "SshKey"
	AccessRequestTypeTotp                  AccessRequestType = "Totp"
)

// String returns the string representation of the AccessRequestType
func (a AccessRequestType) String() string {
	return string(a)
}

// SignalREvent represents the root structure of a SignalR event notification
type SignalREvent struct {
	ApplianceId string    `json:"ApplianceId"`
	Name        string    `json:"Name"`
	Time        time.Time `json:"Time"`
	Message     string    `json:"Message"`
	AuditLogUri *string   `json:"AuditLogUri"`
	Data        EventData `json:"Data"`
}

// EventData represents the Data field of a SignalR event
type EventData struct {
	AccessRequestType           AccessRequestType `json:"AccessRequestType"`
	AccountDistinguishedName    string            `json:"AccountDistinguishedName"`
	AccountDomainName           string            `json:"AccountDomainName"`
	AccountHasTotpAuthenticator bool              `json:"AccountHasTotpAuthenticator"`
	AccountId                   int               `json:"AccountId"`
	AccountName                 string            `json:"AccountName"`
	ActionUserIds               []int             `json:"ActionUserIds"`
	ApproverAccessRequestUri    string            `json:"ApproverAccessRequestUri"`
	AssetId                     int               `json:"AssetId"`
	AssetName                   string            `json:"AssetName"`
	AssetNetworkAddress         string            `json:"AssetNetworkAddress"`
	AssetPlatformType           string            `json:"AssetPlatformType"`
	Comment                     *string           `json:"Comment"`
	DurationInMinutes           int               `json:"DurationInMinutes"`
	OfflineWorkflowMode         bool              `json:"OfflineWorkflowMode"`
	Reason                      *string           `json:"Reason"`
	ReasonCode                  *string           `json:"ReasonCode"`
	Requester                   string            `json:"Requester"`
	RequesterAccessRequestUri   string            `json:"RequesterAccessRequestUri"`
	RequesterId                 int               `json:"RequesterId"`
	RequesterUsername           string            `json:"RequesterUsername"`
	RequestId                   string            `json:"RequestId"`
	RequiredDate                time.Time         `json:"RequiredDate"`
	ReviewerAccessRequestUri    string            `json:"ReviewerAccessRequestUri"`
	SessionSpsNodeIpAddress     *string           `json:"SessionSpsNodeIpAddress"`
	TicketNumber                *string           `json:"TicketNumber"`
	WasCheckedOut               bool              `json:"WasCheckedOut"`
	EventName                   string            `json:"EventName"`
	EventTimestamp              time.Time         `json:"EventTimestamp"`
	ApplianceId                 string            `json:"ApplianceId"`
	EventUserId                 int               `json:"EventUserId"`
	EventUserDisplayName        string            `json:"EventUserDisplayName"`
	EventUserName               string            `json:"EventUserName"`
	EventUserDomainName         *string           `json:"EventUserDomainName"`
	AuditLogUri                 *string           `json:"AuditLogUri"`
	EventDisplayName            string            `json:"EventDisplayName"`
	EventDescription            string            `json:"EventDescription"`
}
