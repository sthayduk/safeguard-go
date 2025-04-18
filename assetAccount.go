package safeguard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// Tag represents metadata that can be attached to an asset account for
// organization and filtering purposes.
type Tag struct {
	Id            int    `json:"Id,omitempty"`
	Name          string `json:"Name,omitempty"`
	Description   string `json:"Description,omitempty"`
	AdminAssigned bool   `json:"AdminAssigned,omitempty"`
}

// Profile represents configuration settings that can be applied to an asset account,
// such as password rules or authentication settings.
type Profile struct {
	Id            int    `json:"Id,omitempty"`
	Name          string `json:"Name,omitempty"`
	EffectiveId   int    `json:"EffectiveId,omitempty"`
	EffectiveName string `json:"EffectiveName,omitempty"`
}

// DiscoveredGroup represents a security group or role that was automatically
// discovered for an asset account during discovery operations.
type DiscoveredGroup struct {
	DiscoveredGroupId                string `json:"DiscoveredGroupId,omitempty"`
	DiscoveredGroupName              string `json:"DiscoveredGroupName,omitempty"`
	DiscoveredGroupDistinguishedName string `json:"DiscoveredGroupDistinguishedName,omitempty"`
}

// DiscoveredProperties contains metadata about when and how an account was discovered,
// including the discovery schedule and discovered group memberships.
type DiscoveredProperties struct {
	AccountDiscoveryScheduleId   int               `json:"AccountDiscoveryScheduleId,omitempty"`
	AccountDiscoveryScheduleName string            `json:"AccountDiscoveryScheduleName,omitempty"`
	DiscoveredUserId             string            `json:"DiscoveredUserId,omitempty"`
	DiscoveredDate               string            `json:"DiscoveredDate,omitempty"`
	DiscoveredGroups             []DiscoveredGroup `json:"DiscoveredGroups,omitempty"`
}

// AssetDirectoryProperties represents directory properties of an asset
type AssetDirectoryProperties struct {
	NetbiosName string `json:"NetbiosName,omitempty"`
	ObjectGuid  string `json:"ObjectGuid,omitempty"`
	ObjectSid   string `json:"ObjectSid,omitempty"`
}

// SyncGroup represents a synchronization group for an asset account
type SyncGroup struct {
	Id       int    `json:"Id,omitempty"`
	Name     string `json:"Name,omitempty"`
	Priority int    `json:"Priority,omitempty"`
	Disabled bool   `json:"Disabled,omitempty"`
}

