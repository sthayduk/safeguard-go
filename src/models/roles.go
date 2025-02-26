package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sthayduk/safeguard-go/client"
)

// Role represents roles in Safeguard made up of members, security scopes, and permissions
type Role struct {
	client *client.SafeguardClient

	Id                          int                         `json:"Id"`
	Name                        string                      `json:"Name"`
	Priority                    int                         `json:"Priority"`
	Description                 string                      `json:"Description"`
	ExpirationDate              time.Time                   `json:"ExpirationDate"`
	IsExpired                   bool                        `json:"IsExpired"`
	HasExpiredPolicies          bool                        `json:"HasExpiredPolicies"`
	HasInvalidPolicies          bool                        `json:"HasInvalidPolicies"`
	CreatedDate                 time.Time                   `json:"CreatedDate"`
	CreatedByUserId             int                         `json:"CreatedByUserId"`
	CreatedByUserDisplayName    string                      `json:"CreatedByUserDisplayName"`
	UserCount                   int                         `json:"UserCount"`
	AccountCount                int                         `json:"AccountCount"`
	AssetCount                  int                         `json:"AssetCount"`
	PolicyCount                 int                         `json:"PolicyCount"`
	HourlyRestrictionProperties HourlyRestrictionProperties `json:"HourlyRestrictionProperties"`
	Members                     []RoleMember                `json:"Members"`
}

// ToJson converts a Role object to its JSON string representation
// Returns:
//   - string: JSON representation of the role
//   - error: An error if marshaling fails, nil otherwise
func (u Role) ToJson() (string, error) {
	userJSON, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(userJSON), nil
}

// RoleMember represents a member of a role
type RoleMember struct {
	DisplayName                       string `json:"DisplayName"`
	Id                                int    `json:"Id"`
	IdentityProviderId                int    `json:"IdentityProviderId"`
	IdentityProviderName              string `json:"IdentityProviderName"`
	IdentityProviderTypeReferenceName string `json:"IdentityProviderTypeReferenceName"`
	IsSystemOwned                     bool   `json:"IsSystemOwned"`
	Name                              string `json:"Name"`
	PrincipalKind                     string `json:"PrincipalKind"`
	EmailAddress                      string `json:"EmailAddress"`
	DomainName                        string `json:"DomainName"`
	FullDisplayName                   string `json:"FullDisplayName"`
}

// HourlyRestrictionProperties represents settings controlling when the policy/role will be effective
type HourlyRestrictionProperties struct {
	EnableHourlyRestrictions bool  `json:"EnableHourlyRestrictions"`
	MondayValidHours         []int `json:"MondayValidHours"`
	TuesdayValidHours        []int `json:"TuesdayValidHours"`
	WednesdayValidHours      []int `json:"WednesdayValidHours"`
	ThursdayValidHours       []int `json:"ThursdayValidHours"`
	FridayValidHours         []int `json:"FridayValidHours"`
	SaturdayValidHours       []int `json:"SaturdayValidHours"`
	SundayValidHours         []int `json:"SundayValidHours"`
}

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

	for i := range userRoles {
		userRoles[i].client = c
	}

	return userRoles, nil
}

// GetEntitlements is an alias for GetRoles that retrieves a list of roles from Safeguard.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - fields: Filter criteria for the request
//
// Returns:
//   - []Role: A slice of roles matching the filter criteria
//   - error: An error if the request fails, nil otherwise
func GetEntitlements(c *client.SafeguardClient, fields client.Filter) ([]Role, error) {
	return GetRoles(c, fields)
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

	userRole.client = c
	return userRole, nil
}

// GetEntitlement is an alias for GetRole that retrieves details for a specific role.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The numeric ID of the role to retrieve
//   - fields: Specific fields to include in the response
//
// Returns:
//   - Role: The requested role object
//   - error: An error if the request fails, nil otherwise
func GetEntitlement(c *client.SafeguardClient, id int, fields client.Fields) (Role, error) {
	return GetRole(c, id, fields)
}

// GetRoleMembers retrieves the list of members belonging to a specific role.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The numeric ID of the role
//   - filter: Filter criteria for the request
//
// Returns:
//   - []ManagedByUser: A slice of users who are members of the role
//   - error: An error if the request fails, nil otherwise
func GetRoleMembers(c *client.SafeguardClient, id int, filter client.Filter) ([]ManagedByUser, error) {
	var members []ManagedByUser

	query := fmt.Sprintf("Roles/%d/Members%s", id, filter.ToQueryString())

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &members); err != nil {
		return nil, err
	}
	return members, nil
}

// GetMembers retrieves the list of members for the current role instance.
// Parameters:
//   - filter: Filter criteria for the request
//
// Returns:
//   - []ManagedByUser: A slice of users who are members of the role
//   - error: An error if the request fails, nil otherwise
func (r Role) GetMembers(filter client.Filter) ([]ManagedByUser, error) {
	return GetRoleMembers(r.client, r.Id, filter)
}

// GetRolePolicies retrieves the list of access policies associated with a specific role.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The numeric ID of the role
//   - filter: Filter criteria for the request
//
// Returns:
//   - []AccessPolicy: A slice of access policies associated with the role
//   - error: An error if the request fails, nil otherwise
func GetRolePolicies(c *client.SafeguardClient, id int, filter client.Filter) ([]AccessPolicy, error) {
	var policies []AccessPolicy

	query := fmt.Sprintf("Roles/%d/Policies%s", id, filter.ToQueryString())

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &policies); err != nil {
		return nil, err
	}

	for i := range policies {
		policies[i].client = c
	}

	return policies, nil
}

// GetPolicies retrieves the list of access policies for the current role instance.
// Parameters:
//   - filter: Filter criteria for the request
//
// Returns:
//   - []AccessPolicy: A slice of access policies associated with the role
//   - error: An error if the request fails, nil otherwise
func (r Role) GetPolicies(filter client.Filter) ([]AccessPolicy, error) {
	return GetRolePolicies(r.client, r.Id, filter)
}
