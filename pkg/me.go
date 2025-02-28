package pkg

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/sthayduk/safeguard-go/client"
)

// AccountEntitlement represents the full account entitlement structure
type AccountEntitlement struct {
	Account        AccountInfo         `json:"Account,omitempty"`
	Asset          AssetInfo           `json:"Asset,omitempty"`
	Policies       []PolicyInfo        `json:"Policies,omitempty"`
	ActiveRequests []ActiveRequestInfo `json:"ActiveRequests,omitempty"`
}

// GetAccountId returns the ID of the account associated with this entitlement
func (m AccountEntitlement) GetAccountId() int {
	return m.Account.Id
}

// GetFilter returns a filter object configured with the account ID
func (m AccountEntitlement) GetFilter() client.Filter {
	var filter client.Filter
	filter.AddFilter("AccountId", "eq", strconv.Itoa(m.GetAccountId()))
	return filter
}

// GetAccessRequestType returns the access request type from the first policy
// If no policies exist, returns an empty string
func (m AccountEntitlement) GetAccessRequestType() AccessRequestType {
	for _, policy := range m.Policies {
		return policy.AccessRequestType
	}
	return ""
}

// AccountInfo represents account information in entitlement response
type AccountInfo struct {
	Id                   int      `json:"Id,omitempty"`
	Name                 string   `json:"Name,omitempty"`
	DomainName           string   `json:"DomainName,omitempty"`
	Description          *string  `json:"Description,omitempty"`
	HasPassword          bool     `json:"HasPassword,omitempty"`
	HasSshKey            bool     `json:"HasSshKey,omitempty"`
	HasApiKey            bool     `json:"HasApiKey,omitempty"`
	HasFile              bool     `json:"HasFile,omitempty"`
	Disabled             bool     `json:"Disabled,omitempty"`
	AssetId              int      `json:"AssetId,omitempty"`
	AssetName            string   `json:"AssetName,omitempty"`
	AssetNetworkAddress  *string  `json:"AssetNetworkAddress,omitempty"`
	AllowPasswordRequest bool     `json:"AllowPasswordRequest,omitempty"`
	AllowSessionRequest  bool     `json:"AllowSessionRequest,omitempty"`
	AllowSshKeyRequest   bool     `json:"AllowSshKeyRequest,omitempty"`
	AllowApiKeyRequest   bool     `json:"AllowApiKeyRequest,omitempty"`
	AllowFileRequest     bool     `json:"AllowFileRequest,omitempty"`
	Tags                 []string `json:"Tags,omitempty"`
}

// AssetInfo represents asset information in entitlement response
type AssetInfo struct {
	Id                  int      `json:"Id,omitempty"`
	Name                string   `json:"Name,omitempty"`
	DomainName          *string  `json:"DomainName,omitempty"`
	Description         *string  `json:"Description,omitempty"`
	NetworkAddress      *string  `json:"NetworkAddress,omitempty"`
	PlatformDisplayName string   `json:"PlatformDisplayName,omitempty"`
	PlatformType        string   `json:"PlatformType,omitempty"`
	Tags                []string `json:"Tags,omitempty"`
}

// PolicyInfo represents policy information in entitlement response
type PolicyInfo struct {
	Id                                   int                         `json:"Id,omitempty"`
	Name                                 string                      `json:"Name,omitempty"`
	Priority                             int                         `json:"Priority,omitempty"`
	RolePriority                         int                         `json:"RolePriority,omitempty"`
	AccessRequestType                    AccessRequestType           `json:"AccessRequestType,omitempty"`
	AllowSimultaneousAccess              bool                        `json:"AllowSimultaneousAccess,omitempty"`
	MaximumSimultaneousReleases          int                         `json:"MaximumSimultaneousReleases,omitempty"`
	RequesterProperties                  RequesterProperties         `json:"RequesterProperties,omitempty"`
	EmergencyAccessProperties            EmergencyAccessProperties   `json:"EmergencyAccessProperties,omitempty"`
	EffectiveExpirationDate              *string                     `json:"EffectiveExpirationDate,omitempty"`
	EffectiveHourlyRestrictionProperties HourlyRestrictionProperties `json:"EffectiveHourlyRestrictionProperties,omitempty"`
	ReasonCodes                          []string                    `json:"ReasonCodes,omitempty"`
}

