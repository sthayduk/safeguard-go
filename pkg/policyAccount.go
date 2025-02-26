package pkg

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/client"
)

// PolicyAccount represents a Safeguard account with its associated policies and properties
type PolicyAccount struct {
	client *client.SafeguardClient

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

func (p PolicyAccount) ToJson() (string, error) {
	policyAccountJSON, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(policyAccountJSON), nil
}

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

	for i := range policyAccounts {
		policyAccounts[i].client = c
	}

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
	policyAccount.client = c
	return policyAccount, nil
}

// LinkToUser links the current PolicyAccount to the specified User.
// It returns a slice of PolicyAccount and an error if any occurs during the linking process.
//
// Parameters:
//
//	user - The User to link the PolicyAccount to.
//
// Returns:
//
//	[]PolicyAccount - A slice containing the linked PolicyAccount.
//	error - An error if any issues occur during the linking process.
func (p PolicyAccount) LinkToUser(user User) ([]PolicyAccount, error) {
	return AddLinkedAccounts(p.client, user, []PolicyAccount{p})
}

// UnlinkFromUser removes the association between the given PolicyAccount and the specified User.
// It returns a slice of PolicyAccount and an error if the operation fails.
//
// Parameters:
//
//	user - The User from whom the PolicyAccount should be unlinked.
//
// Returns:
//
//	[]PolicyAccount - A slice containing the PolicyAccount after unlinking.
//	error - An error if the unlinking operation fails.
func (p PolicyAccount) UnlinkFromUser(user User) ([]PolicyAccount, error) {
	return RemoveLinkedAccounts(p.client, user, []PolicyAccount{p})
}
