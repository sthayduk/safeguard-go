package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/client"
)

// GetPolicyAccounts retrieves a list of policy accounts from the Safeguard API.
// It takes a SafeguardClient and a Filter as parameters, constructs a query string,
// sends a GET request to the Safeguard API, and unmarshals the response into a slice of PolicyAccount.
//
// Parameters:
//   - c: A pointer to a SafeguardClient used to send requests to the Safeguard API.
//   - fields: A Filter object used to construct the query string for filtering the results.
//
// Returns:
//   - A slice of PolicyAccount objects.
//   - An error if the request fails or the response cannot be unmarshaled.
func GetPolicyAccounts(c *client.SafeguardClient, fields client.Filter) ([]PolicyAccount, error) {
	var policyAccounts []PolicyAccount

	query := "PolicyAccounts" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &policyAccounts)
	return policyAccounts, nil
}

// GetPolicyAccount retrieves a PolicyAccount by its ID from the SafeguardClient.
// It takes a SafeguardClient, an integer ID, and optional fields to include in the query.
// It returns the PolicyAccount and an error if the request fails.
//
// Parameters:
//   - c: A pointer to the SafeguardClient used to make the request.
//   - id: The ID of the PolicyAccount to retrieve.
//   - fields: Optional fields to include in the query.
//
// Returns:
//   - PolicyAccount: The retrieved PolicyAccount.
//   - error: An error if the request fails or the response cannot be unmarshaled.
func GetPolicyAccount(c *client.SafeguardClient, id int, fields client.Fields) (PolicyAccount, error) {
	var policyAccount PolicyAccount

	query := fmt.Sprintf("PolicyAccounts/%d", id)
	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return policyAccount, err
	}
	json.Unmarshal(response, &policyAccount)
	return policyAccount, nil
}
