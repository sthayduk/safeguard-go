package models

import (
	"encoding/json"

	"github.com/sthayduk/safeguard-go/src/client"
)

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

// AssetType represents the type of asset
type AssetType string

const (
	AssetTypeComputer      AssetType = "Computer"
	AssetTypeDirectory     AssetType = "Directory"
	AssetTypeDynamicAccess AssetType = "DynamicAccess"
	AssetTypeStarling      AssetType = "Starling"
)

// AssetSshHostKey represents an SSH Host Key used to identify assets
type AssetSshHostKey struct {
	Id            int    `json:"Id"`
	Fingerprint   string `json:"Fingerprint,omitempty"`
	Key           string `json:"Key,omitempty"`
	KeyType       string `json:"KeyType,omitempty"`
	Comment       string `json:"Comment,omitempty"`
	CanBeAccepted bool   `json:"CanBeAccepted,omitempty"`
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
