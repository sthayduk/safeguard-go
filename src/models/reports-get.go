package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/src/client"
)

// GetAccountTaskSchedules retrieves the account task schedules for a given task name and filter.
// It sends a GET request to the Safeguard API and unmarshals the response into a slice of AccountTaskData.
//
// Parameters:
//   - c: A pointer to a SafeguardClient instance used to make the API request.
//   - taskName: The name of the account task to retrieve schedules for.
//   - filter: A Filter instance to apply to the query.
//
// Returns:
//   - A slice of AccountTaskData containing the retrieved account task schedules.
//   - An error if the request fails or the response cannot be unmarshaled.
func GetAccountTaskSchedules(c *client.SafeguardClient, taskName AccountTaskNames, filter client.Filter) ([]AccountTaskData, error) {
	var accountTaskSchedules []AccountTaskData

	query := fmt.Sprintf("Reports/Tasks/AccountTaskSchedules/%s/%s", taskName, filter.ToQueryString())

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &accountTaskSchedules)
	return accountTaskSchedules, nil
}
