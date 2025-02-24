package models

import (
	"encoding/json"
	"time"
)

// Role represents roles in Safeguard made up of members, security scopes, and permissions
type Role struct {
	Id                          int                         `json:"Id"`
	Name                        string                      `json:"Name"`
	Priority                    int                         `json:"Priority"`
	Description                 string                      `json:"Description"`
	ExpirationDate              time.Time                   `json:"ExpirationDate"`
	IsExpired                   bool                        `json:"IsExpired"`
	HasExpiredPolicies          bool                        `json:"HasExpiredPolicies"`
	HasInvalidPolicies          bool                        `json:"HasInvalidPolicies"`
	CreatedDate                 time.Time                   `json:"CreatedDate"`
	CreatedByUserId             int                         `json:"CreatedByUserId"`
	CreatedByUserDisplayName    string                      `json:"CreatedByUserDisplayName"`
	UserCount                   int                         `json:"UserCount"`
	AccountCount                int                         `json:"AccountCount"`
	AssetCount                  int                         `json:"AssetCount"`
	PolicyCount                 int                         `json:"PolicyCount"`
	HourlyRestrictionProperties HourlyRestrictionProperties `json:"HourlyRestrictionProperties"`
	Members                     []RoleMember                `json:"Members"`
}

// RoleMember represents a member of a role
type RoleMember struct {
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

// HourlyRestrictionProperties represents settings controlling when the policy/role will be effective
type HourlyRestrictionProperties struct {
	EnableHourlyRestrictions bool  `json:"EnableHourlyRestrictions"`
	MondayValidHours         []int `json:"MondayValidHours"`
	TuesdayValidHours        []int `json:"TuesdayValidHours"`
	WednesdayValidHours      []int `json:"WednesdayValidHours"`
	ThursdayValidHours       []int `json:"ThursdayValidHours"`
	FridayValidHours         []int `json:"FridayValidHours"`
	SaturdayValidHours       []int `json:"SaturdayValidHours"`
	SundayValidHours         []int `json:"SundayValidHours"`
}

func (u Role) ToJson() (string, error) {
	userJSON, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(userJSON), nil
}
