package pkg

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/client"
)

// PolicyAccount represents a Safeguard account with its associated policies and properties
type PolicyAccount struct {
	Id                          int               `json:"Id"`
	Name                        string            `json:"Name"`
	Description                 string            `json:"Description"`
	HasPassword                 bool              `json:"HasPassword"`
	HasSshKey                   bool              `json:"HasSshKey"`
	HasTotpAuthenticator        bool              `json:"HasTotpAuthenticator"`
	HasApiKeys                  bool              `json:"HasApiKeys"`
	HasFile                     bool              `json:"HasFile"`
	DomainName                  string            `json:"DomainName"`
	DistinguishedName           string            `json:"DistinguishedName"`
	NetBiosName                 string            `json:"NetBiosName"`
	Disabled                    bool              `json:"Disabled"`
	AccountType                 string            `json:"AccountType"`
	IsServiceAccount            bool              `json:"IsServiceAccount"`
	IsApplicationAccount        bool              `json:"IsApplicationAccount"`
	NotifyOwnersOnly            bool              `json:"NotifyOwnersOnly"`
	SuspendAccountWhenCheckedIn bool              `json:"SuspendAccountWhenCheckedIn"`
	DemoteAccountWhenCheckedIn  bool              `json:"DemoteAccountWhenCheckedIn"`
	AltLoginName                string            `json:"AltLoginName"`
	PrivilegeGroupMembership    string            `json:"PrivilegeGroupMembership"`
	LinkedUsersCount            int               `json:"LinkedUsersCount"`
	RequestProperties           RequestProperties `json:"RequestProperties"`
	Platform                    Platform          `json:"Platform"`
	Asset                       Asset             `json:"Asset"`
}

// RequestProperties represents the available request types for an account
type RequestProperties struct {
	AllowPasswordRequest bool `json:"AllowPasswordRequest"`
	AllowSessionRequest  bool `json:"AllowSessionRequest"`
	AllowSshKeyRequest   bool `json:"AllowSshKeyRequest"`
	AllowApiKeyRequest   bool `json:"AllowApiKeyRequest"`
	AllowFileRequest     bool `json:"AllowFileRequest"`
}

// PlatformType represents the type of platform
type PlatformType string

const (
	PlatformTypeUnknown      PlatformType = "Unknown"
	PlatformTypeWindows      PlatformType = "Windows"
	PlatformTypeLinux        PlatformType = "Linux"
	PlatformTypeDirectory    PlatformType = "Directory"
	PlatformTypeLocalhost    PlatformType = "LocalHost"
	PlatformTypeTeamPassword PlatformType = "TeamPassword"
	PlatformTypeOther        PlatformType = "Other"
)

// PlatformFamily represents the family of platform
type PlatformFamily string

const (
	PlatformFamilyNone            PlatformFamily = "None"
	PlatformFamilyUnix            PlatformFamily = "Unix"
	PlatformFamilyActiveDirectory PlatformFamily = "ActiveDirectory"
	PlatformFamilyTeamPassword    PlatformFamily = "TeamPassword"
)

// Platform represents a Safeguard platform configuration
type Platform struct {
	Id                        int            `json:"Id"`
	PlatformType              PlatformType   `json:"PlatformType"`
	DisplayName               string         `json:"DisplayName"`
	IsAcctNameCaseSensitive   bool           `json:"IsAcctNameCaseSensitive"`
	SupportsSessionManagement bool           `json:"SupportsSessionManagement"`
	PlatformFamily            PlatformFamily `json:"PlatformFamily"`
}

// ToJson serializes a PolicyAccount object into a JSON string.
//
// This method converts the PolicyAccount instance into a JSON-formatted string,
// including all defined fields. Empty or zero-valued fields are included in
// the output.
//
// Example:
//
//	account := PolicyAccount{
//	    Name: "webserver-admin",
//	    Description: "Admin account for web servers"
//	}
//	json, err := account.ToJson()
//
// Parameters:
//   - none
//
// Returns:
//   - string: A JSON representation of the PolicyAccount object
//   - error: An error if JSON marshaling fails, nil otherwise
func (p PolicyAccount) ToJson() (string, error) {
	policyAccountJSON, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(policyAccountJSON), nil
}

// GetPolicyAccounts retrieves all policy accounts that match the specified filter criteria.
//
// The method supports filtering accounts based on various properties like Name, Disabled,
// PlatformId etc. Multiple filters can be combined to narrow down results.
//
// Example:
//
//	fields := client.Filter{}
//	fields.AddFilter("Disabled", "eq", "false")
//	fields.AddFilter("PlatformId", "eq", "1")
//	accounts, err := GetPolicyAccounts(fields)
//
// Parameters:
//   - fields: A Filter object containing field comparisons and ordering preferences
//
// Returns:
//   - []PolicyAccount: A slice of PolicyAccount objects matching the filter criteria
//   - error: An error if the request fails or response parsing fails, nil otherwise
func GetPolicyAccounts(fields client.Filter) ([]PolicyAccount, error) {
	var policyAccounts []PolicyAccount

	query := "PolicyAccounts" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &policyAccounts)

	return policyAccounts, nil
}

// GetPolicyAccount retrieves a single policy account by its unique identifier.
//
// The method can include additional related objects in the response based on the
// provided fields parameter.
//
// Example:
//
//	fields := client.Fields{}
//	fields.Add("Asset", "Platform", "Owner")
//	account, err := GetPolicyAccount(123, fields)
//
// Parameters:
//   - id: The unique identifier of the policy account to retrieve
//   - fields: Optional Fields object specifying which related objects to include
//
// Returns:
//   - PolicyAccount: The requested policy account with all specified related objects
//   - error: An error if the account is not found or request fails, nil otherwise
func GetPolicyAccount(id int, fields client.Fields) (PolicyAccount, error) {
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

// LinkToUser creates a relationship between a policy account and a user.
//
// This method establishes a direct link between the account and user, granting
// access based on existing policies. The operation is atomic and transactional.
//
// Example:
//
//	user := User{Id: 456}
//	linkedAccounts, err := account.LinkToUser(user)
//
// Parameters:
//   - user: The User object representing the user to link with
//
// Returns:
//   - []PolicyAccount: A slice containing the updated account after linking
//   - error: An error if the link operation fails, nil otherwise
func (p PolicyAccount) LinkToUser(user User) ([]PolicyAccount, error) {
	return AddLinkedAccounts(user, []PolicyAccount{p})
}

// UnlinkFromUser removes the relationship between a policy account and a user.
//
// This method removes direct access between the account and user. The user may
// still have access through other means (groups, policies etc).
//
// Example:
//
//	user := User{Id: 456}
//	unlinkedAccounts, err := account.UnlinkFromUser(user)
//
// Parameters:
//   - user: The User object representing the user to unlink from
//
// Returns:
//   - []PolicyAccount: A slice containing the updated account after unlinking
//   - error: An error if the unlink operation fails, nil otherwise
func (p PolicyAccount) UnlinkFromUser(user User) ([]PolicyAccount, error) {
	return RemoveLinkedAccounts(user, []PolicyAccount{p})
}
