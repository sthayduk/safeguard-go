package main

import (
	"fmt"

	"github.com/sthayduk/safeguard-go/client"
	"github.com/sthayduk/safeguard-go/examples/common"
	"github.com/sthayduk/safeguard-go/models"
)

func main() {
	sgc, err := common.InitClient()
	if err != nil {
		panic(err)
	}

	// Example 1: Get information about the current user
	fmt.Println("Example 1: Getting current user information")
	filter := client.Filter{}
	filter.AddField("Name")
	filter.AddField("EmailAddress")
	filter.AddField("AdminRoles")

	me, err := models.GetMe(sgc, filter)
	if err != nil {
		fmt.Printf("Error getting current user: %s\n", err)
	} else {
		fmt.Printf("Logged in user: %s (ID: %d)\n", me.Name, me.Id)
		fmt.Printf("Email: %s\n", me.EmailAddress)
		fmt.Printf("Admin Roles: %v\n", me.AdminRoles)
	}
	fmt.Println()

	// Example 2: Get assets available for access requests
	fmt.Println("Example 2: Getting accessible assets")
	assetFilter := client.Filter{}
	assetFilter.AddField("Name")
	assetFilter.AddField("NetworkAddress")
	assetFilter.AddField("Platform.DisplayName")
	assetFilter.AddOrderBy("Name")

	assets, err := models.GetMeAccessRequestAssets(sgc, assetFilter)
	if err != nil {
		fmt.Printf("Error getting accessible assets: %s\n", err)
	} else {
		fmt.Printf("Found %d assets available for access requests:\n", len(assets))
		for i, asset := range assets {
			fmt.Printf("%d. %s (%s) - %s\n", i+1, asset.Name, asset.NetworkAddress, asset.Platform.DisplayName)
			if i >= 4 {
				fmt.Println("   ... (showing first 5 only)")
				break
			}
		}

		// Example 3: Get a specific asset by ID
		if len(assets) > 0 {
			fmt.Println("\nExample 3: Getting a specific asset by ID")
			assetId := fmt.Sprintf("%d", assets[0].Id)
			asset, err := models.GetMeAccessRequestAsset(sgc, assetId)
			if err != nil {
				fmt.Printf("Error getting asset by ID: %s\n", err)
			} else {
				fmt.Printf("Asset Details:\n")
				fmt.Printf("ID: %d\n", asset.Id)
				fmt.Printf("Name: %s\n", asset.Name)
				fmt.Printf("Platform: %s\n", asset.Platform.DisplayName)
				fmt.Printf("Network Address: %s\n", asset.NetworkAddress)
			}
		}
	}
	fmt.Println()

	// Example 4: Get actionable requests with detailed information
	fmt.Println("Example 4: Getting actionable requests with details")
	requestFilter := client.Filter{}
	actionableRequests, err := models.GetMeActionableRequestsDetailed(sgc, requestFilter)
	if err != nil {
		fmt.Printf("Error getting actionable requests: %s\n", err)
	} else {
		fmt.Printf("Found %d total actionable requests across %d roles\n",
			actionableRequests.TotalCount, len(actionableRequests.AvailableRoles))

		// Show counts by role
		fmt.Println("\nCounts by role:")
		for role, count := range actionableRequests.CountByRole {
			fmt.Printf("- %s: %d requests\n", role, count)
		}

		// Show pending requests
		pending := actionableRequests.GetPendingRequests()
		fmt.Printf("\nPending Requests: %d\n", len(pending))
		for i, req := range pending {
			fmt.Printf("- %s: %s (%s) - %s\n",
				req.Id, req.AccountName, req.AssetName, req.State)
			if i >= 4 {
				fmt.Println("... (showing first 5 only)")
				break
			}
		}

		// Example: Filter requests by state
		checkedOutRequests := actionableRequests.FilterRequestsByState(models.StatePasswordCheckedOut)
		fmt.Printf("\nChecked Out Requests: %d\n", len(checkedOutRequests))
		for i, req := range checkedOutRequests {
			fmt.Printf("- %s: %s (%s)\n",
				req.Id, req.AccountName, req.AssetName)
			if i >= 4 {
				fmt.Println("... (showing first 5 only)")
				break
			}
		}

		// Show Admin role requests if available
		if actionableRequests.HasRole(models.AdminRole) {
			adminRequests := actionableRequests.GetRequestsForRole(models.AdminRole)
			fmt.Printf("\nAdmin Role Requests: %d\n", len(adminRequests))
			// ... rest of admin requests display
		}
	}
	fmt.Println()

	// Example 5: Get actionable requests by role
	fmt.Println("Example 5: Getting actionable requests by role (Admin)")
	approverRequests, err := models.GetMeActionableRequestsByRole(sgc, models.AdminRole, requestFilter)
	if err != nil {
		fmt.Printf("Error getting approver requests: %s\n", err)
	} else {
		fmt.Printf("Found %d access requests for role Admin\n", len(approverRequests))
		for i, req := range approverRequests {
			fmt.Printf("\nAccess Request %d:\n", i+1)
			fmt.Printf("ID: %s\n", req.Id)
			fmt.Printf("Type: %s\n", req.AccessRequestType)
			fmt.Printf("Account: %s (%s)\n", req.AccountName, req.AccountDomainName)
			fmt.Printf("Asset: %s (%s)\n", req.AssetName, req.AssetPlatformDisplayName)
			fmt.Printf("Requester: %s (%s)\n", req.RequesterDisplayName, req.RequesterEmailAddress)
			fmt.Printf("State: %s\n", req.State)
			fmt.Printf("Created: %s\n", req.CreatedOn)
			fmt.Printf("Expires: %s\n", req.ExpiresOn)

			if len(req.WorkflowActions) > 0 {
				fmt.Printf("Recent action: %s on %s\n",
					req.WorkflowActions[len(req.WorkflowActions)-1].ActionType,
					req.WorkflowActions[len(req.WorkflowActions)-1].OccurredOn)
			}

			if i >= 2 {
				fmt.Println("\n... (showing first 3 only)")
				break
			}
		}
	}
}
