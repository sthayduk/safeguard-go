package pkg

import (
	"encoding/json"
	"fmt"

	"github.com/sthayduk/safeguard-go/client"
)

// AccountTaskNames defines the supported task names for account tasks
type AccountTaskNames string

const (
	CheckPassword    AccountTaskNames = "CheckPassword"
	ChangePassword   AccountTaskNames = "ChangePassword"
	CheckSshKey      AccountTaskNames = "CheckSshKey"
	ChangeSshKey     AccountTaskNames = "ChangeSshKey"
	DiscoverAccounts AccountTaskNames = "DiscoverAccounts"
	CheckApiKey      AccountTaskNames = "CheckApiKey"
	ChangeApiKey     AccountTaskNames = "ChangeApiKey"
	SuspendAccount   AccountTaskNames = "SuspendAccount"
	RestoreAccount   AccountTaskNames = "RestoreAccount"  // Added
	DiscoverSshKeys  AccountTaskNames = "DiscoverSshKeys" // Added
	InstallSshKey    AccountTaskNames = "InstallSshKey"   // Added
	ElevateAccount   AccountTaskNames = "ElevateAccount"  // Added
	DemoteAccount    AccountTaskNames = "DemoteAccount"   // Added
)

// ScheduledAccountTask represents a scheduled task that runs against accounts
type ScheduledAccountTask struct {
	ID             string           `json:"id"`
	Name           string           `json:"name"`
	Description    string           `json:"description,omitempty"`
	Schedule       string           `json:"schedule"` // Cron expression
	Enabled        bool             `json:"enabled"`
	TaskType       AccountTaskNames `json:"taskType"`
	LastRun        string           `json:"lastRun,omitempty"`
	NextRun        string           `json:"nextRun,omitempty"`
	TaskProperties TaskProperties   `json:"taskProperties"`
}

// TimeOfDayInterval represents a time interval configuration
type TimeOfDayInterval struct {
	StartHour   int `json:"StartHour"`
	StartMinute int `json:"StartMinute"`
	EndHour     int `json:"EndHour"`
	EndMinute   int `json:"EndMinute"`
	Iterations  int `json:"Iterations"`
}

// Schedule represents a task schedule configuration
type Schedule struct {
	Id                        int                 `json:"Id"`
	Name                      string              `json:"Name"`
	Description               string              `json:"Description"`
	NotifyOwnersOnly          bool                `json:"NotifyOwnersOnly"`
	NotifyOwnersOnMismatch    bool                `json:"NotifyOwnersOnMismatch"`
	ResetOnMismatch           bool                `json:"ResetOnMismatch"`
	ScheduleType              string              `json:"ScheduleType"`
	TimeZoneId                string              `json:"TimeZoneId"`
	TimeZoneDisplayName       string              `json:"TimeZoneDisplayName"`
	RepeatInterval            int                 `json:"RepeatInterval"`
	RepeatMonthlyScheduleType string              `json:"RepeatMonthlyScheduleType"`
	RepeatWeekOfMonth         string              `json:"RepeatWeekOfMonth"`
	RepeatDayOfWeek           string              `json:"RepeatDayOfWeek"`
	RepeatDayOfMonth          int                 `json:"RepeatDayOfMonth"`
	RepeatDaysOfWeek          []string            `json:"RepeatDaysOfWeek"`
	TimeOfDayType             string              `json:"TimeOfDayType"`
	StartHour                 int                 `json:"StartHour"`
	StartMinute               int                 `json:"StartMinute"`
	TimeOfDayIntervals        []TimeOfDayInterval `json:"TimeOfDayIntervals"`
}

// AccountTaskData represents platform task information for an asset or directory account
type AccountTaskData struct {
	Id                       int              `json:"Id"`
	Name                     string           `json:"Name"`
	DistinguishedName        string           `json:"DistinguishedName"`
	DomainName               string           `json:"DomainName"`
	Description              string           `json:"Description"`
	Disabled                 bool             `json:"Disabled"`
	Asset                    Asset            `json:"Asset"`
	Platform                 Platform         `json:"Platform"`
	Schedule                 Schedule         `json:"Schedule"`
	TaskProperties           TaskProperties   `json:"TaskProperties"`
	AssetName                string           `json:"assetName,omitempty"`
	AccountName              string           `json:"accountName,omitempty"`
	AccountDomainName        string           `json:"accountDomainName,omitempty"`
	AccountDistinguishedName string           `json:"accountDistinguishedName,omitempty"`
	TaskName                 AccountTaskNames `json:"taskName"`
	Status                   string           `json:"status,omitempty"`
	LastExecuted             string           `json:"lastExecuted,omitempty"`
	NextScheduled            string           `json:"nextScheduled,omitempty"`
	ErrorMessage             string           `json:"errorMessage,omitempty"`
}

func (a AccountTaskData) ToJson() (string, error) {
	userJSON, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(userJSON), nil
}

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

	query := fmt.Sprintf("Reports/Tasks/AccountTaskSchedules/%s%s", taskName, filter.ToQueryString())

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response, &accountTaskSchedules)
	return accountTaskSchedules, nil
}
