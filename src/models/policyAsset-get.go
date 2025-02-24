package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/src/client"
)

// GetPolicyAssets retrieves a list of policy assets from the Safeguard API.
// It takes a SafeguardClient and a Filter as parameters to construct the query.
// The function returns a slice of PolicyAsset and an error if the request fails.
//
// Parameters:
//   - c: A pointer to a SafeguardClient used to make the API request.
//   - fields: A Filter object used to specify query parameters.
//
// Returns:
//   - []PolicyAsset: A slice of PolicyAsset objects retrieved from the API.
//   - error: An error object if the request fails or the response cannot be unmarshaled.
func GetPolicyAssets(c *client.SafeguardClient, fields client.Filter) ([]PolicyAsset, error) {
	var policyAssets []PolicyAsset

	query := "PolicyAssets" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &policyAssets); err != nil {
		return nil, err
	}
	for i := range policyAssets {
		policyAssets[i].client = c
	}
	return policyAssets, nil
}

// GetPolicyAsset retrieves a PolicyAsset by its ID from the SafeguardClient.
// It takes a SafeguardClient, an integer ID, and optional fields to include in the query.
// It returns the PolicyAsset and an error if any occurred during the request.
//
// Parameters:
//   - c: A pointer to the SafeguardClient used to make the request.
//   - id: The ID of the PolicyAsset to retrieve.
//   - fields: Optional fields to include in the query.
//
// Returns:
//   - PolicyAsset: The retrieved PolicyAsset.
//   - error: An error if any occurred during the request.
func GetPolicyAsset(c *client.SafeguardClient, id int, fields client.Fields) (PolicyAsset, error) {
	var policyAsset PolicyAsset
	policyAsset.client = c

	query := fmt.Sprintf("PolicyAssets/%d", id)
	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return policyAsset, err
	}
	if err := json.Unmarshal(response, &policyAsset); err != nil {
		return policyAsset, err
	}
	return policyAsset, nil
}

// GetAssetGroups retrieves the asset groups associated with a specific policy asset.
// It takes a SafeguardClient and a Filter as parameters and returns a slice of AssetGroup and an error.
//
// Parameters:
//   - c: A pointer to the SafeguardClient used to make the request.
//   - fields: A Filter object used to specify query parameters.
//
// Returns:
//   - A slice of AssetGroup objects representing the asset groups associated with the policy asset.
//   - An error if the request fails or the response cannot be unmarshaled.
func (p PolicyAsset) GetAssetGroups(fields client.Filter) ([]AssetGroup, error) {
	var assetGroups []AssetGroup

	query := fmt.Sprintf("PolicyAssets/%d/AssetGroups", p.Id) + fields.ToQueryString()

	response, err := p.client.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &assetGroups); err != nil {
		return nil, err
	}
	return assetGroups, nil
}

// GetDirectoryServiceEntries retrieves the directory service entries associated with the PolicyAsset.
// It takes a SafeguardClient and a Filter as parameters and returns a slice of DirectoryServiceEntry and an error.
//
// Parameters:
//   - c: A pointer to the SafeguardClient used to make the request.
//   - fields: A Filter object used to specify query parameters.
//
// Returns:
//   - A slice of DirectoryServiceEntry objects.
//   - An error if the request fails or the response cannot be unmarshaled.
func (p PolicyAsset) GetDirectoryServiceEntries(fields client.Filter) ([]DirectoryServiceEntry, error) {
	var directoryServiceEntries []DirectoryServiceEntry

	query := fmt.Sprintf("PolicyAssets/%d/DirectoryServiceEntries", p.Id) + fields.ToQueryString()

	response, err := p.client.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &directoryServiceEntries); err != nil {
		return nil, err
	}
	return directoryServiceEntries, nil
}

// GetPolicies retrieves the policies associated with a specific PolicyAsset.
// It sends a GET request to the Safeguard API using the provided SafeguardClient and query fields.
//
// Parameters:
//   - c: A pointer to the SafeguardClient used to send the request.
//   - fields: A Filter object containing query parameters to be appended to the request URL.
//
// Returns:
//   - A slice of AssetPolicy objects representing the policies associated with the PolicyAsset.
//   - An error if the request fails or the response cannot be unmarshaled.
func (p PolicyAsset) GetPolicies(fields client.Filter) ([]AssetPolicy, error) {
	var policies []AssetPolicy

	query := fmt.Sprintf("PolicyAssets/%d/Policies", p.Id) + fields.ToQueryString()

	response, err := p.client.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &policies); err != nil {
		return nil, err
	}
	return policies, nil
}
