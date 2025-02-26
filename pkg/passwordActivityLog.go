package pkg

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sthayduk/safeguard-go/client"
)

// UserLogProperties represents the user properties in a log entry including permissions,
// IP address, and user identification information
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

// LogEntry represents an individual log message with timestamp, status, and message content
type LogEntry struct {
	Timestamp time.Time `json:"Timestamp"`
	Status    string    `json:"Status"`
	Message   string    `json:"Message"`
}

// RequestStatus represents the status and timing information of a password request
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

// CustomScriptParameter represents a parameter used in custom scripts for password management
type CustomScriptParameter struct {
	Name     string `json:"Name"`
	Value    string `json:"Value"`
	Type     string `json:"Type"`
	TaskName string `json:"TaskName"`
}

// PasswordActivityLog represents a comprehensive log entry for password-related activities
// including user actions, asset details, and request status
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

// PasswordChangeSchedule represents a configuration for scheduled password changes
// including timing, repetition, and timezone settings
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

// PasswordCheckSchedule represents a configuration for scheduled password verification
// including timing, repetition, and timezone settings
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

// ScheduleInterval represents a time interval configuration for scheduled tasks
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
// including service account details, credentials, and connection settings
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

// CheckTaskState monitors the state of a password activity task.
// It polls the task status periodically until completion or timeout.
//
// Parameters:
//   - ctx: Context for timeout and cancellation control
//
// Returns:
//   - bool: true if task completed successfully, false if failed
//   - error: An error if monitoring fails or times out
func (p *PasswordActivityLog) CheckTaskState(ctx context.Context) (bool, error) {
	task, err := p.getMatchingAccountTask(ctx)
	if err != nil {
		return false, err
	}

	// Return the error only on the first check, as it indicates that the taskType is missing in the function
	isComplete, err := isTaskCompleteForType(task, p.LogTime, AccountTaskNames(p.Name))
	if err != nil {
		return false, err
	}
	if isComplete {
		return true, nil
	}

	// Return the error only on the first check, as it indicates that the taskType is missing in the function
	isFailed, err := isTaskFailedForType(task, p.LogTime, AccountTaskNames(p.Name))
	if err != nil {
		return false, err
	}
	if isFailed {
		return false, nil
	}

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return false, fmt.Errorf("timeout waiting for task state change")
		case <-ticker.C:
			newTask, err := p.getMatchingAccountTask(ctx)
			if err != nil {
				return false, err
			}

			isComplete, _ := isTaskCompleteForType(newTask, p.LogTime, AccountTaskNames(p.Name))
			if isComplete {
				return true, nil
			}

			isFailed, _ := isTaskFailedForType(newTask, p.LogTime, AccountTaskNames(p.Name))
			if isFailed {
				return false, nil
			}
		}
	}
}

// isTaskCompleteForType checks if a specific task type has completed successfully.
//
// Parameters:
//   - task: The AccountTaskData to check
//   - logTime: The time of the log entry to compare against
//   - taskType: The type of task being checked
//
// Returns:
//   - bool: true if the task is complete, false otherwise
//   - error: An error if the task type is unknown
func isTaskCompleteForType(task AccountTaskData, logTime time.Time, taskType AccountTaskNames) (bool, error) {
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
	case DiscoverSshKeys:
		successTime = task.TaskProperties.LastSuccessSshKeyDiscoveryDate
	case CheckApiKey:
		successCounter = task.TaskProperties.FailedApiKeyCheckAttempts
		return successCounter == 0, nil
	case ChangeApiKey:
		successCounter = task.TaskProperties.FailedApiKeyChangeAttempts
		return successCounter == 0, nil
	case SuspendAccount:
		successTime = task.TaskProperties.LastSuccessSuspendAccountDate
	case RestoreAccount:
		successTime = task.TaskProperties.LastSuccessRestoreAccountDate
	case ElevateAccount:
		successTime = task.TaskProperties.LastSuccessElevateAccountDate
	case DemoteAccount:
		successTime = task.TaskProperties.LastSuccessDemoteAccountDate
	default:
		return false, fmt.Errorf("unknown task type: %v", taskType)
	}

	return !successTime.IsZero() && (successTime.After(logTime) || successTime.Equal(logTime)), nil
}

