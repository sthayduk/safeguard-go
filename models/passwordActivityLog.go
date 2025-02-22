package models

import "time"

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
