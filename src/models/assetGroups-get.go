package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/src/client"
)

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
