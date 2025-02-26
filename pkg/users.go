package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sthayduk/safeguard-go/client"
)

type User struct {
	client *client.SafeguardClient

	Name                                      string                 `json:"Name,omitempty"`
	PrimaryAuthenticationProvider             AuthenticationProvider `json:"PrimaryAuthenticationProvider,omitempty"`
	Preferences                               []Preference           `json:"Preferences,omitempty"`
	Fido2Authenticators                       []Fido2Authenticator   `json:"Fido2Authenticators,omitempty"`
	AdminRoles                                []string               `json:"AdminRoles,omitempty"`
	Id                                        int                    `json:"Id,omitempty"`
	Description                               string                 `json:"Description,omitempty"`
	DisplayName                               string                 `json:"DisplayName,omitempty"`
	LastName                                  string                 `json:"LastName,omitempty"`
	FirstName                                 string                 `json:"FirstName,omitempty"`
	EmailAddress                              string                 `json:"EmailAddress,omitempty"`
	WorkPhone                                 string                 `json:"WorkPhone,omitempty"`
	MobilePhone                               string                 `json:"MobilePhone,omitempty"`
	SecondaryAuthenticationProvider           AuthenticationProvider `json:"SecondaryAuthenticationProvider,omitempty"`
	IdentityProvider                          IdentityProvider       `json:"IdentityProvider,omitempty"`
	Disabled                                  bool                   `json:"Disabled,omitempty"`
	TimeZoneId                                string                 `json:"TimeZoneId,omitempty"`
	TimeZoneDisplayName                       string                 `json:"TimeZoneDisplayName,omitempty"`
	TimeZoneIanaName                          string                 `json:"TimeZoneIanaName,omitempty"`
	IsPartitionOwner                          bool                   `json:"IsPartitionOwner,omitempty"`
	DirectoryProperties                       DirectoryProperties    `json:"DirectoryProperties,omitempty"`
	CloudAssistantApproveEnabled              bool                   `json:"CloudAssistantApproveEnabled,omitempty"`
	CloudAssistantRecipientId                 string                 `json:"CloudAssistantRecipientId,omitempty"`
	AllowPersonalAccounts                     bool                   `json:"AllowPersonalAccounts,omitempty"`
	Locked                                    bool                   `json:"Locked,omitempty"`
	PasswordNeverExpires                      bool                   `json:"PasswordNeverExpires,omitempty"`
	ChangePasswordAtNextLogin                 bool                   `json:"ChangePasswordAtNextLogin,omitempty"`
	Base64PhotoData                           string                 `json:"Base64PhotoData,omitempty"`
	IsSystemOwned                             bool                   `json:"IsSystemOwned,omitempty"`
	IsRequester                               bool                   `json:"IsRequester,omitempty"`
	IsApprover                                bool                   `json:"IsApprover,omitempty"`
	IsReviewer                                bool                   `json:"IsReviewer,omitempty"`
	LastLoginDate                             time.Time              `json:"LastLoginDate,omitempty"`
	LastRequestDate                           time.Time              `json:"LastRequestDate,omitempty"`
	CreatedDate                               time.Time              `json:"CreatedDate,omitempty"`
	CreatedByUserId                           int                    `json:"CreatedByUserId,omitempty"`
	CreatedByUserDisplayName                  string                 `json:"CreatedByUserDisplayName,omitempty"`
	ModifiedDate                              time.Time              `json:"ModifiedDate,omitempty"`
	ModifiedByUserId                          int                    `json:"ModifiedByUserId,omitempty"`
	ModifiedByUserDisplayName                 string                 `json:"ModifiedByUserDisplayName,omitempty"`
	RequireCertificateAuthentication          bool                   `json:"RequireCertificateAuthentication,omitempty"`
	DirectoryRequireCertificateAuthentication bool                   `json:"DirectoryRequireCertificateAuthentication,omitempty"`
	LinkedAccountsCount                       int                    `json:"LinkedAccountsCount,omitempty"`
}

// ToJson converts a User object to its JSON string representation.
// Returns:
//   - string: The JSON string representation of the user
//   - error: An error if marshaling fails, nil otherwise
func (u User) ToJson() (string, error) {
	userJSON, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(userJSON), nil
}

// GetUsers retrieves a list of users from Safeguard.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - fields: Filter criteria for the request
//
// Returns:
//   - []User: A slice of users matching the filter criteria
//   - error: An error if the request fails, nil otherwise
func GetUsers(c *client.SafeguardClient, fields client.Filter) ([]User, error) {
	var users []User

	query := "users" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &users); err != nil {
		return nil, err
	}

	for i := range users {
		users[i].client = c
	}
	return users, nil
}

// GetUser retrieves details for a specific user by their ID.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The numeric ID of the user to retrieve
//   - fields: Specific fields to include in the response
//
// Returns:
//   - User: The requested user object
//   - error: An error if the request fails, nil otherwise
func GetUser(c *client.SafeguardClient, id int, fields client.Fields) (User, error) {
	var user User
	user.client = c

	query := fmt.Sprintf("users/%d", id)
	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return user, err
	}

	if err := json.Unmarshal(response, &user); err != nil {
		return user, err
	}

	return user, nil
}

// GetLinkedAccounts retrieves the policy accounts linked to a specific user ID.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The numeric ID of the user to get linked accounts for
//
// Returns:
//   - []PolicyAccount: A slice of linked policy accounts
//   - error: An error if the request fails, nil otherwise
func GetLinkedAccounts(c *client.SafeguardClient, id string) ([]PolicyAccount, error) {
	var linkedAccounts []PolicyAccount

	query := fmt.Sprintf("users/%s/LinkedPolicyAccounts", id)

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &linkedAccounts); err != nil {
		return nil, err
	}

	return linkedAccounts, nil
}

