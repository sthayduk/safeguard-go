package models

import (
	"encoding/json"
	"time"
)

type User struct {
	Name                                      string                 `json:"Name,omitempty"`
	PrimaryAuthenticationProvider             AuthenticationProvider `json:"PrimaryAuthenticationProvider,omitempty"`
	Preferences                               []Preference           `json:"Preferences,omitempty"`
	Fido2Authenticators                       []Fido2Authenticator   `json:"Fido2Authenticators,omitempty"`
	AdminRoles                                []string               `json:"AdminRoles,omitempty"`
	Id                                        int                    `json:"Id,omitempty"`
	Description                               string                 `json:"Description,omitempty"`
	DisplayName                               string                 `json:"DisplayName,omitempty"`
	LastName                                  string                 `json:"LastName,omitempty"`
	FirstName                                 string                 `json:"FirstName,omitempty"`
	EmailAddress                              string                 `json:"EmailAddress,omitempty"`
	WorkPhone                                 string                 `json:"WorkPhone,omitempty"`
	MobilePhone                               string                 `json:"MobilePhone,omitempty"`
	SecondaryAuthenticationProvider           AuthenticationProvider `json:"SecondaryAuthenticationProvider,omitempty"`
	IdentityProvider                          IdentityProvider       `json:"IdentityProvider,omitempty"`
	Disabled                                  bool                   `json:"Disabled,omitempty"`
	TimeZoneId                                string                 `json:"TimeZoneId,omitempty"`
	TimeZoneDisplayName                       string                 `json:"TimeZoneDisplayName,omitempty"`
	TimeZoneIanaName                          string                 `json:"TimeZoneIanaName,omitempty"`
	IsPartitionOwner                          bool                   `json:"IsPartitionOwner,omitempty"`
	DirectoryProperties                       DirectoryProperties    `json:"DirectoryProperties,omitempty"`
	CloudAssistantApproveEnabled              bool                   `json:"CloudAssistantApproveEnabled,omitempty"`
	CloudAssistantRecipientId                 string                 `json:"CloudAssistantRecipientId,omitempty"`
	AllowPersonalAccounts                     bool                   `json:"AllowPersonalAccounts,omitempty"`
	Locked                                    bool                   `json:"Locked,omitempty"`
	PasswordNeverExpires                      bool                   `json:"PasswordNeverExpires,omitempty"`
	ChangePasswordAtNextLogin                 bool                   `json:"ChangePasswordAtNextLogin,omitempty"`
	Base64PhotoData                           string                 `json:"Base64PhotoData,omitempty"`
	IsSystemOwned                             bool                   `json:"IsSystemOwned,omitempty"`
	IsRequester                               bool                   `json:"IsRequester,omitempty"`
	IsApprover                                bool                   `json:"IsApprover,omitempty"`
	IsReviewer                                bool                   `json:"IsReviewer,omitempty"`
	LastLoginDate                             time.Time              `json:"LastLoginDate,omitempty"`
	LastRequestDate                           time.Time              `json:"LastRequestDate,omitempty"`
	CreatedDate                               time.Time              `json:"CreatedDate,omitempty"`
	CreatedByUserId                           int                    `json:"CreatedByUserId,omitempty"`
	CreatedByUserDisplayName                  string                 `json:"CreatedByUserDisplayName,omitempty"`
	ModifiedDate                              time.Time              `json:"ModifiedDate,omitempty"`
	ModifiedByUserId                          int                    `json:"ModifiedByUserId,omitempty"`
	ModifiedByUserDisplayName                 string                 `json:"ModifiedByUserDisplayName,omitempty"`
	RequireCertificateAuthentication          bool                   `json:"RequireCertificateAuthentication,omitempty"`
	DirectoryRequireCertificateAuthentication bool                   `json:"DirectoryRequireCertificateAuthentication,omitempty"`
	LinkedAccountsCount                       int                    `json:"LinkedAccountsCount,omitempty"`
}

func (u User) ToJson() (string, error) {
	userJSON, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(userJSON), nil
}
