package safeguard

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// TaskNames defines the supported task names for account tasks
type TaskNames string

const (
	Archive                     TaskNames = "Archive"
	ChangeApiKey                TaskNames = "ChangeApiKey"
	ChangeFile                  TaskNames = "ChangeFile"
	ChangePassword              TaskNames = "ChangePassword"
	ChangeSshKey                TaskNames = "ChangeSshKey"
	CheckApiKey                 TaskNames = "CheckApiKey"
	CheckFile                   TaskNames = "CheckFile"
	CheckPassword               TaskNames = "CheckPassword"
	CheckSshKey                 TaskNames = "CheckSshKey"
	DemoteAccount               TaskNames = "DemoteAccount"
	DirectoryAssetDeleteSync    TaskNames = "DirectoryAssetDeleteSync"
	DirectoryAssetSync          TaskNames = "DirectoryAssetSync"
	DirectoryProviderDeleteSync TaskNames = "DirectoryProviderDeleteSync"
	DirectoryProviderSync       TaskNames = "DirectoryProviderSync"
	DiscoverAccounts            TaskNames = "DiscoverAccounts"
	DiscoverAssets              TaskNames = "DiscoverAssets"
	DiscoverServices            TaskNames = "DiscoverServices"
	DiscoverSshHostKey          TaskNames = "DiscoverSshHostKey"
	DiscoverSshKeys             TaskNames = "DiscoverSshKeys"
	ElevateAccount              TaskNames = "ElevateAccount"
	InstallSshKey               TaskNames = "InstallSshKey"
	LocalIdentityProviderSync   TaskNames = "LocalIdentityProviderSync"
	PasswordSyncAccounts        TaskNames = "PasswordSyncAccounts"
	RestoreAccount              TaskNames = "RestoreAccount"
	RetrieveSshHostKey          TaskNames = "RetrieveSshHostKey"
	RevokeSshKey                TaskNames = "RevokeSshKey"
	SshKeySyncAccounts          TaskNames = "SshKeySyncAccounts"
	SuspendAccount              TaskNames = "SuspendAccount"
	TestConnection              TaskNames = "TestConnection"
	UnknownTask                 TaskNames = "Unknown"
	UpdateDependentAsset        TaskNames = "UpdateDependentAsset"
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

// ActivityLog represents a comprehensive log entry for password-related activities
// including user actions, asset details, and request status
type ActivityLog struct {
	apiClient *SafeguardClient

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

func (a ActivityLog) SetClient(c *SafeguardClient) any {
	a.apiClient = c
	return a
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
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
//	defer cancel()
//	success, err := passwordLog.CheckTaskState(ctx)
//
// Parameters:
//   - ctx: Context for timeout and cancellation control. Should include a reasonable timeout.
//
// Returns:
//   - bool: true if task completed successfully, false if task failed or was cancelled
//   - error: Error if monitoring fails, times out, or context is cancelled
func (p *ActivityLog) CheckTaskState(ctx context.Context) (bool, error) {
	task, err := p.getMatchingAccountTask(ctx)
	if err != nil {
		return false, err
	}

	// Return the error only on the first check, as it indicates that the taskType is missing in the function
	isComplete, err := isTaskCompleteForType(task, p.LogTime, TaskNames(p.Name))
	if err != nil {
		return false, err
	}
	if isComplete {
		return true, nil
	}

	// Return the error only on the first check, as it indicates that the taskType is missing in the function
	isFailed, err := isTaskFailedForType(task, p.LogTime, TaskNames(p.Name))
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

			isComplete, _ := isTaskCompleteForType(newTask, p.LogTime, TaskNames(p.Name))
			if isComplete {
				return true, nil
			}

			isFailed, _ := isTaskFailedForType(newTask, p.LogTime, TaskNames(p.Name))
			if isFailed {
				return false, nil
			}
		}
	}
}

// isTaskCompleteForType determines if a specific account task has completed successfully
// by checking the appropriate timestamp or counter for the given task type.
//
// Example:
//
//	complete, err := isTaskCompleteForType(task, logTime, CheckPassword)
//
// Parameters:
//   - task: AccountTaskData containing the task execution details and properties
//   - logTime: Reference time to compare against task completion timestamps
//   - taskType: Specific type of account task being monitored (e.g., CheckPassword, ChangePassword)
//
// Returns:
//   - bool: true if the task has completed successfully after the logTime
//   - error: Error if the taskType is not recognized or supported
func isTaskCompleteForType(task AccountTaskData, logTime time.Time, taskType TaskNames) (bool, error) {
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

// isTaskFailedForType determines if a specific account task has failed
// by checking the appropriate failure timestamp or counter for the given task type.
//
// Example:
//
//	failed, err := isTaskFailedForType(task, logTime, CheckPassword)
//
// Parameters:
//   - task: AccountTaskData containing the task execution details and properties
//   - logTime: Reference time to compare against task failure timestamps
//   - taskType: Specific type of account task being monitored (e.g., CheckPassword, ChangePassword)
//
// Returns:
//   - bool: true if the task has failed after the logTime
//   - error: Error if the taskType is not recognized or supported
func isTaskFailedForType(task AccountTaskData, logTime time.Time, taskType TaskNames) (bool, error) {
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

// getTaskIdForType retrieves the appropriate task identifier for a specific task type
// from the AccountTaskData properties.
//
// Example:
//
//	taskId, err := getTaskIdForType(task, CheckPassword)
//
// Parameters:
//   - task: AccountTaskData containing task identifiers for various operations
//   - taskType: Specific type of account task to get the ID for
//
// Returns:
//   - string: Task identifier for the specified task type
//   - error: Error if the taskType is not recognized or supported
func getTaskIdForType(task AccountTaskData, taskType TaskNames) (string, error) {
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

// getMatchingAccountTask retrieves the AccountTaskData that matches this password activity log.
// It continuously polls until the matching task is found or the context is cancelled.
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
//	defer cancel()
//	taskData, err := passwordLog.getMatchingAccountTask(ctx)
//
// Parameters:
//   - ctx: Context for timeout and cancellation control. Should include a reasonable timeout.
//
// Returns:
//   - AccountTaskData: Task data matching the password activity log's ID and account
//   - error: Error if the task cannot be found, invalid ID format, or context timeout/cancellation
func (p ActivityLog) getMatchingAccountTask(ctx context.Context) (AccountTaskData, error) {
	if p.Id == "" {
		return AccountTaskData{}, fmt.Errorf("invalid task ID")
	}

	_, err := uuid.Parse(p.Id)
	if err != nil {
		return AccountTaskData{}, fmt.Errorf("invalid task ID format: %v", err)
	}

	filter := Filter{}
	filter.AddFilter("Id", OpEqual, fmt.Sprintf("%d", p.AccountId))

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return AccountTaskData{}, fmt.Errorf("timeout waiting for task to become available")
		case <-ticker.C:
			taskData, err := p.apiClient.GetAccountTaskSchedules(TaskNames(p.Name), filter)
			if err != nil {
				continue
			}

			if len(taskData) == 0 {
				continue
			}

			taskType := TaskNames(p.Name)
			for _, task := range taskData {
				taskId, err := getTaskIdForType(task, taskType)
				if err == nil && taskId == p.Id {
					return task, nil
				}
			}
		}
	}
}