// GetLinkedAccounts retrieves the policy accounts linked to this user.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//
// Returns:
//   - []PolicyAccount: A slice of linked policy accounts
//   - error: An error if the request fails, nil otherwise
func (u User) GetLinkedAccounts() ([]PolicyAccount, error) {
	return GetLinkedAccounts(u.client, fmt.Sprintf("%d", u.Id))
}

// GetUserRoles retrieves the roles assigned to a specific user.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The ID of the user
//
// Returns:
//   - []Role: A slice of assigned roles
//   - error: An error if the request fails, nil otherwise
func GetUserRoles(c *client.SafeguardClient, id string) ([]Role, error) {
	var roles []Role

	query := fmt.Sprintf("users/%s/roles", id)

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &roles); err != nil {
		return nil, err
	}
	return roles, nil
}

// GetRoles retrieves the roles assigned to this user.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//
// Returns:
//   - []Role: A slice of assigned roles
//   - error: An error if the request fails, nil otherwise
func (u User) GetRoles() ([]Role, error) {
	return GetUserRoles(u.client, fmt.Sprintf("%d", u.Id))
}

// GetGroups retrieves the groups that a specific user belongs to.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The ID of the user
//
// Returns:
//   - []UserGroup: A slice of user groups
//   - error: An error if the request fails, nil otherwise
func GetGroups(c *client.SafeguardClient, id string) ([]UserGroup, error) {
	var userGroups []UserGroup

	query := fmt.Sprintf("users/%s/UserGroups", id)

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &userGroups); err != nil {
		return nil, err
	}
	return userGroups, nil
}

// GetGroups retrieves the groups that this user belongs to.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//
// Returns:
//   - []UserGroup: A slice of user groups
//   - error: An error if the request fails, nil otherwise
func (u User) GetGroups() ([]UserGroup, error) {
	return GetGroups(u.client, fmt.Sprintf("%d", u.Id))
}

// GetUserPreferences retrieves the preferences for a specific user by ID.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The numeric ID of the user whose preferences to retrieve
//
// Returns:
//   - []Preference: A slice of user preferences
//   - error: An error if the request fails, nil otherwise
func GetUserPreferences(c *client.SafeguardClient, id int) ([]Preference, error) {
	var userPreferences []Preference

	query := fmt.Sprintf("users/%d/preferences", id)

	response, err := c.GetRequest(query)
	if err != nil {
		return userPreferences, err
	}

	if err := json.Unmarshal(response, &userPreferences); err != nil {
		return userPreferences, err
	}
	return userPreferences, nil
}

// GetPreferences retrieves the preferences for the current user instance.
// Returns:
//   - []Preference: A slice of user preferences
//   - error: An error if the request fails, nil otherwise
func (u User) GetPreferences() ([]Preference, error) {
	return GetUserPreferences(u.client, u.Id)
}

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

// AddLinkedAccounts adds linked policy accounts to the current user instance.
// Parameters:
//   - policyAccount: A slice of PolicyAccount objects to be linked to the user
//
// Returns:
//   - []PolicyAccount: A slice of PolicyAccount objects that were successfully linked
//   - error: An error if the operation fails, nil otherwise
func (u User) AddLinkedAccounts(policyAccount []PolicyAccount) ([]PolicyAccount, error) {
	return AddLinkedAccounts(u.client, u, policyAccount)
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

// RemoveLinkedAccounts removes linked policy accounts from the current user instance.
// Parameters:
//   - policyAccount: A slice of PolicyAccount objects to be removed from the user
//
// Returns:
//   - []PolicyAccount: A slice of PolicyAccount objects that were successfully removed
//   - error: An error if the operation fails, nil otherwise
func (u User) RemoveLinkedAccounts(policyAccount []PolicyAccount) ([]PolicyAccount, error) {
	return RemoveLinkedAccounts(u.client, u, policyAccount)
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

// SetAuthenticationProvider sets the primary authentication provider for the user
// and updates the user information in the Safeguard system.
//
// Parameters:
//
//	authProvider - The AuthenticationProvider to be set as the primary authentication provider.
//
// Returns:
//
//	User - The updated User object with the new primary authentication provider.
//	error - An error object if there was an issue updating the user information.
func (u User) SetAuthenticationProvider(authProvider AuthenticationProvider) (User, error) {
	var updatedUser User

	u.PrimaryAuthenticationProvider = authProvider

	updatedUser, err := updateUser(u.client, u)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

// updateUser updates the details of an existing user in the Safeguard system.
// It takes a SafeguardClient and a User object as parameters, and returns the updated User object and an error, if any.
//
// Parameters:
//   - c: A pointer to a SafeguardClient used to make the request.
//   - user: A User object containing the updated user details.
//
// Returns:
//   - User: The updated User object.
//   - error: An error object if an error occurred during the update process, otherwise nil.
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

// DeleteUser removes a user from the Safeguard system by their ID.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The numeric ID of the user to delete
//
// Returns:
//   - error: An error if the deletion fails, nil otherwise
func DeleteUser(c *client.SafeguardClient, id int) error {

	query := fmt.Sprintf("users/%d", id)

	_, err := c.DeleteRequest(query)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes the current user instance from the Safeguard system.
// Returns:
//   - error: An error if the deletion fails, nil otherwise
func (u *User) Delete() error {
	return DeleteUser(u.client, u.Id)
}
