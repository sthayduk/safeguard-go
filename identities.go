package safeguard

import (
	"encoding/json"
	"fmt"
)

// Identity represents the user or identity that manages an asset account,
// including their display name, identity provider details, and contact information.
type Identity struct {
	apiClient *SafeguardClient `json:"-"`

	DisplayName                       string `json:"DisplayName,omitempty"`
	Id                                int    `json:"Id,omitempty"`
	IdentityProviderId                int    `json:"IdentityProviderId,omitempty"`
	IdentityProviderName              string `json:"IdentityProviderName,omitempty"`
	IdentityProviderTypeReferenceName string `json:"IdentityProviderTypeReferenceName,omitempty"`
	IsSystemOwned                     bool   `json:"IsSystemOwned,omitempty"`
	Name                              string `json:"Name,omitempty"`
	PrincipalKind                     string `json:"PrincipalKind,omitempty"`
	EmailAddress                      string `json:"EmailAddress,omitempty"`
	DomainName                        string `json:"DomainName,omitempty"`
	FullDisplayName                   string `json:"FullDisplayName,omitempty"`
}

func (a Identity) SetClient(c *SafeguardClient) any {
	a.apiClient = c
	return a
}

func (c *SafeguardClient) GetIdentities(filter Filter) ([]Identity, error) {

	var identities []Identity

	query := "Identities" + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &identities); err != nil {
		return nil, err
	}

	return addClientToSlice(c, identities), nil

}

func (c *SafeguardClient) GetIdentity(id int, fields Fields) (Identity, error) {
	var identity Identity

	query := fmt.Sprintf("Identities/%d", id)

	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return identity, err
	}

	if err := json.Unmarshal(response, &identity); err != nil {
		return identity, err
	}

	return addClient(c, identity), nil
}

// GetIdentityProvider retrieves the identity provider associated with the given identity.
// It takes a Fields parameter which can be used to specify additional fields to include in the query.
// The function constructs a query string based on the identity's ID and the provided fields,
// sends a GET request to the API, and unmarshals the response into an IdentityProvider object.
// If successful, it returns the IdentityProvider object with the API client added; otherwise, it returns an error.
//
// Parameters:
//   - fields: Fields specifying additional fields to include in the query.
//
// Returns:
//   - IdentityProvider: The identity provider associated with the identity.
//   - error: An error if the request or unmarshalling fails.
func (i Identity) GetIdentityProvider(fields Fields) (IdentityProvider, error) {
	var identityProviders IdentityProvider

	query := fmt.Sprintf("Identities/%d/IdentityProvider", i.Id)

	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := i.apiClient.GetRequest(query)
	if err != nil {
		return identityProviders, err
	}

	if err := json.Unmarshal(response, &identityProviders); err != nil {
		return identityProviders, err
	}

	return addClient(i.apiClient, identityProviders), nil
}

// GetUser retrieves a User associated with the Identity.
// It takes a Fields parameter to specify which fields to include in the query.
// The function constructs a query string using the Identity's Id and the provided fields,
// then sends a GET request to the API client.
// If the request is successful, it unmarshals the response into a User object and returns it.
// If any error occurs during the request or unmarshalling, it returns the error.
//
// Parameters:
//   - fields: Fields specifying which fields to include in the query.
//
// Returns:
//   - User: The User associated with the Identity.
//   - error: An error if the request or unmarshalling fails.
func (i Identity) GetUser(fields Fields) (User, error) {
	var user User

	query := fmt.Sprintf("Identities/%d/User", i.Id)

	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := i.apiClient.GetRequest(query)
	if err != nil {
		return user, err
	}

	if err := json.Unmarshal(response, &user); err != nil {
		return user, err
	}

	return addClient(i.apiClient, user), nil

}

// GetUserGroup retrieves the UserGroup associated with the Identity.
// It takes a Fields parameter which specifies the fields to be included in the query.
// If the fields parameter is not empty, it appends the fields as a query string to the request URL.
// It returns the UserGroup and an error if any occurred during the request or unmarshalling of the response.
//
// Parameters:
//   - fields: Fields specifying the fields to be included in the query.
//
// Returns:
//   - UserGroup: The UserGroup associated with the Identity.
//   - error: An error if any occurred during the request or unmarshalling of the response.
func (i Identity) GetUserGroup(fields Fields) (UserGroup, error) {
	var userGroup UserGroup

	query := fmt.Sprintf("Identities/%d/UserGroup", i.Id)

	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := i.apiClient.GetRequest(query)
	if err != nil {
		return userGroup, err
	}

	if err := json.Unmarshal(response, &userGroup); err != nil {
		return userGroup, err
	}

	return addClient(i.apiClient, userGroup), nil

}
