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
//
// This method serializes all fields of the User object into a JSON-formatted string.
// Empty or zero-valued fields are included in the output.
//
// Example:
//
//	user := User{
//	    Name: "John Smith",
//	    EmailAddress: "john.smith@example.com"
//	}
//	json, err := user.ToJson()
//
// Returns:
//   - string: JSON representation of the user
//   - error: An error if marshaling fails, nil otherwise
func (u User) ToJson() (string, error) {
	userJSON, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(userJSON), nil
}

// GetUsers retrieves a list of users from Safeguard.
//
// This method returns all users matching the specified filter criteria. Common filters
// include Name, EmailAddress, and Disabled.
//
// Example:
//
//	filter := client.Filter{}
//	filter.AddFilter("Disabled", "eq", "false")
//	users, err := GetUsers(filter)
//
// Parameters:
//   - fields: Filter object containing field comparisons and ordering preferences
//
// Returns:
//   - []User: A slice of users matching the filter criteria
//   - error: An error if the request fails, nil otherwise
func GetUsers(fields client.Filter) ([]User, error) {
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

// GetUser retrieves details for a specific user by ID.
//
// This method returns detailed information about a single user, optionally including
// related objects specified in the fields parameter.
//
// Example:
//
//	fields := client.Fields{}
//	fields.Add("LinkedAccounts", "Preferences")
//	user, err := GetUser(123, fields)
//
// Parameters:
//   - id: The unique identifier of the user to retrieve
//   - fields: Optional Fields object specifying which related objects to include
//
// Returns:
//   - User: The requested user with all specified related objects
//   - error: An error if the user is not found or request fails, nil otherwise
func GetUser(id int, fields client.Fields) (User, error) {
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
//
// This method returns all policy accounts that are linked to the specified user.
//
// Example:
//
//	accounts, err := GetLinkedAccounts("123")
//
// Parameters:
//   - id: The string identifier of the user
//
// Returns:
//   - []PolicyAccount: A slice of linked policy accounts
//   - error: An error if the request fails, nil otherwise
func GetLinkedAccounts(id string) ([]PolicyAccount, error) {
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
//
// This method is a convenience wrapper around GetLinkedAccounts that uses the
// current user's ID.
//
// Example:
//
//	accounts, err := user.GetLinkedAccounts()
//
// Returns:
//   - []PolicyAccount: A slice of linked policy accounts
//   - error: An error if the request fails, nil otherwise
func (u User) GetLinkedAccounts() ([]PolicyAccount, error) {
	return GetLinkedAccounts(fmt.Sprintf("%d", u.Id))
}

// GetUserRoles retrieves the roles assigned to a specific user.
//
// This method returns all roles that have been assigned to the specified user.
//
// Example:
//
//	roles, err := GetUserRoles("123")
//
// Parameters:
//   - id: The string identifier of the user
//
// Returns:
//   - []Role: A slice of assigned roles
//   - error: An error if the request fails, nil otherwise
func GetUserRoles(id string) ([]Role, error) {
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
//
// This method is a convenience wrapper around GetUserRoles that uses the
// current user's ID.
//
// Example:
//
//	roles, err := user.GetRoles()
//
// Returns:
//   - []Role: A slice of assigned roles
//   - error: An error if the request fails, nil otherwise
func (u User) GetRoles() ([]Role, error) {
	return GetUserRoles(fmt.Sprintf("%d", u.Id))
}

// GetGroups retrieves the groups that a specific user belongs to.
//
// This method returns all user groups that the specified user is a member of.
//
// Example:
//
//	groups, err := GetGroups("123")
//
// Parameters:
//   - id: The string identifier of the user
//
// Returns:
//   - []UserGroup: A slice of user groups
//   - error: An error if the request fails, nil otherwise
func GetGroups(id string) ([]UserGroup, error) {
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
//
// This method is a convenience wrapper around GetGroups that uses the
// current user's ID.
//
// Example:
//
//	groups, err := user.GetGroups()
//
// Returns:
//   - []UserGroup: A slice of user groups
//   - error: An error if the request fails, nil otherwise
func (u User) GetGroups() ([]UserGroup, error) {
	return GetGroups(fmt.Sprintf("%d", u.Id))
}

// GetUserPreferences retrieves the preferences for a specific user.
//
// This method returns all preferences associated with the specified user ID.
//
// Example:
//
//	prefs, err := GetUserPreferences(123)
//
// Parameters:
//   - id: The unique identifier of the user
//
// Returns:
//   - []Preference: A slice of user preferences
//   - error: An error if the request fails, nil otherwise
func GetUserPreferences(id int) ([]Preference, error) {
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

// GetPreferences retrieves the preferences for this user.
//
// This method is a convenience wrapper around GetUserPreferences that uses the
// current user's ID.
//
// Example:
//
//	prefs, err := user.GetPreferences()
//
// Returns:
//   - []Preference: A slice of user preferences
//   - error: An error if the request fails, nil otherwise
func (u User) GetPreferences() ([]Preference, error) {
	return GetUserPreferences(u.Id)
}

// AddLinkedAccounts adds policy accounts to a user's linked accounts.
//
// This method associates the specified policy accounts with the given user.
//
// Example:
//
//	accounts := []PolicyAccount{{Id: 123}, {Id: 456}}
//	linked, err := AddLinkedAccounts(user, accounts)
//
// Parameters:
//   - user: The user to link accounts to
//   - policyAccount: A slice of policy accounts to link
//
// Returns:
//   - []PolicyAccount: The linked policy accounts
//   - error: An error if the operation fails, nil otherwise
func AddLinkedAccounts(user User, policyAccount []PolicyAccount) ([]PolicyAccount, error) {
	var linkedAccounts []PolicyAccount

	query := fmt.Sprintf("users/%d/LinkedPolicyAccounts/Add", user.Id)
	response, err := fetchAndPostPolicyAccount(policyAccount, query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &linkedAccounts); err != nil {
		return nil, err
	}
	return linkedAccounts, nil
}

// AddLinkedAccounts adds policy accounts to this user's linked accounts.
//
// This method is a convenience wrapper around AddLinkedAccounts that uses the
// current user instance.
//
// Example:
//
//	accounts := []PolicyAccount{{Id: 123}}
//	linked, err := user.AddLinkedAccounts(accounts)
//
// Parameters:
//   - policyAccount: A slice of policy accounts to link
//
// Returns:
//   - []PolicyAccount: The linked policy accounts
//   - error: An error if the operation fails, nil otherwise
func (u User) AddLinkedAccounts(policyAccount []PolicyAccount) ([]PolicyAccount, error) {
	return AddLinkedAccounts(u, policyAccount)
}

// RemoveLinkedAccounts removes policy accounts from a user's linked accounts.
//
// This method removes the association between the specified policy accounts
// and the given user.
//
// Example:
//
//	accounts := []PolicyAccount{{Id: 123}}
//	removed, err := RemoveLinkedAccounts(user, accounts)
//
// Parameters:
//   - user: The user to remove links from
//   - policyAccount: A slice of policy accounts to unlink
//
// Returns:
//   - []PolicyAccount: The unlinked policy accounts
//   - error: An error if the operation fails, nil otherwise
func RemoveLinkedAccounts(user User, policyAccount []PolicyAccount) ([]PolicyAccount, error) {
	var linkedAccounts []PolicyAccount

	query := fmt.Sprintf("users/%d/LinkedPolicyAccounts/Remove", user.Id)
	response, err := fetchAndPostPolicyAccount(policyAccount, query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &linkedAccounts); err != nil {
		return nil, err
	}
	return linkedAccounts, nil
}

// RemoveLinkedAccounts removes policy accounts from this user's linked accounts.
//
// This method is a convenience wrapper around RemoveLinkedAccounts that uses the
// current user instance.
//
// Example:
//
//	accounts := []PolicyAccount{{Id: 123}}
//	removed, err := user.RemoveLinkedAccounts(accounts)
//
// Parameters:
//   - policyAccount: A slice of policy accounts to unlink
//
// Returns:
//   - []PolicyAccount: The unlinked policy accounts
//   - error: An error if the operation fails, nil otherwise
func (u User) RemoveLinkedAccounts(policyAccount []PolicyAccount) ([]PolicyAccount, error) {
	return RemoveLinkedAccounts(u, policyAccount)
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
func fetchAndPostPolicyAccount(policyAccount []PolicyAccount, query string) ([]byte, error) {
	policyAccountJson, err := json.Marshal(policyAccount)
	if err != nil {
		return nil, err
	}

	return c.PostRequest(query, bytes.NewReader(policyAccountJson))
}

// CreateUser creates a new user in Safeguard.
//
// This method creates a new user with the provided user details and returns
// the created user object.
//
// Example:
//
//	newUser := User{
//	    Name: "john.smith",
//	    EmailAddress: "john.smith@example.com"
//	}
//	created, err := CreateUser(newUser)
//
// Parameters:
//   - user: The user object containing the new user's details
//
// Returns:
//   - User: The created user object
//   - error: An error if the creation fails, nil otherwise
func CreateUser(user User) (User, error) {
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

// SetAuthenticationProvider updates the primary authentication provider for this user.
//
// This method updates the user's primary authentication method and saves the
// changes to Safeguard.
//
// Example:
//
//	provider := AuthenticationProvider{Id: 123}
//	updated, err := user.SetAuthenticationProvider(provider)
//
// Parameters:
//   - authProvider: The new authentication provider to set
//
// Returns:
//   - User: The updated user object
//   - error: An error if the update fails, nil otherwise
func (u User) SetAuthenticationProvider(authProvider AuthenticationProvider) (User, error) {
	var updatedUser User

	u.PrimaryAuthenticationProvider = authProvider

	updatedUser, err := updateUser(u)
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
func updateUser(user User) (User, error) {
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

// DeleteUser removes a user from Safeguard.
//
// This method permanently deletes the specified user from the system.
//
// Example:
//
//	err := DeleteUser(123)
//
// Parameters:
//   - id: The unique identifier of the user to delete
//
// Returns:
//   - error: An error if the deletion fails, nil otherwise
func DeleteUser(id int) error {

	query := fmt.Sprintf("users/%d", id)

	_, err := c.DeleteRequest(query)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes this user from Safeguard.
//
// This method is a convenience wrapper around DeleteUser that uses the
// current user's ID.
//
// Example:
//
//	err := user.Delete()
//
// Returns:
//   - error: An error if the deletion fails, nil otherwise
func (u *User) Delete() error {
	return DeleteUser(u.Id)
}
