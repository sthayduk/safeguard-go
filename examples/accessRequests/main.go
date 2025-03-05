package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sthayduk/safeguard-go"
	"github.com/sthayduk/safeguard-go/examples/common"
)

func main() {
	sgc, err := common.InitClient()
	if err != nil {
		panic(err)
	}

	// Example 1: Get information about the current user
	fmt.Println("Example 1: Getting current user information")

	me, err := sgc.GetMe(safeguard.Filter{})
	if err != nil {
		fmt.Printf("Error getting current user: %s\n", err)
	}
	fmt.Printf("Logged in user: %s (ID: %d)\n", me.Name, me.Id)
	fmt.Printf("Email: %s\n", me.EmailAddress)
	fmt.Printf("Admin Roles: %v\n", me.AdminRoles)

	// Example 2: Get all Entitlements
	fmt.Println("Example 2: Getting all entitlements")
	entitlements, err := sgc.GetMeAccountEntitlements(safeguard.AccessRequestTypePassword, false, false, safeguard.Filter{})
	if err != nil {
		fmt.Printf("Error getting entitlements: %s\n", err)
		panic(err)
	}

	fmt.Printf("Found %d entitlements\n", len(entitlements))

	// Print basic info about each entitlement
	for _, entitlement := range entitlements {
		fmt.Printf("(%d) AccountName: %s (AccountDomain: %s)\n", entitlement.Account.Id, entitlement.Account.Name, entitlement.Account.DomainName)
		fmt.Println("Get Access Request for Account")

		accessRequest, err := sgc.GetAccessRequests(entitlement.GetFilter())
		if err != nil {
			fmt.Printf("Error getting access request: %s\n", err)
			panic(err)
		}

		fmt.Printf("Found %d access requests\n", len(accessRequest))
		for _, request := range accessRequest {
			fmt.Println("Access Request ID: ", request.Id)
			fmt.Println("  Account ID:   ", request.AccountId)
			fmt.Println("  Account Name: ", request.AccountName)
			fmt.Println("  Asset ID:     ", request.AccountAssetId)
			fmt.Println("  Asset Name:   ", request.AccountAssetName)
			fmt.Println("  State:        ", request.State)
		}
	}

	// Example 3: New Access Request
	fmt.Println("Example 3: Creating a new access request")
	accessRequest, err := sgc.NewAccessRequests(entitlements, 1200*time.Second)
	if err != nil {
		fmt.Printf("Error creating access request: %s\n", err)
		//panic(err)
	}

	fmt.Println("Access Request ID: ", accessRequest[0].Response.Id)

	// Example 4: Get Password for Access Request
	fmt.Println("Example 4: Getting password for access request")
	for _, entitlement := range entitlements {
		fmt.Printf("(%d) AccountName: %s (AccountDomain: %s)\n", entitlement.Account.Id, entitlement.Account.Name, entitlement.Account.DomainName)
		fmt.Println("Get Access Request for Account")

		accessRequest, err := sgc.GetAccessRequests(entitlement.GetFilter())
		if err != nil {
			fmt.Printf("Error getting access request: %s\n", err)
			panic(err)
		}

		fmt.Printf("Found %d access requests\n", len(accessRequest))
		for _, request := range accessRequest {
			fmt.Println("Access Request ID: ", request.Id)
			fmt.Println("  Account ID:   ", request.AccountId)
			fmt.Println("  Account Name: ", request.AccountName)
			fmt.Println("  Asset ID:     ", request.AccountAssetId)
			fmt.Println("  Asset Name:   ", request.AccountAssetName)
			fmt.Println("  State:        ", request.State)

			if request.State == safeguard.StateCompleted {
				continue
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			password, err := request.CheckOutPassword(ctx, true)
			if err != nil {
				fmt.Printf("Error checking out password: %s\n", err)
				//panic(err)
			} else {
				fmt.Printf("Password: %s\n", password)
			}
		}
	}

	// Example 5: Close Access Requests
	fmt.Println("Example 5: Close an access request")
	for _, entitlement := range entitlements {
		fmt.Printf("(%d) AccountName: %s (AccountDomain: %s)\n", entitlement.Account.Id, entitlement.Account.Name, entitlement.Account.DomainName)
		fmt.Println("Get Access Request for Account")

		accessRequest, err := sgc.GetAccessRequests(entitlement.GetFilter())
		if err != nil {
			fmt.Printf("Error getting access request: %s\n", err)
			panic(err)
		}

		fmt.Printf("Found %d access requests\n", len(accessRequest))
		for _, request := range accessRequest {
			fmt.Println("Access Request ID: ", request.Id)
			fmt.Println("  Account ID:   ", request.AccountId)
			fmt.Println("  Account Name: ", request.AccountName)
			fmt.Println("  Asset ID:     ", request.AccountAssetId)
			fmt.Println("  Asset Name:   ", request.AccountAssetName)
			fmt.Println("  State:        ", request.State)

			_, err := request.Close()
			if err != nil {
				fmt.Printf("Error checkin access request: %s\n", err)
				panic(err)
			}
		}
	}
}
