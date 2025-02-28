package client

import "time"

// RSTSAuthResponse represents the complete authentication response from both RSTS and Safeguard
type RSTSAuthResponse struct {
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
