package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sthayduk/safeguard-go/client"
)

// ClusterMember represents a node in the Safeguard cluster
type ClusterMember struct {
	client *client.SafeguardClient

	Id                 string                 `json:"Id"`
	Name               string                 `json:"Name"`
	NetworkAddress     string                 `json:"NetworkAddress"`
	Description        string                 `json:"Description"`
	IsLeader           bool                   `json:"IsLeader"`
	Version            string                 `json:"Version"`
	PatchVersion       string                 `json:"PatchVersion"`
	State              ClusterOperationState  `json:"State"`
	EnrollmentDate     time.Time              `json:"EnrollmentDate"`
	IsEnrolled         bool                   `json:"IsEnrolled"`
	Health             NodeHealth             `json:"Health"`
	NetworkInformation NodeNetworkInformation `json:"NetworkInformation"`
}

// NodeHealth represents the health status of a cluster node
type NodeHealth struct {
	Status             HealthStatus           `json:"Status"`
	Details            []HealthDetail         `json:"Details"`
	LastUpdateTime     time.Time              `json:"LastUpdateTime"`
	ResourceHealth     NodeResourceHealth     `json:"ResourceHealth"`
	ConnectivityHealth NodeConnectivityHealth `json:"ConnectivityHealth"`
}

// NodeResourceHealth represents resource health information for a node
type NodeResourceHealth struct {
	Status  HealthStatus               `json:"Status"`
	Details []NodeResourceHealthDetail `json:"Details"`
}

// NodeResourceHealthDetail represents detailed resource health information
type NodeResourceHealthDetail struct {
	Name        string       `json:"Name"`
	Status      HealthStatus `json:"Status"`
	Description string       `json:"Description"`
}

// NodeConnectivityHealth represents connectivity health information for a node
type NodeConnectivityHealth struct {
	Status  HealthStatus                      `json:"Status"`
	Details []ClusterConnectivityHealthDetail `json:"Details"`
}

// NodeNetworkInformation represents network configuration for a node
type NodeNetworkInformation struct {
	Ipv4Address    string `json:"Ipv4Address"`
	Ipv6Address    string `json:"Ipv6Address"`
	Netmask        string `json:"Netmask"`
	Gateway        string `json:"Gateway"`
	DnsServers     string `json:"DnsServers"`
	UsingDhcp      bool   `json:"UsingDhcp"`
	InterfaceAlias string `json:"InterfaceAlias"`
}

// HealthDetail represents a specific health status detail
type HealthDetail struct {
	Name        string       `json:"Name"`
	Status      HealthStatus `json:"Status"`
	Description string       `json:"Description"`
}

// ClusterConnectivityHealthDetail represents connectivity status between nodes
type ClusterConnectivityHealthDetail struct {
	Name        string       `json:"Name"`
	Status      HealthStatus `json:"Status"`
	Description string       `json:"Description"`
	Target      string       `json:"Target"`
}

// HealthStatus represents the possible health states of a node or service
type HealthStatus string

const (
	HealthStatusUnknown     HealthStatus = "Unknown"
	HealthStatusError       HealthStatus = "Error"
	HealthStatusWarning     HealthStatus = "Warning"
	HealthStatusHealthy     HealthStatus = "Healthy"
	HealthStatusUnavailable HealthStatus = "Unavailable"
)

// ClusterOperationState represents the possible states of a cluster operation
type ClusterOperationState string

const (
	ClusterOperationStateUnknown      ClusterOperationState = "Unknown"
	ClusterOperationStateInitializing ClusterOperationState = "Initializing"
	ClusterOperationStateReady        ClusterOperationState = "Ready"
	ClusterOperationStateInProgress   ClusterOperationState = "InProgress"
	ClusterOperationStateCompleted    ClusterOperationState = "Completed"
	ClusterOperationStateFailed       ClusterOperationState = "Failed"
)

