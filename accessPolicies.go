package safeguard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
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
	RequiredApprovers int        `json:"RequiredApprovers"`
	Approvers         []Identity `json:"Approvers"`
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
	apiClient *SafeguardClient `json:"-"`

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
	Reviewers                   []Identity                  `json:"Reviewers"`
	NotificationContacts        []NotificationContact       `json:"NotificationContacts"`
	ReasonCodes                 []ReasonCode                `json:"ReasonCodes"`
	ScopeItems                  []PolicyScopeItem           `json:"ScopeItems"`
	ExpirationDate              *time.Time                  `json:"ExpirationDate,omitempty"`
	IsExpired                   bool                        `json:"IsExpired"`
	InvalidConnectionPolicy     bool                        `json:"InvalidConnectionPolicy"`
	HourlyRestrictionProperties HourlyRestrictionProperties `json:"HourlyRestrictionProperties"`
}

func (a AccessPolicy) SetClient(c *SafeguardClient) any {
	a.apiClient = c
	return a
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

// GetApproverSets retrieves the sets of identities that may approve access requests using this policy.
//
// Returns:
//   - []ApproverSet: A slice of ApproverSet objects
//   - error: An error if the request fails
func (a AccessPolicy) GetApproverSets() ([]ApproverSet, error) {
	var approverSets []ApproverSet

	query := fmt.Sprintf("AccessPolicies/%d/ApproverSets", a.Id)

	response, err := a.apiClient.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &approverSets); err != nil {
		return nil, err
	}

	return approverSets, nil
}

// SetApproverSets sets who can approve access requests for this policy.
//
// Parameters:
//   - approverSets: A slice of ApproverSet objects to set as approvers
//
// Returns:
//   - []ApproverSet: The updated ApproverSet objects
//   - error: An error if the request fails
func (a AccessPolicy) SetApproverSets(approverSets []ApproverSet) ([]ApproverSet, error) {
	var updatedApproverSets []ApproverSet

	query := fmt.Sprintf("AccessPolicies/%d/ApproverSets", a.Id)
	data, err := json.Marshal(approverSets)
	if err != nil {
		return nil, err
	}

	response, err := a.apiClient.PutRequest(query, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &updatedApproverSets); err != nil {
		return nil, err
	}

	return updatedApproverSets, nil
}

// ModifyApproverSets adds or removes approvers who can approve access requests for this policy.
//
// Parameters:
//   - operation: The operation to perform (Add or Remove)
//   - approverSets: A slice of ApproverSet objects to modify
//
// Returns:
//   - []ApproverSet: The updated ApproverSet objects
//   - error: An error if the request fails
func (a AccessPolicy) ModifyApproverSets(operation ApiSetOperation, approverSets []ApproverSet) ([]ApproverSet, error) {
	var updatedApproverSets []ApproverSet

	query := fmt.Sprintf("AccessPolicies/%d/ApproverSets/%s", a.Id, operation)
	data, err := json.Marshal(approverSets)
	if err != nil {
		return nil, err
	}

	response, err := a.apiClient.PostRequest(query, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &updatedApproverSets); err != nil {
		return nil, err
	}

	return updatedApproverSets, nil
}

// GetReviewers retrieves the list of reviewers for the access policy.
// It sends a GET request to the API endpoint corresponding to the access policy's reviewers,
// unmarshals the response into a slice of Identity objects, and returns the slice.
// If any error occurs during the request or unmarshalling, it returns the error.
//
// Returns:
//   - A slice of Identity objects representing the reviewers.
//   - An error if the request or unmarshalling fails.
func (a AccessPolicy) GetReviewers() ([]Identity, error) {
	var reviewers []Identity

	query := fmt.Sprintf("AccessPolicies/%d/Reviewers", a.Id)

	response, err := a.apiClient.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &reviewers); err != nil {
		return nil, err
	}

	return addClientToSlice(a.apiClient, reviewers), nil
}

// SetReviewers sets the reviewers for the access policy.
// It takes a slice of Identity objects representing the reviewers and returns
// a slice of modified Identity objects and an error if any occurred.
//
// Parameters:
//
//	reviewers - A slice of Identity objects representing the reviewers to be set.
//
// Returns:
//
//	A slice of modified Identity objects and an error if any occurred during the process.
//
// The function performs the following steps:
//  1. Constructs the query URL using the access policy ID.
//  2. Marshals the reviewers slice into JSON format.
//  3. Sends a PUT request to the API with the marshaled data.
//  4. Unmarshals the response into a slice of modified Identity objects.
//  5. Adds the API client to the modified reviewers slice and returns it.
func (a AccessPolicy) SetReviewers(reviewers []Identity) ([]Identity, error) {
	var modifiedReviewers []Identity

	query := fmt.Sprintf("AccessPolicies/%d/Reviewers", a.Id)
	data, err := json.Marshal(reviewers)
	if err != nil {
		return nil, err
	}

	response, err := a.apiClient.PutRequest(query, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &modifiedReviewers); err != nil {
		return nil, err
	}

	return addClientToSlice(a.apiClient, modifiedReviewers), nil
}

