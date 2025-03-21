package main

import (
	"fmt"
	"log"

	"github.com/sthayduk/safeguard-go"
	"github.com/sthayduk/safeguard-go/examples/common"
)

func main() {
	sgc, err := common.InitClient()
	if err != nil {
		panic(err)
	}

	// Example 1: Get all policy assets
	filter := safeguard.Filter{
		Orderby: []string{"Name"},
		Fields:  []string{"Id", "Name", "AssetType", "AssetPartitionName"},
	}

	policyAssets, err := sgc.GetPolicyAssets(filter)
	if err != nil {
		log.Fatalf("Failed to get policy assets: %v", err)
	}

	fmt.Println("=== All Policy Assets ===")
	for _, asset := range policyAssets {
		fmt.Printf("Asset: %s (ID: %d)\n", asset.Name, asset.Id)
	}

	// Example 2: Get a specific policy asset by ID
	if len(policyAssets) > 0 {
		fields := safeguard.Fields{"Id", "Name", "NetworkAddress", "Platform"}
		policyAsset, err := sgc.GetPolicyAsset(policyAssets[3].Id, fields)
		if err != nil {
			log.Fatalf("Failed to get policy asset: %v", err)
		}

		fmt.Printf("\n=== Specific Policy Asset Details ===\n")
		fmt.Printf("Name: %s\nNetwork Address: %s\nPlatform: %s\n",
			policyAsset.Name,
			policyAsset.NetworkAddress,
			policyAsset.Platform.DisplayName)

		// Example 3: Convert policy asset to JSON
		jsonStr, err := policyAsset.ToJson()
		if err != nil {
			log.Fatalf("Failed to convert to JSON: %v", err)
		}
		fmt.Printf("\n=== Policy Asset JSON ===\n%s\n", jsonStr)

		// Example 4: Get asset groups for the policy asset
		assetGroupFilter := safeguard.Filter{
			Fields: []string{"Id", "Name", "Description"},
		}
		assetGroups, err := policyAsset.GetAssetGroups(assetGroupFilter)
		if err != nil {
			log.Fatalf("Failed to get asset groups: %v", err)
		}

		fmt.Printf("\n=== Asset Groups ===\n")
		for _, group := range assetGroups {
			fmt.Printf("Group: %s (ID: %d)\n", group.Name, group.Id)
		}

		// Example 5: Get directory service entries
		dseFilter := safeguard.Filter{
			Fields: []string{"Name", "DirectoryProperties"},
		}
		entries, err := policyAsset.GetDirectoryServiceEntries(dseFilter)
		if err != nil {
			log.Fatalf("Failed to get directory service entries: %v", err)
		}

		fmt.Printf("\n=== Directory Service Entries ===\n")
		for _, entry := range entries {
			fmt.Printf("Entry: %s (Distinguished Name: %s, Directory ID: %d)\n", entry.Name, entry.DirectoryProperties.DistinguishedName, entry.DirectoryProperties.DirectoryId)
		}

		// Example 6: Get policies for the asset
		policiesFilter := safeguard.Filter{
			Fields: []string{"PolicyId", "PolicyName"},
		}
		policies, err := policyAsset.GetPolicies(policiesFilter)
		if err != nil {
			log.Fatalf("Failed to get policies: %v", err)
		}

		fmt.Printf("\n=== Asset Policies ===\n")
		for _, policy := range policies {
			fmt.Printf("Policy: %s (ID: %d)\n", policy.PolicyName, policy.PolicyId)
		}
	}
}