// GetClusterMembers retrieves the members of a cluster from the Safeguard API.
// It takes a SafeguardClient and a Filter as parameters and returns a slice of ClusterMember and an error.
//
// Parameters:
//   - c: A pointer to a SafeguardClient used to make the API request.
//   - filter: A Filter object used to filter the cluster members.
//
// Returns:
//   - A slice of ClusterMember containing the cluster members.
//   - An error if the request fails or the response cannot be unmarshaled.
func GetClusterMembers(c *client.SafeguardClient, filter client.Filter) ([]ClusterMember, error) {
	var clusterMembers []ClusterMember

	query := "Cluster/Members" + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, &clusterMembers); err != nil {
		return nil, err
	}

	for i := range clusterMembers {
		clusterMembers[i].client = c
	}
	return clusterMembers, nil
}

// GetClusterMember retrieves a cluster member by its ID from the Safeguard API.
// It sends a GET request to the "Cluster/Members/{id}" endpoint and unmarshals
// the response into a ClusterMember struct.
//
// Parameters:
//   - c: A pointer to a SafeguardClient instance used to make the API request.
//   - id: An integer representing the ID of the cluster member to retrieve.
//
// Returns:
//   - ClusterMember: The retrieved cluster member.
//   - error: An error if the request or unmarshalling fails.
func GetClusterMember(c *client.SafeguardClient, id string) (ClusterMember, error) {
	var clusterMember ClusterMember

	query := fmt.Sprintf("Cluster/Members/%s", id)

	response, err := c.GetRequest(query)
	if err != nil {
		return ClusterMember{}, err
	}

	if err := json.Unmarshal(response, &clusterMember); err != nil {
		return ClusterMember{}, err
	}

	clusterMember.client = c
	return clusterMember, nil
}

// GetClusterLeader retrieves the cluster leader from the SafeguardClient.
// It applies a filter to find the cluster member with the "IsLeader" attribute set to true.
//
// Parameters:
//   - c: A pointer to the SafeguardClient instance.
//
// Returns:
//   - ClusterMember: The cluster member that is the leader.
//   - error: An error if no leader is found, more than one leader is found, or if there is an issue retrieving the cluster members.
func GetClusterLeader(c *client.SafeguardClient) (ClusterMember, error) {
	filter := client.Filter{}
	filter.AddFilter("IsLeader", "eq", "true")

	clusterMembers, err := GetClusterMembers(c, filter)
	if err != nil {
		fmt.Println(err)
		return ClusterMember{}, err
	}

	if len(clusterMembers) == 0 {
		return ClusterMember{}, fmt.Errorf("no cluster leader found")
	}

	if len(clusterMembers) > 1 {
		return ClusterMember{}, fmt.Errorf("invalid number of cluster leaders found")
	}

	clusterMembers[0].client = c
	return clusterMembers[0], nil
}

// ForceClusterHealthCheck performs a health check on the cluster members.
// It sends a GET request to the "Cluster/Members/Self" endpoint and unmarshals
// the response into a slice of ClusterMember structs. Each ClusterMember is then
// associated with the provided SafeguardClient.
//
// Parameters:
//   - c: A pointer to a SafeguardClient used to make the GET request.
//
// Returns:
//   - A slice of ClusterMember structs containing the details of each cluster member.
//   - An error if the request or unmarshalling fails.
func ForceClusterHealthCheck(c *client.SafeguardClient) (ClusterMember, error) {
	var clusterMembers ClusterMember
	query := "Cluster/Members/Self"

	response, err := c.GetRequest(query)
	if err != nil {
		return ClusterMember{}, err
	}

	if err := json.Unmarshal(response, &clusterMembers); err != nil {
		return ClusterMember{}, err
	}

	clusterMembers.client = c
	return clusterMembers, nil
}

// IsClusterLeader checks if the current cluster member is the leader.
// It returns true if the member is the leader, otherwise false.
func (c ClusterMember) IsClusterLeader() bool {
	return c.IsLeader
}

// GetHealth returns the health status of the cluster member.
// It returns a NodeHealth object and an error, which is always nil.
func (c ClusterMember) GetHealth() (NodeHealth, error) {
	return c.Health, nil
}
