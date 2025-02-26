package main

import (
	"fmt"

	sg "github.com/sthayduk/safeguard-go"
	"github.com/sthayduk/safeguard-go/client"
	"github.com/sthayduk/safeguard-go/examples/common"
)

func main() {
	sgc, err := common.InitClient()
	if err != nil {
		panic(err)
	}

	// Get all asset groups
	assetGroups, err := sg.GetAssetGroups(sgc, client.Filter{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Found %d asset groups\n", len(assetGroups))

	// Print basic info about each asset group
	for _, group := range assetGroups {
		fmt.Printf("Asset Group: %s (ID: %d)\n", group.Name, group.Id)
	}

	// Get a specific asset group with additional fields
	if len(assetGroups) > 0 {
		fields := client.Fields{"Name", "Description", "CreatedDate"}
		assetGroup, err := sg.GetAssetGroup(sgc, assetGroups[0].Id, fields)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\nDetailed Asset Group Info:\n")
		fmt.Printf("Name: %s\n", assetGroup.Name)
		fmt.Printf("Description: %s\n", assetGroup.Description)
		fmt.Printf("Created: %s\n", assetGroup.CreatedDate)
	}
}
