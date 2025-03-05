package safeguard

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// AccessRequestRole represents the role of a user in an access request
type AccessRequestRole string

const (
	RequestorRole AccessRequestRole = "Requestor"
	ApproverRole  AccessRequestRole = "Approver"
	ReviewerRole  AccessRequestRole = "Reviewer"
	AdminRole     AccessRequestRole = "Admin"
	WatcherRole   AccessRequestRole = "Watcher"
	MonitorRole   AccessRequestRole = "Monitor"
)

// ActionableAccessRequests represents asset requests that the current user can perform some action on
type ActionableAccessRequests struct {
	Count            int               `json:"Count,omitempty"`
	AccessRequests   []AccessRequest   `json:"AccessRequests,omitempty"`
	RequestsToReview []AccessRequest   `json:"RequestsToReview,omitempty"`
	RequestRole      AccessRequestRole `json:"RequestRole,omitempty"`
}

// AccessRequest represents a request for access to an asset or account
type AccessRequest struct {
	apiClient *SafeguardClient `json:"-"`

	Id                                         string                 `json:"Id,omitempty"`
	AccessRequestType                          AccessRequestType      `json:"AccessRequestType,omitempty"`
	AccountId                                  int                    `json:"AccountId,omitempty"`
	AccountName                                string                 `json:"AccountName,omitempty"`
	AccountDomainName                          string                 `json:"AccountDomainName,omitempty"`
	AccountAssetId                             int                    `json:"AccountAssetId,omitempty"`
	AccountAssetName                           string                 `json:"AccountAssetName,omitempty"`
	AccountHasTotpAuthenticator                bool                   `json:"AccountHasTotpAuthenticator,omitempty"`
	AccountRequestType                         string                 `json:"AccountRequestType,omitempty"`
	ApprovedByMe                               bool                   `json:"ApprovedByMe,omitempty"`
	AssetId                                    int                    `json:"AssetId,omitempty"`
	AssetName                                  string                 `json:"AssetName,omitempty"`
	AssetNetworkAddress                        *string                `json:"AssetNetworkAddress,omitempty"`
	AssetSshHostKey                            *string                `json:"AssetSshHostKey,omitempty"`
	AssetSshHostKeyFingerprint                 *string                `json:"AssetSshHostKeyFingerprint,omitempty"`
	AssetSshHostKeyFingerprintSha256           *string                `json:"AssetSshHostKeyFingerprintSha256,omitempty"`
	CreatedOn                                  time.Time              `json:"CreatedOn,omitempty"`
	CurrentApprovalCount                       int                    `json:"CurrentApprovalCount,omitempty"`
	CurrentReviewerCount                       int                    `json:"CurrentReviewerCount,omitempty"`
	DurationInMinutes                          int                    `json:"DurationInMinutes,omitempty"`
	ExpiresOn                                  time.Time              `json:"ExpiresOn,omitempty"`
	IsEmergency                                bool                   `json:"IsEmergency,omitempty"`
	NeedsAcknowledgement                       bool                   `json:"NeedsAcknowledgement,omitempty"`
	RequestAvailability                        []DateTimeInterval     `json:"RequestAvailability,omitempty"`
	ReasonCode                                 *ReasonCodeInfo        `json:"ReasonCode,omitempty"`
	ReasonComment                              *string                `json:"ReasonComment,omitempty"`
	RequestedDurationDays                      int                    `json:"RequestedDurationDays,omitempty"`
	RequestedDurationHours                     int                    `json:"RequestedDurationHours,omitempty"`
	RequestedDurationMinutes                   int                    `json:"RequestedDurationMinutes,omitempty"`
	RequestedFor                               string                 `json:"RequestedFor,omitempty"`
	RequesterDisplayName                       string                 `json:"RequesterDisplayName,omitempty"`
	RequesterEmailAddress                      string                 `json:"RequesterEmailAddress,omitempty"`
	RequesterId                                int                    `json:"RequesterId,omitempty"`
	RequesterUsername                          string                 `json:"RequesterUsername,omitempty"`
	RequiredApprovalCount                      int                    `json:"RequiredApprovalCount,omitempty"`
	RequiredReviewerCount                      int                    `json:"RequiredReviewerCount,omitempty"`
	State                                      AccessRequestState     `json:"State,omitempty"`
	StateChangedOn                             time.Time              `json:"StateChangedOn,omitempty"`
	TicketNumber                               *string                `json:"TicketNumber,omitempty"`
	WasCancelled                               bool                   `json:"WasCancelled,omitempty"`
	WasCheckedOut                              bool                   `json:"WasCheckedOut,omitempty"`
	WasDenied                                  bool                   `json:"WasDenied,omitempty"`
	WasEvicted                                 bool                   `json:"WasEvicted,omitempty"`
	WasExpired                                 bool                   `json:"WasExpired,omitempty"`
	WasRevoked                                 bool                   `json:"WasRevoked,omitempty"`
	WorkflowActions                            []WorkflowAction       `json:"WorkflowActions,omitempty"`
	PolicyId                                   int                    `json:"PolicyId,omitempty"`
	PolicyName                                 string                 `json:"PolicyName,omitempty"`
	RequireReviewerComment                     bool                   `json:"RequireReviewerComment,omitempty"`
	AllowSraSessionLaunch                      bool                   `json:"AllowSraSessionLaunch,omitempty"`
	AllowSessionPasswordRelease                bool                   `json:"AllowSessionPasswordRelease,omitempty"`
	AllowSessionSshKeyRelease                  bool                   `json:"AllowSessionSshKeyRelease,omitempty"`
	IncludePasswordRelease                     bool                   `json:"IncludePasswordRelease,omitempty"`
	IncludeSshKeyRelease                       bool                   `json:"IncludeSshKeyRelease,omitempty"`
	Sessions                                   []AccessRequestSession `json:"Sessions,omitempty"`
	AccountDistinguishedName                   string                 `json:"AccountDistinguishedName,omitempty"`
	AssetPlatformId                            int                    `json:"AssetPlatformId,omitempty"`
	AssetPlatformType                          string                 `json:"AssetPlatformType,omitempty"`
	AssetPlatformDisplayName                   string                 `json:"AssetPlatformDisplayName,omitempty"`
	AllowSubsequentAccessRequestsWithoutReview bool                   `json:"AllowSubsequentAccessRequestsWithoutReview,omitempty"`
	SessionModuleConnectionId                  int                    `json:"SessionModuleConnectionId,omitempty"`
	SessionConnectionPolicyRef                 string                 `json:"SessionConnectionPolicyRef,omitempty"`
	SessionRdpShowWallpaper                    bool                   `json:"SessionRdpShowWallpaper,omitempty"`
}

