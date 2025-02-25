package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sthayduk/safeguard-go/client"
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

// GetAccessPolicies retrieves a list of access policies from the Safeguard API.
// It takes a SafeguardClient and a Filter as parameters and returns a slice of AccessPolicy and an error.
//
// Parameters:
//   - c: A pointer to a SafeguardClient used to make the API request.
//   - filter: A Filter object used to filter the access policies.
//
// Returns:
//   - A slice of AccessPolicy objects.
//   - An error if the request fails or the response cannot be unmarshaled.
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

// GetAccessPolicy retrieves an access policy by its ID from the Safeguard API.
// It takes a SafeguardClient, an integer ID of the access policy, and optional fields to include in the query.
// It returns the AccessPolicy and an error if any occurred during the request or unmarshalling process.
//
// Parameters:
//   - c: A pointer to the SafeguardClient used to make the API request.
//   - id: An integer representing the ID of the access policy to retrieve.
//   - fields: Optional fields to include in the query.
//
// Returns:
//   - AccessPolicy: The retrieved access policy.
//   - error: An error if any occurred during the request or unmarshalling process.
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

// DeleteAccessPolicy deletes an access policy with the given ID using the provided SafeguardClient.
// It constructs a query string with the access policy ID and sends a DELETE request.
// If the request fails, it returns an error.
//
// Parameters:
//   - c: A pointer to a SafeguardClient instance used to send the DELETE request.
//   - id: An integer representing the ID of the access policy to be deleted.
//
// Returns:
//   - error: An error object if the DELETE request fails, otherwise nil.
func DeleteAccessPolicy(c *client.SafeguardClient, id int) error {
	query := fmt.Sprintf("AccessPolicies/%d", id)

	_, err := c.DeleteRequest(query)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes the access policy from the system.
// It calls the DeleteAccessPolicy function with the client and policy ID.
// Returns an error if the deletion fails.
func (a AccessPolicy) Delete() error {
	return DeleteAccessPolicy(a.client, a.Id)
}
