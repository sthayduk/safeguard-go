package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/client"
)

func GetRoles(c *client.SafeguardClient, fields client.Filter) ([]Role, error) {
	var userRoles []Role

	query := "Roles" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &userRoles)
	return userRoles, nil
}

func GetRole(c *client.SafeguardClient, id string, fields client.Fields) (Role, error) {
	var userRole Role

	query := fmt.Sprintf("Roles/%s", id)
	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return userRole, err
	}
	json.Unmarshal(response, &userRole)
	return userRole, nil
}
