package pkg

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sthayduk/safeguard-go/client"
)

// SessionAccessAccountType represents the type of session access account
type SessionAccessAccountType string

const (
	None          SessionAccessAccountType = "None"
	LinkedAccount SessionAccessAccountType = "LinkedAccount"
	Custom        SessionAccessAccountType = "Custom"
)

// TimeOfDay represents the hours in a day (0-23)
type TimeOfDay int

// DayOfWeek represents days of the week
type DayOfWeek string

const (
	Monday    DayOfWeek = "Monday"
	Tuesday   DayOfWeek = "Tuesday"
	Wednesday DayOfWeek = "Wednesday"
	Thursday  DayOfWeek = "Thursday"
	Friday    DayOfWeek = "Friday"
	Saturday  DayOfWeek = "Saturday"
	Sunday    DayOfWeek = "Sunday"
)

// NotificationContact represents contact info for different roles in access policy
type NotificationContact struct {
	Name            string                  `json:"Name"`
	EmailAddress    string                  `json:"EmailAddress,omitempty"`
	ContactType     NotificationContactType `json:"ContactType"`
	UserId          int                     `json:"UserId,omitempty"`
	UserDisplayName string                  `json:"UserDisplayName,omitempty"`
	UserGroupId     int                     `json:"UserGroupId,omitempty"`
	UserGroupName   string                  `json:"UserGroupName,omitempty"`
}

// NotificationContactType represents the type of notification contact
type NotificationContactType string

const (
	Email NotificationContactType = "Email"
	SMS   NotificationContactType = "SMS"
)

// ApproverSet represents a set of identities required to approve an access request
type ApproverSet struct {
	Name        string          `json:"Name"`
	Description string          `json:"Description,omitempty"`
	IsDefault   bool            `json:"IsDefault"`
	Identities  []ManagedByUser `json:"Identities"`
}

// ReasonCode represents a predefined reason for access requests
type ReasonCode struct {
	Id          int    `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"Description,omitempty"`
	Category    string `json:"Category,omitempty"`
}

// AccessPolicy represents security configuration governing the access to assets and accounts
type AccessPolicy struct {
	Id                          int                         `json:"Id"`
	Name                        string                      `json:"Name"`
	Description                 string                      `json:"Description,omitempty"`
	RoleId                      int                         `json:"RoleId"`
	RoleName                    string                      `json:"RoleName"`
	RolePriority                int                         `json:"RolePriority"`
	Priority                    int                         `json:"Priority"`
	AccountCount                int                         `json:"AccountCount"`
	AssetCount                  int                         `json:"AssetCount"`
	AccountGroupCount           int                         `json:"AccountGroupCount"`
	AssetGroupCount             int                         `json:"AssetGroupCount"`
	CreatedDate                 time.Time                   `json:"CreatedDate"`
	CreatedByUserId             int                         `json:"CreatedByUserId"`
	CreatedByUserDisplayName    string                      `json:"CreatedByUserDisplayName"`
	RequesterProperties         RequesterProperties         `json:"RequesterProperties"`
	ApproverProperties          ApproverProperties          `json:"ApproverProperties"`
	ReviewerProperties          ReviewerProperties          `json:"ReviewerProperties"`
	AccessRequestProperties     AccessRequestProperties     `json:"AccessRequestProperties"`
	SessionProperties           *SessionProperties          `json:"SessionProperties,omitempty"`
	EmergencyAccessProperties   EmergencyAccessProperties   `json:"EmergencyAccessProperties"`
	ApproverSets                []ApproverSet               `json:"ApproverSets"`
	Reviewers                   []ManagedByUser             `json:"Reviewers"`
	NotificationContacts        []NotificationContact       `json:"NotificationContacts"`
	ReasonCodes                 []ReasonCode                `json:"ReasonCodes"`
	ScopeItems                  []PolicyScopeItem           `json:"ScopeItems"`
	ExpirationDate              *time.Time                  `json:"ExpirationDate,omitempty"`
	IsExpired                   bool                        `json:"IsExpired"`
	InvalidConnectionPolicy     bool                        `json:"InvalidConnectionPolicy"`
	HourlyRestrictionProperties HourlyRestrictionProperties `json:"HourlyRestrictionProperties"`
}

func (a AccessPolicy) ToJson() (string, error) {
	accessPolicyJSON, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(accessPolicyJSON), nil
}

