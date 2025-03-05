package main

import (
	"fmt"

	"github.com/sthayduk/safeguard-go"
	"github.com/sthayduk/safeguard-go/examples/common"
)

func main() {
	sgc, err := common.InitClient()
	if err != nil {
		panic(err)
	}

	// Example: GetClusterMembers
	filter := safeguard.Filter{}
	clusterMembers, err := sgc.GetClusterMembers(filter)
	if err != nil {
		fmt.Printf("Error getting cluster members: %v\n", err)
	} else {
		fmt.Printf("Cluster Members: %+v\n", clusterMembers)
	}

	// Example: GetClusterMember
	memberID := "46995a16b0b7482899cc6c60f4a0d86d" // Replace with actual member ID
	clusterMember, err := sgc.GetClusterMember(memberID)
	if err != nil {
		fmt.Printf("Error getting cluster member: %v\n", err)
	} else {
		fmt.Printf("Cluster Member: %+v\n", clusterMember)
	}

	// Example: GetClusterLeader
	clusterLeader, err := sgc.GetClusterLeader()
	if err != nil {
		fmt.Printf("Error getting cluster leader: %v\n", err)
	} else {
		fmt.Printf("Cluster Leader: %+v\n", clusterLeader)
	}

	// Example: ForceClusterHealthCheck
	clusterHealth, err := sgc.ForceClusterHealthCheck()
	if err != nil {
		fmt.Printf("Error forcing cluster health check: %v\n", err)
	} else {
		fmt.Printf("Cluster Health: %+v\n", clusterHealth)
	}
}
