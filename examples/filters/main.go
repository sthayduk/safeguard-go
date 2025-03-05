package main

import (
	"fmt"
	"log"

	"github.com/sthayduk/safeguard-go"
	"github.com/sthayduk/safeguard-go/examples/common"
)

func main() {
	// Initialize the Safeguard client
	sgc, err := common.InitClient()
	if err != nil {
		log.Fatalf("Failed to initialize client: %v", err)
	}

	// Example 1: Basic filter with fields
	fmt.Println("\n=== Example 1: Basic filter with fields ===")
	basicFilter()

	// Example 2: Adding filter conditions
	fmt.Println("\n=== Example 2: Adding filter conditions ===")
	filterConditions()

	// Example 3: Complex search filter
	fmt.Println("\n=== Example 3: Complex search filter ===")
	complexSearchFilter()

	// Example 4: Using filters with the API
	fmt.Println("\n=== Example 4: Using filters with the API ===")
	filterWithAPI(sgc)

	// Example 5: Filtering assets by name
	fmt.Println("\n=== Example 5: Filtering assets by name ===")
	filterAssetsByName(sgc, "ITd-Active-Directory")
}

// basicFilter demonstrates how to create a basic filter with fields and ordering
func basicFilter() {
	// Create a new filter
	filter := safeguard.Filter{}

	// Add fields to include in the response
	filter.AddField("Name")
	filter.AddField("Description")
	filter.AddField("CreatedDate")

	// Add ordering by created date in descending order
	filter.AddOrderBy("-CreatedDate")

	// Set count flag to true to include total count in the response
	filter.Count = true

	// Generate the query string
	queryString := filter.ToQueryString()
	fmt.Printf("Basic filter query string: %s\n", queryString)
	// Output: ?fields=Name,Description,CreatedDate&count=true&orderby=-CreatedDate
}

// filterConditions demonstrates how to add filter conditions
func filterConditions() {
	// Create a new filter
	filter := safeguard.Filter{}

	// Add a filter to find items with a specific name
	filter.AddFilter("Name", safeguard.OpEqual, "Administrator")

	// Add a filter to find items created after a certain date
	filter.AddFilter("CreatedDate", safeguard.OpGreaterThan, "2023-01-01")

	// Generate the query string
	queryString := filter.ToQueryString()
	fmt.Printf("Filter conditions query string: %s\n", queryString)
	// Output: ?filter=(Name eq 'Administrator' and CreatedDate gt '2023-01-01')&count=false

	// Create a new filter with OR condition
	orFilter := safeguard.Filter{}
	orConditions := map[string]safeguard.FilterOperator{
		"Name":        safeguard.OpContains,
		"Description": safeguard.OpContains,
	}
	orFilter.AddComplexSearchFilter("admin", orConditions)

	queryString = orFilter.ToQueryString()
	fmt.Printf("OR filter conditions query string: %s\n", queryString)
	// Output: ?filter=(Name contains 'admin' or Description contains 'admin')&count=false
}

// complexSearchFilter demonstrates how to create a complex search filter
func complexSearchFilter() {
	// Create a new filter
	filter := safeguard.Filter{}

	// Add a standard search filter that searches across multiple fields
	filter.AddSearchFilter("database")

	// Generate the query string
	queryString := filter.ToQueryString()
	fmt.Printf("Complex search filter query string: %s\n", queryString)

	// Create a custom complex search filter
	customFilter := safeguard.Filter{}
	customSearchFields := map[string]safeguard.FilterOperator{
		"Name":           safeguard.OpStartsWith,
		"NetworkAddress": safeguard.OpContains,
		"Description":    safeguard.OpIContains,
	}
	customFilter.AddComplexSearchFilter("srv", customSearchFields)

	queryString = customFilter.ToQueryString()
	fmt.Printf("Custom complex search filter query string: %s\n", queryString)
}

// filterWithAPI demonstrates using filters with the Safeguard API
func filterWithAPI(sgc *safeguard.SafeguardClient) {
	// Create a filter to get assets that are Windows machines
	filter := safeguard.Filter{}
	filter.AddField("Name")
	filter.AddField("NetworkAddress")
	filter.AddField("Platform.DisplayName")
	filter.AddFilter("Platform.DisplayName", safeguard.OpContains, "Active Directory")
	filter.AddOrderBy("Name")

	// Use the filter with the API
	assets, err := sgc.GetMeAccessRequestAssets(filter)
	if err != nil {
		fmt.Printf("Error getting Active Directory assets: %s\n", err)
		return
	}

	fmt.Printf("Found %d Active Directory assets:\n", len(assets))
	for i, asset := range assets {
		fmt.Printf("%d. %s (%s) - %s\n", i+1, asset.Name, asset.NetworkAddress, asset.Platform.DisplayName)
		if i >= 4 {
			fmt.Println("   ... (showing first 5 only)")
			break
		}
	}
}

// filterAssetsByName demonstrates filtering assets by a name pattern
func filterAssetsByName(sgc *safeguard.SafeguardClient, searchTerm string) {
	// Create a filter to search for assets by name
	filter := safeguard.Filter{}
	filter.AddField("Name")
	filter.AddField("NetworkAddress")
	filter.AddField("Platform.DisplayName")
	filter.AddField("Description")

	// Add a complex search filter for the asset name patterns
	searchFields := map[string]safeguard.FilterOperator{
		"Name":           safeguard.OpContains,
		"NetworkAddress": safeguard.OpContains,
		"Description":    safeguard.OpContains,
	}
	filter.AddComplexSearchFilter(searchTerm, searchFields)
	filter.AddOrderBy("Name")

	// Use the filter with the API
	assets, err := sgc.GetMeAccessRequestAssets(filter)
	if err != nil {
		fmt.Printf("Error searching for assets with '%s': %s\n", searchTerm, err)
		return
	}

	fmt.Printf("Found %d assets matching '%s':\n", len(assets), searchTerm)
	for i, asset := range assets {
		fmt.Printf("%d. %s (%s) - %s\n", i+1, asset.Name, asset.NetworkAddress, asset.Platform.DisplayName)
		if asset.Description != "" {
			fmt.Printf("   Description: %s\n", asset.Description)
		}
		if i >= 4 {
			fmt.Println("   ... (showing first 5 only)")
			break
		}
	}
}