// TaskProperties represents task properties for an asset account
type TaskProperties struct {
	HasAccountTaskFailure          bool      `json:"HasAccountTaskFailure,omitempty"`
	LastPasswordCheckDate          time.Time `json:"LastPasswordCheckDate,omitempty"`
	LastSuccessPasswordCheckDate   time.Time `json:"LastSuccessPasswordCheckDate,omitempty"`
	LastFailurePasswordCheckDate   time.Time `json:"LastFailurePasswordCheckDate,omitempty"`
	LastPasswordCheckTaskId        string    `json:"LastPasswordCheckTaskId,omitempty"`
	FailedPasswordCheckAttempts    int       `json:"FailedPasswordCheckAttempts,omitempty"`
	NextPasswordCheckDate          time.Time `json:"NextPasswordCheckDate,omitempty"`
	LastPasswordChangeDate         time.Time `json:"LastPasswordChangeDate,omitempty"`
	LastSuccessPasswordChangeDate  time.Time `json:"LastSuccessPasswordChangeDate,omitempty"`
	LastFailurePasswordChangeDate  time.Time `json:"LastFailurePasswordChangeDate,omitempty"`
	LastPasswordChangeTaskId       string    `json:"LastPasswordChangeTaskId,omitempty"`
	FailedPasswordChangeAttempts   int       `json:"FailedPasswordChangeAttempts,omitempty"`
	NextPasswordChangeDate         time.Time `json:"NextPasswordChangeDate,omitempty"`
	LastSshKeyCheckDate            time.Time `json:"LastSshKeyCheckDate,omitempty"`
	LastSuccessSshKeyCheckDate     time.Time `json:"LastSuccessSshKeyCheckDate,omitempty"`
	LastFailureSshKeyCheckDate     time.Time `json:"LastFailureSshKeyCheckDate,omitempty"`
	LastSshKeyCheckTaskId          string    `json:"LastSshKeyCheckTaskId,omitempty"`
	FailedSshKeyCheckAttempts      int       `json:"FailedSshKeyCheckAttempts,omitempty"`
	NextSshKeyCheckDate            time.Time `json:"NextSshKeyCheckDate,omitempty"`
	LastSshKeyChangeDate           time.Time `json:"LastSshKeyChangeDate,omitempty"`
	LastSuccessSshKeyChangeDate    time.Time `json:"LastSuccessSshKeyChangeDate,omitempty"`
	LastFailureSshKeyChangeDate    time.Time `json:"LastFailureSshKeyChangeDate,omitempty"`
	LastSshKeyChangeTaskId         string    `json:"LastSshKeyChangeTaskId,omitempty"`
	FailedSshKeyChangeAttempts     int       `json:"FailedSshKeyChangeAttempts,omitempty"`
	NextSshKeyChangeDate           time.Time `json:"NextSshKeyChangeDate,omitempty"`
	LastSshKeyDiscoveryDate        time.Time `json:"LastSshKeyDiscoveryDate,omitempty"`
	LastSuccessSshKeyDiscoveryDate time.Time `json:"LastSuccessSshKeyDiscoveryDate,omitempty"`
	LastFailureSshKeyDiscoveryDate time.Time `json:"LastFailureSshKeyDiscoveryDate,omitempty"`
	LastSshKeyDiscoveryTaskId      string    `json:"LastSshKeyDiscoveryTaskId,omitempty"`
	FailedSshKeyDiscoveryAttempts  int       `json:"FailedSshKeyDiscoveryAttempts,omitempty"`
	NextSshKeyDiscoveryDate        time.Time `json:"NextSshKeyDiscoveryDate,omitempty"`
	LastSshKeyRevokeDate           time.Time `json:"LastSshKeyRevokeDate,omitempty"`
	LastSuccessSshKeyRevokeDate    time.Time `json:"LastSuccessSshKeyRevokeDate,omitempty"`
	LastFailureSshKeyRevokeDate    time.Time `json:"LastFailureSshKeyRevokeDate,omitempty"`
	LastSshKeyRevokeTaskId         string    `json:"LastSshKeyRevokeTaskId,omitempty"`
	FailedSshKeyRevokeAttempts     int       `json:"FailedSshKeyRevokeAttempts,omitempty"`
	LastSuspendAccountDate         time.Time `json:"LastSuspendAccountDate,omitempty"`
	LastSuccessSuspendAccountDate  time.Time `json:"LastSuccessSuspendAccountDate,omitempty"`
	LastFailureSuspendAccountDate  time.Time `json:"LastFailureSuspendAccountDate,omitempty"`
	LastSuspendAccountTaskId       string    `json:"LastSuspendAccountTaskId,omitempty"`
	FailedSuspendAccountAttempts   int       `json:"FailedSuspendAccountAttempts,omitempty"`
	NextSuspendAccountDate         time.Time `json:"NextSuspendAccountDate,omitempty"`
	LastRestoreAccountDate         time.Time `json:"LastRestoreAccountDate,omitempty"`
	LastSuccessRestoreAccountDate  time.Time `json:"LastSuccessRestoreAccountDate,omitempty"`
	LastFailureRestoreAccountDate  time.Time `json:"LastFailureRestoreAccountDate,omitempty"`
	LastRestoreAccountTaskId       string    `json:"LastRestoreAccountTaskId,omitempty"`
	FailedRestoreAccountAttempts   int       `json:"FailedRestoreAccountAttempts,omitempty"`
	NextRestoreAccountDate         time.Time `json:"NextRestoreAccountDate,omitempty"`
	FailedApiKeyCheckAttempts      int       `json:"FailedApiKeyCheckAttempts,omitempty"`
	FailedApiKeyChangeAttempts     int       `json:"FailedApiKeyChangeAttempts,omitempty"`
	LastFileCheckDate              time.Time `json:"LastFileCheckDate,omitempty"`
	LastSuccessFileCheckDate       time.Time `json:"LastSuccessFileCheckDate,omitempty"`
	LastFailureFileCheckDate       time.Time `json:"LastFailureFileCheckDate,omitempty"`
	LastFileCheckTaskId            string    `json:"LastFileCheckTaskId,omitempty"`
	FailedFileCheckAttempts        int       `json:"FailedFileCheckAttempts,omitempty"`
	LastFileChangeDate             time.Time `json:"LastFileChangeDate,omitempty"`
	LastSuccessFileChangeDate      time.Time `json:"LastSuccessFileChangeDate,omitempty"`
	LastFailureFileChangeDate      time.Time `json:"LastFailureFileChangeDate,omitempty"`
	LastFileChangeTaskId           time.Time `json:"LastFileChangeTaskId,omitempty"`
	FailedFileChangeAttempts       int       `json:"FailedFileChangeAttempts,omitempty"`
	LastDemoteAccountDate          time.Time `json:"LastDemoteAccountDate,omitempty"`
	LastSuccessDemoteAccountDate   time.Time `json:"LastSuccessDemoteAccountDate,omitempty"`
	LastFailureDemoteAccountDate   time.Time `json:"LastFailureDemoteAccountDate,omitempty"`
	LastDemoteAccountTaskId        string    `json:"LastDemoteAccountTaskId,omitempty"`
	FailedDemoteAccountAttempts    int       `json:"FailedDemoteAccountAttempts,omitempty"`
	NextDemoteAccountDate          time.Time `json:"NextDemoteAccountDate,omitempty"`
	LastElevateAccountDate         time.Time `json:"LastElevateAccountDate,omitempty"`
	LastSuccessElevateAccountDate  time.Time `json:"LastSuccessElevateAccountDate,omitempty"`
	LastFailureElevateAccountDate  time.Time `json:"LastFailureElevateAccountDate,omitempty"`
	LastElevateAccountTaskId       string    `json:"LastElevateAccountTaskId,omitempty"`
	FailedElevateAccountAttempts   int       `json:"FailedElevateAccountAttempts,omitempty"`
	NextElevateAccountDate         time.Time `json:"NextElevateAccountDate,omitempty"`
}

