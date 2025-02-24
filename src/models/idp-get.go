package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/src/client"
)

// GetIdentityProviders retrieves a list of identity providers from the Safeguard API.
// It sends a GET request to the "IdentityProviders" endpoint and unmarshals the response
// into a slice of IdentityProvider structs. Each IdentityProvider is then associated
// with the provided SafeguardClient.
//
// Parameters:
//   - client: A pointer to a SafeguardClient instance used to make the API request.
//
// Returns:
//   - A slice of IdentityProvider structs.
//   - An error if the request fails or the response cannot be unmarshaled.
func GetIdentityProviders(c *client.SafeguardClient) ([]IdentityProvider, error) {
	var identityProviders []IdentityProvider

	query := "IdentityProviders"

	response, err := c.GetRequest(query)
	if err != nil {
		return []IdentityProvider{}, err
	}

	if err := json.Unmarshal(response, &identityProviders); err != nil {
		return []IdentityProvider{}, err
	}

	for i := range identityProviders {
		identityProviders[i].client = c
	}

	return identityProviders, nil
}

// GetIdentityProvider retrieves an IdentityProvider by its ID.
//
// Parameters:
//   - client: A pointer to the SafeguardClient used to make the request.
//   - id: The ID of the IdentityProvider to retrieve.
//
// Returns:
//   - IdentityProvider: The retrieved IdentityProvider object.
//   - error: An error object if an error occurred during the request, otherwise nil.
func GetIdentityProvider(c *client.SafeguardClient, id int) (IdentityProvider, error) {
	var identityProvider IdentityProvider
	identityProvider.client = c

	query := fmt.Sprintf("IdentityProviders/%d", id)

	response, err := c.GetRequest(query)
	if err != nil {
		return IdentityProvider{}, err
	}

	if err := json.Unmarshal(response, &identityProvider); err != nil {
		return IdentityProvider{}, err
	}
	return identityProvider, nil
}

// GetDirectoryUsers retrieves a list of directory users from the specified identity provider.
// It sends a GET request to the Safeguard API and unmarshals the response into a slice of User objects.
//
// Parameters:
//   - c: A pointer to the SafeguardClient used to send the request.
//   - id: The ID of the identity provider from which to retrieve directory users.
//   - filter: A Filter object used to apply query parameters to the request.
//
// Returns:
//   - A slice of User objects representing the directory users.
//   - An error if the request fails or the response cannot be unmarshaled.
func GetDirectoryUsers(c *client.SafeguardClient, id int, filter client.Filter) ([]User, error) {
	var directoryUsers []User

	query := fmt.Sprintf("IdentityProviders/%d/DirectoryUsers%s", id, filter.ToQueryString())

	response, err := c.GetRequest(query)
	if err != nil {
		return []User{}, err
	}

	if err := json.Unmarshal(response, &directoryUsers); err != nil {
		return []User{}, err
	}

	for i := range directoryUsers {
		directoryUsers[i].client = c
	}

	return directoryUsers, nil
}

// GetDirectoryUsers retrieves a list of users from the directory associated with the IdentityProvider.
// It accepts a filter parameter to narrow down the search results.
//
// Parameters:
//   - filter: A client.Filter object to specify the criteria for filtering the users.
//
// Returns:
//   - []User: A slice of User objects that match the filter criteria.
//   - error: An error object if there is any issue during the retrieval process.
func (idp IdentityProvider) GetDirectoryUsers(filter client.Filter) ([]User, error) {
	return GetDirectoryUsers(idp.client, idp.Id, filter)
}

// GetDirectoryGroups retrieves the directory groups associated with a specific identity provider.
// It sends a GET request to the Safeguard API and unmarshals the response into a slice of UserGroup.
//
// Parameters:
//   - c: A pointer to a SafeguardClient instance used to send the request.
//   - id: The ID of the identity provider.
//   - filter: A Filter instance used to filter the results.
//
// Returns:
//   - A slice of UserGroup containing the directory groups.
//   - An error if the request fails or the response cannot be unmarshaled.
func GetDirectoryGroups(c *client.SafeguardClient, id int, filter client.Filter) ([]UserGroup, error) {
	var directoryGroups []UserGroup

	query := fmt.Sprintf("IdentityProviders/%d/DirectoryGroups%s", id, filter.ToQueryString())

	response, err := c.GetRequest(query)
	if err != nil {
		return []UserGroup{}, err
	}

	if err := json.Unmarshal(response, &directoryGroups); err != nil {
		return []UserGroup{}, err
	}

	for i := range directoryGroups {
		directoryGroups[i].client = c
	}

	return directoryGroups, nil
}

// GetDirectoryGroups retrieves the directory groups associated with the IdentityProvider.
// It takes a filter parameter to apply specific filtering criteria and returns a slice of UserGroup
// and an error if any occurs during the retrieval process.
//
// Parameters:
//   - filter: A client.Filter object to specify the filtering criteria.
//
// Returns:
//   - []UserGroup: A slice of UserGroup objects that match the filtering criteria.
//   - error: An error object if any issues occur during the retrieval process.
func (idp IdentityProvider) GetDirectoryGroups(filter client.Filter) ([]UserGroup, error) {
	return GetDirectoryGroups(idp.client, idp.Id, filter)
}
