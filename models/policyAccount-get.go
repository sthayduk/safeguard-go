package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/client"
)

func GetPolicyAccounts(c *client.SafeguardClient, fields client.Filter) ([]PolicyAccount, error) {
	var policyAccounts []PolicyAccount

	query := "PolicyAccounts" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &policyAccounts)
	return policyAccounts, nil
}

func GetPolicyAccount(c *client.SafeguardClient, id string, fields client.Fields) (PolicyAccount, error) {
	var policyAccount PolicyAccount

	query := fmt.Sprintf("PolicyAccounts/%s", id)
	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return policyAccount, err
	}
	json.Unmarshal(response, &policyAccount)
	return policyAccount, nil
}
