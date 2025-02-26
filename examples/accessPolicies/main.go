package main

import (
	"fmt"
	"log"

	sg "github.com/sthayduk/safeguard-go"
	"github.com/sthayduk/safeguard-go/client"
	"github.com/sthayduk/safeguard-go/examples/common"
)

func main() {
	// Initialize the Safeguard client
	sgc, err := common.InitClient()
	if err != nil {
		log.Fatalf("Failed to initialize client: %v", err)
	}

	// Example 1: List all access policies
	fmt.Println("\n=== Example 1: Listing all access policies ===")
	listAccessPolicies(sgc)

	// Example 2: Get a specific access policy
	fmt.Println("\n=== Example 2: Get specific access policy ===")
	getSpecificPolicy(sgc, 7) // Replace 1 with an actual policy ID

	// Example 3: Working with reason codes
	fmt.Println("\n=== Example 3: Working with reason codes ===")
	workWithReasonCodes(sgc, 7) // Replace 1 with an actual policy ID

	// Example 4: Delete an access policy
	fmt.Println("\n=== Example 4: Delete access policy ===")
	//deletePolicy(sgc, 1) // Replace 1 with an actual policy ID

	// Example: Dump reason codes for all access policies
	fmt.Println("\n=== Dumping reason codes for all access policies ===")
	dumpAllReasonCodes(sgc)
}

func listAccessPolicies(sgc *client.SafeguardClient) {
	filter := client.Filter{}
	policies, err := sg.GetAccessPolicies(sgc, filter)
	if err != nil {
		log.Printf("Failed to get access policies: %v", err)
		return
	}

	fmt.Printf("Found %d access policies:\n", len(policies))
	for _, policy := range policies {
		fmt.Printf("- %s (ID: %d, Role: %s)\n", policy.Name, policy.Id, policy.RoleName)
	}
}

func getSpecificPolicy(sgc *client.SafeguardClient, policyID int) {
	fields := client.Fields{}
	policy, err := sg.GetAccessPolicy(sgc, policyID, fields)
	if err != nil {
		log.Printf("Failed to get access policy %d: %v", policyID, err)
		return
	}

	fmt.Printf("Policy details:\n")
	fmt.Printf("- Name: %s\n", policy.Name)
	fmt.Printf("- Role: %s\n", policy.RoleName)
	fmt.Printf("- Priority: %d\n", policy.Priority)
	fmt.Printf("- Account Count: %d\n", policy.AccountCount)
	fmt.Printf("- Asset Count: %d\n", policy.AssetCount)
}

func workWithReasonCodes(sgc *client.SafeguardClient, policyID int) {
	fields := client.Fields{}
	policy, err := sg.GetAccessPolicy(sgc, policyID, fields)
	if err != nil {
		log.Printf("Failed to get access policy %d: %v", policyID, err)
		return
	}

	reasonCodes := policy.GetReasonCodes()
	fmt.Printf("Found %d reason codes for policy %s:\n", len(reasonCodes), policy.Name)
	for _, code := range reasonCodes {
		fmt.Printf("- %s (ID: %d)\n", code.Name, code.Id)
		if code.Description != "" {
			fmt.Printf("  Description: %s\n", code.Description)
		}
		if code.Category != "" {
			fmt.Printf("  Category: %s\n", code.Category)
		}
	}
}

func deletePolicy(sgc *client.SafeguardClient, policyID int) {
	fields := client.Fields{}
	policy, err := sg.GetAccessPolicy(sgc, policyID, fields)
	if err != nil {
		log.Printf("Failed to get access policy %d: %v", policyID, err)
		return
	}

	fmt.Printf("Attempting to delete policy: %s (ID: %d)\n", policy.Name, policy.Id)
	if err := policy.Delete(); err != nil {
		log.Printf("Failed to delete policy: %v", err)
		return
	}
	fmt.Printf("Successfully deleted policy %s\n", policy.Name)
}

func dumpAllReasonCodes(sgc *client.SafeguardClient) {
	filter := client.Filter{}
	policies, err := sg.GetAccessPolicies(sgc, filter)
	if err != nil {
		log.Printf("Failed to get access policies: %v", err)
		return
	}

	fmt.Printf("Found %d access policies\n", len(policies))
	for _, policy := range policies {
		fmt.Printf("\nPolicy: %s (ID: %d)\n", policy.Name, policy.Id)
		fmt.Printf("Role: %s (Priority: %d)\n", policy.RoleName, policy.Priority)

		reasonCodes := policy.GetReasonCodes()
		if len(reasonCodes) == 0 {
			fmt.Printf("  No reason codes defined\n")
			continue
		}

		fmt.Printf("  Reason Codes (%d):\n", len(reasonCodes))
		for _, code := range reasonCodes {
			fmt.Printf("  - %s (ID: %d)\n", code.Name, code.Id)
			if code.Description != "" {
				fmt.Printf("    Description: %s\n", code.Description)
			}
			if code.Category != "" {
				fmt.Printf("    Category: %s\n", code.Category)
			}
		}
	}
}
