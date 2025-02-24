package models

import (
	"encoding/json"
	"time"

	"github.com/sthayduk/safeguard-go/src/client"
)

// UserGroup represents a group of users in Safeguard
type UserGroup struct {
	client *client.SafeguardClient

	Id                           int                          `json:"Id"`
	Name                         string                       `json:"Name"`
	Description                  string                       `json:"Description"`
	IdentityProvider             GroupIdentityProvider        `json:"IdentityProvider"`
	IsReadOnly                   bool                         `json:"IsReadOnly"`
	CreatedDate                  time.Time                    `json:"CreatedDate"`
	CreatedByUserId              int                          `json:"CreatedByUserId"`
	CreatedByUserDisplayName     string                       `json:"CreatedByUserDisplayName"`
	ModifiedDate                 time.Time                    `json:"ModifiedDate"`
	ModifiedByUserId             int                          `json:"ModifiedByUserId"`
	ModifiedByUserDisplayName    string                       `json:"ModifiedByUserDisplayName"`
	DirectoryProperties          DirectoryProperties          `json:"DirectoryProperties"`
	Members                      []UserGroupMember            `json:"Members"`
	DirectoryGroupSyncProperties DirectoryGroupSyncProperties `json:"DirectoryGroupSyncProperties"`
}

// GroupIdentityProvider represents identity provider information for a user group
type GroupIdentityProvider struct {
	Id                int    `json:"Id"`
	Name              string `json:"Name"`
	TypeReferenceName string `json:"TypeReferenceName"`
	IdentityId        string `json:"IdentityId"`
}

// DirectoryProperties represents directory-specific properties for a group or user
type UserGroupDirectoryProperties struct {
	DirectoryId       int    `json:"DirectoryId"`
	DirectoryName     string `json:"DirectoryName"`
	DomainName        string `json:"DomainName"`
	NetbiosName       string `json:"NetbiosName"`
	DistinguishedName string `json:"DistinguishedName"`
	ObjectGuid        string `json:"ObjectGuid"`
	ObjectSid         string `json:"ObjectSid"`
}

// UserGroupMember represents a user that is a member of a user group
type UserGroupMember struct {
	AdminRoles                                []string               `json:"AdminRoles"`
	Id                                        int                    `json:"Id"`
	Name                                      string                 `json:"Name"`
	Description                               string                 `json:"Description"`
	DisplayName                               string                 `json:"DisplayName"`
	LastName                                  string                 `json:"LastName"`
	FirstName                                 string                 `json:"FirstName"`
	EmailAddress                              string                 `json:"EmailAddress"`
	WorkPhone                                 string                 `json:"WorkPhone"`
	MobilePhone                               string                 `json:"MobilePhone"`
	PrimaryAuthenticationProvider             AuthenticationProvider `json:"PrimaryAuthenticationProvider"`
	SecondaryAuthenticationProvider           AuthenticationProvider `json:"SecondaryAuthenticationProvider"`
	IdentityProvider                          GroupIdentityProvider  `json:"IdentityProvider"`
	Disabled                                  bool                   `json:"Disabled"`
	TimeZoneId                                string                 `json:"TimeZoneId"`
	TimeZoneDisplayName                       string                 `json:"TimeZoneDisplayName"`
	TimeZoneIanaName                          string                 `json:"TimeZoneIanaName"`
	IsPartitionOwner                          bool                   `json:"IsPartitionOwner"`
	DirectoryProperties                       DirectoryProperties    `json:"DirectoryProperties"`
	CloudAssistantApproveEnabled              bool                   `json:"CloudAssistantApproveEnabled"`
	CloudAssistantRecipientId                 string                 `json:"CloudAssistantRecipientId"`
	AllowPersonalAccounts                     bool                   `json:"AllowPersonalAccounts"`
	Locked                                    bool                   `json:"Locked"`
	PasswordNeverExpires                      bool                   `json:"PasswordNeverExpires"`
	ChangePasswordAtNextLogin                 bool                   `json:"ChangePasswordAtNextLogin"`
	Base64PhotoData                           string                 `json:"Base64PhotoData"`
	IsSystemOwned                             bool                   `json:"IsSystemOwned"`
	IsRequester                               bool                   `json:"IsRequester"`
	IsApprover                                bool                   `json:"IsApprover"`
	IsReviewer                                bool                   `json:"IsReviewer"`
	LastLoginDate                             time.Time              `json:"LastLoginDate"`
	LastRequestDate                           time.Time              `json:"LastRequestDate"`
	CreatedDate                               time.Time              `json:"CreatedDate"`
	CreatedByUserId                           int                    `json:"CreatedByUserId"`
	CreatedByUserDisplayName                  string                 `json:"CreatedByUserDisplayName"`
	ModifiedDate                              time.Time              `json:"ModifiedDate"`
	ModifiedByUserId                          int                    `json:"ModifiedByUserId"`
	ModifiedByUserDisplayName                 string                 `json:"ModifiedByUserDisplayName"`
	RequireCertificateAuthentication          bool                   `json:"RequireCertificateAuthentication"`
	DirectoryRequireCertificateAuthentication bool                   `json:"DirectoryRequireCertificateAuthentication"`
	LinkedAccountsCount                       int                    `json:"LinkedAccountsCount"`
}

// DirectoryGroupSyncProperties represents synchronization properties for directory groups
type DirectoryGroupSyncProperties struct {
	PrimaryAuthenticationProviderId                  int      `json:"PrimaryAuthenticationProviderId"`
	PrimaryAuthenticationProviderTypeReferenceName   string   `json:"PrimaryAuthenticationProviderTypeReferenceName"`
	PrimaryAuthenticationProviderName                string   `json:"PrimaryAuthenticationProviderName"`
	RequireCertificateAuthentication                 bool     `json:"RequireCertificateAuthentication"`
	SecondaryAuthenticationProviderId                int      `json:"SecondaryAuthenticationProviderId"`
	SecondaryAuthenticationProviderTypeReferenceName string   `json:"SecondaryAuthenticationProviderTypeReferenceName"`
	SecondaryAuthenticationProviderName              string   `json:"SecondaryAuthenticationProviderName"`
	LinkDirectoryAccounts                            bool     `json:"LinkDirectoryAccounts"`
	AllowPersonalAccounts                            bool     `json:"AllowPersonalAccounts"`
	AdminRoles                                       []string `json:"AdminRoles"`
}

func (u UserGroup) ToJson() (string, error) {
	userGroupJSON, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(userGroupJSON), nil
}
