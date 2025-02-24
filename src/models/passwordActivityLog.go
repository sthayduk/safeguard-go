package models

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sthayduk/safeguard-go/src/client"
)

// UserLogProperties represents the user properties in a log entry
type UserLogProperties struct {
	ClientIpAddress           string `json:"ClientIpAddress,omitempty"`
	UserName                  string `json:"UserName,omitempty"`
	DomainName                string `json:"DomainName,omitempty"`
	UserDisplayName           string `json:"UserDisplayName,omitempty"`
	UserWasGlobalAdmin        bool   `json:"UserWasGlobalAdmin,omitempty"`
	UserWasDirectoryAdmin     bool   `json:"UserWasDirectoryAdmin,omitempty"`
	UserWasAuditor            bool   `json:"UserWasAuditor,omitempty"`
	UserWasApplicationAuditor bool   `json:"UserWasApplicationAuditor,omitempty"`
	UserWasSystemAuditor      bool   `json:"UserWasSystemAuditor,omitempty"`
	UserWasAssetAdmin         bool   `json:"UserWasAssetAdmin,omitempty"`
	UserWasPartitionOwner     bool   `json:"UserWasPartitionOwner,omitempty"`
	UserWasApplianceAdmin     bool   `json:"UserWasApplianceAdmin,omitempty"`
	UserWasPolicyAdmin        bool   `json:"UserWasPolicyAdmin,omitempty"`
	UserWasUserAdmin          bool   `json:"UserWasUserAdmin,omitempty"`
	UserWasHelpdeskAdmin      bool   `json:"UserWasHelpdeskAdmin,omitempty"`
	UserWasOperationsAdmin    bool   `json:"UserWasOperationsAdmin,omitempty"`
}

// LogEntry represents an individual log message
type LogEntry struct {
	Timestamp time.Time `json:"Timestamp"`
	Status    string    `json:"Status"`
	Message   string    `json:"Message"`
}

// RequestStatus represents the status of a password request
type RequestStatus struct {
	State              string    `json:"State"`
	PercentComplete    int       `json:"PercentComplete"`
	Cancellable        bool      `json:"Cancellable"`
	AcceptedTime       time.Time `json:"AcceptedTime"`
	AcceptanceDuration string    `json:"AcceptanceDuration"`
	StartTime          time.Time `json:"StartTime"`
	QueuedDuration     string    `json:"QueuedDuration"`
	EndTime            time.Time `json:"EndTime"`
	RunningDuration    string    `json:"RunningDuration"`
	TotalDuration      string    `json:"TotalDuration"`
	Message            string    `json:"Message"`
}

// CustomScriptParameter represents a parameter for custom scripts
type CustomScriptParameter struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
	Type  string `json:"Type"`
}

// PasswordActivityLog represents a password activity log entry
type PasswordActivityLog struct {
	client *client.SafeguardClient

	Id                string                  `json:"Id"`
	LogTime           time.Time               `json:"LogTime"`
	UserId            int                     `json:"UserId"`
	UserProperties    UserLogProperties       `json:"UserProperties"`
	ApplianceId       string                  `json:"ApplianceId"`
	ApplianceName     string                  `json:"ApplianceName"`
	EventName         string                  `json:"EventName"`
	EventDisplayName  string                  `json:"EventDisplayName"`
	Name              string                  `json:"Name"`
	AssetId           int                     `json:"AssetId"`
	AssetName         string                  `json:"AssetName"`
	AccountId         int                     `json:"AccountId"`
	AccountName       string                  `json:"AccountName"`
	AccountDomainName string                  `json:"AccountDomainName"`
	NetworkAddress    string                  `json:"NetworkAddress"`
	RequestStatus     RequestStatus           `json:"RequestStatus"`
	Log               []LogEntry              `json:"Log"`
	ConnectionProps   ConnectionProperties    `json:"ConnectionProperties"`
	CustomParams      []CustomScriptParameter `json:"CustomScriptParameters"`
}

// CheckTaskState monitors the state of a password activity task.
// It polls the task status periodically for up to 30 seconds.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//
// Returns:
//   - bool: true if task completed successfully, false if failed
//   - error: An error if monitoring fails or times out
func (p *PasswordActivityLog) CheckTaskState() (bool, error) {
	task, err := p.getMatchingAccountTask()
	if err != nil {
		return false, err
	}

	if isTaskCompleteForType(task, p.LogTime, AccountTaskNames(p.Name)) {
		return true, nil
	}
	if isTaskFailedForType(task, p.LogTime, AccountTaskNames(p.Name)) {
		return false, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return false, fmt.Errorf("timeout waiting for task state change")
		case <-ticker.C:
			newTask, err := p.getMatchingAccountTask()
			if err != nil {
				return false, err
			}

			if isTaskCompleteForType(newTask, p.LogTime, AccountTaskNames(p.Name)) {
				return true, nil
			}
			if isTaskFailedForType(newTask, p.LogTime, AccountTaskNames(p.Name)) {
				return false, nil
			}
		}
	}
}

