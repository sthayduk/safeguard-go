package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/src/client"
)

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
