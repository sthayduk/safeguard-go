// package pkg provides data structures and operations for interacting with Safeguard entities
package pkg

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sthayduk/safeguard-go/client"
)

// UserGroup represents a group of users in Safeguard with associated properties and memberships
type UserGroup struct {
	Id                           int                          `json:"Id"`
	Name                         string                       `json:"Name"`
	Description                  string                       `json:"Description"`
	IdentityProvider             GroupIdentityProvider        `json:"IdentityProvider"`
	IsReadOnly                   bool                         `json:"IsReadOnly"`
	CreatedDate                  time.Time                    `json:"CreatedDate"`
	CreatedByUserId              int                          `json:"CreatedByUserId"`
	CreatedByUserDisplayName     string                       `json:"CreatedByUserDisplayName"`
	ModifiedDate                 time.Time                    `json:"ModifiedDate"`
	ModifiedByUserId             int                          `json:"ModifiedByUserId"`
	ModifiedByUserDisplayName    string                       `json:"ModifiedByUserDisplayName"`
	DirectoryProperties          DirectoryProperties          `json:"DirectoryProperties"`
	Members                      []UserGroupMember            `json:"Members"`
	DirectoryGroupSyncProperties DirectoryGroupSyncProperties `json:"DirectoryGroupSyncProperties"`
}

// GroupIdentityProvider represents authentication provider information for a user group
type GroupIdentityProvider struct {
	Id                int    `json:"Id"`
	Name              string `json:"Name"`
	TypeReferenceName string `json:"TypeReferenceName"`
	IdentityId        string `json:"IdentityId"`
}

// DirectoryProperties represents directory-specific properties for a group or user
type UserGroupDirectoryProperties struct {
	DirectoryId       int    `json:"DirectoryId"`
	DirectoryName     string `json:"DirectoryName"`
	DomainName        string `json:"DomainName"`
	NetbiosName       string `json:"NetbiosName"`
	DistinguishedName string `json:"DistinguishedName"`
	ObjectGuid        string `json:"ObjectGuid"`
	ObjectSid         string `json:"ObjectSid"`
}

// UserGroupMember represents a user that belongs to a Safeguard user group,
// including their roles and authentication configuration
type UserGroupMember struct {
	AdminRoles                                []string               `json:"AdminRoles"`
	Id                                        int                    `json:"Id"`
	Name                                      string                 `json:"Name"`
	Description                               string                 `json:"Description"`
	DisplayName                               string                 `json:"DisplayName"`
	LastName                                  string                 `json:"LastName"`
	FirstName                                 string                 `json:"FirstName"`
	EmailAddress                              string                 `json:"EmailAddress"`
	WorkPhone                                 string                 `json:"WorkPhone"`
	MobilePhone                               string                 `json:"MobilePhone"`
	PrimaryAuthenticationProvider             AuthenticationProvider `json:"PrimaryAuthenticationProvider"`
	SecondaryAuthenticationProvider           AuthenticationProvider `json:"SecondaryAuthenticationProvider"`
	IdentityProvider                          GroupIdentityProvider  `json:"IdentityProvider"`
	Disabled                                  bool                   `json:"Disabled"`
	TimeZoneId                                string                 `json:"TimeZoneId"`
	TimeZoneDisplayName                       string                 `json:"TimeZoneDisplayName"`
	TimeZoneIanaName                          string                 `json:"TimeZoneIanaName"`
	IsPartitionOwner                          bool                   `json:"IsPartitionOwner"`
	DirectoryProperties                       DirectoryProperties    `json:"DirectoryProperties"`
	CloudAssistantApproveEnabled              bool                   `json:"CloudAssistantApproveEnabled"`
	CloudAssistantRecipientId                 string                 `json:"CloudAssistantRecipientId"`
	AllowPersonalAccounts                     bool                   `json:"AllowPersonalAccounts"`
	Locked                                    bool                   `json:"Locked"`
	PasswordNeverExpires                      bool                   `json:"PasswordNeverExpires"`
	ChangePasswordAtNextLogin                 bool                   `json:"ChangePasswordAtNextLogin"`
	Base64PhotoData                           string                 `json:"Base64PhotoData"`
	IsSystemOwned                             bool                   `json:"IsSystemOwned"`
	IsRequester                               bool                   `json:"IsRequester"`
	IsApprover                                bool                   `json:"IsApprover"`
	IsReviewer                                bool                   `json:"IsReviewer"`
	LastLoginDate                             time.Time              `json:"LastLoginDate"`
	LastRequestDate                           time.Time              `json:"LastRequestDate"`
	CreatedDate                               time.Time              `json:"CreatedDate"`
	CreatedByUserId                           int                    `json:"CreatedByUserId"`
	CreatedByUserDisplayName                  string                 `json:"CreatedByUserDisplayName"`
	ModifiedDate                              time.Time              `json:"ModifiedDate"`
	ModifiedByUserId                          int                    `json:"ModifiedByUserId"`
	ModifiedByUserDisplayName                 string                 `json:"ModifiedByUserDisplayName"`
	RequireCertificateAuthentication          bool                   `json:"RequireCertificateAuthentication"`
	DirectoryRequireCertificateAuthentication bool                   `json:"DirectoryRequireCertificateAuthentication"`
	LinkedAccountsCount                       int                    `json:"LinkedAccountsCount"`
}

