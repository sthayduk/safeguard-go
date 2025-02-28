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
	UserToken         string    `json:"UserToken"`
	Status            string    `json:"Status"`
	IdentityProvider  string    `json:"IdentityProvider"`
	AuthorizationCode string    `json:"-"` // Used internally for OAuth flow
	AuthTime          time.Time `json:"-"` // Time when the token was received
}
