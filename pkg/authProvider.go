package pkg

import (
	"encoding/json"
	"fmt"
)

type TypeReferenceName string

const (
	Unknown          TypeReferenceName = "Unknown"
	LocalMachine     TypeReferenceName = "LocalMachine"
	Certificate      TypeReferenceName = "Certificate"
	DirectoryAccount TypeReferenceName = "DirectoryAccount"
	ExternalFed      TypeReferenceName = "ExternalFederation"
	Radius           TypeReferenceName = "Radius"
	OneLoginMfa      TypeReferenceName = "OneLoginMfa"
	Fido2            TypeReferenceName = "Fido2"
	Starling         TypeReferenceName = "Starling"
)

type AuthenticationProvider struct {
	Id                 int    `json:"Id,omitempty"`
	Name               string `json:"Name,omitempty"`
	TypeReferenceName  string `json:"TypeReferenceName,omitempty"`
	IdentityProviderId int    `json:"IdentityProviderId,omitempty"`
	Identity           string `json:"Identity"`
	RstsProviderId     string `json:"RstsProviderId,omitempty"`
	RstsProviderScope  string `json:"RstsProviderScope,omitempty"`
	IsDefault          bool   `json:"ForceAsDefault,omitempty"`
}

// ToJson converts an AuthenticationProvider instance to a JSON string.
// This is useful for serializing the provider data for transmission or storage.
//
// Returns:
//   - string: A JSON-encoded string representation of the authentication provider
//   - error: An error if JSON marshaling encounters any issues
func (a AuthenticationProvider) ToJson() (string, error) {
	userJSON, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(userJSON), nil
}

// GetAuthenticationProviders retrieves all authentication providers configured in Safeguard.
// This includes all provider types like LDAP, RADIUS, certificate-based, etc.
//
// Returns:
//   - []AuthenticationProvider: A slice containing all configured authentication providers
//   - error: An error if the API request fails or the response cannot be parsed
func GetAuthenticationProviders() ([]AuthenticationProvider, error) {
	var authProviders []AuthenticationProvider

	query := "AuthenticationProviders"

	response, err := c.GetRequest(query)
	if err != nil {
		return []AuthenticationProvider{}, err
	}

	if err := json.Unmarshal(response, &authProviders); err != nil {
		return []AuthenticationProvider{}, err
	}

	return authProviders, nil
}

// GetAuthenticationProvider retrieves a specific authentication provider by its ID.
// Use this to get detailed information about a single provider configuration.
//
// Parameters:
//   - id: The unique identifier of the authentication provider to retrieve
//
// Returns:
//   - AuthenticationProvider: The requested authentication provider's configuration
//   - error: An error if the provider cannot be found or the request fails
func GetAuthenticationProvider(id int) (AuthenticationProvider, error) {
	var authProvider AuthenticationProvider

	query := fmt.Sprintf("AuthenticationProviders/%d", id)

	response, err := c.GetRequest(query)
	if err != nil {
		return AuthenticationProvider{}, err
	}

	if err := json.Unmarshal(response, &authProvider); err != nil {
		return AuthenticationProvider{}, err
	}
	return authProvider, nil
}

// ClearDefaultAuthProvider removes the current default authentication provider setting.
// After calling this, no authentication provider will be marked as default.
//
// Returns:
//   - error: An error if the operation fails or the API request is unsuccessful
func ClearDefaultAuthProvider() error {
	query := "AuthenticationProviders/ClearDefault"

	_, err := c.PostRequest(query, nil)
	if err != nil {
		return err
	}

	return nil
}

// ForceAsDefaultAuthProvider sets a specific authentication provider as the system default.
// Only one provider can be the default at any time.
//
// Parameters:
//   - id: The unique identifier of the authentication provider to set as default
//
// Returns:
//   - AuthenticationProvider: The updated authentication provider configuration
//   - error: An error if the operation fails or the provider cannot be found
func ForceAsDefaultAuthProvider(id int) (AuthenticationProvider, error) {
	var authProvider AuthenticationProvider
	query := fmt.Sprintf("AuthenticationProviders/%d/ForceAsDefault", id)

	response, err := c.PostRequest(query, nil)
	if err != nil {
		return AuthenticationProvider{}, err
	}

	err = json.Unmarshal(response, &authProvider)
	if err != nil {
		return AuthenticationProvider{}, err
	}

	return authProvider, nil
}

// ForceAsDefault marks this authentication provider instance as the system default.
// This is a convenience method that calls ForceAsDefaultAuthProvider with this instance's ID.
//
// Returns:
//   - AuthenticationProvider: The updated authentication provider configuration
//   - error: An error if the operation fails or the API request is unsuccessful
func (a AuthenticationProvider) ForceAsDefault() (AuthenticationProvider, error) {
	return ForceAsDefaultAuthProvider(a.Id)
}
