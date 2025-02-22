package models

import "encoding/json"

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
	HasAccountTaskFailure          bool   `json:"HasAccountTaskFailure"`
	LastPasswordCheckDate          string `json:"LastPasswordCheckDate"`
	LastSuccessPasswordCheckDate   string `json:"LastSuccessPasswordCheckDate"`
	LastFailurePasswordCheckDate   string `json:"LastFailurePasswordCheckDate"`
	LastPasswordCheckTaskId        string `json:"LastPasswordCheckTaskId"`
	FailedPasswordCheckAttempts    int    `json:"FailedPasswordCheckAttempts"`
	NextPasswordCheckDate          string `json:"NextPasswordCheckDate"`
	LastPasswordChangeDate         string `json:"LastPasswordChangeDate"`
	LastSuccessPasswordChangeDate  string `json:"LastSuccessPasswordChangeDate"`
	LastFailurePasswordChangeDate  string `json:"LastFailurePasswordChangeDate"`
	LastPasswordChangeTaskId       string `json:"LastPasswordChangeTaskId"`
	FailedPasswordChangeAttempts   int    `json:"FailedPasswordChangeAttempts"`
	NextPasswordChangeDate         string `json:"NextPasswordChangeDate"`
	LastSshKeyCheckDate            string `json:"LastSshKeyCheckDate"`
	LastSuccessSshKeyCheckDate     string `json:"LastSuccessSshKeyCheckDate"`
	LastFailureSshKeyCheckDate     string `json:"LastFailureSshKeyCheckDate"`
	LastSshKeyCheckTaskId          string `json:"LastSshKeyCheckTaskId"`
	FailedSshKeyCheckAttempts      int    `json:"FailedSshKeyCheckAttempts"`
	NextSshKeyCheckDate            string `json:"NextSshKeyCheckDate"`
	LastSshKeyChangeDate           string `json:"LastSshKeyChangeDate"`
	LastSuccessSshKeyChangeDate    string `json:"LastSuccessSshKeyChangeDate"`
	LastFailureSshKeyChangeDate    string `json:"LastFailureSshKeyChangeDate"`
	LastSshKeyChangeTaskId         string `json:"LastSshKeyChangeTaskId"`
	FailedSshKeyChangeAttempts     int    `json:"FailedSshKeyChangeAttempts"`
	NextSshKeyChangeDate           string `json:"NextSshKeyChangeDate"`
	LastSshKeyDiscoveryDate        string `json:"LastSshKeyDiscoveryDate"`
	LastSuccessSshKeyDiscoveryDate string `json:"LastSuccessSshKeyDiscoveryDate"`
	LastFailureSshKeyDiscoveryDate string `json:"LastFailureSshKeyDiscoveryDate"`
	LastSshKeyDiscoveryTaskId      string `json:"LastSshKeyDiscoveryTaskId"`
	FailedSshKeyDiscoveryAttempts  int    `json:"FailedSshKeyDiscoveryAttempts"`
	NextSshKeyDiscoveryDate        string `json:"NextSshKeyDiscoveryDate"`
	LastSshKeyRevokeDate           string `json:"LastSshKeyRevokeDate"`
	LastSuccessSshKeyRevokeDate    string `json:"LastSuccessSshKeyRevokeDate"`
	LastFailureSshKeyRevokeDate    string `json:"LastFailureSshKeyRevokeDate"`
	LastSshKeyRevokeTaskId         string `json:"LastSshKeyRevokeTaskId"`
	FailedSshKeyRevokeAttempts     int    `json:"FailedSshKeyRevokeAttempts"`
	LastSuspendAccountDate         string `json:"LastSuspendAccountDate"`
	LastSuccessSuspendAccountDate  string `json:"LastSuccessSuspendAccountDate"`
	LastFailureSuspendAccountDate  string `json:"LastFailureSuspendAccountDate"`
	LastSuspendAccountTaskId       string `json:"LastSuspendAccountTaskId"`
	FailedSuspendAccountAttempts   int    `json:"FailedSuspendAccountAttempts"`
	NextSuspendAccountDate         string `json:"NextSuspendAccountDate"`
	LastRestoreAccountDate         string `json:"LastRestoreAccountDate"`
	LastSuccessRestoreAccountDate  string `json:"LastSuccessRestoreAccountDate"`
	LastFailureRestoreAccountDate  string `json:"LastFailureRestoreAccountDate"`
	LastRestoreAccountTaskId       string `json:"LastRestoreAccountTaskId"`
	FailedRestoreAccountAttempts   int    `json:"FailedRestoreAccountAttempts"`
	NextRestoreAccountDate         string `json:"NextRestoreAccountDate"`
	FailedApiKeyCheckAttempts      int    `json:"FailedApiKeyCheckAttempts"`
	FailedApiKeyChangeAttempts     int    `json:"FailedApiKeyChangeAttempts"`
	LastFileCheckDate              string `json:"LastFileCheckDate"`
	LastSuccessFileCheckDate       string `json:"LastSuccessFileCheckDate"`
	LastFailureFileCheckDate       string `json:"LastFailureFileCheckDate"`
	LastFileCheckTaskId            string `json:"LastFileCheckTaskId"`
	FailedFileCheckAttempts        int    `json:"FailedFileCheckAttempts"`
	LastFileChangeDate             string `json:"LastFileChangeDate"`
	LastSuccessFileChangeDate      string `json:"LastSuccessFileChangeDate"`
	LastFailureFileChangeDate      string `json:"LastFailureFileChangeDate"`
	LastFileChangeTaskId           string `json:"LastFileChangeTaskId"`
	FailedFileChangeAttempts       int    `json:"FailedFileChangeAttempts"`
	LastDemoteAccountDate          string `json:"LastDemoteAccountDate"`
	LastSuccessDemoteAccountDate   string `json:"LastSuccessDemoteAccountDate"`
	LastFailureDemoteAccountDate   string `json:"LastFailureDemoteAccountDate"`
	LastDemoteAccountTaskId        string `json:"LastDemoteAccountTaskId"`
	FailedDemoteAccountAttempts    int    `json:"FailedDemoteAccountAttempts"`
	NextDemoteAccountDate          string `json:"NextDemoteAccountDate"`
	LastElevateAccountDate         string `json:"LastElevateAccountDate"`
	LastSuccessElevateAccountDate  string `json:"LastSuccessElevateAccountDate"`
	LastFailureElevateAccountDate  string `json:"LastFailureElevateAccountDate"`
	LastElevateAccountTaskId       string `json:"LastElevateAccountTaskId"`
	FailedElevateAccountAttempts   int    `json:"FailedElevateAccountAttempts"`
	NextElevateAccountDate         string `json:"NextElevateAccountDate"`
}

type AssetAccount struct {
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
