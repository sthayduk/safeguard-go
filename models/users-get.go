package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/client"
)

// GetUsers retrieves a list of users from Safeguard.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - fields: Filter criteria for the request
//
// Returns:
//   - []User: A slice of users matching the filter criteria
//   - error: An error if the request fails, nil otherwise
func GetUsers(c *client.SafeguardClient, fields client.Filter) ([]User, error) {
	var users []User

	query := "users" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &users)
	return users, nil
}

// GetUser retrieves details for a specific user by their ID.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The numeric ID of the user to retrieve
//   - fields: Specific fields to include in the response
//
// Returns:
//   - User: The requested user object
//   - error: An error if the request fails, nil otherwise
func GetUser(c *client.SafeguardClient, id string, fields client.Fields) (User, error) {
	var user User

	query := fmt.Sprintf("users/%s", id)
	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return user, err
	}
	json.Unmarshal(response, &user)
	return user, nil
}

// GetLinkedAccounts retrieves the policy accounts linked to a specific user ID.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The numeric ID of the user to get linked accounts for
//
// Returns:
//   - []PolicyAccount: A slice of linked policy accounts
//   - error: An error if the request fails, nil otherwise
func GetLinkedAccounts(c *client.SafeguardClient, id string) ([]PolicyAccount, error) {
	var linkedAccounts []PolicyAccount

	query := fmt.Sprintf("users/%s/LinkedPolicyAccounts", id)

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &linkedAccounts)
	return linkedAccounts, nil
}

// GetLinkedAccounts retrieves the policy accounts linked to this user.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//
// Returns:
//   - []PolicyAccount: A slice of linked policy accounts
//   - error: An error if the request fails, nil otherwise
func (u User) GetLinkedAccounts(c *client.SafeguardClient) ([]PolicyAccount, error) {
	return GetLinkedAccounts(c, fmt.Sprintf("%d", u.Id))
}

// GetUserRoles retrieves the roles assigned to a specific user.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The ID of the user
//
// Returns:
//   - []Role: A slice of assigned roles
//   - error: An error if the request fails, nil otherwise
func GetUserRoles(c *client.SafeguardClient, id string) ([]Role, error) {
	var roles []Role

	query := fmt.Sprintf("users/%s/roles", id)

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &roles)
	return roles, nil
}

// GetRoles retrieves the roles assigned to this user.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//
// Returns:
//   - []Role: A slice of assigned roles
//   - error: An error if the request fails, nil otherwise
func (u User) GetRoles(c *client.SafeguardClient) ([]Role, error) {
	return GetUserRoles(c, fmt.Sprintf("%d", u.Id))
}

// GetGroups retrieves the groups that a specific user belongs to.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The ID of the user
//
// Returns:
//   - []UserGroup: A slice of user groups
//   - error: An error if the request fails, nil otherwise
func GetGroups(c *client.SafeguardClient, id string) ([]UserGroup, error) {
	var userGroups []UserGroup

	query := fmt.Sprintf("users/%s/UserGroups", id)

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &userGroups)
	return userGroups, nil
}

// GetGroups retrieves the groups that this user belongs to.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//
// Returns:
//   - []UserGroup: A slice of user groups
//   - error: An error if the request fails, nil otherwise
func (u User) GetGroups(c *client.SafeguardClient) ([]UserGroup, error) {
	return GetGroups(c, fmt.Sprintf("%d", u.Id))
}