func isTaskCompleteForType(task AccountTaskData, logTime time.Time, taskType AccountTaskNames) bool {
	var successTime time.Time
	var successCounter int

	switch taskType {
	case CheckPassword:
		successTime = task.TaskProperties.LastSuccessPasswordCheckDate
	case ChangePassword:
		successTime = task.TaskProperties.LastSuccessPasswordChangeDate
	case CheckSshKey:
		successTime = task.TaskProperties.LastSuccessSshKeyCheckDate
	case ChangeSshKey:
		successTime = task.TaskProperties.LastSuccessSshKeyChangeDate
	case DiscoverAccounts:
		successTime = task.TaskProperties.LastSuccessSshKeyDiscoveryDate
	case CheckApiKey:
		successCounter = task.TaskProperties.FailedApiKeyCheckAttempts
		return successCounter == 0
	case ChangeApiKey:
		successCounter = task.TaskProperties.FailedApiKeyChangeAttempts
		return successCounter == 0
	default:
		return false
	}

	return !successTime.IsZero() && (successTime.After(logTime) || successTime.Equal(logTime))
}

func isTaskFailedForType(task AccountTaskData, logTime time.Time, taskType AccountTaskNames) bool {
	var failureTime time.Time
	var failureCounter int

	switch taskType {
	case CheckPassword:
		failureTime = task.TaskProperties.LastFailurePasswordCheckDate
	case ChangePassword:
		failureTime = task.TaskProperties.LastFailurePasswordChangeDate
	case CheckSshKey:
		failureTime = task.TaskProperties.LastFailureSshKeyCheckDate
	case ChangeSshKey:
		failureTime = task.TaskProperties.LastFailureSshKeyChangeDate
	case DiscoverAccounts:
		failureTime = task.TaskProperties.LastFailureSshKeyDiscoveryDate
	case CheckApiKey:
		failureCounter = task.TaskProperties.FailedApiKeyCheckAttempts
		return failureCounter > 0
	case ChangeApiKey:
		failureCounter = task.TaskProperties.FailedApiKeyChangeAttempts
		return failureCounter > 0
	default:
		return false
	}

	return !failureTime.IsZero() && (failureTime.After(logTime) || failureTime.Equal(logTime))
}

func getTaskIdForType(task AccountTaskData, taskType AccountTaskNames) string {
	switch taskType {
	case CheckPassword:
		return task.TaskProperties.LastPasswordCheckTaskId
	case ChangePassword:
		return task.TaskProperties.LastPasswordChangeTaskId
	case CheckSshKey:
		return task.TaskProperties.LastSshKeyCheckTaskId
	case ChangeSshKey:
		return task.TaskProperties.LastSshKeyChangeTaskId
	default:
		return ""
	}
}

func (p PasswordActivityLog) getMatchingAccountTask() (AccountTaskData, error) {
	if p.Id == "" {
		return AccountTaskData{}, fmt.Errorf("invalid task ID")
	}

	_, err := uuid.Parse(p.Id)
	if err != nil {
		return AccountTaskData{}, fmt.Errorf("invalid task ID format: %v", err)
	}

	filter := client.Filter{}
	filter.AddFilter("Id", "eq", fmt.Sprintf("%d", p.AccountId))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return AccountTaskData{}, fmt.Errorf("timeout waiting for task to become available")
		case <-ticker.C:
			taskData, err := GetAccountTaskSchedules(p.client, AccountTaskNames(p.Name), filter)
			if err != nil {
				continue
			}

			if len(taskData) == 0 {
				continue
			}

			taskType := AccountTaskNames(p.Name)
			for _, task := range taskData {
				if getTaskIdForType(task, taskType) == p.Id {
					return task, nil
				}
			}
		}
	}
}

// PasswordChangeSchedule represents a schedule used by a partition profile to change passwords
type PasswordChangeSchedule struct {
	Id                  int       `json:"Id,omitempty"`
	Name                string    `json:"Name,omitempty"`
	ScheduleType        string    `json:"ScheduleType,omitempty"`
	TimeZoneId          string    `json:"TimeZoneId,omitempty"`
	Description         string    `json:"Description,omitempty"`
	StartDate           time.Time `json:"StartDate,omitempty"`
	RepeatInterval      int       `json:"RepeatInterval,omitempty"`
	RepeatIntervalUnit  string    `json:"RepeatIntervalUnit,omitempty"`
	MonthlyScheduleType string    `json:"MonthlyScheduleType,omitempty"`
	DayOfMonth          int       `json:"DayOfMonth,omitempty"`
	DayOfWeek           string    `json:"DayOfWeek,omitempty"`
	WeekOfMonth         string    `json:"WeekOfMonth,omitempty"`
	TimeOfDayType       string    `json:"TimeOfDayType,omitempty"`
	TimeOfDay           string    `json:"TimeOfDay,omitempty"`
	NoEndDate           bool      `json:"NoEndDate,omitempty"`
	EndDate             time.Time `json:"EndDate,omitempty"`
}

