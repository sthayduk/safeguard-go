package models

import (
	"encoding/json"
	"time"
)

// AssetPartition represents a collection of assets and accounts along with management configuration
type AssetPartition struct {
	Id                       int        `json:"Id"`
	Name                     string     `json:"Name"`
	Description              string     `json:"Description"`
	CreatedDate              time.Time  `json:"CreatedDate"`
	CreatedByUserId          int        `json:"CreatedByUserId"`
	CreatedByUserDisplayName string     `json:"CreatedByUserDisplayName"`
	ManagedBy                []Identity `json:"ManagedBy"`
	DefaultProfileId         int        `json:"DefaultProfileId"`
	DefaultProfileName       string     `json:"DefaultProfileName"`
	DefaultSshKeyProfileId   int        `json:"DefaultSshKeyProfileId"`
	DefaultSshKeyProfileName string     `json:"DefaultSshKeyProfileName"`
}

// Identity represents an identity that can manage the partition
type Identity struct {
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

func (u AssetPartition) ToJson() (string, error) {
	assetPartitionJSON, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(assetPartitionJSON), nil
}
