package pkg

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sthayduk/safeguard-go/client"
)

// Asset represents a Safeguard asset
type Asset struct {
	Id                           int                          `json:"Id,omitempty"`
	Name                         string                       `json:"Name,omitempty"`
	NetworkAddress               string                       `json:"NetworkAddress,omitempty"`
	Description                  string                       `json:"Description,omitempty"`
	PlatformId                   int                          `json:"PlatformId,omitempty"`
	PlatformDisplayName          string                       `json:"PlatformDisplayName,omitempty"`
	AssetPartitionId             int                          `json:"AssetPartitionId,omitempty"`
	AssetPartitionName           string                       `json:"AssetPartitionName,omitempty"`
	LicenseClass                 string                       `json:"LicenseClass,omitempty"`
	IsDirectory                  bool                         `json:"IsDirectory,omitempty"`
	ManagedNetworkId             int                          `json:"ManagedNetworkId,omitempty"`
	ManagedNetworkName           string                       `json:"ManagedNetworkName,omitempty"`
	CreatedDate                  time.Time                    `json:"CreatedDate,omitempty"`
	CreatedByUserId              int                          `json:"CreatedByUserId,omitempty"`
	CreatedByUserDisplayName     string                       `json:"CreatedByUserDisplayName,omitempty"`
	Platform                     Platform                     `json:"Platform,omitempty"`
	Tags                         []Tag                        `json:"Tags,omitempty"`
	ManagedBy                    []ManagedByUser              `json:"ManagedBy,omitempty"`
	DiscoveredGroups             []DiscoveredGroup            `json:"DiscoveredGroups,omitempty"`
	TaskProperties               AssetTaskProperties          `json:"TaskProperties,omitempty"`
	ConnectionProperties         AssetConnectionProperties    `json:"ConnectionProperties,omitempty"`
	SessionAccessProperties      AssetSessionAccessProperties `json:"SessionAccessProperties,omitempty"`
	SshHostKey                   AssetSshHostKey              `json:"SshHostKey,omitempty"`
	Disabled                     bool                         `json:"Disabled,omitempty"`
	AssetType                    AssetType                    `json:"AssetType,omitempty"`
	DirectoryProperties          DirectoryProperties          `json:"DirectoryProperties,omitempty"`
	DirectoryAssetProperties     DirectoryAssetProperties     `json:"DirectoryAssetProperties,omitempty"`
	StarlingAssetProperties      StarlingAssetProperties      `json:"StarlingAssetProperties,omitempty"`
	AssetDiscoveryJobId          int                          `json:"AssetDiscoveryJobId,omitempty"`
	AssetDiscoveryJobName        string                       `json:"AssetDiscoveryJobName,omitempty"`
	AccountDiscoveryScheduleId   int                          `json:"AccountDiscoveryScheduleId,omitempty"`
	AccountDiscoveryScheduleName string                       `json:"AccountDiscoveryScheduleName,omitempty"`
	DependentSystemIds           []int                        `json:"DependentSystemIds,omitempty"`
	CustomScriptParameters       []CustomScriptParameter      `json:"CustomScriptParameters,omitempty"`
	PasswordProfile              Profile                      `json:"PasswordProfile,omitempty"`
	SshKeyProfile                Profile                      `json:"SshKeyProfile,omitempty"`
	RegisteredConnector          RegisteredConnector          `json:"RegisteredConnector,omitempty"`
}

// StarlingAssetProperties represents properties specific to Starling assets
type StarlingAssetProperties struct {
	UniqueId           string `json:"UniqueId,omitempty"`
	HostName           string `json:"HostName,omitempty"`
	NetworkAddress     string `json:"NetworkAddress,omitempty"`
	ServiceAccountName string `json:"ServiceAccountName,omitempty"`
}

