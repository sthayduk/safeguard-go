package safeguard

import (
	"encoding/json"
	"fmt"
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
	DomainName                     string               `json:"DomainName,omitempty"`
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
	DirectoryId                    int                  `json:"DirectoryId"`
	DirectoryName                  string               `json:"DirectoryName"`
	NetbiosName                    string               `json:"NetbiosName"`
	DistinguishedName              string               `json:"DistinguishedName"`
	ObjectGuid                     string               `json:"ObjectGuid"`
	ObjectSid                      string               `json:"ObjectSid"`
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
	Name           string `json:"Name"`
	Port           int    `json:"Port"`
	NetworkAddress string `json:"NetworkAddress,omitempty"`
	DomainName     string `json:"DomainName,omitempty"`
	IsWritable     bool   `json:"IsWritable,omitempty"`
	ServerType     string `json:"ServerType,omitempty"`
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

// GetIdentityProviders retrieves all configured identity providers from Safeguard.
//
// This function returns all authentication sources configured in the system, including:
// - Directory services (Active Directory, LDAP)
// - Federation providers (SAML, OAuth)
// - Other authentication methods (RADIUS, Starling, etc.)
//
// Returns:
//   - []IdentityProvider: A slice of all configured identity providers
//   - error: An error if the API request fails or response cannot be parsed
func GetIdentityProviders() ([]IdentityProvider, error) {
	var identityProviders []IdentityProvider

	query := "IdentityProviders"

	response, err := c.GetRequest(query)
	if err != nil {
		return []IdentityProvider{}, err
	}

	if err := json.Unmarshal(response, &identityProviders); err != nil {
		return []IdentityProvider{}, err
	}

	return identityProviders, nil
}

// GetIdentityProvider retrieves a specific identity provider by its ID.
//
// This function fetches detailed configuration information for a single identity
// provider, including all its type-specific properties and settings.
//
// Parameters:
//   - id: The unique identifier of the identity provider
//
// Returns:
//   - IdentityProvider: The requested identity provider's complete configuration
//   - error: An error if the provider cannot be found or the request fails
func GetIdentityProvider(id int) (IdentityProvider, error) {
	var identityProvider IdentityProvider

	query := fmt.Sprintf("IdentityProviders/%d", id)

	response, err := c.GetRequest(query)
	if err != nil {
		return IdentityProvider{}, err
	}

	if err := json.Unmarshal(response, &identityProvider); err != nil {
		return IdentityProvider{}, err
	}
	return identityProvider, nil
}

// GetDirectoryUsers retrieves users from a specific identity provider's directory.
//
// This function only works with identity providers that are directories (IsDirectory = true).
// It supports pagination and filtering through the filter parameter.
//
// Parameters:
//   - identityProviderId: The ID of the directory identity provider
//   - filter: Query parameters to filter the results (e.g., search text, limit, offset)
//
// Returns:
//   - []User: A slice of directory users matching the filter criteria
//   - error: An error if the directory cannot be queried or the request fails
func GetDirectoryUsers(identityProviderId int, filter Filter) ([]User, error) {
	var directoryUsers []User

	query := fmt.Sprintf("IdentityProviders/%d/DirectoryUsers%s", identityProviderId, filter.ToQueryString())

	response, err := c.GetRequest(query)
	if err != nil {
		return []User{}, err
	}

	if err := json.Unmarshal(response, &directoryUsers); err != nil {
		return []User{}, err
	}

	for i := range directoryUsers {
		directoryUsers[i].client = c
	}

	return directoryUsers, nil
}

// GetDirectoryUsers retrieves users from this identity provider's directory.
//
// This method is a convenience wrapper around the package-level GetDirectoryUsers
// function, automatically using this identity provider's ID.
//
// Parameters:
//   - filter: Query parameters to filter the results (e.g., search text, limit, offset)
//
// Returns:
//   - []User: A slice of directory users matching the filter criteria
//   - error: An error if the directory cannot be queried or the request fails
func (idp IdentityProvider) GetDirectoryUsers(filter Filter) ([]User, error) {
	return GetDirectoryUsers(idp.Id, filter)
}

// GetDirectoryGroups retrieves groups from a specific identity provider's directory.
//
// This function only works with identity providers that are directories (IsDirectory = true).
// It supports pagination and filtering through the filter parameter.
//
// Parameters:
//   - id: The ID of the directory identity provider
//   - filter: Query parameters to filter the results (e.g., search text, limit, offset)
//
// Returns:
//   - []UserGroup: A slice of directory groups matching the filter criteria
//   - error: An error if the directory cannot be queried or the request fails
func GetDirectoryGroups(id int, filter Filter) ([]UserGroup, error) {
	var directoryGroups []UserGroup

	query := fmt.Sprintf("IdentityProviders/%d/DirectoryGroups%s", id, filter.ToQueryString())

	response, err := c.GetRequest(query)
	if err != nil {
		return []UserGroup{}, err
	}

	if err := json.Unmarshal(response, &directoryGroups); err != nil {
		return []UserGroup{}, err
	}

	return directoryGroups, nil
}

// GetDirectoryGroups retrieves groups from this identity provider's directory.
//
// This method is a convenience wrapper around the package-level GetDirectoryGroups
// function, automatically using this identity provider's ID.
//
// Parameters:
//   - filter: Query parameters to filter the results (e.g., search text, limit, offset)
//
// Returns:
//   - []UserGroup: A slice of directory groups matching the filter criteria
//   - error: An error if the directory cannot be queried or the request fails
func (idp IdentityProvider) GetDirectoryGroups(filter Filter) ([]UserGroup, error) {
	return GetDirectoryGroups(idp.Id, filter)
}