// AssetAccount represents a privileged account managed by Safeguard,
// containing its configuration, credentials and relationships.
type AssetAccount struct {
	apiClient *SafeguardClient `json:"-"`

	Id                           int                  `json:"Id,omitempty"`
	Name                         string               `json:"Name,omitempty"`
	DistinguishedName            string               `json:"DistinguishedName,omitempty"`
	DomainName                   string               `json:"DomainName,omitempty"`
	AccountNamespace             string               `json:"AccountNamespace,omitempty"`
	Description                  string               `json:"Description,omitempty"`
	AltLoginName                 string               `json:"AltLoginName,omitempty"`
	PrivilegeGroupMembershipList []string             `json:"PrivilegeGroupMembershipList,omitempty"`
	CreatedDate                  string               `json:"CreatedDate,omitempty"`
	CreatedByUserId              int                  `json:"CreatedByUserId,omitempty"`
	CreatedByUserDisplayName     string               `json:"CreatedByUserDisplayName,omitempty"`
	ManagedBy                    []Identity           `json:"ManagedBy,omitempty"`
	Disabled                     bool                 `json:"Disabled,omitempty"`
	IsServiceAccount             bool                 `json:"IsServiceAccount,omitempty"`
	IsApplicationAccount         bool                 `json:"IsApplicationAccount,omitempty"`
	SharedServiceAccount         bool                 `json:"SharedServiceAccount,omitempty"`
	Tags                         []Tag                `json:"Tags,omitempty"`
	Asset                        Asset                `json:"Asset,omitempty"`
	PasswordProfile              Profile              `json:"PasswordProfile,omitempty"`
	SshKeyProfile                Profile              `json:"SshKeyProfile,omitempty"`
	RequestProperties            RequestProperties    `json:"RequestProperties,omitempty"`
	Platform                     Platform             `json:"Platform,omitempty"`
	DiscoveredProperties         DiscoveredProperties `json:"DiscoveredProperties,omitempty"`
	DirectoryProperties          DirectoryProperties  `json:"DirectoryProperties,omitempty"`
	SyncGroup                    SyncGroup            `json:"SyncGroup,omitempty"`
	SshKeySyncGroup              SyncGroup            `json:"SshKeySyncGroup,omitempty"`
	HasPassword                  bool                 `json:"HasPassword,omitempty"`
	HasSshKey                    bool                 `json:"HasSshKey,omitempty"`
	HasTotpAuthenticator         bool                 `json:"HasTotpAuthenticator,omitempty"`
	HasApiKeys                   bool                 `json:"HasApiKeys,omitempty"`
	HasFile                      bool                 `json:"HasFile,omitempty"`
	TaskProperties               TaskProperties       `json:"TaskProperties,omitempty"`
}

