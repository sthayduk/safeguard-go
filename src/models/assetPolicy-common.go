package models

import "encoding/json"

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
	AccessRequestTypePassword AccessRequestType = "Password"
	AccessRequestTypeSSH      AccessRequestType = "SSH"
	AccessRequestTypeAPI      AccessRequestType = "API"
	AccessRequestTypeRDP      AccessRequestType = "RDP"
)

// ToJson converts an AssetPolicy to its JSON string representation
func (a AssetPolicy) ToJson() (string, error) {
	assetPolicyJSON, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(assetPolicyJSON), nil
}
