package models

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/client"
)

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
