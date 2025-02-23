package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/client"
)

// GetRoles retrieves a list of roles from Safeguard.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - fields: Filter criteria for the request
//
// Returns:
//   - []Role: A slice of roles matching the filter criteria
//   - error: An error if the request fails, nil otherwise
func GetRoles(c *client.SafeguardClient, fields client.Filter) ([]Role, error) {
	var userRoles []Role

	query := "Roles" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &userRoles)
	return userRoles, nil
}

// GetRole retrieves details for a specific role by ID.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The numeric ID of the role to retrieve
//   - fields: Specific fields to include in the response
//
// Returns:
//   - Role: The requested role object
//   - error: An error if the request fails, nil otherwise
func GetRole(c *client.SafeguardClient, id string, fields client.Fields) (Role, error) {
	var userRole Role

	query := fmt.Sprintf("Roles/%s", id)
	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return userRole, err
	}
	json.Unmarshal(response, &userRole)
	return userRole, nil
}
