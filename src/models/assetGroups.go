package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sthayduk/safeguard-go/src/client"
)

// AssetGroup represents a group of assets on the appliance for use in session policy.
// Only assets that support session access are allowed.
type AssetGroup struct {
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

func GetAssetGroups(c *client.SafeguardClient, fields client.Filter) ([]AssetGroup, error) {
	var assetGroup []AssetGroup

	query := "AssetGroups" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &assetGroup)
	return assetGroup, nil
}

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
	return assetGroup, nil
}
