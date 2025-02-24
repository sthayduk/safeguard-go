package models

import (
	"fmt"

	"github.com/sthayduk/safeguard-go/src/client"
)

func DeleteUser(c *client.SafeguardClient, id int) error {

	query := fmt.Sprintf("users/%d", id)

	_, err := c.DeleteRequest(query)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Delete() error {
	return DeleteUser(u.client, u.Id)
}
