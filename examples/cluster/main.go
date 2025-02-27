package main

import (
	"fmt"

	safeguard "github.com/sthayduk/safeguard-go"
	"github.com/sthayduk/safeguard-go/client"
	"github.com/sthayduk/safeguard-go/examples/common"
)

func main() {
	err := common.InitClient()
	if err != nil {
		panic(err)
	}

	// Example: GetClusterMembers
	filter := client.Filter{}
	clusterMembers, err := safeguard.GetClusterMembers(filter)
	if err != nil {
		fmt.Printf("Error getting cluster members: %v\n", err)
	} else {
		fmt.Printf("Cluster Members: %+v\n", clusterMembers)
	}

	// Example: GetClusterMember
	memberID := "46995a16b0b7482899cc6c60f4a0d86d" // Replace with actual member ID
	clusterMember, err := safeguard.GetClusterMember(memberID)
	if err != nil {
		fmt.Printf("Error getting cluster member: %v\n", err)
	} else {
		fmt.Printf("Cluster Member: %+v\n", clusterMember)
	}

	// Example: GetClusterLeader
	clusterLeader, err := safeguard.GetClusterLeader()
	if err != nil {
		fmt.Printf("Error getting cluster leader: %v\n", err)
	} else {
		fmt.Printf("Cluster Leader: %+v\n", clusterLeader)
	}

	// Example: ForceClusterHealthCheck
	clusterHealth, err := safeguard.ForceClusterHealthCheck()
	if err != nil {
		fmt.Printf("Error forcing cluster health check: %v\n", err)
	} else {
		fmt.Printf("Cluster Health: %+v\n", clusterHealth)
	}
}
