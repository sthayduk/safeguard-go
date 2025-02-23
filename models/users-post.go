package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/client"
)

// AddLinkedAccounts adds linked policy accounts to a user in Safeguard.
// It takes a SafeguardClient, a User, and a slice of PolicyAccount as parameters.
// It returns a slice of PolicyAccount representing the linked accounts and an error if any occurred.
//
// Parameters:
//   - c: A pointer to a SafeguardClient used to make API requests.
//   - user: A User object representing the user to whom the policy accounts will be linked.
//   - policyAccount: A slice of PolicyAccount objects to be linked to the user.
//
// Returns:
//   - []PolicyAccount: A slice of PolicyAccount objects that were successfully linked to the user.
//   - error: An error object if an error occurred during the operation, otherwise nil.
func AddLinkedAccounts(c *client.SafeguardClient, user User, policyAccount []PolicyAccount) ([]PolicyAccount, error) {
	var linkedAccounts []PolicyAccount

	query := fmt.Sprintf("users/%d/LinkedPolicyAccounts/Add", user.Id)
	response, err := fetchAndPostPolicyAccount(c, policyAccount, query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &linkedAccounts); err != nil {
		return nil, err
	}
	return linkedAccounts, nil
}

func (u User) AddLinkedAccounts(c *client.SafeguardClient, policyAccount []PolicyAccount) ([]PolicyAccount, error) {
	return AddLinkedAccounts(c, u, policyAccount)
}

// RemoveLinkedAccounts removes linked policy accounts from a user in Safeguard.
// It takes a SafeguardClient, a User, and a slice of PolicyAccount as parameters.
// It returns a slice of PolicyAccount representing the linked accounts that were removed, and an error if any occurred.
//
// Parameters:
//   - c: A pointer to a SafeguardClient used to make API requests.
//   - user: A User object representing the user from whom linked accounts will be removed.
//   - policyAccount: A slice of PolicyAccount objects representing the accounts to be removed.
//
// Returns:
//   - []PolicyAccount: A slice of PolicyAccount objects that were removed.
//   - error: An error object if an error occurred, otherwise nil.
func RemoveLinkedAccounts(c *client.SafeguardClient, user User, policyAccount []PolicyAccount) ([]PolicyAccount, error) {
	var linkedAccounts []PolicyAccount

	query := fmt.Sprintf("users/%d/LinkedPolicyAccounts/Remove", user.Id)
	response, err := fetchAndPostPolicyAccount(c, policyAccount, query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &linkedAccounts); err != nil {
		return nil, err
	}
	return linkedAccounts, nil
}

func (u User) RemoveLinkedAccounts(c *client.SafeguardClient, policyAccount []PolicyAccount) ([]PolicyAccount, error) {
	return RemoveLinkedAccounts(c, u, policyAccount)
}

// fetchAndPostPolicyAccount sends a POST request to the Safeguard API with the given policy accounts and query.
// It takes a SafeguardClient, a slice of PolicyAccount, and a query string as parameters.
// It returns the response as a byte slice and an error if any occurred.
//
// Parameters:
//   - c: A pointer to a SafeguardClient used to make API requests.
//   - policyAccount: A slice of PolicyAccount objects to be sent in the request body.
//   - query: A string representing the API endpoint to which the request will be sent.
//
// Returns:
//   - []byte: The response from the API as a byte slice.
//   - error: An error object if an error occurred during the operation, otherwise nil.
func fetchAndPostPolicyAccount(c *client.SafeguardClient, policyAccount []PolicyAccount, query string) ([]byte, error) {
	policyAccountJson, err := json.Marshal(policyAccount)
	if err != nil {
		return nil, err
	}

	return c.PostRequest(query, bytes.NewReader(policyAccountJson))
}
