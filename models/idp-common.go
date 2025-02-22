package models

import (
	"time"
)

// IdentityProvider represents the structure for the given JSON array
type IdentityProvider struct {
	Id                       int                   `json:"Id,omitempty"`
	TypeReferenceName        string                `json:"TypeReferenceName,omitempty"`
	Name                     string                `json:"Name,omitempty"`
	Description              string                `json:"Description,omitempty"`
	NetworkAddress           string                `json:"NetworkAddress,omitempty"`
	IsSystemOwned            bool                  `json:"IsSystemOwned,omitempty"`
	IsDirectory              bool                  `json:"IsDirectory,omitempty"`
	RstsProviderId           string                `json:"RstsProviderId,omitempty"`
	RstsProviderScope        string                `json:"RstsProviderScope,omitempty"`
	StarlingProperties       StarlingProperties    `json:"StarlingProperties,omitempty"`
	RadiusProperties         RadiusProperties      `json:"RadiusProperties,omitempty"`
	ExternalFederation       ExternalFederation    `json:"ExternalFederationProperties,omitempty"`
	Fido2Properties          Fido2Properties       `json:"Fido2Properties,omitempty"`
	OneLoginMfa              OneLoginMfaProperties `json:"OneLoginMfaProperties,omitempty"`
	ScimProperties           ScimProperties        `json:"ScimProperties,omitempty"`
	DirectoryProperties      DirectoryProperties   `json:"DirectoryProperties,omitempty"`
	CreatedDate              time.Time             `json:"CreatedDate,omitempty"`
	CreatedByUserId          int                   `json:"CreatedByUserId,omitempty"`
	CreatedByUserDisplayName string                `json:"CreatedByUserDisplayName,omitempty"`
}

type StarlingProperties struct {
	HasApiKey bool `json:"HasApiKey,omitempty"`
}

type RadiusProperties struct {
	ServerAddress1                      string `json:"ServerAddress1,omitempty"`
	ServerAddress2                      string `json:"ServerAddress2,omitempty"`
	ServerPort                          int    `json:"ServerPort,omitempty"`
	SharedSecret                        string `json:"SharedSecret,omitempty"`
	Timeout                             int    `json:"Timeout,omitempty"`
	Retries                             int    `json:"Retries,omitempty"`
	PreAuthenticateForChallengeResponse bool   `json:"PreAuthenticateForChallengeResponse,omitempty"`
	AlwaysMaskUserInput                 bool   `json:"AlwaysMaskUserInput,omitempty"`
}

type ExternalFederation struct {
	Realm                  string `json:"Realm,omitempty"`
	FederationMetadata     string `json:"FederationMetadata,omitempty"`
	AuthnContextClasses    string `json:"AuthnContextClasses,omitempty"`
	AuthnContextComparison string `json:"AuthnContextComparison,omitempty"`
	NameIDFormat           string `json:"NameIDFormat,omitempty"`
	RequireAuthentication  bool   `json:"RequireAuthentication,omitempty"`
	ApplicationIdOverride  string `json:"ApplicationIdOverride,omitempty"`
}

type Fido2Properties struct {
	DomainSuffix string `json:"DomainSuffix,omitempty"`
}

type OneLoginMfaProperties struct {
	DnsHostName  string `json:"DnsHostName,omitempty"`
	ClientId     string `json:"ClientId,omitempty"`
	ClientSecret string `json:"ClientSecret,omitempty"`
}

type ScimProperties struct {
	UserTemplate      UserTemplate `json:"UserTemplate,omitempty"`
	TenantUrl         string       `json:"TenantUrl,omitempty"`
	HasToken          bool         `json:"HasToken,omitempty"`
	TokenCreationDate time.Time    `json:"TokenCreationDate,omitempty"`
}

type UserTemplate struct {
	PrimaryAuthenticationProviderId     int      `json:"PrimaryAuthenticationProviderId,omitempty"`
	PrimaryAuthenticationProviderType   string   `json:"PrimaryAuthenticationProviderTypeReferenceName,omitempty"`
	PrimaryAuthenticationProviderName   string   `json:"PrimaryAuthenticationProviderName,omitempty"`
	RequireCertificateAuthentication    bool     `json:"RequireCertificateAuthentication,omitempty"`
	SecondaryAuthenticationProviderId   int      `json:"SecondaryAuthenticationProviderId,omitempty"`
	SecondaryAuthenticationProviderType string   `json:"SecondaryAuthenticationProviderTypeReferenceName,omitempty"`
	SecondaryAuthenticationProviderName string   `json:"SecondaryAuthenticationProviderName,omitempty"`
	AllowPersonalAccounts               bool     `json:"AllowPersonalAccounts,omitempty"`
	AdminRoles                          []string `json:"AdminRoles,omitempty"`
}

