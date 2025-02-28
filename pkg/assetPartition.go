package pkg

import (
	"encoding/json"
	"fmt"
	"time"
)

// AssetPartition represents a collection of assets and accounts along with management configuration.
// The partition defines boundaries for asset management and access control within Safeguard.
type AssetPartition struct {
	Id                       int             `json:"Id"`
	Name                     string          `json:"Name"`
	Description              string          `json:"Description"`
	CreatedDate              time.Time       `json:"CreatedDate"`
	CreatedByUserId          int             `json:"CreatedByUserId"`
	CreatedByUserDisplayName string          `json:"CreatedByUserDisplayName"`
	ManagedBy                []ManagedByUser `json:"ManagedBy"`
	DefaultProfileId         int             `json:"DefaultProfileId"`
	DefaultProfileName       string          `json:"DefaultProfileName"`
	DefaultSshKeyProfileId   int             `json:"DefaultSshKeyProfileId"`
	DefaultSshKeyProfileName string          `json:"DefaultSshKeyProfileName"`
}

// ToJson serializes the AssetPartition instance into a JSON string representation.
//
// Returns:
//   - (string): JSON representation of the asset partition
//   - (error): An error if JSON marshaling fails
func (u AssetPartition) ToJson() (string, error) {
	assetPartitionJSON, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(assetPartitionJSON), nil
}

// AccountPasswordRule defines the requirements and constraints for generating and validating
// account passwords within an asset partition. It specifies character requirements, restrictions,
// and other password complexity rules.
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

// Assign associates this password rule with the specified asset account.
// This operation updates the asset account's password profile with the current rule.
//
// Parameters:
//   - assetAccount: The asset account to modify
//
// Returns:
//   - (AssetAccount): The updated asset account
//   - (error): An error if the assignment fails
func (r AccountPasswordRule) Assign(assetAccount AssetAccount) (AssetAccount, error) {
	return UpdatePasswordProfile(assetAccount, r)
}

// ToJson serializes the AccountPasswordRule instance into a JSON string representation.
//
// Returns:
//   - (string): JSON representation of the password rule
//   - (error): An error if JSON marshaling fails
func (r AccountPasswordRule) ToJson() (string, error) {
	ruleJSON, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(ruleJSON), nil
}

// GetAssetPartitions retrieves all asset partitions matching the specified filter criteria.
//
// Parameters:
//   - filter: Query parameters to filter the results
//
// Returns:
//   - ([]AssetPartition): Slice of matching asset partitions
//   - (error): An error if the API request fails
func GetAssetPartitions(filter Filter) ([]AssetPartition, error) {
	var AssetPartitions []AssetPartition

	query := "AssetPartitions" + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &AssetPartitions)
	return AssetPartitions, nil
}

// GetAssetPartition retrieves a single asset partition by its ID.
//
// Parameters:
//   - id: Unique identifier of the asset partition
//   - fields: Optional fields to include in the response
//
// Returns:
//   - (AssetPartition): The requested asset partition
//   - (error): An error if the API request fails
func GetAssetPartition(id int, fields Fields) (AssetPartition, error) {
	var AssetPartition AssetPartition

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

// GetPasswordRules retrieves all password rules associated with this asset partition.
//
// Returns:
//   - ([]AccountPasswordRule): Slice of password rules for this partition
//   - (error): An error if the API request fails
func (a AssetPartition) GetPasswordRules() ([]AccountPasswordRule, error) {
	return GetPasswordRules(a, Filter{})
}

// GetPasswordRules retrieves password rules for the specified asset partition.
//
// Parameters:
//   - assetPartition: The partition to get rules for
//   - filter: Query parameters to filter the results
//
// Returns:
//   - ([]AccountPasswordRule): Slice of matching password rules
//   - (error): An error if the API request fails or no rules are found
func GetPasswordRules(assetPartition AssetPartition, filter Filter) ([]AccountPasswordRule, error) {
	var PasswordRules []AccountPasswordRule

	query := fmt.Sprintf("AssetPartitions/%d/Profiles", assetPartition.Id) + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &PasswordRules); err != nil {
		return nil, err
	}

	if len(PasswordRules) == 0 {
		return PasswordRules, fmt.Errorf("no password rules found for asset partition %d", assetPartition.Id)
	}

	return PasswordRules, nil
}

// DeleteAssetPartition removes an asset partition from the system.
//
// Parameters:
//   - id: Unique identifier of the asset partition to delete
//
// Returns:
//   - (error): An error if the deletion fails
func DeleteAssetPartition(id int) error {
	query := fmt.Sprintf("AssetPartitions/%d", id)

	_, err := c.DeleteRequest(query)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes this asset partition from the system.
//
// Returns:
//   - (error): An error if the deletion fails
func (a AssetPartition) Delete() error {
	return DeleteAssetPartition(a.Id)
}
