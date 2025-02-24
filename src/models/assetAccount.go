package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sthayduk/safeguard-go/src/client"
)

type ManagedByUser struct {
	DisplayName                       string `json:"DisplayName"`
	Id                                int    `json:"Id"`
	IdentityProviderId                int    `json:"IdentityProviderId"`
	IdentityProviderName              string `json:"IdentityProviderName"`
	IdentityProviderTypeReferenceName string `json:"IdentityProviderTypeReferenceName"`
	IsSystemOwned                     bool   `json:"IsSystemOwned"`
	Name                              string `json:"Name"`
	PrincipalKind                     string `json:"PrincipalKind"`
	EmailAddress                      string `json:"EmailAddress"`
	DomainName                        string `json:"DomainName"`
	FullDisplayName                   string `json:"FullDisplayName"`
}

type Tag struct {
	Id            int    `json:"Id"`
	Name          string `json:"Name"`
	Description   string `json:"Description"`
	AdminAssigned bool   `json:"AdminAssigned"`
}

type Profile struct {
	Id            int    `json:"Id"`
	Name          string `json:"Name"`
	EffectiveId   int    `json:"EffectiveId"`
	EffectiveName string `json:"EffectiveName"`
}

type DiscoveredGroup struct {
	DiscoveredGroupId                string `json:"DiscoveredGroupId"`
	DiscoveredGroupName              string `json:"DiscoveredGroupName"`
	DiscoveredGroupDistinguishedName string `json:"DiscoveredGroupDistinguishedName"`
}

type DiscoveredProperties struct {
	AccountDiscoveryScheduleId   int               `json:"AccountDiscoveryScheduleId"`
	AccountDiscoveryScheduleName string            `json:"AccountDiscoveryScheduleName"`
	DiscoveredUserId             string            `json:"DiscoveredUserId"`
	DiscoveredDate               string            `json:"DiscoveredDate"`
	DiscoveredGroups             []DiscoveredGroup `json:"DiscoveredGroups"`
}

type AssetDirectoryProperties struct {
	NetbiosName string `json:"NetbiosName"`
	ObjectGuid  string `json:"ObjectGuid"`
	ObjectSid   string `json:"ObjectSid"`
}

type SyncGroup struct {
	Id       int    `json:"Id"`
	Name     string `json:"Name"`
	Priority int    `json:"Priority"`
	Disabled bool   `json:"Disabled"`
}

