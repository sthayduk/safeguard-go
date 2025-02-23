package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/client"
)

// GetUserGroups retrieves a list of user groups from Safeguard.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - fields: Filter criteria for the request
//
// Returns:
//   - []UserGroup: A slice of user groups matching the filter criteria
//   - error: An error if the request fails, nil otherwise
func GetUserGroups(c *client.SafeguardClient, fields client.Filter) ([]UserGroup, error) {
	var userGroups []UserGroup

	query := "UserGroups" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &userGroups)
	return userGroups, nil
}

// GetUserGroup retrieves details for a specific user group by ID.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The numeric ID of the user group to retrieve
//   - fields: Specific fields to include in the response
//
// Returns:
//   - UserGroup: The requested user group object
//   - error: An error if the request fails, nil otherwise
func GetUserGroup(c *client.SafeguardClient, id string, fields client.Fields) (UserGroup, error) {
	var userGroup UserGroup

	query := fmt.Sprintf("UserGroups/%s", id)
	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return userGroup, err
	}
	json.Unmarshal(response, &userGroup)
	return userGroup, nil
}
