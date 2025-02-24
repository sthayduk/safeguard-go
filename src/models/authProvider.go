package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/src/client"
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
	client *client.SafeguardClient

	Id                 int    `json:"Id,omitempty"`
	Name               string `json:"Name,omitempty"`
	TypeReferenceName  string `json:"TypeReferenceName,omitempty"`
	IdentityProviderId int    `json:"IdentityProviderId,omitempty"`
	Identity           string `json:"Identity"`
	RstsProviderId     string `json:"RstsProviderId,omitempty"`
	RstsProviderScope  string `json:"RstsProviderScope,omitempty"`
	IsDefault          bool   `json:"ForceAsDefault,omitempty"`
}

func (a AuthenticationProvider) ToJson() (string, error) {
	userJSON, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(userJSON), nil
}

// GetAuthenticationProviders retrieves a list of authentication providers from the Safeguard API.
// It sends a GET request to the "AuthenticationProviders" endpoint and unmarshals the response
// into a slice of AuthenticationProvider structs. Each AuthenticationProvider is then associated
// with the provided SafeguardClient.
//
// Parameters:
//   - client: A pointer to a SafeguardClient instance used to make the API request.
//
// Returns:
//   - A slice of AuthenticationProvider structs.
//   - An error if the request fails or the response cannot be unmarshaled.
func GetAuthenticationProviders(c *client.SafeguardClient) ([]AuthenticationProvider, error) {
	var authProviders []AuthenticationProvider

	query := "AuthenticationProviders"

	response, err := c.GetRequest(query)
	if err != nil {
		return []AuthenticationProvider{}, err
	}

	if err := json.Unmarshal(response, &authProviders); err != nil {
		return []AuthenticationProvider{}, err
	}

	for i := range authProviders {
		authProviders[i].client = c
	}

	return authProviders, nil
}

// GetAuthenticationProvider retrieves an authentication provider by its ID.
//
// Parameters:
//   - client: A pointer to the SafeguardClient used to make the request.
//   - id: The ID of the authentication provider to retrieve.
//
// Returns:
//   - AuthenticationProvider: The retrieved authentication provider.
//   - error: An error object if an error occurred during the request, otherwise nil.
func GetAuthenticationProvider(c *client.SafeguardClient, id int) (AuthenticationProvider, error) {
	var authProvider AuthenticationProvider
	authProvider.client = c

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

// ClearDefaultAuthProvider clears the default authentication provider in the Safeguard system.
// It sends a POST request to the "AuthenticationProviders/ClearDefault" endpoint.
//
// Parameters:
//   - c: A pointer to a SafeguardClient instance.
//
// Returns:
//   - error: An error object if the request fails, otherwise nil.
func ClearDefaultAuthProvider(c *client.SafeguardClient) error {
	query := "AuthenticationProviders/ClearDefault"

	_, err := c.PostRequest(query, nil)
	if err != nil {
		return err
	}

	return nil
}

// ForceAsDefaultAuthProvider sets the specified authentication provider as the default one.
//
// Parameters:
//   - c: A pointer to the SafeguardClient instance used to make the request.
//   - id: The ID of the authentication provider to be set as default.
//
// Returns:
//   - AuthenticationProvider: The updated authentication provider object.
//   - error: An error object if an error occurred, otherwise nil.
func ForceAsDefaultAuthProvider(c *client.SafeguardClient, id int) (AuthenticationProvider, error) {
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

	authProvider.client = c
	return authProvider, nil
}

// ForceAsDefault sets the current AuthenticationProvider instance as the default authentication provider.
// It returns the updated AuthenticationProvider instance and an error if the operation fails.
func (a AuthenticationProvider) ForceAsDefault() (AuthenticationProvider, error) {
	return ForceAsDefaultAuthProvider(a.client, a.Id)
}