// AssetTaskProperties represents task properties and history for an asset
type AssetTaskProperties struct {
	HasAssetTaskFailure                   bool      `json:"HasAssetTaskFailure,omitempty"`
	LastAccountDiscoveryDate              time.Time `json:"LastAccountDiscoveryDate,omitempty"`
	LastSuccessAccountDiscoveryDate       time.Time `json:"LastSuccessAccountDiscoveryDate,omitempty"`
	LastFailureAccountDiscoveryDate       time.Time `json:"LastFailureAccountDiscoveryDate,omitempty"`
	FailedAccountDiscoveryAttempts        int       `json:"FailedAccountDiscoveryAttempts,omitempty"`
	NextAccountDiscoveryDate              time.Time `json:"NextAccountDiscoveryDate,omitempty"`
	LastAccountDiscoveryTaskId            string    `json:"LastAccountDiscoveryTaskId,omitempty"`
	LastServiceDiscoveryDate              time.Time `json:"LastServiceDiscoveryDate,omitempty"`
	LastSuccessServiceDiscoveryDate       time.Time `json:"LastSuccessServiceDiscoveryDate,omitempty"`
	LastFailureServiceDiscoveryDate       time.Time `json:"LastFailureServiceDiscoveryDate,omitempty"`
	FailedServiceDiscoveryAttempts        int       `json:"FailedServiceDiscoveryAttempts,omitempty"`
	NextServiceDiscoveryDate              time.Time `json:"NextServiceDiscoveryDate,omitempty"`
	LastServiceDiscoveryTaskId            string    `json:"LastServiceDiscoveryTaskId,omitempty"`
	LastTestConnectionDate                time.Time `json:"LastTestConnectionDate,omitempty"`
	LastSuccessTestConnectionDate         time.Time `json:"LastSuccessTestConnectionDate,omitempty"`
	LastFailureTestConnectionDate         time.Time `json:"LastFailureTestConnectionDate,omitempty"`
	FailedTestConnectionAttempts          int       `json:"FailedTestConnectionAttempts,omitempty"`
	NextTestConnectionDate                time.Time `json:"NextTestConnectionDate,omitempty"`
	LastTestConnectionTaskId              string    `json:"LastTestConnectionTaskId,omitempty"`
	LastDependentServiceUpdateDate        time.Time `json:"LastDependentServiceUpdateDate,omitempty"`
	LastSuccessDependentServiceUpdateDate time.Time `json:"LastSuccessDependentServiceUpdateDate,omitempty"`
	LastFailureDependentServiceUpdateDate time.Time `json:"LastFailureDependentServiceUpdateDate,omitempty"`
	FailedDependentServiceUpdateAttempts  int       `json:"FailedDependentServiceUpdateAttempts,omitempty"`
	NextDependentServiceUpdateDate        time.Time `json:"NextDependentServiceUpdateDate,omitempty"`
	LastDependentServiceUpdateTaskId      string    `json:"LastDependentServiceUpdateTaskId,omitempty"`
}

