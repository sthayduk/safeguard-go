package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/client"
)

// AssetPolicy represents a policy that an asset belongs to plus how that membership was granted
type AssetPolicy struct {
	PolicyId                int                     `json:"PolicyId"`
	PolicyName              string                  `json:"PolicyName"`
	AccessRequestType       AccessRequestType       `json:"AccessRequestType"`
	RoleId                  int                     `json:"RoleId"`
	RoleName                string                  `json:"RoleName"`
	AssetId                 int                     `json:"AssetId"`
	AssetName               string                  `json:"AssetName"`
	PolicyAccountCount      int                     `json:"PolicyAccountCount"`
	PolicyAccountGroupCount int                     `json:"PolicyAccountGroupCount"`
	PolicyAssetCount        int                     `json:"PolicyAssetCount"`
	PolicyAssetGroupCount   int                     `json:"PolicyAssetGroupCount"`
	Membership              []AssetPolicyMembership `json:"Membership"`
}

// AssetPolicyMembership represents details about how an asset is assigned to a policy
type AssetPolicyMembership struct {
	PolicyId                   int    `json:"PolicyId"`
	AssetId                    int    `json:"AssetId"`
	PolicyMemberId             int    `json:"PolicyMemberId"`
	PolicyMemberName           string `json:"PolicyMemberName"`
	PolicyMemberIsAssetGroup   bool   `json:"PolicyMemberIsAssetGroup"`
	PolicyMemberIsAccountGroup bool   `json:"PolicyMemberIsAccountGroup"`
}

// AccessRequestType represents the type of access request
type AccessRequestType string

const (
	AccessRequestTypePassword       AccessRequestType = "Password"
	AccessRequestTypeRDPFile        AccessRequestType = "RemoteDesktop"
	AccessRequestTypeSSHFile        AccessRequestType = "SSH"
	AccessRequestTypeSSHKey         AccessRequestType = "SSHKey"
	AccessRequestTypeAPIKey         AccessRequestType = "APIKey"
	AccessRequestTypeRDP            AccessRequestType = "RemoteDesktop"
	AccessRequestTypeTelnet         AccessRequestType = "Telnet"
	AccessRequestTypeRDPApplication AccessRequestType = "RemoteDesktopApplication"
	AccessRequestTypeFile           AccessRequestType = "File"
)

// ToJson converts an AssetPolicy to its JSON string representation
func (a AssetPolicy) ToJson() (string, error) {
	assetPolicyJSON, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(assetPolicyJSON), nil
}

// PolicyAsset represents an remote asset available for request.
// A PolicyAsset is an alternate view of an asset that is used for AccessPolicies, AssetGroups,
// and UserFavorites. The asset must have AllowSessionRequests set to true in order to be
// used in UserFavorites or to be able to request a session on the asset.
type PolicyAsset struct {
	client *client.SafeguardClient

	Id                 int                     `json:"Id"`
	Name               string                  `json:"Name"`
	AssetType          AssetType               `json:"AssetType"`
	NetworkAddress     string                  `json:"NetworkAddress"`
	Description        string                  `json:"Description"`
	AssetPartitionId   int                     `json:"AssetPartitionId"`
	AssetPartitionName string                  `json:"AssetPartitionName"`
	DomainName         string                  `json:"DomainName"`
	Disabled           bool                    `json:"Disabled"`
	Platform           PolicyAssetPlatform     `json:"Platform"`
	SshHostKey         AssetSshHostKey         `json:"SshHostKey,omitempty"`
	SessionAccess      SessionAccessProperties `json:"SessionAccessProperties"`
}

func (p PolicyAsset) ToJson() (string, error) {
	policyAssetJSON, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(policyAssetJSON), nil
}

// AssetSshHostKey represents an SSH Host Key used to identify assets
type AssetSshHostKey struct {
	Id                int    `json:"Id"`
	Fingerprint       string `json:"Fingerprint,omitempty"`
	Key               string `json:"Key,omitempty"`
	KeyType           string `json:"KeyType,omitempty"`
	Comment           string `json:"Comment,omitempty"`
	CanBeAccepted     bool   `json:"CanBeAccepted,omitempty"`
	SshHostKey        string `json:"SshHostKey"`
	FingerprintSha256 string `json:"FingerprintSha256"`
}

// PolicyAssetPlatform represents platform information specific to policy assets
type PolicyAssetPlatform struct {
	Id                        int          `json:"Id"`
	PlatformType              PlatformType `json:"PlatformType"`
	DisplayName               string       `json:"DisplayName"`
	SupportsSessionManagement bool         `json:"SupportsSessionManagement"`
}