func (ar AccessRequest) SetClient(c *SafeguardClient) any {
	ar.apiClient = c
	return ar
}

func (ar AccessRequest) GetState() AccessRequestState {
	return ar.State
}

// String returns the string representation of the AccessRequestType
func (a AccessRequestType) String() string {
	return string(a)
}

// AccessRequestSession represents information about sessions initialized using this request
type AccessRequestSession struct {
	Id                        string     `json:"Id,omitempty"`
	AccessRequestId           string     `json:"AccessRequestId,omitempty"`
	ApiVersion                string     `json:"ApiVersion,omitempty"`
	LaunchedByUserId          int        `json:"LaunchedByUserId,omitempty"`
	LaunchedByUserDisplayName string     `json:"LaunchedByUserDisplayName,omitempty"`
	SessionStarted            time.Time  `json:"SessionStarted,omitempty"`
	SessionEnd                *time.Time `json:"SessionEnd,omitempty"`
	ApplianceId               string     `json:"ApplianceId,omitempty"`
	ApplianceName             string     `json:"ApplianceName,omitempty"`
	ApplianceAddress          string     `json:"ApplianceAddress,omitempty"`
	SessionKey                string     `json:"SessionKey,omitempty"`
	PSMKey                    string     `json:"PSMKey,omitempty"`
	AccountName               string     `json:"AccountName,omitempty"`
	AccountDomainName         string     `json:"AccountDomainName,omitempty"`
	AssetName                 string     `json:"AssetName,omitempty"`
	NodeName                  string     `json:"NodeName,omitempty"`
	IsPlaybackAvailable       bool       `json:"IsPlaybackAvailable,omitempty"`
	IsTerminated              bool       `json:"IsTerminated,omitempty"`
	TerminatedByUserId        *int       `json:"TerminatedByUserId,omitempty"`
	TerminatedByUserName      *string    `json:"TerminatedByUserName,omitempty"`
	ErrorMessage              *string    `json:"ErrorMessage,omitempty"`
	SessionId                 int        `json:"SessionId,omitempty"`
	InitializedDate           time.Time  `json:"InitializedDate,omitempty"`
	ConnectedDate             time.Time  `json:"ConnectedDate,omitempty"`
	TerminatedDate            time.Time  `json:"TerminatedDate,omitempty"`
	State                     string     `json:"State,omitempty"`
	HasRecording              bool       `json:"HasRecording,omitempty"`
}

