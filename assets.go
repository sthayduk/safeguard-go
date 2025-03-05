package safeguard

import (
	"encoding/json"
	"fmt"
	"time"
)

// Asset represents a Safeguard asset
type Asset struct {
	apiClient *SafeguardClient

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

func (a Asset) SetClient(c *SafeguardClient) any {
	a.apiClient = c
	return a
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

// ToJson serializes the Asset instance into a JSON string representation.
//
// Returns:
//   - (string): JSON representation of the asset
//   - (error): An error if JSON marshaling fails
func (a Asset) ToJson() (string, error) {
	assetJSON, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(assetJSON), nil
}

// GetAssets retrieves all assets matching the specified filter criteria.
//
// Parameters:
//   - filter: Query parameters to filter the results
//
// Returns:
//   - ([]Asset): Slice of matching assets
//   - (error): An error if the API request fails
func (c *SafeguardClient) GetAssets(fields Filter) ([]Asset, error) {
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

	return addClientToSlice(c, assets), nil
}

// GetAsset retrieves a single asset by its ID.
//
// Parameters:
//   - id: Unique identifier of the asset
//   - fields: Optional fields to include in the response
//
// Returns:
//   - (Asset): The requested asset
//   - (error): An error if the API request fails
func (c *SafeguardClient) GetAsset(id int, fields Fields) (Asset, error) {
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

	return addClient(c, asset), nil
}

// GetAssetDirectoryAccounts retrieves all directory accounts associated with the specified asset.
//
// Parameters:
//   - assetId: Unique identifier of the asset
//   - filter: Query parameters to filter the results
//
// Returns:
//   - ([]AssetAccount): Slice of matching directory accounts
//   - (error): An error if the API request fails
func (c *SafeguardClient) GetAssetDirectoryAccounts(assetId int, filter Filter) ([]AssetAccount, error) {
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

	return addClientToSlice(c, accounts), nil
}

// GetDirectoryAccounts retrieves all directory accounts associated with this asset.
//
// Parameters:
//   - filter: Query parameters to filter the results
//
// Returns:
//   - ([]AssetAccount): Slice of matching directory accounts
//   - (error): An error if the API request fails
func (a Asset) GetDirectoryAccounts(filter Filter) ([]AssetAccount, error) {
	return a.apiClient.GetAssetDirectoryAccounts(a.Id, filter)
}

// GetAssetDirectoryAssets retrieves all directory assets associated with the specified asset.
//
// Parameters:
//   - assetId: Unique identifier of the asset
//   - filter: Query parameters to filter the results
//
// Returns:
//   - ([]Asset): Slice of matching directory assets
//   - (error): An error if the API request fails
func (c *SafeguardClient) GetAssetDirectoryAssets(assetId int, filter Filter) ([]Asset, error) {
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

	return addClientToSlice(c, assets), nil
}

// GetDirectoryAssets retrieves all directory assets associated with this asset.
//
// Parameters:
//   - filter: Query parameters to filter the results
//
// Returns:
//   - ([]Asset): Slice of matching directory assets
//   - (error): An error if the API request fails
func (a Asset) GetDirectoryAssets(filter Filter) ([]Asset, error) {
	return a.apiClient.GetAssetDirectoryAssets(a.Id, filter)
}

// GetAssetDirectoryServiceEntries retrieves all directory service entries associated with the specified asset.
//
// Parameters:
//   - assetId: Unique identifier of the asset
//   - filter: Query parameters to filter the results
//
// Returns:
//   - ([]DirectoryServiceEntry): Slice of matching directory service entries
//   - (error): An error if the API request fails
func (c *SafeguardClient) GetAssetDirectoryServiceEntries(assetId int, filter Filter) ([]DirectoryServiceEntry, error) {
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

	return addClientToSlice(c, entries), nil
}

// GetDirectoryServiceEntries retrieves all directory service entries associated with this asset.
//
// Parameters:
//   - filter: Query parameters to filter the results
//
// Returns:
//   - ([]DirectoryServiceEntry): Slice of matching directory service entries
//   - (error): An error if the API request fails
func (a Asset) GetDirectoryServiceEntries(filter Filter) ([]DirectoryServiceEntry, error) {
	return a.apiClient.GetAssetDirectoryServiceEntries(a.Id, filter)
}
