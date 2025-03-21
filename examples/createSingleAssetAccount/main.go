package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sthayduk/safeguard-go"

	"github.com/fatih/color"
	"github.com/sthayduk/safeguard-go/examples/common"
)

func main() {
	start := time.Now()
	adId := 464
	samAccountName := "mustermann"

	// Initialize colored output
	success := color.New(color.FgGreen).SprintFunc()
	info := color.New(color.FgCyan).SprintFunc()
	warning := color.New(color.FgYellow).SprintFunc()

	// Initialize the logger
	logger := log.New(os.Stdout, info("[CreateUser] "), log.LstdFlags|log.Lshortfile)

	// Initialize the Safeguard client
	logger.Println("Initializing Safeguard client...")
	sgc, err := common.InitClient()
	if err != nil {
		logger.Fatalf("%s Failed to initialize client: %v", warning("ERROR:"), err)
	}

	// Get the Active Directory
	logger.Printf("Getting Active Directory with ID: %d", adId)
	ad, err := sgc.GetAsset(adId, safeguard.Fields{"Id", "Name"})
	if err != nil {
		logger.Fatalf("Failed to get Active Directory: %v", err)
	}

	// Search for the user
	logger.Printf("Searching for user: %s", samAccountName)
	filter := safeguard.Filter{}
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
		logger.Fatalf("%s Failed to create user: %v", warning("ERROR:"), err)
	}

	// Print user details in a formatted box
	fmt.Printf("\n%s\n", info("╭─────────── User Details ───────────╮"))
	fmt.Printf("│ ID:                %-15d │\n", createdUser.Id)
	fmt.Printf("│ Name:              %-15s │\n", createdUser.Name)
	fmt.Printf("│ DistinguishedName: %-15s │\n", createdUser.DistinguishedName)
	fmt.Printf("│ Created by:        %-15s │\n", createdUser.CreatedByUserDisplayName)
	fmt.Printf("%s\n\n", info("╰──────────────────────────────────╯"))

	// Suspend User
	logger.Println("Suspending user...")
	task, err := createdUser.Suspend()
	if err != nil {
		logger.Fatalf("Failed to suspend user: %v", err)
	}

	// Wait for the task to complete
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	logger.Println("Waiting for task to complete...")
	if _, err := task.CheckTaskState(ctx); err != nil {
		logger.Fatalf("Failed to check task state: %v", err)
	}

	// Update Password Profile
	logger.Println("Updating password profile...")
	assetPartition, err := sgc.GetAssetPartition(1, safeguard.Fields{"Id", "Name"})
	if err != nil {
		logger.Fatalf("Failed to get asset partition: %v", err)
	}

	filter = safeguard.Filter{}
	filter.AddFilter("Name", "eq", "ITdesign Profile Suspend")
	passwordProfile, err := sgc.GetPasswordRules(assetPartition, filter)
	if err != nil {
		logger.Fatalf("Failed to get password profile: %v", err)
	}

	updatedUser, err := createdUser.UpdatePasswordProfile(passwordProfile[0])
	if err != nil {
		logger.Fatalf("Failed to update user password profile: %v", err)
	}

	// Print updated user details in a formatted box
	fmt.Printf("\n%s\n", info("╭─────────── Updated User Details ───────────╮"))
	fmt.Printf("│ ID:                %-15d │\n", updatedUser.Id)
	fmt.Printf("│ Name:              %-15s │\n", updatedUser.Name)
	fmt.Printf("│ DistinguishedName: %-15s │\n", updatedUser.DistinguishedName)
	fmt.Printf("│ Created by:        %-15s │\n", updatedUser.CreatedByUserDisplayName)
	fmt.Printf("%s\n\n", info("╰───────────────────────────────────────────╯"))
	// Print password profile details in a formatted box
	fmt.Printf("\n%s\n", info("╭─────────── Password Profile Details ───────────╮"))
	fmt.Printf("│ ID:                %-15d │\n", updatedUser.PasswordProfile.Id)
	fmt.Printf("│ Name:              %-15s │\n", updatedUser.PasswordProfile.Name)
	fmt.Printf("%s\n\n", info("╰────────────────────────────────────────────╯"))

	// Update User Password
	logger.Println("Updating user password...")
	taskState, err := updatedUser.ChangePassword()
	if err != nil {
		logger.Fatalf("Failed to change user password: %v", err)
	}

	// Wait for the task to complete
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger.Println("Waiting for task to complete...")
	if _, err := taskState.CheckTaskState(ctx); err != nil {
		logger.Fatalf("Failed to check task state: %v", err)
	}

	// Check Users Password
	logger.Println("Checking user password...")
	passwordStatus, err := updatedUser.CheckPassword()
	if err != nil {
		logger.Fatalf("Failed to check user password: %v", err)
	}

	// Print password status in a formatted way
	fmt.Printf("\n%s\n", info("╭─────────── Password Status ───────────╮"))
	fmt.Printf("│ Status:    %-25s │\n", passwordStatus.RequestStatus.State)
	fmt.Printf("│ Message:   %-25s │\n", passwordStatus.RequestStatus.Message)
	fmt.Printf("│ Duration:  %-25s │\n", passwordStatus.RequestStatus.TotalDuration)
	fmt.Printf("%s\n", info("╰────────────────────────────────────╯"))

	// Print total execution time
	duration := time.Since(start)
	fmt.Printf("\n%s Total execution time: %s\n", success("✓"), duration.Round(time.Millisecond))
}