type DirectoryProperties struct {
	ForestRootDomain               string               `json:"ForestRootDomain,omitempty"`
	SynchronizationIntervalMinutes int                  `json:"SynchronizationIntervalMinutes,omitempty"`
	LastSynchronizedDate           time.Time            `json:"LastSynchronizedDate,omitempty"`
	NextSynchronizedDate           time.Time            `json:"NextSynchronizedDate,omitempty"`
	DeleteSyncIntervalMinutes      int                  `json:"DeleteSyncIntervalMinutes,omitempty"`
	LastDeleteSyncDate             time.Time            `json:"LastDeleteSyncDate,omitempty"`
	NextDeleteSyncDate             time.Time            `json:"NextDeleteSyncDate,omitempty"`
	LastSuccessSynchronizedDate    time.Time            `json:"LastSuccessSynchronizedDate,omitempty"`
	LastFailureSynchronizedDate    time.Time            `json:"LastFailureSynchronizedDate,omitempty"`
	FailedSyncAttempts             int                  `json:"FailedSyncAttempts,omitempty"`
	LastSuccessDeleteSyncDate      time.Time            `json:"LastSuccessDeleteSyncDate,omitempty"`
	LastFailureDeleteSyncDate      time.Time            `json:"LastFailureDeleteSyncDate,omitempty"`
	FailedDeleteSyncAttempts       int                  `json:"FailedDeleteSyncAttempts,omitempty"`
	Domains                        []Domain             `json:"Domains,omitempty"`
	DomainControllers              []DomainController   `json:"DomainControllers,omitempty"`
	SchemaProperties               SchemaProperties     `json:"SchemaProperties,omitempty"`
	ConnectionProperties           ConnectionProperties `json:"ConnectionProperties,omitempty"`
}

type Domain struct {
	DomainName     string `json:"DomainName,omitempty"`
	NetBiosName    string `json:"NetBiosName,omitempty"`
	DomainUniqueId string `json:"DomainUniqueId,omitempty"`
	NamingContext  string `json:"NamingContext,omitempty"`
	IsVisible      bool   `json:"IsVisible,omitempty"`
	IsForestRoot   bool   `json:"IsForestRoot,omitempty"`
}

type DomainController struct {
	NetworkAddress string `json:"NetworkAddress,omitempty"`
	DomainName     string `json:"DomainName,omitempty"`
	IsWritable     bool   `json:"IsWritable,omitempty"`
	ServerType     string `json:"ServerType,omitempty"`
}

type SchemaProperties struct {
	UserProperties  UserProperties  `json:"UserProperties,omitempty"`
	GroupProperties GroupProperties `json:"GroupProperties,omitempty"`
}

type UserProperties struct {
	UserClassType                                                  []string `json:"UserClassType,omitempty"`
	UserNameAttribute                                              string   `json:"UserNameAttribute,omitempty"`
	FirstNameAttribute                                             string   `json:"FirstNameAttribute,omitempty"`
	LastNameAttribute                                              string   `json:"LastNameAttribute,omitempty"`
	DescriptionAttribute                                           string   `json:"DescriptionAttribute,omitempty"`
	MailAttribute                                                  string   `json:"MailAttribute,omitempty"`
	PhoneAttribute                                                 string   `json:"PhoneAttribute,omitempty"`
	MobileAttribute                                                string   `json:"MobileAttribute,omitempty"`
	DirectoryGroupSyncAttributeForExternalFederationAuthentication string   `json:"DirectoryGroupSyncAttributeForExternalFederationAuthentication,omitempty"`
	DirectoryGroupSyncAttributeForRadiusAuthentication             string   `json:"DirectoryGroupSyncAttributeForRadiusAuthentication,omitempty"`
	DirectoryGroupSyncAttributeForManagedObjects                   string   `json:"DirectoryGroupSyncAttributeForManagedObjects,omitempty"`
}

type GroupProperties struct {
	GroupClassType       []string `json:"GroupClassType,omitempty"`
	MemberAttribute      string   `json:"MemberAttribute,omitempty"`
	NameAttribute        string   `json:"NameAttribute,omitempty"`
	DescriptionAttribute string   `json:"DescriptionAttribute,omitempty"`
}

type ConnectionProperties struct {
	UseSslEncryption                         bool   `json:"UseSslEncryption,omitempty"`
	VerifySslCertificate                     bool   `json:"VerifySslCertificate,omitempty"`
	ServiceAccountUniqueObjectId             string `json:"ServiceAccountUniqueObjectId,omitempty"`
	ServiceAccountSecurityId                 string `json:"ServiceAccountSecurityId,omitempty"`
	ServiceAccountId                         int    `json:"ServiceAccountId,omitempty"`
	ServiceAccountName                       string `json:"ServiceAccountName,omitempty"`
	EffectiveServiceAccountName              string `json:"EffectiveServiceAccountName,omitempty"`
	ServiceAccountDomainName                 string `json:"ServiceAccountDomainName,omitempty"`
	ServiceAccountDistinguishedName          string `json:"ServiceAccountDistinguishedName,omitempty"`
	EffectiveServiceAccountDistinguishedName string `json:"EffectiveServiceAccountDistinguishedName,omitempty"`
	ServiceAccountCredentialType             string `json:"ServiceAccountCredentialType,omitempty"`
	ServiceAccountPassword                   string `json:"ServiceAccountPassword,omitempty"`
	ServiceAccountHasPassword                bool   `json:"ServiceAccountHasPassword,omitempty"`
	ServiceAccountSshKey                     SshKey `json:"ServiceAccountSshKey,omitempty"`
	ServiceAccountHasSshKey                  bool   `json:"ServiceAccountHasSshKey,omitempty"`
	Port                                     int    `json:"Port,omitempty"`
	ServiceAccountAssetId                    int    `json:"ServiceAccountAssetId,omitempty"`
	ServiceAccountAssetName                  string `json:"ServiceAccountAssetName,omitempty"`
	ServiceAccountAssetPlatformId            int    `json:"ServiceAccountAssetPlatformId,omitempty"`
	ServiceAccountAssetPlatformType          string `json:"ServiceAccountAssetPlatformType,omitempty"`
	ServiceAccountAssetPlatformDisplayName   string `json:"ServiceAccountAssetPlatformDisplayName,omitempty"`
	ServiceAccountNetbiosName                string `json:"ServiceAccountNetbiosName,omitempty"`
}

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
