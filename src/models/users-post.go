package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/src/client"
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

// CreateUser creates a new user in the Safeguard system.
// It takes a SafeguardClient, a User object, an IdentityProvider, and an AuthenticationProvider as parameters.
// It returns the created User object and an error if any occurred during the process.
//
// Parameters:
//   - c: A pointer to the SafeguardClient used to make the request.
//   - user: The User object containing the details of the user to be created.
//   - identityProvider: The IdentityProvider associated with the user.
//   - authProvider: The AuthenticationProvider associated with the user.
//
// Returns:
//   - User: The created User object.
//   - error: An error if any occurred during the user creation process.
func CreateUser(c *client.SafeguardClient, user User) (User, error) {
	var createdUser User

	userJson, err := json.Marshal(user)
	if err != nil {
		return createdUser, err
	}

	response, err := c.PostRequest("users", bytes.NewReader(userJson))
	if err != nil {
		return createdUser, err
	}

	if err := json.Unmarshal(response, &createdUser); err != nil {
		return createdUser, err
	}

	createdUser.client = c
	return createdUser, nil
}

func (u User) SetAuthenticationProvider(c *client.SafeguardClient, authProvider AuthenticationProvider) (User, error) {
	var updatedUser User

	u.PrimaryAuthenticationProvider = authProvider

	updatedUser, err := updateUser(c, u)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func updateUser(c *client.SafeguardClient, user User) (User, error) {
	var updatedUser User

	query := fmt.Sprintf("users/%d", user.Id)

	userJson, err := json.Marshal(user)
	if err != nil {
		return updatedUser, err
	}

	response, err := c.PutRequest(query, bytes.NewReader(userJson))
	if err != nil {
		return updatedUser, err
	}

	if err := json.Unmarshal(response, &updatedUser); err != nil {
		return updatedUser, err
	}

	updatedUser.client = c
	return updatedUser, nil
}