// DateTimeInterval represents a time period with begin and end times
type DateTimeInterval struct {
	Begin time.Time `json:"Begin,omitempty"`
	End   time.Time `json:"End,omitempty"`
}

// WorkflowAction represents an action taken to modify an access request
type WorkflowAction struct {
	ActionType string    `json:"ActionType,omitempty"`
	Comment    *string   `json:"Comment,omitempty"`
	NewState   string    `json:"NewState,omitempty"`
	OccurredOn time.Time `json:"OccurredOn,omitempty"`
	OldState   string    `json:"OldState,omitempty"`
	User       UserInfo  `json:"User,omitempty"`
	SessionId  *string   `json:"SessionId,omitempty"`
}

// UserInfo represents information about a user that performed an action
type UserInfo struct {
	DisplayName                       string  `json:"DisplayName,omitempty"`
	Id                                int     `json:"Id,omitempty"`
	IdentityProviderId                int     `json:"IdentityProviderId,omitempty"`
	IdentityProviderName              string  `json:"IdentityProviderName,omitempty"`
	IdentityProviderTypeReferenceName string  `json:"IdentityProviderTypeReferenceName,omitempty"`
	IsSystemOwned                     bool    `json:"IsSystemOwned,omitempty"`
	Name                              string  `json:"Name,omitempty"`
	PrincipalKind                     string  `json:"PrincipalKind,omitempty"`
	EmailAddress                      *string `json:"EmailAddress,omitempty"`
	DomainName                        *string `json:"DomainName,omitempty"`
	FullDisplayName                   string  `json:"FullDisplayName,omitempty"`
}

// AccessRequestBatchResponse represents the response for a batch access request
type AccessRequestBatchResponse struct {
	apiClient *SafeguardClient `json:"-"`

	Response         AccessRequest `json:"Response,omitempty"`
	StatusCode       string        `json:"StatusCode,omitempty"`
	StatusCodeNumber int           `json:"StatusCodeNumber,omitempty"`
	IsSuccess        bool          `json:"IsSuccess,omitempty"`
	Error            ApiError      `json:"Error,omitempty"`
	Request          BatchRequest  `json:"Request,omitempty"`
}

func (a AccessRequestBatchResponse) SetClient(c *SafeguardClient) any {
	a.apiClient = c
	return a
}

// hasError checks if the AccessRequestBatchResponse indicates a successful operation.
// If the operation was successful, it returns nil. Otherwise, it returns an error
// containing the error message from the response.
//
// Returns:
//   - error: nil if the operation was successful, otherwise an error with the response message.
func (ab AccessRequestBatchResponse) hasError() error {
	if ab.IsSuccess {
		return nil
	}

	return fmt.Errorf("error: %s", ab.Error.Message)
}

// ApiError represents error information returned by the API
type ApiError struct {
	Code       int    `json:"Code,omitempty"`
	Message    string `json:"Message,omitempty"`
	InnerError string `json:"InnerError,omitempty"`
}

// BatchRequest represents the request portion of a batch response
type BatchRequest struct {
	AccountId                int               `json:"AccountId,omitempty"`
	AssetId                  int               `json:"AssetId,omitempty"`
	AccessRequestType        AccessRequestType `json:"AccessRequestType,omitempty"`
	IsEmergency              bool              `json:"IsEmergency,omitempty"`
	ReasonCodeId             int               `json:"ReasonCodeId,omitempty"`
	ReasonComment            string            `json:"ReasonComment,omitempty"`
	RequestedDurationDays    int               `json:"RequestedDurationDays,omitempty"`
	RequestedDurationHours   int               `json:"RequestedDurationHours,omitempty"`
	RequestedDurationMinutes int               `json:"RequestedDurationMinutes,omitempty"`
	RequestedFor             string            `json:"RequestedFor,omitempty"`
	TicketNumber             string            `json:"TicketNumber,omitempty"`
	AllowSraSessionLaunch    bool              `json:"AllowSraSessionLaunch,omitempty"`
}

