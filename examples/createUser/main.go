package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sthayduk/safeguard-go"
	"github.com/sthayduk/safeguard-go/examples/common"
)

func main() {
	idpID := 449
	authProviderId := 448
	username := "pam.test"

	// Initialize the logger
	logger := log.New(os.Stdout, "[CreateUser] ", log.LstdFlags|log.Lshortfile)

	// Initialize the Safeguard client
	logger.Println("Initializing Safeguard client...")
	sgc, err := common.InitClient()
	if err != nil {
		logger.Fatalf("Failed to initialize client: %v", err)
	}

	// Get the Identity Provider
	logger.Printf("Getting Identity Provider with ID: %d", idpID)
	idp, err := sgc.GetIdentityProvider(idpID)
	if err != nil {
		logger.Fatalf("Failed to get Identity Provider: %v", err)
	}

	// Search for the user
	logger.Printf("Searching for user: %s", username)
	filter := safeguard.Filter{}
	filter.AddFilter("Name", "eq", username)
	users, err := idp.GetDirectoryUsers(filter)
	if err != nil {
		logger.Fatalf("Failed to get directory users: %v", err)
	}

	// Validate search results
	if len(users) == 0 {
		logger.Fatalf("No user found with name: %s", username)
	}
	if len(users) > 1 {
		logger.Fatalf("Multiple users found with name: %s", username)
	}

	// Create the user
	logger.Println("Creating user in Safeguard...")
	response, err := sgc.CreateUser(users[0])
	if err != nil {
		logger.Fatalf("Failed to create user: %v", err)
	}

	// Print success message
	logger.Printf("Successfully created user with ID: %d", response.Id)
	fmt.Printf("\nUser Details:\n"+
		"ID: %d\n"+
		"Name: %s\n"+
		"DisplayName: %s\n"+
		"EmailAddress: %s\n",
		response.Id, response.Name, response.DisplayName, response.EmailAddress)

	// Update Authentication Provider
	authProvider, err := sgc.GetAuthenticationProvider(authProviderId)
	if err != nil {
		logger.Fatalf("Failed to get Authentication Provider: %v", err)
	}
	response, err = response.SetAuthenticationProvider(authProvider)
	if err != nil {
		logger.Fatalf("Failed to set Authentication Provider: %v", err)
	}

	logger.Printf("Successfully set Authentication Provider for user with ID: %d", response.Id)
	fmt.Printf("Authentication Provider ID: %d\n", response.PrimaryAuthenticationProvider.Id)

}
