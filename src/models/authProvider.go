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
	ForceAsDefault     bool   `json:"ForceAsDefault,omitempty"`
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
