package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/client"
)

func GetAssetPartitions(c *client.SafeguardClient, fields client.Filter) ([]AssetPartition, error) {
	var AssetPartitions []AssetPartition

	query := "AssetPartitions" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &AssetPartitions)
	return AssetPartitions, nil
}

func GetAssetPartition(c *client.SafeguardClient, id string, fields client.Fields) (AssetPartition, error) {
	var AssetPartition AssetPartition

	query := fmt.Sprintf("AssetPartitions/%s", id)
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

func GetPasswordRules(c *client.SafeguardClient, AssetPartitionId int, filter client.Filter) ([]AccountPasswordRule, error) {
	var PasswordRules []AccountPasswordRule

	query := fmt.Sprintf("AssetPartitions/%d/PasswordRules", AssetPartitionId) + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &PasswordRules)
	return PasswordRules, nil
}

func (a AssetPartition) GetPasswordRules(c *client.SafeguardClient) ([]AccountPasswordRule, error) {
	return GetPasswordRules(c, a.Id, client.Filter{})
}
