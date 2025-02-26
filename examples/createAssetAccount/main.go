package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sthayduk/safeguard-go/client"
	"github.com/sthayduk/safeguard-go/examples/common"
	"github.com/sthayduk/safeguard-go/models"
)

func main() {
	adId := 464
	samAccountName := "mustermann"

	// Initialize the logger
	logger := log.New(os.Stdout, "[CreateUser] ", log.LstdFlags|log.Lshortfile)

	// Initialize the Safeguard client
	logger.Println("Initializing Safeguard client...")
	sgc, err := common.InitClient()
	if err != nil {
		logger.Fatalf("Failed to initialize client: %v", err)
	}

	// Get the Active Directory
	logger.Printf("Getting Active Directory with ID: %d", adId)
	ad, err := models.GetAsset(sgc, adId, client.Fields{"Id", "Name"})
	if err != nil {
		logger.Fatalf("Failed to get Active Directory: %v", err)
	}

	// Search for the user
	logger.Printf("Searching for user: %s", samAccountName)
	filter := client.Filter{}
	filter.AddFilter("Name", "eq", samAccountName)
	users, err := ad.GetDirectoryAccounts(filter)
	if err != nil {
		logger.Fatalf("Failed to get directory users: %v", err)
	}

	// Validate search results
	if len(users) == 0 {
		logger.Fatalf("No user found with name: %s", samAccountName)
	}
	if len(users) > 1 {
		logger.Fatalf("Multiple users found with name: %s", samAccountName)
	}

	// Create the user
	logger.Println("Creating user in Safeguard...")
	createdUser, err := users[0].Create()
	if err != nil {
		logger.Fatalf("Failed to create user: %v", err)
	}

	// Print success message
	logger.Printf("Successfully created user with ID: %d", createdUser.Id)
	fmt.Printf("\nUser Details:\n"+
		"ID:                %d\n"+
		"Name:              %s\n"+
		"DistinguishedName: %s\n"+
		"Created by User:   %s\n",
		createdUser.Id, createdUser.Name, createdUser.DistinguishedName, createdUser.CreatedByUserDisplayName)

	// Update User Password
	taskState, err := createdUser.ChangePassword()
	if err != nil {
		logger.Fatalf("Failed to change user password: %v", err)
	}

	// Wait for the task to complete
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := taskState.CheckTaskState(ctx); err != nil {
		logger.Fatalf("Failed to check task state: %v", err)
	}

	// Check Users Password
	passwordStatus, err := createdUser.CheckPassword()
	if err != nil {
		logger.Fatalf("Failed to check user password: %v", err)
	}
	logger.Printf("User's password status: %s", passwordStatus.RequestStatus.State)
	logger.Printf("   Message: %s      ", passwordStatus.RequestStatus.Message)
	logger.Printf("   TotalDuration: %s", passwordStatus.RequestStatus.TotalDuration)

}