// ReasonCodeInfo represents a reason code with additional information
type ReasonCodeInfo struct {
	Id          int    `json:"Id,omitempty"`
	Name        string `json:"Name,omitempty"`
	Description string `json:"Description,omitempty"`
}

// AccessRequestApprovalBatchResponse represents the response for a batch approval request
type AccessRequestApprovalBatchResponse struct {
	Request     AccessRequest `json:"Request"`
	Status      string        `json:"Status,omitempty"`
	Message     string        `json:"Message,omitempty"`
	Comment     string        `json:"Comment,omitempty"`
	IsEmergency bool          `json:"IsEmergency,omitempty"`
}

// AccessRequestDenyBatchResponse represents the response for a batch deny request
type AccessRequestDenyBatchResponse struct {
	Request AccessRequest `json:"Request"`
	Status  string        `json:"Status,omitempty"`
	Message string        `json:"Message,omitempty"`
	Comment string        `json:"Comment,omitempty"`
}

// AccessRequestReviewBatchResponse represents the response for a batch review request
type AccessRequestReviewBatchResponse struct {
	Request AccessRequest `json:"Request"`
	Status  string        `json:"Status,omitempty"`
	Message string        `json:"Message,omitempty"`
	Comment string        `json:"Comment,omitempty"`
}

// GetAccessRequests retrieves access requests filtered by the provided criteria.
// The requests are sorted by creation date in descending order.
//
// Parameters:
//   - filter: Filter criteria for the requests
//
// Returns:
//   - []AccessRequest: Matching access requests
//   - error: API or unmarshalling errors
func (c *SafeguardClient) GetAccessRequests(filter Filter) ([]AccessRequest, error) {

	query := "AccessRequests" + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return []AccessRequest{}, err
	}

	var accessRequests []AccessRequest
	if err := json.Unmarshal(response, &accessRequests); err != nil {
		return []AccessRequest{}, err
	}

	return addClientToSlice(c, accessRequests), err
}

// GetAccessRequest retrieves a specific access request by its ID.
//
// Parameters:
//   - id: The unique identifier of the access request.
//   - fields: Optional fields to include in the response.
//
// Returns:
//   - AccessRequest: The retrieved access request.
//   - error: An error if the request fails or unmarshalling fails.
func (c *SafeguardClient) GetAccessRequest(id string, fields Fields) (AccessRequest, error) {

	query := "AccessRequests/" + id
	if fields != nil {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return AccessRequest{}, err
	}

	var accessRequest AccessRequest
	if err := json.Unmarshal(response, &accessRequest); err != nil {
		return AccessRequest{}, err
	}

	return addClient(c, accessRequest), err
}

// NewAccessRequests creates multiple access requests in a single batch operation.
//
// Parameters:
//   - accountEntitlements: Slice of account entitlements to request access for.
//
// Returns:
//   - []AccessRequestBatchResponse: Responses for each request in the batch.
//   - error: An error if the batch operation fails.
func (c *SafeguardClient) NewAccessRequests(accountEntitlements []AccountEntitlement, requestDuration time.Duration) ([]AccessRequestBatchResponse, error) {
	var accessRequests []batchAccessRequest

	// Reduce properties to the required ones for the request
	for i := range accountEntitlements {
		accessRequestType := accountEntitlements[i].GetAccessRequestType()

		accessRequest := constructAccessRequest(accessRequestType, accountEntitlements[i].Account.Id, accountEntitlements[i].Asset.Id, "", requestDuration, "", "")
		accessRequests = append(accessRequests, accessRequest)
	}

	// TODO: Check if accountEntitlements[0].Policies[0].RequesterProperties.MaximumReleaseDurationDays, Hours, Minutes is less that requestDuration
	// 		 If requestDuration is to long, set it to the maximum!
	//		 This should be maybe be solved in a program and not in an API

	accessBatchRequests, err := c.batchCreateAccessRequest(accessRequests)
	return addClientToSlice(c, accessBatchRequests), err
}

