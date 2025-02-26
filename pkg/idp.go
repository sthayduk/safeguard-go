package pkg

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sthayduk/safeguard-go/client"
)

// IdentityProvider represents the structure for the given JSON array
type IdentityProvider struct {
	client *client.SafeguardClient

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

// GetIdentityProviders retrieves a list of identity providers from the Safeguard API.
// It sends a GET request to the "IdentityProviders" endpoint and unmarshals the response
// into a slice of IdentityProvider structs. Each IdentityProvider is then associated
// with the provided SafeguardClient.
//
// Parameters:
//   - client: A pointer to a SafeguardClient instance used to make the API request.
//
// Returns:
//   - A slice of IdentityProvider structs.
//   - An error if the request fails or the response cannot be unmarshaled.
func GetIdentityProviders(c *client.SafeguardClient) ([]IdentityProvider, error) {
	var identityProviders []IdentityProvider

	query := "IdentityProviders"

	response, err := c.GetRequest(query)
	if err != nil {
		return []IdentityProvider{}, err
	}

	if err := json.Unmarshal(response, &identityProviders); err != nil {
		return []IdentityProvider{}, err
	}

	for i := range identityProviders {
		identityProviders[i].client = c
	}

	return identityProviders, nil
}

// GetIdentityProvider retrieves an IdentityProvider by its ID.
//
// Parameters:
//   - client: A pointer to the SafeguardClient used to make the request.
//   - id: The ID of the IdentityProvider to retrieve.
//
// Returns:
//   - IdentityProvider: The retrieved IdentityProvider object.
//   - error: An error object if an error occurred during the request, otherwise nil.
func GetIdentityProvider(c *client.SafeguardClient, id int) (IdentityProvider, error) {
	var identityProvider IdentityProvider
	identityProvider.client = c

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

// GetDirectoryUsers retrieves a list of directory users from the specified identity provider.
// It sends a GET request to the Safeguard API and unmarshals the response into a slice of User objects.
//
// Parameters:
//   - c: A pointer to the SafeguardClient used to send the request.
//   - id: The ID of the identity provider from which to retrieve directory users.
//   - filter: A Filter object used to apply query parameters to the request.
//
// Returns:
//   - A slice of User objects representing the directory users.
//   - An error if the request fails or the response cannot be unmarshaled.
func GetDirectoryUsers(c *client.SafeguardClient, identityProviderId int, filter client.Filter) ([]User, error) {
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

// GetDirectoryUsers retrieves a list of users from the directory associated with the IdentityProvider.
// It accepts a filter parameter to narrow down the search results.
//
// Parameters:
//   - filter: A client.Filter object to specify the criteria for filtering the users.
//
// Returns:
//   - []User: A slice of User objects that match the filter criteria.
//   - error: An error object if there is any issue during the retrieval process.
func (idp IdentityProvider) GetDirectoryUsers(filter client.Filter) ([]User, error) {
	return GetDirectoryUsers(idp.client, idp.Id, filter)
}

// GetDirectoryGroups retrieves the directory groups associated with a specific identity provider.
// It sends a GET request to the Safeguard API and unmarshals the response into a slice of UserGroup.
//
// Parameters:
//   - c: A pointer to a SafeguardClient instance used to send the request.
//   - id: The ID of the identity provider.
//   - filter: A Filter instance used to filter the results.
//
// Returns:
//   - A slice of UserGroup containing the directory groups.
//   - An error if the request fails or the response cannot be unmarshaled.
func GetDirectoryGroups(c *client.SafeguardClient, id int, filter client.Filter) ([]UserGroup, error) {
	var directoryGroups []UserGroup

	query := fmt.Sprintf("IdentityProviders/%d/DirectoryGroups%s", id, filter.ToQueryString())

	response, err := c.GetRequest(query)
	if err != nil {
		return []UserGroup{}, err
	}

	if err := json.Unmarshal(response, &directoryGroups); err != nil {
		return []UserGroup{}, err
	}

	for i := range directoryGroups {
		directoryGroups[i].client = c
	}

	return directoryGroups, nil
}

// GetDirectoryGroups retrieves the directory groups associated with the IdentityProvider.
// It takes a filter parameter to apply specific filtering criteria and returns a slice of UserGroup
// and an error if any occurs during the retrieval process.
//
// Parameters:
//   - filter: A client.Filter object to specify the filtering criteria.
//
// Returns:
//   - []UserGroup: A slice of UserGroup objects that match the filtering criteria.
//   - error: An error object if any issues occur during the retrieval process.
func (idp IdentityProvider) GetDirectoryGroups(filter client.Filter) ([]UserGroup, error) {
	return GetDirectoryGroups(idp.client, idp.Id, filter)
}