// AssetConnectionProperties represents connection settings for an asset
type AssetConnectionProperties struct {
	ServiceAccountId                         int                          `json:"ServiceAccountId,omitempty"`
	ServiceAccountName                       string                       `json:"ServiceAccountName,omitempty"`
	EffectiveServiceAccountName              string                       `json:"EffectiveServiceAccountName,omitempty"`
	ServiceAccountDomainName                 string                       `json:"ServiceAccountDomainName,omitempty"`
	ServiceAccountDistinguishedName          string                       `json:"ServiceAccountDistinguishedName,omitempty"`
	EffectiveServiceAccountDistinguishedName string                       `json:"EffectiveServiceAccountDistinguishedName,omitempty"`
	ServiceAccountCredentialType             ServiceAccountCredentialType `json:"ServiceAccountCredentialType,omitempty"`
	ServiceAccountPassword                   string                       `json:"ServiceAccountPassword,omitempty"`
	ServiceAccountHasPassword                bool                         `json:"ServiceAccountHasPassword,omitempty"`
	ServiceAccountSshKey                     SshKey                       `json:"ServiceAccountSshKey,omitempty"`
	ServiceAccountHasSshKey                  bool                         `json:"ServiceAccountHasSshKey,omitempty"`
	ServiceAccountApiKey                     string                       `json:"ServiceAccountApiKey,omitempty"`
	ServiceAccountHasApiKey                  bool                         `json:"ServiceAccountHasApiKey,omitempty"`
	Port                                     int                          `json:"Port,omitempty"`
	ServiceAccountAssetId                    int                          `json:"ServiceAccountAssetId,omitempty"`
	ServiceAccountAssetName                  string                       `json:"ServiceAccountAssetName,omitempty"`
	ServiceAccountAssetPlatformId            int                          `json:"ServiceAccountAssetPlatformId,omitempty"`
	ServiceAccountAssetPlatformType          string                       `json:"ServiceAccountAssetPlatformType,omitempty"`
	ServiceAccountAssetPlatformDisplayName   string                       `json:"ServiceAccountAssetPlatformDisplayName,omitempty"`
	ServiceAccountNetbiosName                string                       `json:"ServiceAccountNetbiosName,omitempty"`
	ServiceAccountUniqueObjectId             string                       `json:"ServiceAccountUniqueObjectId,omitempty"`
	ServiceAccountSecurityId                 string                       `json:"ServiceAccountSecurityId,omitempty"`
	ServiceAccountProfileId                  int                          `json:"ServiceAccountProfileId,omitempty"`
	ServiceAccountProfileName                string                       `json:"ServiceAccountProfileName,omitempty"`
	ServiceAccountSshKeyProfileId            int                          `json:"ServiceAccountSshKeyProfileId,omitempty"`
	ServiceAccountSshKeyProfileName          string                       `json:"ServiceAccountSshKeyProfileName,omitempty"`
	EnablePassword                           string                       `json:"EnablePassword,omitempty"`
	EnableHasPassword                        bool                         `json:"EnableHasPassword,omitempty"`
	CommandTimeout                           int                          `json:"CommandTimeout,omitempty"`
	WorkstationId                            string                       `json:"WorkstationId,omitempty"`
	ClientId                                 int                          `json:"ClientId,omitempty"`
	UseSslEncryption                         bool                         `json:"UseSslEncryption,omitempty"`
	VerifySslCertificate                     bool                         `json:"VerifySslCertificate,omitempty"`
	Instance                                 string                       `json:"Instance,omitempty"`
	ServiceName                              string                       `json:"ServiceName,omitempty"`
	SslThumbprint                            string                       `json:"SslThumbprint,omitempty"`
	PrivilegeElevationCommand                string                       `json:"PrivilegeElevationCommand,omitempty"`
	AccessKeyId                              string                       `json:"AccessKeyId,omitempty"`
	SecretKey                                string                       `json:"SecretKey,omitempty"`
	HasSecretKey                             bool                         `json:"HasSecretKey,omitempty"`
	OraclePrivileges                         string                       `json:"OraclePrivileges,omitempty"`
	HideAlterUserCommand                     bool                         `json:"HideAlterUserCommand,omitempty"`
	UseServiceAccountUserNameOnly            bool                         `json:"UseServiceAccountUserNameOnly,omitempty"`
	UseNamedPipeForServiceAccountConnection  bool                         `json:"UseNamedPipeForServiceAccountConnection,omitempty"`
	RegisteredConnectorId                    int                          `json:"RegisteredConnectorId,omitempty"`
	TacacsSecret                             string                       `json:"TacacsSecret,omitempty"`
	HasTacacsSecret                          bool                         `json:"HasTacacsSecret,omitempty"`
	UseTopSecretInterval                     bool                         `json:"UseTopSecretInterval,omitempty"`
	UseHttpProxy                             bool                         `json:"UseHttpProxy,omitempty"`
	AlternateConnectionProperties            map[string]string            `json:"AlternateConnectionProperties,omitempty"`
}

// AssetSessionAccessProperties represents session access configuration for an asset
type AssetSessionAccessProperties struct {
	AllowSessionRequests     bool                   `json:"AllowSessionRequests,omitempty"`
	SshSessionPort           int                    `json:"SshSessionPort,omitempty"`
	RemoteDesktopSessionPort int                    `json:"RemoteDesktopSessionPort,omitempty"`
	TelnetSessionPort        int                    `json:"TelnetSessionPort,omitempty"`
	ProtocolId               int                    `json:"ProtocolId,omitempty"`
	ProtocolName             string                 `json:"ProtocolName,omitempty"`
	ApplicationProperties    map[string]interface{} `json:"ApplicationProperties,omitempty"`
}

// RegisteredConnector represents a Starling connector registration
type RegisteredConnector struct {
	Id                             int              `json:"Id,omitempty"`
	RegisteredConnectorId          string           `json:"RegisteredConnectorId,omitempty"`
	RegisteredConnectorDisplayName string           `json:"RegisteredConnectorDisplayName,omitempty"`
	DisplayName                    string           `json:"DisplayName,omitempty"`
	StarlingConnectorId            string           `json:"StarlingConnectorId,omitempty"`
	StarlingConnectorVersion       string           `json:"StarlingConnectorVersion,omitempty"`
	Platform                       Platform         `json:"Platform,omitempty"`
	VisibleToAllPartitions         bool             `json:"VisibleToAllPartitions,omitempty"`
	VisibleToPartitions            []AssetPartition `json:"VisibleToPartitions,omitempty"`
}