func (a AssetAccount) SetClient(c *SafeguardClient) any {
	a.apiClient = c
	return a
}

// SshKey represents SSH key information as specified in swagger.json
type SshKey struct {
	PrivateKey        string `json:"PrivateKey,omitempty"`
	Passphrase        string `json:"Passphrase,omitempty"`
	PublicKey         string `json:"PublicKey,omitempty"`
	Comment           string `json:"Comment,omitempty"`
	Fingerprint       string `json:"Fingerprint,omitempty"`
	FingerprintSha256 string `json:"FingerprintSha256,omitempty"`
	KeyType           string `json:"KeyType,omitempty"`
	KeyLength         int    `json:"KeyLength,omitempty"`
}

// SshKeyFormat specifies supported SSH key formats
type SshKeyFormat string

const (
	SshKeyFormatUnknown   SshKeyFormat = "Unknown"
	SshKeyFormatOpenSsh   SshKeyFormat = "OpenSsh"
	SshKeyFormatSSH2      SshKeyFormat = "SSH2"
	SshKeyFormatPuttyPPK  SshKeyFormat = "PuttyPPK"
	SshKeyFormatSecureCRT SshKeyFormat = "SecureCRT"
)

// SshKeyType specifies supported SSH key types
type SshKeyType string

const (
	SshKeyTypeUnknown SshKeyType = "Unknown"
	SshKeyTypeDSA     SshKeyType = "DSA"
	SshKeyTypeRSA     SshKeyType = "RSA"
	SshKeyTypeECDSA   SshKeyType = "ECDSA"
	SshKeyTypeED25519 SshKeyType = "ED25519"
)

// ToJson serializes an AssetAccount object into its JSON string representation.
// Returns the JSON string and any error that occurred during marshalling.
func (a AssetAccount) ToJson() (string, error) {
	assetAccountJSON, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(assetAccountJSON), nil
}

// AssetAccountBatchResponse represents a single item in the batch response array
type AssetAccountBatchResponse struct {
	apiClient *SafeguardClient `json:"-"`

	Response         AssetAccount `json:"Response,omitempty"`
	StatusCode       string       `json:"StatusCode,omitempty"`
	StatusCodeNumber int          `json:"StatusCodeNumber,omitempty"`
	IsSuccess        bool         `json:"IsSuccess,omitempty"`
	Error            ApiError     `json:"Error,omitempty"`
	Request          AssetAccount `json:"Request,omitempty"`
}

func (a AssetAccountBatchResponse) SetClient(c *SafeguardClient) any {
	a.apiClient = c
	return a
}

func (ab AssetAccountBatchResponse) hasError() error {
	if ab.IsSuccess {
		return nil
	}

	return fmt.Errorf("error: %s", ab.Error.Message)
}

