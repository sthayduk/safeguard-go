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
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - fields: Filter criteria for the request
//
// Returns:
//   - User: The user information for the authenticated user
//   - error: An error if the request fails, nil otherwise
func GetMe(c *client.SafeguardClient, filter client.Filter) (User, error) {
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

// GetMeAccessRequestAssets retrieves all assets that the current user can request access to
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - filter: Filter criteria for the request
//
// Returns:
//   - []PolicyAsset: A slice of assets the user can request access to
//   - error: An error if the request fails, nil otherwise
func GetMeAccessRequestAssets(c *client.SafeguardClient, filter client.Filter) ([]PolicyAsset, error) {
	var assets []PolicyAsset

	query := "me/AccessRequestAssets" + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return []PolicyAsset{}, err
	}

	if err := json.Unmarshal(response, &assets); err != nil {
		return []PolicyAsset{}, err
	}

	for i := range assets {
		assets[i].client = c
	}

	return assets, nil
}

// GetMeAccessRequestAsset retrieves a specific asset that the current user can request access to
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - assetId: The ID of the asset to retrieve
//
// Returns:
//   - PolicyAsset: The requested asset information
//   - error: An error if the request fails, nil otherwise
func GetMeAccessRequestAsset(c *client.SafeguardClient, assetId string) (PolicyAsset, error) {
	query := "me/AccessRequestAssets/" + assetId

	response, err := c.GetRequest(query)
	if err != nil {
		return PolicyAsset{}, err
	}

	var asset PolicyAsset
	if err := json.Unmarshal(response, &asset); err != nil {
		return PolicyAsset{}, err
	}

	asset.client = c
	return asset, nil
}

// GetMeActionableRequests retrieves access requests that the current user can perform actions on
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - filter: Filter criteria for the request
//
// Returns:
//   - map[AccessRequestRole][]AccessRequest: Access requests grouped by role that the user can act on
//   - error: An error if the request fails, nil otherwise
func GetMeActionableRequests(c *client.SafeguardClient, filter client.Filter) (map[AccessRequestRole][]AccessRequest, error) {
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

// GetMeActionableRequestsByRole retrieves access requests for a specific role that the current user can perform actions on
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - role: The access request role to filter by
//   - filter: Filter criteria for the request
//
// Returns:
//   - []AccessRequest: Access requests for the specified role that the user can act on
//   - error: An error if the request fails, nil otherwise
func GetMeActionableRequestsByRole(c *client.SafeguardClient, role AccessRequestRole, filter client.Filter) ([]AccessRequest, error) {
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

// GetMeActionableRequestsDetailed retrieves and processes access requests that the current user can perform actions on
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - filter: Filter criteria for the request
//
// Returns:
//   - ActionableRequestsResult: Processed access requests with additional helper information
//   - error: An error if the request fails, nil otherwise
func GetMeActionableRequestsDetailed(c *client.SafeguardClient, filter client.Filter) (ActionableRequestsResult, error) {
	requests, err := GetMeActionableRequests(c, filter)
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

// FilterRequestsByState returns all requests matching the given state
func (r *ActionableRequestsResult) FilterRequestsByState(state AccessRequestState) []AccessRequest {
	var filtered []AccessRequest
	for _, req := range r.AllRequests {
		if AccessRequestState(req.State) == state {
			filtered = append(filtered, req)
		}
	}
	return filtered
}

// GetPendingRequests returns all requests that require action
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

// HasRole checks if there are any requests for the given role
func (r *ActionableRequestsResult) HasRole(role AccessRequestRole) bool {
	_, exists := r.RequestsByRole[role]
	return exists
}

// GetRequestsForRole returns all requests for a specific role
func (r *ActionableRequestsResult) GetRequestsForRole(role AccessRequestRole) []AccessRequest {
	return r.RequestsByRole[role]
}

// GetMeAccountEntitlements retrieves the account entitlements for the current user.
//
// Parameters:
//   - c: A pointer to the SafeguardClient used to make the request.
//   - accessRequestType: The type of access request to filter by.
//   - includeActiveRequests: A boolean indicating whether to include active requests in the response.
//   - filterByCredential: A boolean indicating whether to filter by credential.
//   - filter: A client.Filter object to apply additional filtering.
//
// Returns:
//   - A slice of AccountEntitlement objects.
//   - An error if the request fails or the response cannot be unmarshaled.
func GetMeAccountEntitlements(c *client.SafeguardClient, accessRequestType AccessRequestType, includeActiveRequests bool, filterByCredential bool, filter client.Filter) ([]AccountEntitlement, error) {
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
