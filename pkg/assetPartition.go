package pkg

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sthayduk/safeguard-go/client"
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

// ToJson converts the AssetPartition struct to a JSON string representation.
// Returns:
//   - string: JSON string of the asset partition
//   - error: An error if JSON marshaling fails, nil otherwise
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

// Assign applies this password rule to the specified asset account.
// Parameters:
//   - assetAccount: The asset account to which the password rule should be applied
//
// Returns:
//   - AssetAccount: The updated asset account with the applied password rule
//   - error: An error if the assignment fails, nil otherwise
func (r AccountPasswordRule) Assign(assetAccount AssetAccount) (AssetAccount, error) {
	return UpdatePasswordProfile(assetAccount, r)
}

// ToJson converts the AccountPasswordRule struct to a JSON string representation.
// Returns:
//   - string: JSON string of the password rule
//   - error: An error if JSON marshaling fails, nil otherwise
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
//   - filter: Criteria to filter the returned partitions
//
// Returns:
//   - []AssetPartition: A slice of asset partitions matching the filter criteria
//   - error: An error if the request fails, nil otherwise
func GetAssetPartitions(filter client.Filter) ([]AssetPartition, error) {
	var AssetPartitions []AssetPartition

	query := "AssetPartitions" + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &AssetPartitions)
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
func GetAssetPartition(id int, fields client.Fields) (AssetPartition, error) {
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

// GetPasswordRules retrieves all password rules defined for this asset partition.
// It uses the current client instance to make the API request.
//
// Returns:
//   - []AccountPasswordRule: A slice of password rules defined for this partition
//   - error: An error if the request fails, nil otherwise
func (a AssetPartition) GetPasswordRules() ([]AccountPasswordRule, error) {
	return GetPasswordRules(a, client.Filter{})
}

// GetPasswordRules retrieves all password rules for a specific asset partition.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - assetPartition: The asset partition to get rules for
//   - filter: Criteria to filter the returned rules
//
// Returns:
//   - []AccountPasswordRule: A slice of password rules matching the filter
//   - error: An error if the request fails or no rules are found
func GetPasswordRules(assetPartition AssetPartition, filter client.Filter) ([]AccountPasswordRule, error) {
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

// DeleteAssetPartition deletes an asset partition with the specified ID.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The ID of the asset partition to delete
//
// Returns:
//   - error: An error if the deletion fails, nil otherwise
func DeleteAssetPartition(id int) error {
	query := fmt.Sprintf("AssetPartitions/%d", id)

	_, err := c.DeleteRequest(query)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes this asset partition from the system.
// Uses the DeleteAssetPartition function with the current client instance.
//
// Returns:
//   - error: An error if the deletion fails, nil otherwise
func (a AssetPartition) Delete() error {
	return DeleteAssetPartition(a.Id)
}
