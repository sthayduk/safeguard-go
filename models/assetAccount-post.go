package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/client"
)

// ChangePassword initiates a password change for the asset account.
// Returns:
//   - PasswordActivityLog: Details of the password change activity
//   - error: An error if the request fails, nil otherwise
func (a AssetAccount) ChangePassword(c *client.SafeguardClient) (PasswordActivityLog, error) {
	var log PasswordActivityLog

	query := fmt.Sprintf("AssetAccounts/%d/ChangePassword", a.Id)

	response, err := c.PostRequest(query)
	if err != nil {
		return log, err
	}

	json.Unmarshal(response, &log)
	return log, nil
}

// CheckPassword verifies the current password of the asset account.
// Returns:
//   - PasswordActivityLog: Details of the password check activity
//   - error: An error if the request fails, nil otherwise
func (a AssetAccount) CheckPassword(c *client.SafeguardClient) (PasswordActivityLog, error) {
	var log PasswordActivityLog

	query := fmt.Sprintf("AssetAccounts/%d/CheckPassword", a.Id)

	response, err := c.PostRequest(query)
	if err != nil {
		return log, err
	}

	json.Unmarshal(response, &log)
	return log, nil
}