type TaskProperties struct {
	HasAccountTaskFailure          bool      `json:"HasAccountTaskFailure"`
	LastPasswordCheckDate          time.Time `json:"LastPasswordCheckDate"`
	LastSuccessPasswordCheckDate   time.Time `json:"LastSuccessPasswordCheckDate"`
	LastFailurePasswordCheckDate   time.Time `json:"LastFailurePasswordCheckDate"`
	LastPasswordCheckTaskId        string    `json:"LastPasswordCheckTaskId"`
	FailedPasswordCheckAttempts    int       `json:"FailedPasswordCheckAttempts"`
	NextPasswordCheckDate          time.Time `json:"NextPasswordCheckDate"`
	LastPasswordChangeDate         time.Time `json:"LastPasswordChangeDate"`
	LastSuccessPasswordChangeDate  time.Time `json:"LastSuccessPasswordChangeDate"`
	LastFailurePasswordChangeDate  time.Time `json:"LastFailurePasswordChangeDate"`
	LastPasswordChangeTaskId       string    `json:"LastPasswordChangeTaskId"`
	FailedPasswordChangeAttempts   int       `json:"FailedPasswordChangeAttempts"`
	NextPasswordChangeDate         time.Time `json:"NextPasswordChangeDate"`
	LastSshKeyCheckDate            time.Time `json:"LastSshKeyCheckDate"`
	LastSuccessSshKeyCheckDate     time.Time `json:"LastSuccessSshKeyCheckDate"`
	LastFailureSshKeyCheckDate     time.Time `json:"LastFailureSshKeyCheckDate"`
	LastSshKeyCheckTaskId          string    `json:"LastSshKeyCheckTaskId"`
	FailedSshKeyCheckAttempts      int       `json:"FailedSshKeyCheckAttempts"`
	NextSshKeyCheckDate            time.Time `json:"NextSshKeyCheckDate"`
	LastSshKeyChangeDate           time.Time `json:"LastSshKeyChangeDate"`
	LastSuccessSshKeyChangeDate    time.Time `json:"LastSuccessSshKeyChangeDate"`
	LastFailureSshKeyChangeDate    time.Time `json:"LastFailureSshKeyChangeDate"`
	LastSshKeyChangeTaskId         string    `json:"LastSshKeyChangeTaskId"`
	FailedSshKeyChangeAttempts     int       `json:"FailedSshKeyChangeAttempts"`
	NextSshKeyChangeDate           time.Time `json:"NextSshKeyChangeDate"`
	LastSshKeyDiscoveryDate        time.Time `json:"LastSshKeyDiscoveryDate"`
	LastSuccessSshKeyDiscoveryDate time.Time `json:"LastSuccessSshKeyDiscoveryDate"`
	LastFailureSshKeyDiscoveryDate time.Time `json:"LastFailureSshKeyDiscoveryDate"`
	LastSshKeyDiscoveryTaskId      string    `json:"LastSshKeyDiscoveryTaskId"`
	FailedSshKeyDiscoveryAttempts  int       `json:"FailedSshKeyDiscoveryAttempts"`
	NextSshKeyDiscoveryDate        time.Time `json:"NextSshKeyDiscoveryDate"`
	LastSshKeyRevokeDate           time.Time `json:"LastSshKeyRevokeDate"`
	LastSuccessSshKeyRevokeDate    time.Time `json:"LastSuccessSshKeyRevokeDate"`
	LastFailureSshKeyRevokeDate    time.Time `json:"LastFailureSshKeyRevokeDate"`
	LastSshKeyRevokeTaskId         string    `json:"LastSshKeyRevokeTaskId"`
	FailedSshKeyRevokeAttempts     int       `json:"FailedSshKeyRevokeAttempts"`
	LastSuspendAccountDate         time.Time `json:"LastSuspendAccountDate"`
	LastSuccessSuspendAccountDate  time.Time `json:"LastSuccessSuspendAccountDate"`
	LastFailureSuspendAccountDate  time.Time `json:"LastFailureSuspendAccountDate"`
	LastSuspendAccountTaskId       string    `json:"LastSuspendAccountTaskId"`
	FailedSuspendAccountAttempts   int       `json:"FailedSuspendAccountAttempts"`
	NextSuspendAccountDate         time.Time `json:"NextSuspendAccountDate"`
	LastRestoreAccountDate         time.Time `json:"LastRestoreAccountDate"`
	LastSuccessRestoreAccountDate  time.Time `json:"LastSuccessRestoreAccountDate"`
	LastFailureRestoreAccountDate  time.Time `json:"LastFailureRestoreAccountDate"`
	LastRestoreAccountTaskId       string    `json:"LastRestoreAccountTaskId"`
	FailedRestoreAccountAttempts   int       `json:"FailedRestoreAccountAttempts"`
	NextRestoreAccountDate         time.Time `json:"NextRestoreAccountDate"`
	FailedApiKeyCheckAttempts      int       `json:"FailedApiKeyCheckAttempts"`
	FailedApiKeyChangeAttempts     int       `json:"FailedApiKeyChangeAttempts"`
	LastFileCheckDate              time.Time `json:"LastFileCheckDate"`
	LastSuccessFileCheckDate       time.Time `json:"LastSuccessFileCheckDate"`
	LastFailureFileCheckDate       time.Time `json:"LastFailureFileCheckDate"`
	LastFileCheckTaskId            string    `json:"LastFileCheckTaskId"`
	FailedFileCheckAttempts        int       `json:"FailedFileCheckAttempts"`
	LastFileChangeDate             time.Time `json:"LastFileChangeDate"`
	LastSuccessFileChangeDate      time.Time `json:"LastSuccessFileChangeDate"`
	LastFailureFileChangeDate      time.Time `json:"LastFailureFileChangeDate"`
	LastFileChangeTaskId           time.Time `json:"LastFileChangeTaskId"`
	FailedFileChangeAttempts       int       `json:"FailedFileChangeAttempts"`
	LastDemoteAccountDate          time.Time `json:"LastDemoteAccountDate"`
	LastSuccessDemoteAccountDate   time.Time `json:"LastSuccessDemoteAccountDate"`
	LastFailureDemoteAccountDate   time.Time `json:"LastFailureDemoteAccountDate"`
	LastDemoteAccountTaskId        string    `json:"LastDemoteAccountTaskId"`
	FailedDemoteAccountAttempts    int       `json:"FailedDemoteAccountAttempts"`
	NextDemoteAccountDate          time.Time `json:"NextDemoteAccountDate"`
	LastElevateAccountDate         time.Time `json:"LastElevateAccountDate"`
	LastSuccessElevateAccountDate  time.Time `json:"LastSuccessElevateAccountDate"`
	LastFailureElevateAccountDate  time.Time `json:"LastFailureElevateAccountDate"`
	LastElevateAccountTaskId       string    `json:"LastElevateAccountTaskId"`
	FailedElevateAccountAttempts   int       `json:"FailedElevateAccountAttempts"`
	NextElevateAccountDate         time.Time `json:"NextElevateAccountDate"`
}