// GetReasonCodes returns the list of reason codes assigned to this policy.
// If no reason codes are assigned, it returns an empty slice.
//
// Returns:
//   - A slice of ReasonCode objects assigned to this policy.
func (a AccessPolicy) GetReasonCodes() []ReasonCode {
	if a.ReasonCodes == nil {
		return []ReasonCode{}
	}
	return a.ReasonCodes
}

// AccessRequestProperties represents configuration governing access requests
type AccessRequestProperties struct {
	AccessRequestType                AccessRequestType        `json:"AccessRequestType"`
	AllowSimultaneousAccess          bool                     `json:"AllowSimultaneousAccess"`
	MaximumSimultaneousReleases      int                      `json:"MaximumSimultaneousReleases"`
	ChangePasswordAfterCheckin       bool                     `json:"ChangePasswordAfterCheckin"`
	ChangeSshKeyAfterCheckin         bool                     `json:"ChangeSshKeyAfterCheckin"`
	AllowSessionPasswordRelease      bool                     `json:"AllowSessionPasswordRelease"`
	AllowSessionSshKeyRelease        bool                     `json:"AllowSessionSshKeyRelease"`
	IncludePasswordRelease           bool                     `json:"IncludePasswordRelease"`
	IncludeSshKeyRelease             bool                     `json:"IncludeSshKeyRelease"`
	SessionAccessAccountType         SessionAccessAccountType `json:"SessionAccessAccountType"`
	SessionAccessAccounts            []int                    `json:"SessionAccessAccounts"`
	TerminateExpiredSessions         bool                     `json:"TerminateExpiredSessions"`
	AllowLinkedAccountPasswordAccess bool                     `json:"AllowLinkedAccountPasswordAccess"`
	PassphraseProtectSshKey          bool                     `json:"PassphraseProtectSshKey"`
	UseAltLoginName                  bool                     `json:"UseAltLoginName"`
	LinkedAccountScopeFiltering      bool                     `json:"LinkedAccountScopeFiltering"`
}

// SessionProperties represents session-specific configuration
type SessionProperties struct {
	SessionModuleConnectionId          int                                 `json:"SessionModuleConnectionId"`
	SessionConnectionPolicyRef         string                              `json:"SessionConnectionPolicyRef"`
	RdpShowWallpaper                   bool                                `json:"RdpShowWallpaper"`
	RemoteDesktopApplicationProperties *RemoteDesktopApplicationProperties `json:"RemoteDesktopApplicationProperties"`
}

// RemoteDesktopApplicationProperties represents RDP application-specific settings
type RemoteDesktopApplicationProperties struct {
	ApplicationHostAssetId      *int    `json:"ApplicationHostAssetId"`
	ApplicationHostAsset        *Asset  `json:"ApplicationHostAsset"`
	ApplicationHostAccountId    *int    `json:"ApplicationHostAccountId"`
	ApplicationHostLoginAccount *string `json:"ApplicationHostLoginAccount"`
	ApplicationDisplayName      *string `json:"ApplicationDisplayName"`
	ApplicationAlias            *string `json:"ApplicationAlias"`
	ApplicationProgram          *string `json:"ApplicationProgram"`
	ApplicationCmdLine          *string `json:"ApplicationCmdLine"`
	ApplicationHostUserSupplied bool    `json:"ApplicationHostUserSupplied"`
}

// PolicyApproverProperties represents settings related to approving an access request
type PolicyApproverProperties struct {
	RequireApproval        bool `json:"RequireApproval"`
	RequireReapproval      bool `json:"RequireReapproval"`
	AutoApproveRequests    bool `json:"AutoApproveRequests"`
	AllowSelfApproval      bool `json:"AllowSelfApproval"`
	RequiredApprovers      int  `json:"RequiredApprovers"`
	RequireTimeRestriction bool `json:"RequireTimeRestriction"`
	MaximumTimeRestriction int  `json:"MaximumTimeRestriction"`
}

// PolicyReviewerProperties represents settings related to reviewing a password request
type PolicyReviewerProperties struct {
	RequireReview     bool `json:"RequireReview"`
	AllowSelfReview   bool `json:"AllowSelfReview"`
	RequiredReviewers int  `json:"RequiredReviewers"`
}

// PolicyRequesterProperties represents settings for requesting asset/accounts
type PolicyRequesterProperties struct {
	AllowEmergencyAccess       bool `json:"AllowEmergencyAccess"`
	AllowUseRequestComments    bool `json:"AllowUseRequestComments"`
	RequireUseRequestComments  bool `json:"RequireUseRequestComments"`
	MaximumDaysUntilExpiration int  `json:"MaximumDaysUntilExpiration"`
}

