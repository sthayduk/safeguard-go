package safeguard

import (
	"encoding/json"
	"fmt"
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

// ToJson serializes an AssetPolicy object into a JSON string.
//
// This method converts the AssetPolicy instance into a JSON-formatted string,
// including all defined fields. Empty or zero-valued fields are included in
// the output.
//
// Example:
//
//	policy := AssetPolicy{
//	    PolicyName: "Linux Servers",
//	    AccessRequestType: AccessRequestTypeSSH
//	}
//	json, err := policy.ToJson()
//
// Parameters:
//   - none
//
// Returns:
//   - string: A JSON representation of the AssetPolicy object
//   - error: An error if JSON marshaling fails, nil otherwise
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

// ToJson converts a PolicyAsset to its JSON string representation.
//
// Example:
//
//	asset := PolicyAsset{...}
//	json, err := asset.ToJson()
//
// Returns:
//   - string: A JSON-formatted string containing all non-empty fields of the PolicyAsset
//   - error: An error if JSON marshaling fails, nil otherwise
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

// GetPolicyAssets retrieves policy assets based on filter criteria.
//
// This method returns assets that match all specified filter conditions. Commonly
// used filters include Disabled, PlatformId, and AssetPartitionId.
//
// Example:
//
//	filter := Filter{}
//	filter.AddFilter("Disabled", "eq", "false")
//	filter.AddFilter("PlatformId", "eq", "1")
//	assets, err := GetPolicyAssets(filter)
//
// Parameters:
//   - fields: A Filter object containing field comparisons and ordering preferences
//
// Returns:
//   - []PolicyAsset: A slice of PolicyAsset objects matching the filter criteria
//   - error: An error if the request or response parsing fails, nil otherwise
func GetPolicyAssets(fields Filter) ([]PolicyAsset, error) {
	var policyAssets []PolicyAsset

	query := "PolicyAssets" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &policyAssets); err != nil {
		return nil, err
	}

	return policyAssets, nil
}

// GetPolicyAsset retrieves a single policy asset by its unique identifier.
//
// The method supports including additional related objects based on the fields parameter.
// Common fields include Platform, SessionAccess, and SshHostKey.
//
// Example:
//
//	fields := Fields{}
//	fields.Add("Platform", "SessionAccess")
//	asset, err := GetPolicyAsset(123, fields)
//
// Parameters:
//   - id: The unique identifier of the policy asset to retrieve
//   - fields: Optional Fields object specifying which related objects to include
//
// Returns:
//   - PolicyAsset: The requested policy asset with all specified related objects
//   - error: An error if the asset is not found or request fails, nil otherwise
func GetPolicyAsset(id int, fields Fields) (PolicyAsset, error) {
	var policyAsset PolicyAsset

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

// GetAssetGroups retrieves all asset groups containing this policy asset.
//
// Returns both direct group memberships and nested group memberships if any exist.
// The results can be filtered using the fields parameter.
//
// Example:
//
//	filter := Filter{}
//	filter.AddFilter("Disabled", "eq", "false")
//	groups, err := asset.GetAssetGroups(filter)
//
// Parameters:
//   - fields: A Filter object to restrict which groups are returned
//
// Returns:
//   - []AssetGroup: A slice of AssetGroup objects this asset belongs to
//   - error: An error if the request or response parsing fails, nil otherwise
func (p PolicyAsset) GetAssetGroups(fields Filter) ([]AssetGroup, error) {
	var assetGroups []AssetGroup

	query := fmt.Sprintf("PolicyAssets/%d/AssetGroups", p.Id) + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &assetGroups); err != nil {
		return nil, err
	}
	return assetGroups, nil
}

// GetDirectoryServiceEntries retrieves directory entries for directory assets.
//
// This method is primarily used with directory server assets to list their
// contained directory entries. Not applicable for non-directory assets.
//
// Example:
//
//	filter := Filter{}
//	entries, err := directoryAsset.GetDirectoryServiceEntries(filter)
//
// Parameters:
//   - fields: A Filter object to restrict which entries are returned
//
// Returns:
//   - []DirectoryServiceEntry: A slice of directory entries from this asset
//   - error: An error if the request or response parsing fails, nil otherwise
func (p PolicyAsset) GetDirectoryServiceEntries(fields Filter) ([]DirectoryServiceEntry, error) {
	var directoryServiceEntries []DirectoryServiceEntry

	query := fmt.Sprintf("PolicyAssets/%d/DirectoryServiceEntries", p.Id) + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &directoryServiceEntries); err != nil {
		return nil, err
	}
	return directoryServiceEntries, nil
}

// GetPolicies retrieves all access policies affecting this policy asset.
//
// Returns policies that grant access to this asset, either directly or through
// asset group membership. Includes details about how access was granted.
//
// Example:
//
//	filter := Filter{}
//	policies, err := asset.GetPolicies(filter)
//
// Parameters:
//   - fields: A Filter object to restrict which policies are returned
//
// Returns:
//   - []AssetPolicy: A slice of policies granting access to this asset
//   - error: An error if the request or response parsing fails, nil otherwise
func (p PolicyAsset) GetPolicies(fields Filter) ([]AssetPolicy, error) {
	var policies []AssetPolicy

	query := fmt.Sprintf("PolicyAssets/%d/Policies", p.Id) + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &policies); err != nil {
		return nil, err
	}
	return policies, nil
}