// SessionAccessProperties represents the session access configuration for a policy asset
type SessionAccessProperties struct {
	AllowSessionRequests     bool `json:"AllowSessionRequests"`
	SshSessionPort           int  `json:"SshSessionPort,omitempty"`
	RemoteDesktopSessionPort int  `json:"RemoteDesktopSessionPort,omitempty"`
	TelnetSessionPort        int  `json:"TelnetSessionPort,omitempty"`
}

// GetPolicyAssets retrieves a list of policy assets from the Safeguard API.
// It takes a SafeguardClient and a Filter as parameters to construct the query.
// The function returns a slice of PolicyAsset and an error if the request fails.
//
// Parameters:
//   - c: A pointer to a SafeguardClient used to make the API request.
//   - fields: A Filter object used to specify query parameters.
//
// Returns:
//   - []PolicyAsset: A slice of PolicyAsset objects retrieved from the API.
//   - error: An error object if the request fails or the response cannot be unmarshaled.
func GetPolicyAssets(c *client.SafeguardClient, fields client.Filter) ([]PolicyAsset, error) {
	var policyAssets []PolicyAsset

	query := "PolicyAssets" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &policyAssets); err != nil {
		return nil, err
	}
	for i := range policyAssets {
		policyAssets[i].client = c
	}
	return policyAssets, nil
}

// GetPolicyAsset retrieves a PolicyAsset by its ID from the SafeguardClient.
// It takes a SafeguardClient, an integer ID, and optional fields to include in the query.
// It returns the PolicyAsset and an error if any occurred during the request.
//
// Parameters:
//   - c: A pointer to the SafeguardClient used to make the request.
//   - id: The ID of the PolicyAsset to retrieve.
//   - fields: Optional fields to include in the query.
//
// Returns:
//   - PolicyAsset: The retrieved PolicyAsset.
//   - error: An error if any occurred during the request.
func GetPolicyAsset(c *client.SafeguardClient, id int, fields client.Fields) (PolicyAsset, error) {
	var policyAsset PolicyAsset
	policyAsset.client = c

	query := fmt.Sprintf("PolicyAssets/%d", id)
	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return policyAsset, err
	}
	if err := json.Unmarshal(response, &policyAsset); err != nil {
		return policyAsset, err
	}
	return policyAsset, nil
}

// GetAssetGroups retrieves the asset groups associated with a specific policy asset.
// It takes a SafeguardClient and a Filter as parameters and returns a slice of AssetGroup and an error.
//
// Parameters:
//   - c: A pointer to the SafeguardClient used to make the request.
//   - fields: A Filter object used to specify query parameters.
//
// Returns:
//   - A slice of AssetGroup objects representing the asset groups associated with the policy asset.
//   - An error if the request fails or the response cannot be unmarshaled.
func (p PolicyAsset) GetAssetGroups(fields client.Filter) ([]AssetGroup, error) {
	var assetGroups []AssetGroup

	query := fmt.Sprintf("PolicyAssets/%d/AssetGroups", p.Id) + fields.ToQueryString()

	response, err := p.client.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &assetGroups); err != nil {
		return nil, err
	}
	return assetGroups, nil
}

// GetDirectoryServiceEntries retrieves the directory service entries associated with the PolicyAsset.
// It takes a SafeguardClient and a Filter as parameters and returns a slice of DirectoryServiceEntry and an error.
//
// Parameters:
//   - c: A pointer to the SafeguardClient used to make the request.
//   - fields: A Filter object used to specify query parameters.
//
// Returns:
//   - A slice of DirectoryServiceEntry objects.
//   - An error if the request fails or the response cannot be unmarshaled.
func (p PolicyAsset) GetDirectoryServiceEntries(fields client.Filter) ([]DirectoryServiceEntry, error) {
	var directoryServiceEntries []DirectoryServiceEntry

	query := fmt.Sprintf("PolicyAssets/%d/DirectoryServiceEntries", p.Id) + fields.ToQueryString()

	response, err := p.client.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &directoryServiceEntries); err != nil {
		return nil, err
	}
	return directoryServiceEntries, nil
}

// GetPolicies retrieves the policies associated with a specific PolicyAsset.
// It sends a GET request to the Safeguard API using the provided SafeguardClient and query fields.
//
// Parameters:
//   - c: A pointer to the SafeguardClient used to send the request.
//   - fields: A Filter object containing query parameters to be appended to the request URL.
//
// Returns:
//   - A slice of AssetPolicy objects representing the policies associated with the PolicyAsset.
//   - An error if the request fails or the response cannot be unmarshaled.
func (p PolicyAsset) GetPolicies(fields client.Filter) ([]AssetPolicy, error) {
	var policies []AssetPolicy

	query := fmt.Sprintf("PolicyAssets/%d/Policies", p.Id) + fields.ToQueryString()

	response, err := p.client.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &policies); err != nil {
		return nil, err
	}
	return policies, nil
}