// DirectoryAssetProperties represents extended properties specific to directory assets
type DirectoryAssetProperties struct {
	DirectoryConnectionProperties  DirectoryConnectionProperties `json:"DirectoryConnectionProperties,omitempty"`
	DirectoryObjectProperties      DirectoryObjectProperties     `json:"DirectoryObjectProperties,omitempty"`
	ForestRootDomain               string                        `json:"ForestRootDomain,omitempty"`
	DomainName                     string                        `json:"DomainName,omitempty"`
	AllowSharedSearch              bool                          `json:"AllowSharedSearch,omitempty"`
	UsePasswordHash                bool                          `json:"UsePasswordHash,omitempty"`
	SynchronizationIntervalMinutes int                           `json:"SynchronizationIntervalMinutes,omitempty"`
	DeleteSyncIntervalMinutes      int                           `json:"DeleteSyncIntervalMinutes,omitempty"`
	Domains                        []Domain                      `json:"Domains,omitempty"`
	DomainControllers              []DirectoryDomainController   `json:"DomainControllers,omitempty"`
	LastSynchronizedDate           time.Time                     `json:"LastSynchronizedDate,omitempty"`
	LastSuccessSynchronizedDate    time.Time                     `json:"LastSuccessSynchronizedDate,omitempty"`
	LastFailureSynchronizedDate    time.Time                     `json:"LastFailureSynchronizedDate,omitempty"`
	FailedSyncAttempts             int                           `json:"FailedSyncAttempts,omitempty"`
	LastDirectorySyncTaskId        string                        `json:"LastDirectorySyncTaskId,omitempty"`
	NextSynchronizedDate           time.Time                     `json:"NextSynchronizedDate,omitempty"`
	LastDeleteSyncDate             time.Time                     `json:"LastDeleteSyncDate,omitempty"`
	LastSuccessDeleteSyncDate      time.Time                     `json:"LastSuccessDeleteSyncDate,omitempty"`
	LastFailureDeleteSyncDate      time.Time                     `json:"LastFailureDeleteSyncDate,omitempty"`
	FailedDeleteSyncAttempts       int                           `json:"FailedDeleteSyncAttempts,omitempty"`
	LastDirectoryDeleteSyncTaskId  string                        `json:"LastDirectoryDeleteSyncTaskId,omitempty"`
	NextDeleteSyncDate             time.Time                     `json:"NextDeleteSyncDate,omitempty"`
	SchemaProperties               SchemaProperties              `json:"SchemaProperties,omitempty"`
}

// DirectoryDomainController represents a directory domain controller
type DirectoryDomainController struct {
	NetworkAddress string `json:"NetworkAddress,omitempty"`
	DomainName     string `json:"DomainName,omitempty"`
	IsWritable     bool   `json:"IsWritable,omitempty"`
	ServerType     string `json:"ServerType,omitempty"`
}

// SchemaProperties represents directory schema mappings
type SchemaProperties struct {
	UserProperties     UserSchemaProperties     `json:"UserProperties,omitempty"`
	GroupProperties    GroupSchemaProperties    `json:"GroupProperties,omitempty"`
	ComputerProperties ComputerSchemaProperties `json:"ComputerProperties,omitempty"`
}

// UserSchemaProperties represents directory attribute mappings for users
type UserSchemaProperties struct {
	UserClassType         []string `json:"UserClassType,omitempty"`
	UserNameAttribute     string   `json:"UserNameAttribute,omitempty"`
	PasswordAttribute     string   `json:"PasswordAttribute,omitempty"`
	DescriptionAttribute  string   `json:"DescriptionAttribute,omitempty"`
	MemberOfAttribute     string   `json:"MemberOfAttribute,omitempty"`
	AltLoginNameAttribute string   `json:"AltLoginNameAttribute,omitempty"`
}

// GroupSchemaProperties represents directory attribute mappings for groups
type GroupSchemaProperties struct {
	GroupClassType  []string `json:"GroupClassType,omitempty"`
	MemberAttribute string   `json:"MemberAttribute,omitempty"`
	NameAttribute   string   `json:"NameAttribute,omitempty"`
}

