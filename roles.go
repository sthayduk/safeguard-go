package safeguard

import (
	"encoding/json"
	"fmt"
	"time"
)

// Role represents roles in Safeguard made up of members, security scopes, and permissions
type Role struct {
	apiClient *SafeguardClient

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

func (a Role) SetClient(c *SafeguardClient) any {
	a.apiClient = c
	return a
}

// ToJson converts a Role object to its JSON string representation.
//
// Example:
//
//	role := Role{
//	    Name: "Administrator",
//	    Description: "Full system access"
//	}
//	json, err := role.ToJson()
//
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
//
// This method returns all roles matching the specified filter criteria. Common filters
// include Name, IsExpired, and CreatedDate.
//
// Example:
//
//	filter := Filter{}
//	filter.AddFilter("IsExpired", "eq", "false")
//	roles, err := GetRoles(filter)
//
// Parameters:
//   - fields: Filter object containing field comparisons and ordering preferences
//
// Returns:
//   - []Role: A slice of roles matching the filter criteria
//   - error: An error if the request fails, nil otherwise
func (c *SafeguardClient) GetRoles(fields Filter) ([]Role, error) {
	var userRoles []Role

	query := "Roles" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &userRoles); err != nil {
		return nil, err
	}

	return addClientToSlice(c, userRoles), nil
}

// GetEntitlements is an alias for GetRoles that retrieves a list of roles from Safeguard.
//
// This method provides compatibility with systems that use the term "entitlements"
// instead of "roles". It has identical functionality to GetRoles.
//
// Example:
//
//	filter := Filter{}
//	entitlements, err := GetEntitlements(filter)
//
// Parameters:
//   - fields: Filter object containing field comparisons and ordering preferences
//
// Returns:
//   - []Role: A slice of roles matching the filter criteria
//   - error: An error if the request fails, nil otherwise
func (c *SafeguardClient) GetEntitlements(fields Filter) ([]Role, error) {
	return c.GetRoles(fields)
}

// GetRole retrieves details for a specific role by ID.
//
// This method returns detailed information about a single role, optionally including
// related objects specified in the fields parameter.
//
// Example:
//
//	fields := Fields{}
//	fields.Add("Members", "Policies")
//	role, err := GetRole(123, fields)
//
// Parameters:
//   - id: The unique identifier of the role to retrieve
//   - fields: Optional Fields object specifying which related objects to include
//
// Returns:
//   - Role: The requested role with all specified related objects
//   - error: An error if the role is not found or request fails, nil otherwise
func (c *SafeguardClient) GetRole(id int, fields Fields) (Role, error) {
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

	return addClient(c, userRole), nil
}

// GetEntitlement is an alias for GetRole that retrieves details for a specific role.
//
// This method provides compatibility with systems that use the term "entitlement"
// instead of "role". It has identical functionality to GetRole.
//
// Example:
//
//	fields := Fields{}
//	entitlement, err := GetEntitlement(123, fields)
//
// Parameters:
//   - id: The unique identifier of the role to retrieve
//   - fields: Optional Fields object specifying which related objects to include
//
// Returns:
//   - Role: The requested role with all specified related objects
//   - error: An error if the role is not found or request fails, nil otherwise
func (c *SafeguardClient) GetEntitlement(id int, fields Fields) (Role, error) {
	return c.GetRole(id, fields)
}

// GetRoleMembers retrieves the list of members belonging to a specific role.
//
// This method returns all users who are members of the specified role, with optional
// filtering to restrict the results.
//
// Example:
//
//	filter := Filter{}
//	filter.AddFilter("PrincipalKind", "eq", "User")
//	members, err := GetRoleMembers(123, filter)
//
// Parameters:
//   - id: The unique identifier of the role
//   - filter: Filter object to restrict which members are returned
//
// Returns:
//   - []ManagedByUser: A slice of users who are members of the role
//   - error: An error if the request fails, nil otherwise
func (c *SafeguardClient) GetRoleMembers(id int, filter Filter) ([]ManagedByUser, error) {
	var members []ManagedByUser

	query := fmt.Sprintf("Roles/%d/Members%s", id, filter.ToQueryString())

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &members); err != nil {
		return nil, err
	}
	return addClientToSlice(c, members), nil
}

// GetMembers retrieves the list of members for the current role instance.
//
// This method is a convenience wrapper around GetRoleMembers that uses the
// current role's ID.
//
// Example:
//
//	filter := Filter{}
//	members, err := role.GetMembers(filter)
//
// Parameters:
//   - filter: Filter object to restrict which members are returned
//
// Returns:
//   - []ManagedByUser: A slice of users who are members of the role
//   - error: An error if the request fails, nil otherwise
func (r Role) GetMembers(filter Filter) ([]ManagedByUser, error) {
	return r.apiClient.GetRoleMembers(r.Id, filter)
}

// GetRolePolicies retrieves the list of access policies associated with a specific role.
//
// This method returns all access policies that are linked to the specified role,
// with optional filtering to restrict the results.
//
// Example:
//
//	filter := Filter{}
//	filter.AddFilter("IsExpired", "eq", "false")
//	policies, err := GetRolePolicies(123, filter)
//
// Parameters:
//   - id: The unique identifier of the role
//   - filter: Filter object to restrict which policies are returned
//
// Returns:
//   - []AccessPolicy: A slice of access policies associated with the role
//   - error: An error if the request fails, nil otherwise
func (c *SafeguardClient) GetRolePolicies(id int, filter Filter) ([]AccessPolicy, error) {
	var policies []AccessPolicy

	query := fmt.Sprintf("Roles/%d/Policies%s", id, filter.ToQueryString())

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &policies); err != nil {
		return nil, err
	}

	return addClientToSlice(c, policies), nil
}

// GetPolicies retrieves the list of access policies for the current role instance.
//
// This method is a convenience wrapper around GetRolePolicies that uses the
// current role's ID.
//
// Example:
//
//	filter := Filter{}
//	policies, err := role.GetPolicies(filter)
//
// Parameters:
//   - filter: Filter object to restrict which policies are returned
//
// Returns:
//   - []AccessPolicy: A slice of access policies associated with the role
//   - error: An error if the request fails, nil otherwise
func (r Role) GetPolicies(filter Filter) ([]AccessPolicy, error) {
	return r.apiClient.GetRolePolicies(r.Id, filter)
}
