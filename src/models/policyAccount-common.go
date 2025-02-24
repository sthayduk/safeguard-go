package models

import (
	"encoding/json"

	"github.com/sthayduk/safeguard-go/src/client"
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

// Asset represents a Safeguard asset
type Asset struct {
	Id                 int    `json:"Id"`
	Name               string `json:"Name"`
	NetworkAddress     string `json:"NetworkAddress"`
	AssetPartitionId   int    `json:"AssetPartitionId"`
	AssetPartitionName string `json:"AssetPartitionName"`
}

func (p PolicyAccount) ToJson() (string, error) {
	policyAccountJSON, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(policyAccountJSON), nil
}