// ComputerSchemaProperties represents directory attribute mappings for computers
type ComputerSchemaProperties struct {
	ComputerClassType               []string `json:"ComputerClassType,omitempty"`
	NameAttribute                   string   `json:"NameAttribute,omitempty"`
	DescriptionAttribute            string   `json:"DescriptionAttribute,omitempty"`
	NetworkAddressAttribute         string   `json:"NetworkAddressAttribute,omitempty"`
	OperatingSystemAttribute        string   `json:"OperatingSystemAttribute,omitempty"`
	OperatingSystemVersionAttribute string   `json:"OperatingSystemVersionAttribute,omitempty"`
	MemberOfAttribute               string   `json:"MemberOfAttribute,omitempty"`
}

// AssetType represents the type of asset
type AssetType string

const (
	AssetTypeComputer      AssetType = "Computer"
	AssetTypeDirectory     AssetType = "Directory"
	AssetTypeDynamicAccess AssetType = "DynamicAccess"
	AssetTypeStarling      AssetType = "Starling"
	AssetTypeServer        AssetType = "Server"
	AssetTypeOther         AssetType = "Other"
)

// ServiceAccountCredentialType represents the type of credential used for the service account
type ServiceAccountCredentialType string

const (
	ServiceAccountCredentialTypePassword ServiceAccountCredentialType = "Password"
	ServiceAccountCredentialTypeSshKey   ServiceAccountCredentialType = "SshKey"
	ServiceAccountCredentialTypeNone     ServiceAccountCredentialType = "None"
)

// DirectoryConnectionProperties represents connection properties for a directory
type DirectoryConnectionProperties struct {
	ServerType        DirectoryServerType `json:"ServerType,omitempty"`
	UseSSL            bool                `json:"UseSSL,omitempty"`
	SslCertificateId  int                 `json:"SslCertificateId,omitempty"`
	DomainControllers []DomainController  `json:"DomainControllers,omitempty"`
	Domain            string              `json:"Domain,omitempty"`
	Port              int                 `json:"Port,omitempty"`
}

// DirectoryServerType represents the type of directory server
type DirectoryServerType string

const (
	DirectoryServerTypeActiveDirectory DirectoryServerType = "ActiveDirectory"
	DirectoryServerTypeLdap            DirectoryServerType = "Ldap"
	DirectoryServerTypeOther           DirectoryServerType = "Other"
)

// DirectoryObjectProperties represents object properties for a directory
type DirectoryObjectProperties struct {
	Container         string `json:"Container,omitempty"`
	ForestName        string `json:"ForestName,omitempty"`
	NetBiosName       string `json:"NetBiosName,omitempty"`
	ObjectClass       string `json:"ObjectClass,omitempty"`
	ObjectClassGroups string `json:"ObjectClassGroups,omitempty"`
	ObjectClassUsers  string `json:"ObjectClassUsers,omitempty"`
}

// ToJson converts an Asset to its JSON string representation
func (a Asset) ToJson() (string, error) {
	assetJSON, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(assetJSON), nil
}

// GetAssets retrieves a list of assets from Safeguard
// GetAssets retrieves a list of assets from the Safeguard API based on the provided filter fields.
// It sends a GET request to the "Assets" endpoint with the query parameters specified in the fields.
//
// Parameters:
//   - c: A pointer to a SafeguardClient instance used to make the API request.
//   - fields: A Filter object containing the query parameters for filtering the assets.
//
// Returns:
//   - A slice of Asset objects retrieved from the API.
//   - An error if the request fails or if there is an issue unmarshalling the response.
//
// Each Asset object in the returned slice will have its client field set to the provided SafeguardClient instance.
func GetAssets(fields client.Filter) ([]Asset, error) {
	var assets []Asset

	query := "Assets" + fields.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &assets)
	if err != nil {
		return nil, err
	}

	return assets, nil
}

// GetAsset retrieves a specific asset by ID from Safeguard
// GetAsset retrieves an asset by its ID from the Safeguard API.
// It takes a SafeguardClient, an asset ID, and optional fields to include in the query.
// It returns the Asset and an error, if any.
//
// Parameters:
//   - c: A pointer to a SafeguardClient instance used to make the API request.
//   - id: An integer representing the ID of the asset to retrieve.
//   - fields: A Fields object specifying additional fields to include in the query.
//
// Returns:
//   - Asset: The retrieved Asset object.
//   - error: An error object if the request or unmarshalling fails.
func GetAsset(id int, fields client.Fields) (Asset, error) {
	var asset Asset

	query := fmt.Sprintf("Assets/%d", id)
	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return asset, err
	}

	err = json.Unmarshal(response, &asset)
	if err != nil {
		return asset, err
	}

	return asset, nil
}