type AssetAccount struct {
	client *client.SafeguardClient

	Id                           int                  `json:"Id"`
	Name                         string               `json:"Name"`
	DistinguishedName            string               `json:"DistinguishedName"`
	DomainName                   string               `json:"DomainName"`
	AccountNamespace             string               `json:"AccountNamespace"`
	Description                  string               `json:"Description"`
	AltLoginName                 string               `json:"AltLoginName"`
	PrivilegeGroupMembershipList []string             `json:"PrivilegeGroupMembershipList"`
	CreatedDate                  string               `json:"CreatedDate"`
	CreatedByUserId              int                  `json:"CreatedByUserId"`
	CreatedByUserDisplayName     string               `json:"CreatedByUserDisplayName"`
	ManagedBy                    []ManagedByUser      `json:"ManagedBy"`
	Disabled                     bool                 `json:"Disabled"`
	IsServiceAccount             bool                 `json:"IsServiceAccount"`
	IsApplicationAccount         bool                 `json:"IsApplicationAccount"`
	SharedServiceAccount         bool                 `json:"SharedServiceAccount"`
	Tags                         []Tag                `json:"Tags"`
	Asset                        Asset                `json:"Asset"`
	PasswordProfile              Profile              `json:"PasswordProfile"`
	SshKeyProfile                Profile              `json:"SshKeyProfile"`
	RequestProperties            RequestProperties    `json:"RequestProperties"`
	Platform                     Platform             `json:"Platform"`
	DiscoveredProperties         DiscoveredProperties `json:"DiscoveredProperties"`
	DirectoryProperties          DirectoryProperties  `json:"DirectoryProperties"`
	SyncGroup                    SyncGroup            `json:"SyncGroup"`
	SshKeySyncGroup              SyncGroup            `json:"SshKeySyncGroup"`
	HasPassword                  bool                 `json:"HasPassword"`
	HasSshKey                    bool                 `json:"HasSshKey"`
	HasTotpAuthenticator         bool                 `json:"HasTotpAuthenticator"`
	HasApiKeys                   bool                 `json:"HasApiKeys"`
	HasFile                      bool                 `json:"HasFile"`
	TaskProperties               TaskProperties       `json:"TaskProperties"`
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
func GetAssetAccounts(c *client.SafeguardClient, fields client.Filter) ([]AssetAccount, error) {
	var users []AssetAccount

	query := "AssetAccounts" + fields.ToQueryString()

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
	user.client = c

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
	return user, nil
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
