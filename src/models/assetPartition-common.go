package models

import (
	"encoding/json"
	"time"

	"github.com/sthayduk/safeguard-go/src/client"
)

// AssetPartition represents a collection of assets and accounts along with management configuration
type AssetPartition struct {
	client *client.SafeguardClient

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

// AccountPasswordRule represents a password rule used to generate account passwords
type AccountPasswordRule struct {
	Id                                      int       `json:"Id"`
	IsSystemOwned                           bool      `json:"IsSystemOwned"`
	AssetPartitionId                        int       `json:"AssetPartitionId"`
	AssetPartitionName                      string    `json:"AssetPartitionName"`
	CreatedDate                             time.Time `json:"CreatedDate"`
	CreatedByUserId                         int       `json:"CreatedByUserId"`
	CreatedByUserDisplayName                string    `json:"CreatedByUserDisplayName"`
	Name                                    string    `json:"Name"`
	Description                             string    `json:"Description"`
	MaxCharacters                           int       `json:"MaxCharacters"`
	MinCharacters                           int       `json:"MinCharacters"`
	AllowUppercaseCharacters                bool      `json:"AllowUppercaseCharacters"`
	MinUppercaseCharacters                  int       `json:"MinUppercaseCharacters"`
	InvalidUppercaseCharacters              []string  `json:"InvalidUppercaseCharacters"`
	MaxConsecutiveUppercaseCharacters       int       `json:"MaxConsecutiveUppercaseCharacters"`
	AllowLowercaseCharacters                bool      `json:"AllowLowercaseCharacters"`
	MinLowercaseCharacters                  int       `json:"MinLowercaseCharacters"`
	InvalidLowercaseCharacters              []string  `json:"InvalidLowercaseCharacters"`
	MaxConsecutiveLowercaseCharacters       int       `json:"MaxConsecutiveLowercaseCharacters"`
	AllowNumericCharacters                  bool      `json:"AllowNumericCharacters"`
	MinNumericCharacters                    int       `json:"MinNumericCharacters"`
	InvalidNumericCharacters                []string  `json:"InvalidNumericCharacters"`
	MaxConsecutiveNumericCharacters         int       `json:"MaxConsecutiveNumericCharacters"`
	AllowNonAlphaNumericCharacters          bool      `json:"AllowNonAlphaNumericCharacters"`
	MinNonAlphaNumericCharacters            int       `json:"MinNonAlphaNumericCharacters"`
	NonAlphaNumericRestrictionType          string    `json:"NonAlphaNumericRestrictionType"`
	AllowedNonAlphaNumericCharacters        []string  `json:"AllowedNonAlphaNumericCharacters"`
	InvalidNonAlphaNumericCharacters        []string  `json:"InvalidNonAlphaNumericCharacters"`
	MaxConsecutiveNonAlphaNumericCharacters int       `json:"MaxConsecutiveNonAlphaNumericCharacters"`
	AllowedFirstCharacterType               string    `json:"AllowedFirstCharacterType"`
	AllowedLastCharacterType                string    `json:"AllowedLastCharacterType"`
	MaxConsecutiveAlphabeticCharacters      int       `json:"MaxConsecutiveAlphabeticCharacters"`
	MaxConsecutiveAlphaNumericCharacters    int       `json:"MaxConsecutiveAlphaNumericCharacters"`
	RepeatedCharacterRestriction            string    `json:"RepeatedCharacterRestriction"`
}

func (r AccountPasswordRule) ToJson() (string, error) {
	ruleJSON, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(ruleJSON), nil
}