// GetAssetDirectoryAccounts retrieves the directory accounts associated with a specific asset.
// It sends a GET request to the Safeguard API to fetch the accounts.
//
// Parameters:
//   - c: A pointer to a SafeguardClient instance used to make the API request.
//   - assetId: An integer representing the ID of the asset.
//   - filter: A Filter object used to apply query parameters to the request.
//
// Returns:
//   - A slice of AssetAccount objects representing the directory accounts.
//   - An error if the request fails or the response cannot be unmarshaled.
func GetAssetDirectoryAccounts(assetId int, filter client.Filter) ([]AssetAccount, error) {
	var accounts []AssetAccount

	query := fmt.Sprintf("Assets/%d/DirectoryAccounts%s", assetId, filter.ToQueryString())

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// GetDirectoryAccounts retrieves the directory accounts associated with the asset.
// It accepts a filter parameter to narrow down the results based on specific criteria.
// Returns a slice of AssetAccount and an error if the operation fails.
//
// Parameters:
//   - filter: A client.Filter object to apply filtering criteria.
//
// Returns:
//   - []AssetAccount: A slice of AssetAccount objects that match the filter criteria.
//   - error: An error object if there is an issue retrieving the directory accounts.
func (a Asset) GetDirectoryAccounts(filter client.Filter) ([]AssetAccount, error) {
	return GetAssetDirectoryAccounts(a.Id, filter)
}

// GetAssetDirectoryAssets retrieves a list of directory assets associated with a given asset ID.
// It sends a GET request to the Safeguard API and parses the response into a slice of Asset objects.
//
// Parameters:
//   - c: A pointer to a SafeguardClient instance used to make the API request.
//   - assetId: An integer representing the ID of the asset whose directory assets are to be retrieved.
//   - filter: A Filter object containing fields to filter the directory assets.
//
// Returns:
//   - A slice of Asset objects representing the directory assets.
//   - An error if the request fails or the response cannot be unmarshaled.
func GetAssetDirectoryAssets(assetId int, filter client.Filter) ([]Asset, error) {
	var assets []Asset

	query := fmt.Sprintf("Assets/%d/DirectoryAssets", assetId)
	if len(filter.Fields) > 0 {
		query += filter.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &assets)
	if err != nil {
		return nil, err
	}

	return assets, nil
}

// GetDirectoryAssets retrieves a list of directory assets based on the provided filter.
// It returns a slice of Asset and an error if any occurs during the retrieval process.
//
// Parameters:
//   - filter: A client.Filter object that specifies the criteria for filtering the assets.
//
// Returns:
//   - []Asset: A slice of Asset objects that match the filter criteria.
//   - error: An error object if an error occurs, otherwise nil.
func (a Asset) GetDirectoryAssets(filter client.Filter) ([]Asset, error) {
	return GetAssetDirectoryAssets(a.Id, filter)
}

// GetAssetDirectoryServiceEntries retrieves directory service entries for a specific asset
// It returns a slice of DirectoryServiceEntry objects and any error encountered
//
// Parameters:
//   - c: A pointer to a SafeguardClient instance used to make the API request.
//   - assetId: The ID of the asset to retrieve directory service entries for
//   - filter: A Filter object containing the query parameters for filtering the entries.
//
// Returns:
//   - A slice of DirectoryServiceEntry objects retrieved from the API.
//   - An error if the request fails or if there is an issue unmarshalling the response.
func GetAssetDirectoryServiceEntries(assetId int, filter client.Filter) ([]DirectoryServiceEntry, error) {
	var entries []DirectoryServiceEntry

	query := fmt.Sprintf("Assets/%d/DirectoryServiceEntries", assetId)
	if len(filter.Fields) > 0 {
		query += filter.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

// GetDirectoryServiceEntries retrieves directory service entries for this asset
// It returns a slice of DirectoryServiceEntry objects and any error encountered
//
// Parameters:
//   - filter: A Filter object containing the query parameters for filtering the entries.
//
// Returns:
//   - A slice of DirectoryServiceEntry objects retrieved from the API.
//   - An error if the request fails or if there is an issue unmarshalling the response.
func (a Asset) GetDirectoryServiceEntries(filter client.Filter) ([]DirectoryServiceEntry, error) {
	return GetAssetDirectoryServiceEntries(a.Id, filter)
}
