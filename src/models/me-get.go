package models

import (
	"encoding/json"

	"github.com/sthayduk/safeguard-go/src/client"
)

// GetMe retrieves information about the currently authenticated user.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - fields: Filter criteria for the request
//
// Returns:
//   - User: The user information for the authenticated user
//   - error: An error if the request fails, nil otherwise
func GetMe(c *client.SafeguardClient, fields client.Filter) (User, error) {
	query := "me" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return User{}, err
	}

	var me User
	if err := json.Unmarshal(response, &me); err != nil {
		return User{}, err
	}

	return me, nil
}