// RequesterProperties represents requester properties in policy
type RequesterProperties struct {
	DefaultReleaseDurationDays    int  `json:"DefaultReleaseDurationDays,omitempty"`
	DefaultReleaseDurationHours   int  `json:"DefaultReleaseDurationHours,omitempty"`
	DefaultReleaseDurationMinutes int  `json:"DefaultReleaseDurationMinutes,omitempty"`
	MaximumReleaseDurationDays    int  `json:"MaximumReleaseDurationDays,omitempty"`
	MaximumReleaseDurationHours   int  `json:"MaximumReleaseDurationHours,omitempty"`
	MaximumReleaseDurationMinutes int  `json:"MaximumReleaseDurationMinutes,omitempty"`
	AllowCustomDuration           bool `json:"AllowCustomDuration,omitempty"`
	RequireReasonCode             bool `json:"RequireReasonCode,omitempty"`
	RequireReasonComment          bool `json:"RequireReasonComment,omitempty"`
	RequireServiceTicket          bool `json:"RequireServiceTicket,omitempty"`
}

// GetDefaultReleaseDuration calculates and returns the default release duration as a time.Duration
func (r RequesterProperties) GetDefaultReleaseDuration() time.Duration {
	return time.Duration(r.DefaultReleaseDurationDays*24*60*60+r.DefaultReleaseDurationHours*60*60+r.DefaultReleaseDurationMinutes*60) * time.Second
}

// GetMaximumReleaseDuration calculates and returns the maximum release duration as a time.Duration
func (r RequesterProperties) GetMaximumReleaseDuration() time.Duration {
	return time.Duration(r.MaximumReleaseDurationDays*24*60*60+r.MaximumReleaseDurationHours*60*60+r.MaximumReleaseDurationMinutes*60) * time.Second
}

// EmergencyAccessProperties represents emergency access settings
type EmergencyAccessProperties struct {
	AllowEmergencyAccess     bool `json:"AllowEmergencyAccess,omitempty"`
	IgnoreHourlyRestrictions bool `json:"IgnoreHourlyRestrictions,omitempty"`
}

// ActiveRequestInfo represents information about an active request
type ActiveRequestInfo struct {
	Id                string             `json:"Id,omitempty"`
	AccessRequestType AccessRequestType  `json:"AccessRequestType,omitempty"`
	State             AccessRequestState `json:"State,omitempty"`
	ExpiresOn         time.Time          `json:"ExpiresOn,omitempty"`
	CreatedOn         time.Time          `json:"CreatedOn,omitempty"`
}

// AccessRequestState represents possible states of an access request
type AccessRequestState string

const (
	StateNew                    AccessRequestState = "New"
	StatePendingApproval        AccessRequestState = "PendingApproval"
	StatePendingTimeRequested   AccessRequestState = "PendingTimeRequested"
	StatePendingAccountRestored AccessRequestState = "PendingAccountRestored"
	StatePendingAccountElevated AccessRequestState = "PendingAccountElevated"
	StateRequestAvailable       AccessRequestState = "RequestAvailable"
	StatePasswordCheckedOut     AccessRequestState = "PasswordCheckedOut"
	StatePasswordCheckedIn      AccessRequestState = "PasswordCheckedIn"
	StatePendingReview          AccessRequestState = "PendingReview"
	StatePendingPasswordReset   AccessRequestState = "PendingPasswordReset"
	StateExpired                AccessRequestState = "Expired"
	StateDenied                 AccessRequestState = "Denied"
	StateCanceled               AccessRequestState = "Canceled"
	StateRevoked                AccessRequestState = "Revoked"
	StatePendingAcknowledgment  AccessRequestState = "PendingAcknowledgment"
	StateAcknowledged           AccessRequestState = "Acknowledged"
	StateCompleted              AccessRequestState = "Complete"
	StatePending                AccessRequestState = "Pending"
)

