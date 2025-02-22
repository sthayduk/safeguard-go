package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/client"
)

func GetUserGroups(c *client.SafeguardClient, fields client.Filter) ([]UserGroup, error) {
	var userGroups []UserGroup

	query := "UserGroups" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &userGroups)
	return userGroups, nil
}

func GetUserGroup(c *client.SafeguardClient, id string, fields client.Fields) (UserGroup, error) {
	var userGroup UserGroup

	query := fmt.Sprintf("UserGroups/%s", id)
	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return userGroup, err
	}
	json.Unmarshal(response, &userGroup)
	return userGroup, nil
}