// constructAccessRequest creates a new batchAccessRequest with the provided parameters.
//
// Parameters:
//   - c: A pointer to a SafeguardClient instance.
//   - accessRequestType: The type of access request.
//   - accountId: The ID of the account for which access is being requested.
//   - assetId: The ID of the asset for which access is being requested.
//   - requesterUsername: The username of the requester.
//   - requestedDurationMinutes: The duration of the requested access in minutes.
//   - reasonCode: The reason code for the access request.
//   - reasonComment: Additional comments for the access request.
//
// Returns:
//
//	A batchAccessRequest instance populated with the provided parameters.
func constructAccessRequest(accessRequestType AccessRequestType, accountId int, assetId int, requesterUsername string, requestDuration time.Duration, reasonCode string, reasonComment string) batchAccessRequest {

	// Split time duration into days, hours, and minutes
	// The Access Request API only accepts days, hours, and minutes
	requestedDurationMinutes := int(requestDuration.Minutes()) % 60
	requestedDurationHours := int(requestDuration.Hours()) % 24
	requestedDurationDays := int(requestDuration.Hours() / 24)

	return batchAccessRequest{
		AccessRequestType:        accessRequestType,
		AccountId:                accountId,
		AssetId:                  assetId,
		RequestedDurationDays:    requestedDurationDays,
		RequestedDurationMinutes: requestedDurationMinutes,
		RequestedDurationHours:   requestedDurationHours,
		RequesterUsername:        requesterUsername,
		ReasonCode:               &reasonCode,
		ReasonComment:            &reasonComment,
	}
}

type batchAccessRequest struct {
	AccessRequestType        AccessRequestType `json:"AccessRequestType,omitempty"`
	AccountId                int               `json:"AccountId,omitempty"`
	AssetId                  int               `json:"AssetId,omitempty"`
	RequestedDurationMinutes int               `json:"RequestedDurationMinutes,omitempty"`
	RequestedDurationHours   int               `json:"RequestedDurationHours,omitempty"`
	RequestedDurationDays    int               `json:"RequestedDurationDays,omitempty"`
	RequesterUsername        string            `json:"RequesterUsername,omitempty"`
	ReasonCode               *string           `json:"ReasonCode,omitempty"`
	ReasonComment            *string           `json:"ReasonComment,omitempty"`
	IsEmergency              bool              `json:"IsEmergency,omitempty"`
}

// batchCreateAccessRequest sends a batch of access requests to the Safeguard API for creation.
// It takes a SafeguardClient and a slice of batchAccessRequest as input, and returns a slice of
// AccessRequestBatchResponse and an error.
//
// Parameters:
//   - c: A pointer to a SafeguardClient used to make the API request.
//   - accessRequests: A slice of batchAccessRequest containing the access requests to be created.
//
// Returns:
//   - A slice of AccessRequestBatchResponse containing the responses for each access request.
//   - An error if any occurred during the process.
//
// The function performs the following steps:
//  1. Marshals the accessRequests slice into JSON.
//  2. Sends a POST request to the "AccessRequests/BatchCreate" endpoint with the JSON payload.
//  3. Unmarshals the response into a slice of AccessRequestBatchResponse.
//  4. Adds the client to each AccessRequest in the response.
//  5. Collects any errors encountered during the process and returns them along with the responses.
func (c *SafeguardClient) batchCreateAccessRequest(accessRequests []batchAccessRequest) ([]AccessRequestBatchResponse, error) {
	requestBody, err := json.Marshal(accessRequests)
	if err != nil {
		return []AccessRequestBatchResponse{}, err
	}
	response, err := c.PostRequest("AccessRequests/BatchCreate", bytes.NewReader(requestBody))
	if err != nil {
		return []AccessRequestBatchResponse{}, err
	}

	var createdAccessRequests []AccessRequestBatchResponse
	if err := json.Unmarshal(response, &createdAccessRequests); err != nil {
		return []AccessRequestBatchResponse{}, err
	}

	var collectedErrors error
	for i := range createdAccessRequests {
		if err := createdAccessRequests[i].hasError(); err != nil {
			collectedErrors = fmt.Errorf("%v\n%v", collectedErrors, err)
		}
	}

	if collectedErrors != nil {
		return addClientToSlice(c, createdAccessRequests), collectedErrors
	}

	return addClientToSlice(c, createdAccessRequests), nil
}