// PolicyEmergencyAccessProperties represents settings related to emergency access
type PolicyEmergencyAccessProperties struct {
	RequireEmergencyTicketNumber bool   `json:"RequireEmergencyTicketNumber"`
	EmergencyTicketSystem        string `json:"EmergencyTicketSystem"`
}

// PolicyScopeItem represents requestable items governed by policy
type PolicyScopeItem struct {
	Id                 int    `json:"Id"`
	Name               string `json:"Name"`
	Description        string `json:"Description"`
	AssetPartitionId   int    `json:"AssetPartitionId"`
	AssetPartitionName string `json:"AssetPartitionName"`
	Type               string `json:"Type"`
}

// ApproverProperties represents settings related to approving access requests
type ApproverProperties struct {
	RequireApproval                                bool `json:"RequireApproval"`
	PendingApprovalEscalationEnabled               bool `json:"PendingApprovalEscalationEnabled"`
	PendingApprovalDurationBeforeEscalationDays    int  `json:"PendingApprovalDurationBeforeEscalationDays"`
	PendingApprovalDurationBeforeEscalationHours   int  `json:"PendingApprovalDurationBeforeEscalationHours"`
	PendingApprovalDurationBeforeEscalationMinutes int  `json:"PendingApprovalDurationBeforeEscalationMinutes"`
}

// ReviewerProperties represents settings related to reviewing access requests
type ReviewerProperties struct {
	RequiredReviewers                            int  `json:"RequiredReviewers"`
	RequireReviewerComment                       bool `json:"RequireReviewerComment"`
	AllowSubsequentAccessRequestsWithoutReview   bool `json:"AllowSubsequentAccessRequestsWithoutReview"`
	PendingReviewEscalationEnabled               bool `json:"PendingReviewEscalationEnabled"`
	PendingReviewDurationBeforeEscalationDays    int  `json:"PendingReviewDurationBeforeEscalationDays"`
	PendingReviewDurationBeforeEscalationHours   int  `json:"PendingReviewDurationBeforeEscalationHours"`
	PendingReviewDurationBeforeEscalationMinutes int  `json:"PendingReviewDurationBeforeEscalationMinutes"`
}

// GetAccessPolicies retrieves a list of access policies from the Safeguard API.
// It takes a Filter as parameter and uses the global client reference to make the API request.
//
// Parameters:
//   - filter: A Filter object used to filter the access policies.
//
// Returns:
//   - A slice of AccessPolicy objects.
//   - An error if the request fails or the response cannot be unmarshaled.
func GetAccessPolicies(filter client.Filter) ([]AccessPolicy, error) {
	var accessPolicies []AccessPolicy

	query := "AccessPolicies" + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &accessPolicies); err != nil {
		return nil, err
	}

	return accessPolicies, nil
}

// GetAccessPolicy retrieves an access policy by its ID from the Safeguard API.
// It uses the global client reference to make the API request.
//
// Parameters:
//   - id: An integer representing the ID of the access policy to retrieve.
//   - fields: Optional fields to include in the query.
//
// Returns:
//   - AccessPolicy: The retrieved access policy.
//   - error: An error if any occurred during the request or unmarshalling process.
func GetAccessPolicy(id int, fields client.Fields) (AccessPolicy, error) {
	var accessPolicy AccessPolicy

	query := fmt.Sprintf("AccessPolicies/%d", id)

	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return accessPolicy, err
	}
	if err := json.Unmarshal(response, &accessPolicy); err != nil {
		return accessPolicy, err
	}

	return accessPolicy, nil
}

// DeleteAccessPolicy deletes an access policy with the given ID.
// It uses the global client reference to make the API request.
//
// Parameters:
//   - id: An integer representing the ID of the access policy to be deleted.
//
// Returns:
//   - error: An error object if the DELETE request fails, otherwise nil.
func DeleteAccessPolicy(id int) error {
	query := fmt.Sprintf("AccessPolicies/%d", id)

	_, err := c.DeleteRequest(query)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes the access policy from the system.
// It calls the DeleteAccessPolicy function with the client and policy ID.
// Returns an error if the deletion fails.
func (a AccessPolicy) Delete() error {
	return DeleteAccessPolicy(a.Id)
}