// PasswordCheckSchedule represents a schedule used by a partition profile to check passwords
type PasswordCheckSchedule struct {
	Id                  int       `json:"Id,omitempty"`
	Name                string    `json:"Name,omitempty"`
	ScheduleType        string    `json:"ScheduleType,omitempty"`
	TimeZoneId          string    `json:"TimeZoneId,omitempty"`
	Description         string    `json:"Description,omitempty"`
	StartDate           time.Time `json:"StartDate,omitempty"`
	RepeatInterval      int       `json:"RepeatInterval,omitempty"`
	RepeatIntervalUnit  string    `json:"RepeatIntervalUnit,omitempty"`
	MonthlyScheduleType string    `json:"MonthlyScheduleType,omitempty"`
	DayOfMonth          int       `json:"DayOfMonth,omitempty"`
	DayOfWeek           string    `json:"DayOfWeek,omitempty"`
	WeekOfMonth         string    `json:"WeekOfMonth,omitempty"`
	TimeOfDayType       string    `json:"TimeOfDayType,omitempty"`
	TimeOfDay           string    `json:"TimeOfDay,omitempty"`
	NoEndDate           bool      `json:"NoEndDate,omitempty"`
	EndDate             time.Time `json:"EndDate,omitempty"`
}

// ScheduleInterval represents interval of time in which to execute tasks
type ScheduleInterval struct {
	RepeatInterval     int    `json:"RepeatInterval,omitempty"`
	RepeatIntervalUnit string `json:"RepeatIntervalUnit,omitempty"`
}

// Constants for schedule types and intervals
const (
	// Schedule Types
	ScheduleTypeOnce    = "Once"
	ScheduleTypeDaily   = "Daily"
	ScheduleTypeWeekly  = "Weekly"
	ScheduleTypeMonthly = "Monthly"

	// Time of Day Types
	TimeOfDayTypeAny   = "Any"
	TimeOfDayTypeExact = "Exact"

	// Monthly Schedule Types
	MonthlyDayOfMonth = "DayOfMonth"
	MonthlyDayOfWeek  = "DayOfWeek"

	// Interval Units
	IntervalUnitMinutes = "Minutes"
	IntervalUnitHours   = "Hours"
	IntervalUnitDays    = "Days"
	IntervalUnitWeeks   = "Weeks"
	IntervalUnitMonths  = "Months"
)

// ConnectionProperties represents connection-specific properties for various services
type ConnectionProperties struct {
	// Service Account Properties
	ServiceAccountUniqueObjectId             string `json:"ServiceAccountUniqueObjectId,omitempty"`
	ServiceAccountSecurityId                 string `json:"ServiceAccountSecurityId,omitempty"`
	ServiceAccountId                         int    `json:"ServiceAccountId,omitempty"`
	ServiceAccountName                       string `json:"ServiceAccountName,omitempty"`
	ServiceAccountDomainName                 string `json:"ServiceAccountDomainName,omitempty"`
	ServiceAccountDistinguishedName          string `json:"ServiceAccountDistinguishedName,omitempty"`
	ServiceAccountNetbiosName                string `json:"ServiceAccountNetbiosName,omitempty"`
	EffectiveServiceAccountName              string `json:"EffectiveServiceAccountName,omitempty"`
	EffectiveServiceAccountDistinguishedName string `json:"EffectiveServiceAccountDistinguishedName,omitempty"`

	// Credential Properties
	ServiceAccountCredentialType string `json:"ServiceAccountCredentialType,omitempty"`
	ServiceAccountPassword       string `json:"ServiceAccountPassword,omitempty"`
	ServiceAccountHasPassword    bool   `json:"ServiceAccountHasPassword,omitempty"`
	ServiceAccountSshKey         SshKey `json:"ServiceAccountSshKey,omitempty"` // Changed from SshKeyData to SshKey
	ServiceAccountHasSshKey      bool   `json:"ServiceAccountHasSshKey,omitempty"`

	// Connection Settings
	UseSslEncryption     bool `json:"UseSslEncryption,omitempty"`
	VerifySslCertificate bool `json:"VerifySslCertificate,omitempty"`
	Port                 int  `json:"Port,omitempty"`

	// Asset Properties
	ServiceAccountAssetId                  int    `json:"ServiceAccountAssetId,omitempty"`
	ServiceAccountAssetName                string `json:"ServiceAccountAssetName,omitempty"`
	ServiceAccountAssetPlatformId          int    `json:"ServiceAccountAssetPlatformId,omitempty"`
	ServiceAccountAssetPlatformType        string `json:"ServiceAccountAssetPlatformType,omitempty"`
	ServiceAccountAssetPlatformDisplayName string `json:"ServiceAccountAssetPlatformDisplayName,omitempty"`
}