// GetAssetAccounts retrieves accounts matching the provided filter.
//
// Parameters:
//   - filter: Query parameters for filtering accounts
//
// Returns:
//   - []AssetAccount: Matching accounts
//   - error: API or unmarshalling errors
func (c *SafeguardClient) GetAssetAccounts(filter Filter) ([]AssetAccount, error) {
	var users []AssetAccount

	query := "AssetAccounts" + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &users); err != nil {
		return nil, err
	}

	return addClientToSlice(c, users), nil
}

// GetAssetAccount retrieves a specific asset account by ID from Safeguard.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The ID of the asset account to retrieve
//   - fields: Specific fields to include in the response
//
// Returns:
//   - AssetAccount: The requested asset account
//   - error: An error if the request fails, nil otherwise
func (c *SafeguardClient) GetAssetAccount(id int, fields Fields) (AssetAccount, error) {
	var user AssetAccount

	query := fmt.Sprintf("AssetAccounts/%d", id)
	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return user, err
	}
	if err := json.Unmarshal(response, &user); err != nil {
		return user, err
	}

	return addClient(c, user), nil
}

// DeleteAssetAccount deletes an asset account from Safeguard.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The ID of the asset account to delete
//
// Returns:
//   - error: An error if the deletion fails, nil otherwise
func (c *SafeguardClient) DeleteAssetAccount(id int) error {
	query := fmt.Sprintf("AssetAccounts/%d", id)

	_, err := c.DeleteRequest(query)
	if err != nil {
		return err
	}

	return nil
}

// Delete permanently removes the account.
//
// Returns:
//   - error: Deletion errors
func (a AssetAccount) Delete() error {
	return a.apiClient.DeleteAssetAccount(a.Id)
}

// ChangePassword initiates a password change operation for the asset account.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//
// Returns:
//   - PasswordActivityLog: Log details of the password change activity
//   - error: An error if the password change fails or cannot be initiated
func (a AssetAccount) ChangePassword() (ActivityLog, error) {
	var log ActivityLog

	query := fmt.Sprintf("AssetAccounts/%d/ChangePassword", a.Id)

	response, err := a.apiClient.PostRequest(query, nil)
	if err != nil {
		return log, err
	}

	if json.Unmarshal(response, &log) != nil {
		return log, err
	}

	return addClient(a.apiClient, log), nil
}

// CheckPassword verifies if the current password for the asset account is valid.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//
// Returns:
//   - PasswordActivityLog: Log details of the password check activity
//   - error: An error if the password check fails or cannot be initiated
func (a AssetAccount) CheckPassword() (ActivityLog, error) {
	var log ActivityLog

	query := fmt.Sprintf("AssetAccounts/%d/CheckPassword", a.Id)

	response, err := a.apiClient.PostRequest(query, nil)
	if err != nil {
		return log, err
	}

	if json.Unmarshal(response, &log) != nil {
		return log, err
	}

	return addClient(a.apiClient, log), nil
}

// CreateAssetAccount creates a new asset account in Safeguard.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - assetAccount: The AssetAccount object containing the account details to create
//
// Returns:
//   - AssetAccount: The newly created asset account with updated fields
//   - error: An error if the creation fails, nil otherwise
func (c *SafeguardClient) CreateAssetAccount(assetAccount AssetAccount) (AssetAccount, error) {
	assetAccountsBatch, err := c.batchCreateAssetAccounts([]AssetAccount{assetAccount})
	if err != nil {
		return AssetAccount{}, err
	}

	if len(assetAccountsBatch) > 1 {
		return AssetAccount{}, fmt.Errorf("expected 1 response, got %d", len(assetAccountsBatch))
	}

	return addClient(c, assetAccountsBatch[0].Response), nil
}

