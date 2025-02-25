package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sthayduk/safeguard-go/client"
)

// AssetGroup represents a group of assets on the appliance for use in session policy.
// Only assets that support session access are allowed.
type AssetGroup struct {
	client *client.SafeguardClient

	Id                       int               `json:"Id"`
	Name                     string            `json:"Name"`
	Description              string            `json:"Description"`
	IsDynamic                bool              `json:"IsDynamic"`
	Assets                   []PolicyAsset     `json:"Assets"`
	AssetGroupingRule        AssetGroupingRule `json:"AssetGroupingRule"`
	CreatedDate              time.Time         `json:"CreatedDate"`
	CreatedByUserId          int               `json:"CreatedByUserId"`
	CreatedByUserDisplayName string            `json:"CreatedByUserDisplayName"`
}

// AssetGroupingRule represents rules for automatically grouping assets
type AssetGroupingRule struct {
	Description        string             `json:"Description"`
	Enabled            bool               `json:"Enabled"`
	RuleConditionGroup RuleConditionGroup `json:"RuleConditionGroup"`
}

// RuleConditionGroup represents a group of conditions for asset grouping rules
type RuleConditionGroup struct {
	LogicalJoinType string                 `json:"LogicalJoinType"`
	Children        []RuleConditionOrGroup `json:"Children"`
}

// RuleConditionOrGroup represents either a condition or a group of conditions
type RuleConditionOrGroup struct {
	TaggingGroupingCondition      *TaggingGroupingCondition `json:"TaggingGroupingCondition,omitempty"`
	TaggingGroupingConditionGroup string                    `json:"TaggingGroupingConditionGroup,omitempty"`
}

// TaggingGroupingCondition represents a single condition for grouping
type TaggingGroupingCondition struct {
	ObjectAttribute string `json:"ObjectAttribute"`
	CompareType     string `json:"CompareType"`
	CompareValue    string `json:"CompareValue"`
}

// ToJson converts an AssetGroup to its JSON string representation
func (a AssetGroup) ToJson() (string, error) {
	assetGroupJSON, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(assetGroupJSON), nil
}

// GetAssetGroups retrieves a list of asset groups from the Safeguard API based on the provided filter fields.
// It sends a GET request to the "AssetGroups" endpoint with the query parameters specified in the fields.
//
// Parameters:
//   - c: A pointer to a SafeguardClient instance used to send the request.
//   - fields: A Filter object containing the query parameters for filtering the asset groups.
//
// Returns:
//   - A slice of AssetGroup objects retrieved from the API.
//   - An error if the request fails or the response cannot be unmarshaled.
func GetAssetGroups(c *client.SafeguardClient, filter client.Filter) ([]AssetGroup, error) {
	var assetGroup []AssetGroup

	query := "AssetGroups" + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &assetGroup)

	for i := range assetGroup {
		assetGroup[i].client = c
	}

	return assetGroup, nil
}

// GetAssetGroup retrieves an AssetGroup by its ID from the Safeguard API.
// It takes a SafeguardClient, an integer ID, and optional fields to include in the query.
// It returns the AssetGroup and an error, if any occurred during the request.
//
// Parameters:
//   - c: A pointer to the SafeguardClient used to make the request.
//   - id: The ID of the AssetGroup to retrieve.
//   - fields: Optional fields to include in the query.
//
// Returns:
//   - AssetGroup: The retrieved AssetGroup.
//   - error: An error if the request failed or the response could not be unmarshaled.
func GetAssetGroup(c *client.SafeguardClient, id int, fields client.Fields) (AssetGroup, error) {
	var assetGroup AssetGroup

	query := fmt.Sprintf("AssetGroups/%d", id)
	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return assetGroup, err
	}
	json.Unmarshal(response, &assetGroup)

	assetGroup.client = c
	return assetGroup, nil
}

// UpdateAssetGroup updates an existing asset group identified by the given ID.
// It sends a PUT request to the Safeguard API with the updated asset group data.
//
// Parameters:
//   - c: A pointer to the SafeguardClient used to make the API request.
//   - id: The ID of the asset group to be updated.
//   - assetGroup: The AssetGroup object containing the updated data.
//
// Returns:
//   - AssetGroup: The updated AssetGroup object returned by the API.
//   - error: An error object if an error occurred during the update process, otherwise nil.
func UpdateAssetGroup(c *client.SafeguardClient, id int, assetGroup AssetGroup) (AssetGroup, error) {
	query := fmt.Sprintf("AssetGroups/%d", id)

	assetGroupJSON, err := assetGroup.ToJson()
	if err != nil {
		return AssetGroup{}, err
	}

	response, err := c.PutRequest(query, bytes.NewReader([]byte(assetGroupJSON)))
	if err != nil {
		return AssetGroup{}, err
	}

	err = json.Unmarshal(response, &assetGroup)
	if err != nil {
		return AssetGroup{}, err
	}

	return assetGroup, nil
}

// Update updates the current AssetGroup instance in the database.
// It returns the updated AssetGroup and an error if the update fails.
//
// Returns:
//   - (AssetGroup): The updated AssetGroup instance.
//   - (error): An error if the update operation fails.
func (a AssetGroup) Update() (AssetGroup, error) {
	return UpdateAssetGroup(a.client, a.Id, a)
}

// DeleteAssetGroup deletes an asset group identified by the given ID using the provided SafeguardClient.
//
// Parameters:
//   - c: A pointer to a SafeguardClient instance used to make the delete request.
//   - id: An integer representing the ID of the asset group to be deleted.
//
// Returns:
//   - error: An error object if the delete request fails, otherwise nil.
func DeleteAssetGroup(c *client.SafeguardClient, id int) error {
	query := fmt.Sprintf("AssetGroups/%d", id)

	_, err := c.DeleteRequest(query)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes the AssetGroup from the system.
// It calls the DeleteAssetGroup function with the client and Id of the AssetGroup.
// Returns an error if the deletion fails.
func (a AssetGroup) Delete() error {
	return DeleteAssetGroup(a.client, a.Id)
}
