package models

import (
	"encoding/json"
	"fmt"
)

// ChangePassword initiates a password change operation for the asset account.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//
// Returns:
//   - PasswordActivityLog: Log details of the password change activity
//   - error: An error if the password change fails or cannot be initiated
func (a AssetAccount) ChangePassword() (PasswordActivityLog, error) {
	var log PasswordActivityLog
	log.client = a.client

	query := fmt.Sprintf("AssetAccounts/%d/ChangePassword", a.Id)

	response, err := a.client.PostRequest(query, nil)
	if err != nil {
		return log, err
	}

	json.Unmarshal(response, &log)
	return log, nil
}

// CheckPassword verifies if the current password for the asset account is valid.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//
// Returns:
//   - PasswordActivityLog: Log details of the password check activity
//   - error: An error if the password check fails or cannot be initiated
func (a AssetAccount) CheckPassword() (PasswordActivityLog, error) {
	var log PasswordActivityLog
	log.client = a.client

	query := fmt.Sprintf("AssetAccounts/%d/CheckPassword", a.Id)

	response, err := a.client.PostRequest(query, nil)
	if err != nil {
		return log, err
	}

	json.Unmarshal(response, &log)
	return log, nil
}
