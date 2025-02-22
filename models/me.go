package models

import (
	"encoding/json"

	"itdesign.at/safeguard-go/client"
)

func GetMe(c *client.SafeguardClient, fields client.Filter) (User, error) {
	query := "me" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return User{}, err
	}

	var me User
	if err := json.Unmarshal(response, &me); err != nil {
		return User{}, err
	}

	return me, nil
}
