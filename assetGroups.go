package safeguard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// AssetGroup represents a group of assets on the appliance for use in session policy.
// Only assets that support session access are allowed.
type AssetGroup struct {
	apiClient *SafeguardClient `json:"-"`

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

func (a AssetGroup) SetClient(c *SafeguardClient) any {
	a.apiClient = c
	return a
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

// ToJson serializes the AssetGroup instance into a JSON string representation.
//
// Returns:
//   - (string): JSON representation of the asset group
//   - (error): An error if JSON marshaling fails
func (a AssetGroup) ToJson() (string, error) {
	assetGroupJSON, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(assetGroupJSON), nil
}

// GetAssetGroups retrieves all asset groups matching the specified filter criteria.
//
// Parameters:
//   - filter: Query parameters to filter the results
//
// Returns:
//   - ([]AssetGroup): Slice of matching asset groups
//   - (error): An error if the API request fails
func (c *SafeguardClient) GetAssetGroups(filter Filter) ([]AssetGroup, error) {
	var assetGroup []AssetGroup

	query := "AssetGroups" + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &assetGroup)

	return addClientToSlice(c, assetGroup), nil
}

// GetAssetGroup retrieves a single asset group by its ID.
//
// Parameters:
//   - id: Unique identifier of the asset group
//   - fields: Optional fields to include in the response
//
// Returns:
//   - (AssetGroup): The requested asset group
//   - (error): An error if the API request fails
func (c *SafeguardClient) GetAssetGroup(id int, fields Fields) (AssetGroup, error) {
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

	return addClient(c, assetGroup), nil
}

// UpdateAssetGroup modifies an existing asset group.
//
// Parameters:
//   - id: Unique identifier of the asset group to update
//   - assetGroup: Modified asset group data
//
// Returns:
//   - (AssetGroup): The updated asset group
//   - (error): An error if the update fails
func (c *SafeguardClient) UpdateAssetGroup(id int, assetGroup AssetGroup) (AssetGroup, error) {
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

	return addClient(c, assetGroup), nil
}

// Update persists any changes made to this AssetGroup instance.
//
// Returns:
//   - (AssetGroup): The updated asset group
//   - (error): An error if the update fails
func (a AssetGroup) Update() (AssetGroup, error) {
	return a.apiClient.UpdateAssetGroup(a.Id, a)
}

// DeleteAssetGroup removes an asset group from the system.
//
// Parameters:
//   - id: Unique identifier of the asset group to delete
//
// Returns:
//   - (error): An error if the deletion fails
func (c *SafeguardClient) DeleteAssetGroup(id int) error {
	query := fmt.Sprintf("AssetGroups/%d", id)

	_, err := c.DeleteRequest(query)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes this asset group from the system.
//
// Returns:
//   - (error): An error if the deletion fails
func (a AssetGroup) Delete() error {
	return a.apiClient.DeleteAssetGroup(a.Id)
}