// CreateAssetAccounts creates multiple asset accounts in a single batch request.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - assetAccounts: A slice of AssetAccount objects to create
//
// Returns:
//   - []AssetAccount: A slice of the newly created asset accounts
//   - error: An error if any of the creations fail, nil otherwise
func (c *SafeguardClient) CreateAssetAccounts(assetAccounts []AssetAccount) ([]AssetAccount, error) {
	batchCreatedAccounts, err := c.batchCreateAssetAccounts(assetAccounts)
	if err != nil {
		return []AssetAccount{}, err
	}

	var createdAssetAccounts []AssetAccount
	for i := range batchCreatedAccounts {
		createdAssetAccounts = append(createdAssetAccounts, batchCreatedAccounts[i].Response)
	}

	return addClientToSlice(c, createdAssetAccounts), nil
}

// Create adds this account to Safeguard.
// Required fields must be populated before calling.
//
// Returns:
//   - AssetAccount: Created account with server-assigned fields
//   - error: Creation errors
func (a AssetAccount) Create() (AssetAccount, error) {
	return a.apiClient.CreateAssetAccount(a)
}

// UpdateAssetAccount updates an existing asset account in Safeguard.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - assetAccount: The AssetAccount object containing the updated account details
//
// Returns:
//   - AssetAccount: The updated asset account with current fields
//   - error: An error if the update fails, nil otherwise
func (c *SafeguardClient) UpdateAssetAccount(assetAccount AssetAccount) (AssetAccount, error) {
	query := fmt.Sprintf("AssetAccounts/%d", assetAccount.Id)

	assetAccountJSON, err := json.Marshal(assetAccount)
	if err != nil {
		return AssetAccount{}, err
	}

	response, err := c.PutRequest(query, bytes.NewReader(assetAccountJSON))
	if err != nil {
		return AssetAccount{}, err
	}

	var updatedAssetAccount AssetAccount
	err = json.Unmarshal(response, &updatedAssetAccount)
	if err != nil {
		return AssetAccount{}, err
	}

	return addClient(c, updatedAssetAccount), nil
}

// Update modifies the account in Safeguard.
// Only modifiable fields are updated.
//
// Returns:
//   - AssetAccount: Updated account state
//   - error: Update errors
func (a AssetAccount) Update() (AssetAccount, error) {
	return a.apiClient.UpdateAssetAccount(a)
}

// UpdatePasswordProfile updates the password profile for an asset account.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - assetAccount: The AssetAccount object to update
//   - passwordPolicy: The AccountPasswordRule to apply to the account
//
// Returns:
//   - AssetAccount: The updated asset account with the new password profile
//   - error: An error if the update fails, nil otherwise
func (c *SafeguardClient) UpdatePasswordProfile(assetAccount AssetAccount, passwordPolicy AccountPasswordRule) (AssetAccount, error) {
	var passwordProfile Profile
	passwordProfile.Id = passwordPolicy.Id
	passwordProfile.Name = passwordPolicy.Name
	passwordProfile.EffectiveId = passwordPolicy.Id

	assetAccount.PasswordProfile = passwordProfile

	updatedAssetAccount, err := c.UpdateAssetAccount(assetAccount)
	if err != nil {
		return AssetAccount{}, err
	}

	return addClient(c, updatedAssetAccount), nil
}

// UpdatePasswordProfile updates the password profile for this asset account.
// It uses the UpdatePasswordProfile function with the current client.
// Parameters:
//   - passwordPolicy: The AccountPasswordRule to apply to this account
//
// Returns:
//   - AssetAccount: The updated asset account with the new password profile
//   - error: An error if the update fails, nil otherwise
func (a AssetAccount) UpdatePasswordProfile(passwordPolicy AccountPasswordRule) (AssetAccount, error) {
	return a.apiClient.UpdatePasswordProfile(a, passwordPolicy)
}

// DisableAssetAccount disables an asset account in Safeguard.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - assetAccount: The AssetAccount to disable
//
// Returns:
//   - AssetAccount: The updated asset account reflecting the disabled state
//   - error: An error if the disable operation fails, nil otherwise
func (c *SafeguardClient) DisableAssetAccount(assetAccount AssetAccount) (AssetAccount, error) {
	query := fmt.Sprintf("AssetAccounts/%d/Disable", assetAccount.Id)

	response, err := c.PostRequest(query, nil)
	if err != nil {
		return AssetAccount{}, err
	}
	var updatedAssetAccount AssetAccount
	err = json.Unmarshal(response, &updatedAssetAccount)
	if err != nil {
		return AssetAccount{}, err
	}

	return addClient(c, updatedAssetAccount), nil
}

