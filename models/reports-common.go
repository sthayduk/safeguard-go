package models

import "encoding/json"

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