// ModifyReviewers modifies the reviewers of an access policy based on the specified operation.
// It sends a POST request to the API with the updated list of reviewers and returns the updated list.
//
// Parameters:
//   - operation: The operation to perform on the reviewers (e.g., add or remove).
//   - reviewers: A slice of Identity objects representing the reviewers to be modified.
//
// Returns:
//   - A slice of Identity objects representing the updated list of reviewers.
//   - An error if the operation fails or if there is an issue with the API request.
//
// Example:
//
//	updatedReviewers, err := accessPolicy.ModifyReviewers(ApiSetOperationAdd, reviewers)
//	if err != nil {
//	    log.Fatalf("Failed to modify reviewers: %v", err)
//	}
func (a AccessPolicy) ModifyReviewers(operation ApiSetOperation, reviewers []Identity) ([]Identity, error) {
	var updatedReviewers []Identity

	query := fmt.Sprintf("AccessPolicies/%d/Reviewers/%s", a.Id, operation)
	data, err := json.Marshal(reviewers)
	if err != nil {
		return nil, err
	}

	response, err := a.apiClient.PostRequest(query, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &updatedReviewers); err != nil {
		return nil, err
	}

	return addClientToSlice(a.apiClient, updatedReviewers), nil
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
func (c *SafeguardClient) GetAccessPolicies(filter Filter) ([]AccessPolicy, error) {
	var accessPolicies []AccessPolicy

	query := "AccessPolicies" + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &accessPolicies); err != nil {
		return nil, err
	}

	return addClientToSlice(c, accessPolicies), nil
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
func (c *SafeguardClient) GetAccessPolicy(id int, fields Fields) (AccessPolicy, error) {
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

	return addClient(c, accessPolicy), nil
}

// DeleteAccessPolicy deletes an access policy with the given ID.
// It uses the global client reference to make the API request.
//
// Parameters:
//   - id: An integer representing the ID of the access policy to be deleted.
//
// Returns:
//   - error: An error object if the DELETE request fails, otherwise nil.
func (c *SafeguardClient) DeleteAccessPolicy(id int) error {
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
	return a.apiClient.DeleteAccessPolicy(a.Id)
}

// UpdateAccessPolicy updates an existing access policy with the provided details.
// It takes the ID of the access policy to update and an AccessPolicy object containing the updated values.
// The method makes a PUT request to the Safeguard API, updates the access policy, and returns the updated
// AccessPolicy object with the client reference attached.
//
// Parameters:
//   - id: The unique identifier of the access policy to update
//   - updatedAccessPolicy: AccessPolicy object containing the updated values
//
// Returns:
//   - AccessPolicy: The updated access policy object
//   - error: An error if the update operation fails
func (c *SafeguardClient) UpdateAccessPolicy(id int, updatedAccessPolicy AccessPolicy) (AccessPolicy, error) {
	var accessPolicy AccessPolicy

	query := fmt.Sprintf("AccessPolicies/%d", id)
	accessPolicyJSON, err := json.Marshal(updatedAccessPolicy)
	if err != nil {
		return accessPolicy, err
	}

	response, err := c.PutRequest(query, bytes.NewReader(accessPolicyJSON))
	if err != nil {
		return accessPolicy, err
	}
	if err := json.Unmarshal(response, &accessPolicy); err != nil {
		return accessPolicy, err
	}

	return addClient(c, accessPolicy), nil
}

// Update updates this AccessPolicy with the provided updates.
// It sends a request to update the AccessPolicy identified by this AccessPolicy's ID
// with the details from updatedAccessPolicy.
//
// Parameters:
//   - updatedAccessPolicy: The AccessPolicy object containing the updated fields
//
// Returns:
//   - AccessPolicy: The updated AccessPolicy object
//   - error: An error if the update operation fails, nil otherwise
func (a AccessPolicy) Update(updatedAccessPolicy AccessPolicy) (AccessPolicy, error) {
	return a.apiClient.UpdateAccessPolicy(a.Id, updatedAccessPolicy)
}