// Close attempts to close the request based on its current state.
// For checked out passwords, checks them back in.
// For pending/available requests, cancels them.
// Returns error for invalid states.
//
// Returns:
//   - AccessRequest: Updated request state
//   - error: Any errors during close
func (ar AccessRequest) Close() (AccessRequest, error) {
	switch ar.State {
	case "PasswordCheckedOut":
		return ar.CheckIn()
	case "Pending":
		return ar.Cancel()
	case "RequestAvailable":
		return ar.Cancel()
	case "PendingAccountRestored":
		return ar.Cancel()
	case "Complete":
		return ar, nil
	default:
		return AccessRequest{}, fmt.Errorf("cannot close access request in state: %s", ar.State)
	}
}

// CancelAccessRequest cancels an access request with the given ID using the provided SafeguardClient.
// It sends a POST request to the "AccessRequests/{id}/Cancel" endpoint and unmarshals the response
// into an AccessRequest object.
//
// Parameters:
//   - c: A pointer to a SafeguardClient used to make the request.
//   - id: The ID of the access request to be canceled.
//
// Returns:
//   - AccessRequest: The canceled access request object.
//   - error: An error object if the request fails or if there is an issue unmarshaling the response.
func (c *SafeguardClient) CancelAccessRequest(id string) (AccessRequest, error) {

	response, err := c.PostRequest("AccessRequests/"+id+"/Cancel", nil)
	if err != nil {
		return AccessRequest{}, err
	}

	var accessRequest AccessRequest
	if err := json.Unmarshal(response, &accessRequest); err != nil {
		return AccessRequest{}, err
	}

	return addClient(c, accessRequest), err
}

// Cancel cancels a pending or available access request.
//
// Returns:
//   - AccessRequest: The updated access request showing canceled state.
//   - error: An error if the cancellation fails.
func (ar AccessRequest) Cancel() (AccessRequest, error) {
	return ar.apiClient.CancelAccessRequest(ar.Id)
}

// CheckInAccessRequest checks in an access request with the given ID using the provided SafeguardClient.
// It sends a POST request to the "AccessRequests/{id}/CheckIn" endpoint and unmarshals the response into an AccessRequest object.
//
// Parameters:
//   - c: A pointer to a SafeguardClient used to make the request.
//   - id: The ID of the access request to check in.
//
// Returns:
//   - AccessRequest: The checked-in access request object.
//   - error: An error object if an error occurred during the request or unmarshalling.
func (c *SafeguardClient) CheckInAccessRequest(id string) (AccessRequest, error) {

	response, err := c.PostRequest("AccessRequests/"+id+"/CheckIn", nil)
	if err != nil {
		return AccessRequest{}, err
	}

	var accessRequest AccessRequest
	if err := json.Unmarshal(response, &accessRequest); err != nil {
		return AccessRequest{}, err
	}

	return addClient(c, accessRequest), err
}

// CheckIn checks in a checked-out password access request.
//
// Returns:
//   - AccessRequest: The updated access request showing checked-in state.
//   - error: An error if the check-in fails.
func (ar AccessRequest) CheckIn() (AccessRequest, error) {
	return ar.apiClient.CheckInAccessRequest(ar.Id)
}

// CheckOutPassword retrieves the password for the request, optionally waiting
// for pending requests to become available.
//
// Parameters:
//   - ctx: Context for cancellation/timeout
//   - waitForPending: Whether to wait for pending state
//
// Returns:
//   - string: The password if successful
//   - error: Checkout or timeout errors
func (ar AccessRequest) CheckOutPassword(ctx context.Context, waitForPending bool) (string, error) {
	return ar.apiClient.CheckOutPassword(ctx, ar, waitForPending)
}

