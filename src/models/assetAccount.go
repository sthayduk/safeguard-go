package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sthayduk/safeguard-go/client"
)

// ManagedByUser represents a user who manages the asset account
type ManagedByUser struct {
	DisplayName                       string `json:"DisplayName,omitempty"`
	Id                                int    `json:"Id,omitempty"`
	IdentityProviderId                int    `json:"IdentityProviderId,omitempty"`
	IdentityProviderName              string `json:"IdentityProviderName,omitempty"`
	IdentityProviderTypeReferenceName string `json:"IdentityProviderTypeReferenceName,omitempty"`
	IsSystemOwned                     bool   `json:"IsSystemOwned,omitempty"`
	Name                              string `json:"Name,omitempty"`
	PrincipalKind                     string `json:"PrincipalKind,omitempty"`
	EmailAddress                      string `json:"EmailAddress,omitempty"`
	DomainName                        string `json:"DomainName,omitempty"`
	FullDisplayName                   string `json:"FullDisplayName,omitempty"`
}

// Tag represents a tag associated with an asset account
type Tag struct {
	Id            int    `json:"Id,omitempty"`
	Name          string `json:"Name,omitempty"`
	Description   string `json:"Description,omitempty"`
	AdminAssigned bool   `json:"AdminAssigned,omitempty"`
}

// Profile represents a profile associated with an asset account
type Profile struct {
	Id            int    `json:"Id,omitempty"`
	Name          string `json:"Name,omitempty"`
	EffectiveId   int    `json:"EffectiveId,omitempty"`
	EffectiveName string `json:"EffectiveName,omitempty"`
}

// DiscoveredGroup represents a discovered group associated with an asset account
type DiscoveredGroup struct {
	DiscoveredGroupId                string `json:"DiscoveredGroupId,omitempty"`
	DiscoveredGroupName              string `json:"DiscoveredGroupName,omitempty"`
	DiscoveredGroupDistinguishedName string `json:"DiscoveredGroupDistinguishedName,omitempty"`
}

// DiscoveredProperties represents properties discovered for an asset account
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

// AssetAccount represents an asset account in Safeguard
type AssetAccount struct {
	client *client.SafeguardClient

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
	ManagedBy                    []ManagedByUser      `json:"ManagedBy,omitempty"`
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

// ToJson converts an AssetAccount to its JSON string representation
func (a AssetAccount) ToJson() (string, error) {
	assetAccountJSON, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(assetAccountJSON), nil
}

// GetAssetAccounts retrieves a list of asset accounts from Safeguard.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - fields: Filter criteria for the request
//
// Returns:
//   - []AssetAccount: A slice of asset accounts matching the filter criteria
//   - error: An error if the request fails, nil otherwise
func GetAssetAccounts(c *client.SafeguardClient, filter client.Filter) ([]AssetAccount, error) {
	var users []AssetAccount

	query := "AssetAccounts" + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &users); err != nil {
		return nil, err
	}

	for u := range users {
		users[u].client = c
	}
	return users, nil
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
func GetAssetAccount(c *client.SafeguardClient, id int, fields client.Fields) (AssetAccount, error) {
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

	user.client = c
	return user, nil
}

// DeleteAssetAccount deletes an asset account from Safeguard.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//   - id: The ID of the asset account to delete
//
// Returns:
//   - error: An error if the deletion fails, nil otherwise
func DeleteAssetAccount(c *client.SafeguardClient, id int) error {
	query := fmt.Sprintf("AssetAccounts/%d", id)

	_, err := c.DeleteRequest(query)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes the AssetAccount from the system.
// It calls the DeleteAssetAccount function with the client and Id of the AssetAccount.
// Returns an error if the deletion fails.
func (a AssetAccount) Delete() error {
	return DeleteAssetAccount(a.client, a.Id)
}

// ChangePassword initiates a password change operation for the asset account.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//
// Returns:
//   - PasswordActivityLog: Log details of the password change activity
//   - error: An error if the password change fails or cannot be initiated
func (a AssetAccount) ChangePassword() (PasswordActivityLog, error) {
	var log PasswordActivityLog
	log.client = a.client

	query := fmt.Sprintf("AssetAccounts/%d/ChangePassword", a.Id)

	response, err := a.client.PostRequest(query, nil)
	if err != nil {
		return log, err
	}

	json.Unmarshal(response, &log)
	return log, nil
}

// CheckPassword verifies if the current password for the asset account is valid.
// Parameters:
//   - c: The SafeguardClient instance for making API requests
//
// Returns:
//   - PasswordActivityLog: Log details of the password check activity
//   - error: An error if the password check fails or cannot be initiated
func (a AssetAccount) CheckPassword() (PasswordActivityLog, error) {
	var log PasswordActivityLog
	log.client = a.client

	query := fmt.Sprintf("AssetAccounts/%d/CheckPassword", a.Id)

	response, err := a.client.PostRequest(query, nil)
	if err != nil {
		return log, err
	}

	json.Unmarshal(response, &log)
	return log, nil
}

func CreateAssetAccount(c *client.SafeguardClient, assetAccount AssetAccount) (AssetAccount, error) {
	// https://spp-itd-01.itdesign.at/service/core/v4/Assets/464/DirectoryAccounts?searchScope=Subtree&searchName=da-hayduk

	query := "AssetAccounts"

	assetAccountJSON, err := json.Marshal(assetAccount)
	if err != nil {
		return AssetAccount{}, err
	}

	response, err := c.PostRequest(query, bytes.NewReader(assetAccountJSON))
	if err != nil {
		return AssetAccount{}, err
	}

	var createdAssetAccount AssetAccount
	err = json.Unmarshal(response, &createdAssetAccount)
	if err != nil {
		return AssetAccount{}, err
	}

	createdAssetAccount.client = c
	return createdAssetAccount, nil
}

func (a AssetAccount) Create() (AssetAccount, error) {
	return CreateAssetAccount(a.client, a)
}