// DirectoryGroupSyncProperties represents synchronization properties for groups synced from a directory
type DirectoryGroupSyncProperties struct {
	PrimaryAuthenticationProviderId                  int      `json:"PrimaryAuthenticationProviderId"`
	PrimaryAuthenticationProviderTypeReferenceName   string   `json:"PrimaryAuthenticationProviderTypeReferenceName"`
	PrimaryAuthenticationProviderName                string   `json:"PrimaryAuthenticationProviderName"`
	RequireCertificateAuthentication                 bool     `json:"RequireCertificateAuthentication"`
	SecondaryAuthenticationProviderId                int      `json:"SecondaryAuthenticationProviderId"`
	SecondaryAuthenticationProviderTypeReferenceName string   `json:"SecondaryAuthenticationProviderTypeReferenceName"`
	SecondaryAuthenticationProviderName              string   `json:"SecondaryAuthenticationProviderName"`
	LinkDirectoryAccounts                            bool     `json:"LinkDirectoryAccounts"`
	AllowPersonalAccounts                            bool     `json:"AllowPersonalAccounts"`
	AdminRoles                                       []string `json:"AdminRoles"`
}

// ToJson serializes a UserGroup object into a JSON string.
//
// This method converts the UserGroup instance into a JSON-formatted string,
// including all defined fields. Empty or zero-valued fields are included in
// the output.
//
// Example:
//
//	group := UserGroup{
//	    Name: "Administrators",
//	    Description: "System administrators group"
//	}
//	json, err := group.ToJson()
//
// Returns:
//   - string: A JSON representation of the UserGroup object
//   - error: An error if JSON marshaling fails, nil otherwise
func (u UserGroup) ToJson() (string, error) {
	userGroupJSON, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(userGroupJSON), nil
}

// GetUserGroups retrieves all user groups that match the specified filter criteria.
//
// The method supports filtering groups based on various properties like Name, IsReadOnly,
// CreatedDate etc. Multiple filters can be combined to narrow down results.
//
// Example:
//
//	fields := client.Filter{}
//	fields.AddFilter("IsReadOnly", "eq", "false")
//	fields.AddFilter("Name", "contains", "admin")
//	groups, err := GetUserGroups(fields)
//
// Parameters:
//   - fields: A Filter object containing field comparisons and ordering preferences
//
// Returns:
//   - []UserGroup: A slice of UserGroup objects matching the filter criteria
//   - error: An error if the request fails or response parsing fails, nil otherwise
func GetUserGroups(fields client.Filter) ([]UserGroup, error) {
	var userGroups []UserGroup

	query := "UserGroups" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &userGroups); err != nil {
		return nil, err
	}
	return userGroups, nil
}

// GetUserGroup retrieves a single user group by its unique identifier.
//
// The method can include additional related objects in the response based on the
// provided fields parameter.
//
// Example:
//
//	fields := client.Fields{}
//	fields.Add("Members", "DirectoryProperties")
//	group, err := GetUserGroup(123, fields)
//
// Parameters:
//   - id: The unique identifier of the user group to retrieve
//   - fields: Optional Fields object specifying which related objects to include
//
// Returns:
//   - UserGroup: The requested user group with all specified related objects
//   - error: An error if the group is not found or request fails, nil otherwise
func GetUserGroup(id int, fields client.Fields) (UserGroup, error) {
	var userGroup UserGroup

	query := fmt.Sprintf("UserGroups/%d", id)
	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return userGroup, err
	}
	if err := json.Unmarshal(response, &userGroup); err != nil {
		return userGroup, err
	}
	return userGroup, nil
}
