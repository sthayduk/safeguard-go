package models

import (
	"encoding/json"
	"fmt"
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

// GetAssetPartitions retrieves a list of asset partitions from Safeguard.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - fields: Filter criteria for the request
//
// Returns:
//   - []AssetPartition: A slice of asset partitions matching the filter criteria
//   - error: An error if the request fails, nil otherwise
func GetAssetPartitions(c *client.SafeguardClient, fields client.Filter) ([]AssetPartition, error) {
	var AssetPartitions []AssetPartition

	query := "AssetPartitions" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &AssetPartitions)
	for i := range AssetPartitions {
		AssetPartitions[i].client = c
	}
	return AssetPartitions, nil
}

// GetAssetPartition retrieves details for a specific asset partition by ID.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The numeric ID of the asset partition to retrieve
//   - fields: Specific fields to include in the response
//
// Returns:
//   - AssetPartition: The requested asset partition object
//   - error: An error if the request fails, nil otherwise
func GetAssetPartition(c *client.SafeguardClient, id int, fields client.Fields) (AssetPartition, error) {
	var AssetPartition AssetPartition
	AssetPartition.client = c

	query := fmt.Sprintf("AssetPartitions/%d", id)
	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return AssetPartition, err
	}
	json.Unmarshal(response, &AssetPartition)
	return AssetPartition, nil
}

// GetPasswordRules retrieves the password rules for this asset partition.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//
// Returns:
//   - []AccountPasswordRule: A slice of password rules
//   - error: An error if the request fails, nil otherwise
func (a AssetPartition) GetPasswordRules() ([]AccountPasswordRule, error) {
	return GetPasswordRules(a.client, a.Id, client.Filter{})
}

// GetPasswordRules retrieves the password rules for a specific asset partition.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - AssetPartitionId: The ID of the asset partition
//   - filter: Filter criteria for the request
//
// Returns:
//   - []AccountPasswordRule: A slice of password rules
//   - error: An error if the request fails, nil otherwise
func GetPasswordRules(c *client.SafeguardClient, AssetPartitionId int, filter client.Filter) ([]AccountPasswordRule, error) {
	var PasswordRules []AccountPasswordRule

	query := fmt.Sprintf("AssetPartitions/%d/PasswordRules", AssetPartitionId) + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &PasswordRules); err != nil {
		return nil, err
	}
	return PasswordRules, nil
}
