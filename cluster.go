package safeguard

import (
	"encoding/json"
	"fmt"
	"time"
)

// ClusterMember represents a node in the Safeguard cluster
type ClusterMember struct {
	apiClient *SafeguardClient `json:"-"`

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

func (a ClusterMember) SetClient(c *SafeguardClient) any {
	a.apiClient = c
	return a
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

// GetClusterMembers retrieves all members that are part of the Safeguard cluster.
// Use filters to narrow down the results based on specific criteria.
//
// Parameters:
//   - filter: A Filter object containing query parameters to filter the results
//
// Returns:
//   - []ClusterMember: A slice of cluster members matching the filter criteria
//   - error: An error if the API request fails or the response cannot be parsed
func (c *SafeguardClient) GetClusterMembers(filter Filter) ([]ClusterMember, error) {
	var clusterMembers []ClusterMember

	query := "Cluster/Members" + filter.ToQueryString()

	response, err := c.GetRequest(query)
	if err != nil {
		return []ClusterMember{}, err
	}

	if err := json.Unmarshal(response, &clusterMembers); err != nil {
		return nil, err
	}

	return addClientToSlice(c, clusterMembers), nil
}

// GetClusterMember retrieves detailed information about a specific cluster member.
//
// Parameters:
//   - id: The unique identifier (GUID) of the cluster member to retrieve
//
// Returns:
//   - ClusterMember: The requested cluster member's configuration and status, or nil if not found
//   - error: An error if the member cannot be found or the request fails
func (c *SafeguardClient) GetClusterMember(id string) (ClusterMember, error) {
	var clusterMember ClusterMember

	query := fmt.Sprintf("Cluster/Members/%s", id)

	response, err := c.GetRequest(query)
	if err != nil {
		return ClusterMember{}, err
	}

	if err := json.Unmarshal(response, &clusterMember); err != nil {
		return ClusterMember{}, err
	}

	return addClient(c, clusterMember), nil
}

// GetClusterLeader identifies and retrieves the current leader of the Safeguard cluster.
//
// A healthy cluster should have exactly one leader at any given time. The leader
// is responsible for coordinating cluster-wide operations and maintaining consistency.
//
// Returns:
//   - ClusterMember: The cluster member that is currently the leader, or nil if no leader is found
//   - error: An error if no leader is found, multiple leaders are detected, or the request fails
func (c *SafeguardClient) GetClusterLeader() (ClusterMember, error) {
	filter := Filter{}
	filter.AddFilter("IsLeader", "eq", "true")

	clusterMembers, err := c.GetClusterMembers(filter)
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

	return addClient(c, clusterMembers[0]), nil
}

// ForceClusterHealthCheck triggers an immediate health check of the cluster.
//
// This operation initiates a comprehensive health assessment of the current node,
// including resource utilization, connectivity, and service status checks.
//
// Returns:
//   - ClusterMember: The cluster member representing the current node with updated health status
//   - error: An error if the health check fails to complete or the response cannot be parsed
func (c *SafeguardClient) ForceClusterHealthCheck() (ClusterMember, error) {
	var clusterMembers ClusterMember
	query := "Cluster/Members/Self"

	response, err := c.GetRequest(query)
	if err != nil {
		return ClusterMember{}, err
	}

	if err := json.Unmarshal(response, &clusterMembers); err != nil {
		return ClusterMember{}, err
	}

	return addClient(c, clusterMembers), nil
}

// IsClusterLeader checks whether this cluster member is currently the cluster leader.
//
// This is a convenience method that checks the leadership status of the current node
// without requiring a full cluster state query.
//
// Returns:
//   - bool: true if this member is the leader, false otherwise
//   - error: An error if the leadership status cannot be determined
func (c ClusterMember) IsClusterLeader() (bool, error) {
	return c.IsLeader, nil
}

// GetHealth retrieves the current health status information for this cluster member.
//
// The health status includes resource utilization, connectivity status, and
// any active warnings or errors affecting the node.
//
// Returns:
//   - *NodeHealth: The current health status and detailed health information
//   - error: An error if the health status cannot be retrieved
func (c ClusterMember) GetHealth() (*NodeHealth, error) {
	return &c.Health, nil
}
