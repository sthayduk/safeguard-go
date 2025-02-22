package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/client"
)

// GetAssetPartitions retrieves a list of asset partitions from Safeguard.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - fields: Filter criteria for the request
//
// Returns:
//   - []AssetPartition: A slice of asset partitions matching the filter criteria
//   - error: An error if the request fails, nil otherwise
func GetAssetPartitions(c *client.SafeguardClient, fields client.Filter) ([]AssetPartition, error) {
	var AssetPartitions []AssetPartition

	query := "AssetPartitions" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &AssetPartitions)
	return AssetPartitions, nil
}

// GetAssetPartition retrieves a specific asset partition by ID from Safeguard.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The ID of the asset partition to retrieve
//   - fields: Specific fields to include in the response
//
// Returns:
//   - AssetPartition: The requested asset partition
//   - error: An error if the request fails, nil otherwise
func GetAssetPartition(c *client.SafeguardClient, id string, fields client.Fields) (AssetPartition, error) {
	var AssetPartition AssetPartition

	query := fmt.Sprintf("AssetPartitions/%s", id)
	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return AssetPartition, err
	}
	json.Unmarshal(response, &AssetPartition)
	return AssetPartition, nil
}

// GetPasswordRules retrieves the password rules for a specific asset partition.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - AssetPartitionId: The ID of the asset partition
//   - filter: Filter criteria for the request
//
// Returns:
//   - []AccountPasswordRule: A slice of password rules
//   - error: An error if the request fails, nil otherwise
func GetPasswordRules(c *client.SafeguardClient, AssetPartitionId int, filter client.Filter) ([]AccountPasswordRule, error) {
	var PasswordRules []AccountPasswordRule

	query := fmt.Sprintf("AssetPartitions/%d/PasswordRules", AssetPartitionId) + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &PasswordRules)
	return PasswordRules, nil
}

// GetPasswordRules retrieves the password rules for this asset partition.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//
// Returns:
//   - []AccountPasswordRule: A slice of password rules
//   - error: An error if the request fails, nil otherwise
func (a AssetPartition) GetPasswordRules(c *client.SafeguardClient) ([]AccountPasswordRule, error) {
	return GetPasswordRules(c, a.Id, client.Filter{})
}