// CheckOutPassword checks out the password for the access request.
// It returns the password as a string and an error if the operation fails.
//
// Parameters:
//   - ctx: The context for the operation, which can be used to cancel the request.
//   - c: The SafeguardClient instance for making API requests.
//   - accessRequest: The access request for which the password is being checked out.
//   - waitForPending: A boolean indicating whether to wait for the access request to become valid if it is in a pending state.
//
// Returns:
//   - string: The checked-out password.
//   - error: An error if the password checkout fails.
func (c *SafeguardClient) CheckOutPassword(ctx context.Context, accessRequest AccessRequest, shouldWaitForPending bool) (string, error) {
	if accessRequest.IsInvalid() {
		return "", fmt.Errorf("cannot check out password for access request in state: %s", accessRequest.State)
	}

	if accessRequest.IsPending() {
		if !shouldWaitForPending {
			return "", fmt.Errorf("cannot check out password for access request in state: %s", accessRequest.State)
		}

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

	outerLoop:
		for {
			select {
			case <-ctx.Done():
				return "", fmt.Errorf("password request timed out")
			case <-ticker.C:
				accessRequest, err := c.GetAccessRequest(accessRequest.Id, nil)
				if err != nil {
					return "", err
				}

				if accessRequest.IsValid() {
					break outerLoop
				}
			}
		}
	}

	return c.getPasswordforAccessRequest(accessRequest)
}

// IsPending checks if the access request is in any pending state.
//
// Returns:
//   - bool: true if request is pending approval, review, or other pending states.
func (ar AccessRequest) IsPending() bool {
	return isAccessRequestPending(ar)
}

// IsValid checks if the access request is in a valid state for password checkout.
//
// Returns:
//   - bool: true if password can be checked out.
func (ar AccessRequest) IsValid() bool {
	return isAccessRequestValid(ar)
}

// IsInvalid checks if the access request is in a terminal or invalid state.
//
// Returns:
//   - bool: true if request is completed, expired, denied, canceled or revoked.
func (ar AccessRequest) IsInvalid() bool {
	return isAccessRequestInvalid(ar)
}

// isAccessRequestPending checks if the access request is in a pending state.
//
// Parameters:
//   - accessRequest: The access request to check.
//
// Returns:
//   - bool: True if the access request is in a pending state, false otherwise.
func isAccessRequestPending(accessRequest AccessRequest) bool {
	pendingStates := map[AccessRequestState]bool{
		StatePending:                true,
		StatePendingApproval:        true,
		StatePendingTimeRequested:   true,
		StatePendingAccountRestored: true,
		StatePendingAccountElevated: true,
		StatePendingReview:          true,
		StatePendingPasswordReset:   true,
		StatePendingAcknowledgment:  true,
	}

	return pendingStates[accessRequest.State]
}

// isAccessRequestInvalid checks if the access request is in an invalid state.
//
// Parameters:
//   - accessRequest: The access request to check.
//
// Returns:
//   - bool: True if the access request is in an invalid state, false otherwise.
func isAccessRequestInvalid(accessRequest AccessRequest) bool {
	invalidStates := map[AccessRequestState]bool{
		StateCompleted: true,
		StateExpired:   true,
		StateDenied:    true,
		StateCanceled:  true,
		StateRevoked:   true,
	}

	return invalidStates[accessRequest.State]
}

// isAccessRequestValid checks if the access request is in a valid state for password checkout.
//
// Parameters:
//   - accessRequest: The access request to check.
//
// Returns:
//   - bool: True if the access request is in a valid state, false otherwise.
func isAccessRequestValid(accessRequest AccessRequest) bool {
	validStates := map[AccessRequestState]bool{
		StatePasswordCheckedOut: true,
		StateRequestAvailable:   true,
		StateAcknowledged:       true,
	}

	return validStates[accessRequest.State]
}

// getPasswordforAccessRequest retrieves the password for the given access request.
// It sends a POST request to the "AccessRequests/{id}/CheckOutPassword" endpoint.
//
// Parameters:
//   - c: The SafeguardClient instance for making API requests.
//   - accessRequest: The access request for which the password is being checked out.
//
// Returns:
//   - string: The checked-out password.
//   - error: An error if the password checkout fails.
func (c *SafeguardClient) getPasswordforAccessRequest(accessRequest AccessRequest) (string, error) {
	query := fmt.Sprintf("AccessRequests/%s/CheckOutPassword", accessRequest.Id)

	response, err := c.PostRequest(query, nil)
	if err != nil {
		return "", err
	}

	return string(response), err
}

// RefreshState updates the request with its current state from the server.
//
// Returns:
//   - AccessRequest: Updated request state
//   - error: API or unmarshalling errors
func (ar AccessRequest) RefreshState() (AccessRequest, error) {
	return ar.apiClient.GetAccessRequest(ar.Id, nil)
}