// Disable deactivates the account for password management.
//
// Returns:
//   - AssetAccount: Updated account showing disabled
//   - error: Disable operation errors
func (a AssetAccount) Disable() (AssetAccount, error) {
	return a.apiClient.DisableAssetAccount(a)
}

// EnableAssetAccount enables a previously disabled asset account in Safeguard.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - assetAccount: The AssetAccount to enable
//
// Returns:
//   - AssetAccount: The updated asset account reflecting the enabled state
//   - error: An error if the enable operation fails, nil otherwise
func (c *SafeguardClient) EnableAssetAccount(assetAccount AssetAccount) (AssetAccount, error) {
	query := fmt.Sprintf("AssetAccounts/%d/Enable", assetAccount.Id)

	response, err := c.PostRequest(query, nil)
	if err != nil {
		return AssetAccount{}, err
	}
	var updatedAssetAccount AssetAccount
	err = json.Unmarshal(response, &updatedAssetAccount)
	if err != nil {
		return AssetAccount{}, err
	}

	return addClient(c, updatedAssetAccount), nil
}

// Enable activates the account for password management.
//
// Returns:
//   - AssetAccount: Updated account showing enabled
//   - error: Enable operation errors
func (a AssetAccount) Enable() (AssetAccount, error) {
	return a.apiClient.EnableAssetAccount(a)
}

// batchCreateAssetAccounts handles the batch creation of multiple asset accounts.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - accessRequests: A slice of AssetAccount objects to create in batch
//
// Returns:
//   - []AssetAccountBatchResponse: A slice of responses for each account creation attempt
//   - error: An error if the batch request fails or if any individual creation fails
func (c *SafeguardClient) batchCreateAssetAccounts(accessRequests []AssetAccount) ([]AssetAccountBatchResponse, error) {
	requestBody, err := json.Marshal(accessRequests)
	if err != nil {
		return []AssetAccountBatchResponse{}, err
	}
	response, err := c.PostRequest("AssetAccounts/BatchCreate", bytes.NewReader(requestBody))
	if err != nil {
		return []AssetAccountBatchResponse{}, err
	}

	var createdAssetAccounts []AssetAccountBatchResponse
	if err := json.Unmarshal(response, &createdAssetAccounts); err != nil {
		return []AssetAccountBatchResponse{}, err
	}

	var collectedErrors error
	for i := range createdAssetAccounts {
		if err := createdAssetAccounts[i].hasError(); err != nil {
			collectedErrors = fmt.Errorf("%v\n%v", collectedErrors, err)
		}
	}

	if collectedErrors != nil {
		return createdAssetAccounts, collectedErrors
	}

	return addClientToSlice(c, createdAssetAccounts), nil
}

// SuspendAssetAccount suspends an asset account in Safeguard.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - a: The AssetAccount to suspend
//
// Returns:
//   - PasswordActivityLog: Log details of the suspend activity
//   - error: An error if the suspend operation fails, nil otherwise
func (c *SafeguardClient) SuspendAssetAccount(a AssetAccount) (ActivityLog, error) {
	query := fmt.Sprintf("AssetAccounts/%d/SuspendAccount", a.Id)

	response, err := c.PostRequest(query, nil)
	if err != nil {
		return ActivityLog{}, err
	}

	var log ActivityLog
	if err := json.Unmarshal(response, &log); err != nil {
		return ActivityLog{}, err
	}

	return addClient(c, log), nil
}

// Suspend temporarily disables the account on target system.
//
// Returns:
//   - PasswordActivityLog: Suspension details
//   - error: Suspend operation errors
func (a AssetAccount) Suspend() (ActivityLog, error) {
	return a.apiClient.SuspendAssetAccount(a)
}
