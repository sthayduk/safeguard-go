package main

import (
	"fmt"
	"log"

	safeguard "github.com/sthayduk/safeguard-go"
	"github.com/sthayduk/safeguard-go/examples/common"
)

func main() {
	// Initialize Safeguard client
	// Replace with your Safeguard appliance address
	sgc, err := common.InitClient()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully authenticated to Safeguard")

	// Example 1: Get all identities
	fmt.Println("\n=== Example 1: Get all identities ===")
	identities, err := sgc.GetIdentities(safeguard.Filter{})
	if err != nil {
		log.Fatalf("Failed to get identities: %v", err)
	}
	fmt.Printf("Found %d identities\n", len(identities))

	// Print first 67 identities or all if less than 67
	displayCount := min(67, len(identities))
	for i := 0; i < displayCount; i++ {
		fmt.Printf("Identity %d: %s (ID: %d)\n", i+1, identities[i].DisplayName, identities[i].Id)
	}

	// Example 2: Get identities with filtering
	fmt.Println("\n=== Example 2: Get identities with filtering ===")
	filter := safeguard.Filter{}
	filter.AddField("Id")
	filter.AddField("IdentityProviderName")
	filter.AddField("DisplayName")
	filter.AddField("DomainName")
	filter.AddField("EmailAddress")
	filter.AddField("PrincipalKind")
	filter.AddFilter("PrincipalKind", safeguard.OpEqual, "User")
	filter.AddOrderBy("DisplayName")

	userIdentities, err := sgc.GetIdentities(filter)
	if err != nil {
		log.Fatalf("Failed to get filtered identities: %v", err)
	}
	fmt.Printf("Found %d user identities\n", len(userIdentities))
	for i, identity := range userIdentities {
		if i >= 3 {
			break // Display only the first 3
		}
		fmt.Printf("User Identity %d: %s (ID: %d)\n", i+1, identity.DisplayName, identity.Id)
	}

	// Example 3: Get a specific identity by ID
	// For this example, we'll use the first identity from our previous query
	if len(userIdentities) == 0 {
		fmt.Println("No identities found to use for specific identity example")
		return
	}

	fmt.Println("\n=== Example 3: Get a specific identity by ID ===")
	targetID := userIdentities[51].Id
	identity, err := sgc.GetIdentity(targetID, safeguard.Fields{})
	if err != nil {
		log.Fatalf("Failed to get identity with ID %d: %v", targetID, err)
	}
	fmt.Printf("Retrieved identity: %s (ID: %d)\n", identity.DisplayName, identity.Id)
	fmt.Printf("  Idenitiy Provider Name: %s\n", identity.IdentityProviderName)
	fmt.Printf("  Email: %s\n", identity.EmailAddress)
	fmt.Printf("  Identity Provider: %s (ID: %d)\n", identity.IdentityProviderName, identity.IdentityProviderId)

	// Example 4: Get identity provider for an identity
	fmt.Println("\n=== Example 4: Get identity provider for an identity ===")
	provider, err := identity.GetIdentityProvider(safeguard.Fields{})
	if err != nil {
		log.Printf("Failed to get identity provider: %v", err)
	} else {
		fmt.Printf("Identity Provider: %s (ID: %d)\n", provider.Name, provider.Id)
		fmt.Printf("  Type: %s\n", provider.TypeReferenceName)
	}

	// Example 5: Get user associated with an identity
	fmt.Println("\n=== Example 5: Get user associated with an identity ===")
	user, err := identity.GetUser(safeguard.Fields{})
	if err != nil {
		log.Printf("Failed to get user: %v", err)
	} else {
		fmt.Printf("User: %s (ID: %d)\n", user.Name, user.Id)
		fmt.Printf("  DisplayName Name: %s\n", user.DisplayName)
	}

	// Example 6: Get user group associated with an identity
	fmt.Println("\n=== Example 6: Get user group associated with an identity ===")
	userGroup, err := identity.GetUserGroup(safeguard.Fields{})
	if err != nil {
		log.Printf("Failed to get user group: %v", err)
	} else {
		fmt.Printf("User Group: %s (ID: %d)\n", userGroup.Name, userGroup.Id)
		fmt.Printf("  Description: %s\n", userGroup.Description)
	}

	fmt.Println("\nExamples completed")
}

// min returns the smaller of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
