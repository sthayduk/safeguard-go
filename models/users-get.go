package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/client"
)

func GetUsers(c *client.SafeguardClient, fields client.Filter) ([]User, error) {
	var users []User

	query := "users" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &users)
	return users, nil
}

func GetUser(c *client.SafeguardClient, id string, fields client.Fields) (User, error) {
	var user User

	query := fmt.Sprintf("users/%s", id)
	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return user, err
	}
	json.Unmarshal(response, &user)
	return user, nil
}

func GetLinkedAccounts(c *client.SafeguardClient, id string) ([]PolicyAccount, error) {
	var linkedAccounts []PolicyAccount

	query := fmt.Sprintf("users/%s/LinkedPolicyAccounts", id)

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &linkedAccounts)
	return linkedAccounts, nil
}

func (u User) GetLinkedAccounts(c *client.SafeguardClient) ([]PolicyAccount, error) {
	return GetLinkedAccounts(c, fmt.Sprintf("%d", u.Id))
}

func GetUserRoles(c *client.SafeguardClient, id string) ([]Role, error) {
	var roles []Role

	query := fmt.Sprintf("users/%s/roles", id)

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &roles)
	return roles, nil
}

func (u User) GetRoles(c *client.SafeguardClient) ([]Role, error) {
	return GetUserRoles(c, fmt.Sprintf("%d", u.Id))
}

func GetGroups(c *client.SafeguardClient, id string) ([]UserGroup, error) {
	var userGroups []UserGroup

	query := fmt.Sprintf("users/%s/UserGroups", id)

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &userGroups)
	return userGroups, nil
}

func (u User) GetGroups(c *client.SafeguardClient) ([]UserGroup, error) {
	return GetGroups(c, fmt.Sprintf("%d", u.Id))
}
