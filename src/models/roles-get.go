package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/src/client"
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

	if err := json.Unmarshal(response, &userRoles); err != nil {
		return nil, err
	}
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
func GetRole(c *client.SafeguardClient, id int, fields client.Fields) (Role, error) {
	var userRole Role

	query := fmt.Sprintf("Roles/%d", id)
	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return userRole, err
	}
	if err := json.Unmarshal(response, &userRole); err != nil {
		return userRole, err
	}
	return userRole, nil
}