type EntitlementType string

const (
	PasswordEntitlement EntitlementType = "Password"
	SessionEntitlement  EntitlementType = "Session"
	SshKeyEntitlement   EntitlementType = "SshKey"
	ApiKeyEntitlement   EntitlementType = "ApiKey"
	FileEntitlement     EntitlementType = "File"
)

// GetMe retrieves information about the currently authenticated user.
//
// Returns:
//   - User: The user information for the authenticated user
//   - error: An error if the request fails or the response cannot be parsed
func GetMe(filter client.Filter) (User, error) {
	query := "me" + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return User{}, err
	}

	var me User
	if err := json.Unmarshal(response, &me); err != nil {
		return User{}, err
	}

	return me, nil
}

// GetMeAccessRequestAssets retrieves all assets that the current user can request access to.
//
// Parameters:
//   - filter: Filter criteria to narrow down the results
//
// Returns:
//   - []PolicyAsset: A slice of assets the user can request access to
//   - error: An error if the request fails or the response cannot be parsed
func GetMeAccessRequestAssets(filter client.Filter) ([]PolicyAsset, error) {
	var assets []PolicyAsset

	query := "me/AccessRequestAssets" + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return []PolicyAsset{}, err
	}

	if err := json.Unmarshal(response, &assets); err != nil {
		return []PolicyAsset{}, err
	}

	return assets, nil
}

// GetMeAccessRequestAsset retrieves a specific asset that the current user can request access to.
//
// Parameters:
//   - assetId: The ID of the asset to retrieve information for
//
// Returns:
//   - PolicyAsset: The requested asset's information
//   - error: An error if the asset cannot be found or the request fails
func GetMeAccessRequestAsset(assetId string) (PolicyAsset, error) {
	query := "me/AccessRequestAssets/" + assetId

	response, err := c.GetRequest(query)
	if err != nil {
		return PolicyAsset{}, err
	}

	var asset PolicyAsset
	if err := json.Unmarshal(response, &asset); err != nil {
		return PolicyAsset{}, err
	}

	return asset, nil
}

// GetMeActionableRequests retrieves access requests that require action from the current user.
//
// Parameters:
//   - filter: Filter criteria to narrow down the results
//
// Returns:
//   - map[AccessRequestRole][]AccessRequest: Access requests grouped by role
//   - error: An error if the request fails or the response cannot be parsed
func GetMeActionableRequests(filter client.Filter) (map[AccessRequestRole][]AccessRequest, error) {
	query := "me/ActionableRequests" + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	var requests map[AccessRequestRole][]AccessRequest
	if err := json.Unmarshal(response, &requests); err != nil {
		return nil, err
	}

	return requests, nil
}

// GetMeActionableRequestsByRole retrieves access requests for a specific role that require action.
//
// Parameters:
//   - role: The specific role to filter requests by
//   - filter: Additional filter criteria to narrow down the results
//
// Returns:
//   - []AccessRequest: Access requests for the specified role
//   - error: An error if the request fails or the response cannot be parsed
func GetMeActionableRequestsByRole(role AccessRequestRole, filter client.Filter) ([]AccessRequest, error) {
	query := "me/ActionableRequests/" + string(role) + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	var requests []AccessRequest
	if err := json.Unmarshal(response, &requests); err != nil {
		return nil, err
	}

	return requests, nil
}

// ActionableRequestsResult represents the processed result of GetMeActionableRequests
type ActionableRequestsResult struct {
	AllRequests    []AccessRequest
	RequestsByRole map[AccessRequestRole][]AccessRequest
	TotalCount     int
	CountByRole    map[AccessRequestRole]int
	AvailableRoles []AccessRequestRole
}

