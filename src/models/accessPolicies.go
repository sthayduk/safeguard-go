package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sthayduk/safeguard-go/src/client"
)

// AccessPolicy represents security configuration governing the access to assets and accounts
type AccessPolicy struct {
	client *client.SafeguardClient

	Id                        int                             `json:"Id"`
	Name                      string                          `json:"Name"`
	Priority                  int                             `json:"Priority"`
	RoleId                    int                             `json:"RoleId"`
	RoleName                  string                          `json:"RoleName"`
	AccessRequestProperties   AccessRequestProperties         `json:"AccessRequestProperties"`
	ApproverProperties        PolicyApproverProperties        `json:"ApproverProperties"`
	ReviewerProperties        PolicyReviewerProperties        `json:"ReviewerProperties"`
	RequesterProperties       PolicyRequesterProperties       `json:"RequesterProperties"`
	EmergencyAccessProperties PolicyEmergencyAccessProperties `json:"EmergencyAccessProperties"`
	ExpirationDate            time.Time                       `json:"ExpirationDate"`
	IsExpired                 bool                            `json:"IsExpired"`
	IsValid                   bool                            `json:"IsValid"`
	CreatedDate               time.Time                       `json:"CreatedDate"`
	CreatedByUserId           int                             `json:"CreatedByUserId"`
	CreatedByUserDisplayName  string                          `json:"CreatedByUserDisplayName"`
	ScopeItems                []PolicyScopeItem               `json:"ScopeItems"`
}

func (a AccessPolicy) ToJson() (string, error) {
	accessPolicyJSON, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(accessPolicyJSON), nil
}

// AccessRequestProperties represents configuration governing access requests
type AccessRequestProperties struct {
	AllowSimultaneousAccess    bool     `json:"AllowSimultaneousAccess"`
	ChangePasswordAfterCheckin bool     `json:"ChangePasswordAfterCheckin"`
	MaximumDuration            int      `json:"MaximumDuration"`
	AllowedRequestTypes        []string `json:"AllowedRequestTypes"`
}

// PolicyApproverProperties represents settings related to approving an access request
type PolicyApproverProperties struct {
	RequireApproval        bool `json:"RequireApproval"`
	RequireReapproval      bool `json:"RequireReapproval"`
	AutoApproveRequests    bool `json:"AutoApproveRequests"`
	AllowSelfApproval      bool `json:"AllowSelfApproval"`
	RequiredApprovers      int  `json:"RequiredApprovers"`
	RequireTimeRestriction bool `json:"RequireTimeRestriction"`
	MaximumTimeRestriction int  `json:"MaximumTimeRestriction"`
}

// PolicyReviewerProperties represents settings related to reviewing a password request
type PolicyReviewerProperties struct {
	RequireReview     bool `json:"RequireReview"`
	AllowSelfReview   bool `json:"AllowSelfReview"`
	RequiredReviewers int  `json:"RequiredReviewers"`
}

// PolicyRequesterProperties represents settings for requesting asset/accounts
type PolicyRequesterProperties struct {
	AllowEmergencyAccess       bool `json:"AllowEmergencyAccess"`
	AllowUseRequestComments    bool `json:"AllowUseRequestComments"`
	RequireUseRequestComments  bool `json:"RequireUseRequestComments"`
	MaximumDaysUntilExpiration int  `json:"MaximumDaysUntilExpiration"`
}

// PolicyEmergencyAccessProperties represents settings related to emergency access
type PolicyEmergencyAccessProperties struct {
	RequireEmergencyTicketNumber bool   `json:"RequireEmergencyTicketNumber"`
	EmergencyTicketSystem        string `json:"EmergencyTicketSystem"`
}

// PolicyScopeItem represents requestable items governed by policy
type PolicyScopeItem struct {
	Id                 int    `json:"Id"`
	Name               string `json:"Name"`
	Description        string `json:"Description"`
	AssetPartitionId   int    `json:"AssetPartitionId"`
	AssetPartitionName string `json:"AssetPartitionName"`
	Type               string `json:"Type"`
}

func GetAccessPolicies(c *client.SafeguardClient, filter client.Filter) ([]AccessPolicy, error) {
	var accessPolicies []AccessPolicy

	query := "AccessPolicies" + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &accessPolicies); err != nil {
		return nil, err
	}

	for i := range accessPolicies {
		accessPolicies[i].client = c
	}

	return accessPolicies, nil
}

func GetAccessPolicy(c *client.SafeguardClient, id int, fields client.Fields) (AccessPolicy, error) {
	var accessPolicy AccessPolicy

	query := fmt.Sprintf("AccessPolicies/%d", id)

	if len(fields) > 0 {
		query += fields.ToQueryString()
	}

	response, err := c.GetRequest(query)
	if err != nil {
		return accessPolicy, err
	}
	if err := json.Unmarshal(response, &accessPolicy); err != nil {
		return accessPolicy, err
	}

	accessPolicy.client = c
	return accessPolicy, nil
}