// isTaskFailedForType checks if a specific task type has failed.
//
// Parameters:
//   - task: The AccountTaskData to check
//   - logTime: The time of the log entry to compare against
//   - taskType: The type of task being checked
//
// Returns:
//   - bool: true if the task has failed, false otherwise
//   - error: An error if the task type is unknown
func isTaskFailedForType(task AccountTaskData, logTime time.Time, taskType AccountTaskNames) (bool, error) {
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
	case DiscoverSshKeys:
		failureTime = task.TaskProperties.LastFailureSshKeyDiscoveryDate
	case CheckApiKey:
		failureCounter = task.TaskProperties.FailedApiKeyCheckAttempts
		return failureCounter > 0, nil
	case ChangeApiKey:
		failureCounter = task.TaskProperties.FailedApiKeyChangeAttempts
		return failureCounter > 0, nil
	case SuspendAccount:
		failureTime = task.TaskProperties.LastFailureSuspendAccountDate
	case RestoreAccount:
		failureTime = task.TaskProperties.LastFailureRestoreAccountDate
	case ElevateAccount:
		failureTime = task.TaskProperties.LastFailureElevateAccountDate
	case DemoteAccount:
		failureTime = task.TaskProperties.LastFailureDemoteAccountDate
	default:
		return false, fmt.Errorf("unknown task type: %v", taskType)
	}

	return !failureTime.IsZero() && (failureTime.After(logTime) || failureTime.Equal(logTime)), nil
}

// getTaskIdForType retrieves the task ID for a specific task type.
//
// Parameters:
//   - task: The AccountTaskData containing the task information
//   - taskType: The type of task to get the ID for
//
// Returns:
//   - string: The task ID
//   - error: An error if the task type is unknown
func getTaskIdForType(task AccountTaskData, taskType AccountTaskNames) (string, error) {
	switch taskType {
	case CheckPassword:
		return task.TaskProperties.LastPasswordCheckTaskId, nil
	case ChangePassword:
		return task.TaskProperties.LastPasswordChangeTaskId, nil
	case CheckSshKey:
		return task.TaskProperties.LastSshKeyCheckTaskId, nil
	case ChangeSshKey:
		return task.TaskProperties.LastSshKeyChangeTaskId, nil
	case DiscoverSshKeys:
		return task.TaskProperties.LastSshKeyDiscoveryTaskId, nil
	case SuspendAccount:
		return task.TaskProperties.LastSuspendAccountTaskId, nil
	case RestoreAccount:
		return task.TaskProperties.LastRestoreAccountTaskId, nil
	case ElevateAccount:
		return task.TaskProperties.LastElevateAccountTaskId, nil
	case DemoteAccount:
		return task.TaskProperties.LastDemoteAccountTaskId, nil
	default:
		return "", fmt.Errorf("unknown task type: %v", taskType)
	}
}

// getMatchingAccountTask retrieves the account task data matching the password activity log.
//
// Parameters:
//   - ctx: Context for timeout and cancellation control
//
// Returns:
//   - AccountTaskData: The matching task data
//   - error: An error if the task cannot be found or if there's a timeout
func (p PasswordActivityLog) getMatchingAccountTask(ctx context.Context) (AccountTaskData, error) {
	if p.Id == "" {
		return AccountTaskData{}, fmt.Errorf("invalid task ID")
	}

	_, err := uuid.Parse(p.Id)
	if err != nil {
		return AccountTaskData{}, fmt.Errorf("invalid task ID format: %v", err)
	}

	filter := client.Filter{}
	filter.AddFilter("Id", "eq", fmt.Sprintf("%d", p.AccountId))

	ticker := time.NewTicker(500 * time.Millisecond)
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
				taskId, err := getTaskIdForType(task, taskType)
				if err == nil && taskId == p.Id {
					return task, nil
				}
			}
		}
	}
}