// GetMeActionableRequestsDetailed provides a detailed analysis of actionable access requests.
// This is a convenience method that processes the results from GetMeActionableRequests
// and provides additional helper information.
//
// Parameters:
//   - filter: Filter criteria to narrow down the results
//
// Returns:
//   - ActionableRequestsResult: Processed access requests with additional metadata
//   - error: An error if the request fails or the response cannot be parsed
func GetMeActionableRequestsDetailed(filter client.Filter) (ActionableRequestsResult, error) {
	requests, err := GetMeActionableRequests(filter)
	if err != nil {
		return ActionableRequestsResult{}, err
	}

	result := ActionableRequestsResult{
		RequestsByRole: requests,
		CountByRole:    make(map[AccessRequestRole]int),
	}

	// Process the requests
	for role, roleRequests := range requests {
		result.AllRequests = append(result.AllRequests, roleRequests...)
		result.CountByRole[role] = len(roleRequests)
		result.AvailableRoles = append(result.AvailableRoles, role)
		result.TotalCount += len(roleRequests)
	}

	return result, nil
}

// FilterRequestsByState returns all requests matching the specified state.
//
// Parameters:
//   - state: The AccessRequestState to filter by
//
// Returns:
//   - []AccessRequest: A slice of access requests in the specified state
func (r *ActionableRequestsResult) FilterRequestsByState(state AccessRequestState) []AccessRequest {
	var filtered []AccessRequest
	for _, req := range r.AllRequests {
		if AccessRequestState(req.State) == state {
			filtered = append(filtered, req)
		}
	}
	return filtered
}

// GetPendingRequests returns all requests that require action.
// This includes requests in New, PendingApproval, and PendingReview states.
//
// Returns:
//   - []AccessRequest: A slice of pending access requests
func (r *ActionableRequestsResult) GetPendingRequests() []AccessRequest {
	var pending []AccessRequest
	pendingStates := map[AccessRequestState]bool{
		StateNew:             true,
		StatePendingApproval: true,
		StatePendingReview:   true,
	}

	for _, req := range r.AllRequests {
		if pendingStates[AccessRequestState(req.State)] {
			pending = append(pending, req)
		}
	}
	return pending
}

// HasRole checks if there are any requests for the specified role.
//
// Parameters:
//   - role: The AccessRequestRole to check for
//
// Returns:
//   - bool: true if there are requests for the role, false otherwise
func (r *ActionableRequestsResult) HasRole(role AccessRequestRole) bool {
	_, exists := r.RequestsByRole[role]
	return exists
}

// GetRequestsForRole returns all requests for a specific role.
//
// Parameters:
//   - role: The AccessRequestRole to get requests for
//
// Returns:
//   - []AccessRequest: A slice of access requests for the specified role
func (r *ActionableRequestsResult) GetRequestsForRole(role AccessRequestRole) []AccessRequest {
	return r.RequestsByRole[role]
}

// GetMeAccountEntitlements retrieves the account entitlements for the current user.
//
// Parameters:
//   - accessRequestType: Optional type of access request to filter by
//   - includeActiveRequests: If true, includes currently active requests in the response
//   - filterByCredential: If true, filters results by credential type
//   - filter: Additional filter criteria to narrow down the results
//
// Returns:
//   - []AccountEntitlement: A slice of account entitlements for the user
//   - error: An error if the request fails or the response cannot be parsed
func GetMeAccountEntitlements(accessRequestType AccessRequestType, includeActiveRequests bool, filterByCredential bool, filter client.Filter) ([]AccountEntitlement, error) {
	var entitlements []AccountEntitlement

	query := "me/AccountEntitlements" + filter.ToQueryString() +
		"&includeActiveRequests=" + strconv.FormatBool(includeActiveRequests) +
		"&filterByCredential=" + strconv.FormatBool(filterByCredential)

	if accessRequestType != "" {
		query += "&accessRequestType=" + string(accessRequestType)
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return []AccountEntitlement{}, err
	}

	if err := json.Unmarshal(response, &entitlements); err != nil {
		return []AccountEntitlement{}, err
	}

	return entitlements, nil
}
