package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/client"
)

// GetAssetAccounts retrieves a list of asset accounts from Safeguard.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - fields: Filter criteria for the request
//
// Returns:
//   - []AssetAccount: A slice of asset accounts matching the filter criteria
//   - error: An error if the request fails, nil otherwise
func GetAssetAccounts(c *client.SafeguardClient, fields client.Filter) ([]AssetAccount, error) {
	var users []AssetAccount

	query := "AssetAccounts" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &users)
	return users, nil
}

// GetAssetAccount retrieves a specific asset account by ID from Safeguard.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The ID of the asset account to retrieve
//   - fields: Specific fields to include in the response
//
// Returns:
//   - AssetAccount: The requested asset account
//   - error: An error if the request fails, nil otherwise
func GetAssetAccount(c *client.SafeguardClient, id string, fields client.Fields) (AssetAccount, error) {
	var user AssetAccount

	query := fmt.Sprintf("AssetAccounts/%s", id)
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
