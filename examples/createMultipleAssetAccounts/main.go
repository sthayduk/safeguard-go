package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/sthayduk/safeguard-go/client"
	"github.com/sthayduk/safeguard-go/examples/common"
	"github.com/sthayduk/safeguard-go/models"
)

func main() {
	start := time.Now()
	adId := 464
	samAccountName := "muster"

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
	ad, err := models.GetAsset(sgc, adId, client.Fields{"Id", "Name"})
	if err != nil {
		logger.Fatalf("Failed to get Active Directory: %v", err)
	}

	// Search for the user
	logger.Printf("Searching for user: %s", samAccountName)
	filter := client.Filter{}
	filter.AddFilter("Name", "contains", samAccountName)
	users, err := ad.GetDirectoryAccounts(filter)
	if err != nil {
		logger.Fatalf("Failed to get directory users: %v", err)
	}

	for _, user := range users {
		logger.Printf("Found user: %s", user.Name)
	}

	createdUsers, err := models.CreateAssetAccounts(sgc, users)
	if err != nil {
		logger.Fatalf("Failed to create asset accounts: %s", err)
	}

	var wg sync.WaitGroup
	for _, createdUser := range createdUsers {
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

		// Update and check password in a goroutine
		wg.Add(1)
		go updateAndCheckPassword(&wg, sgc, logger, createdUser)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	fmt.Printf("\n%s All users' passwords successfully updated\n", success("✓"))

	// Print total execution time
	duration := time.Since(start)
	fmt.Printf("\n%s Total execution time: %s\n", success("✓"), duration.Round(time.Millisecond))

}

func updateAndCheckPassword(wg *sync.WaitGroup, sgc *client.SafeguardClient, logger *log.Logger, createdUser models.AssetAccount) {
	// Initialize colored output
	info := color.New(color.FgCyan).SprintFunc()

	// Update Password Profile
	logger.Println("Updating password profile...")
	assetPartition, err := models.GetAssetPartition(sgc, 1, client.Fields{"Id", "Name"})
	if err != nil {
		logger.Fatalf("Failed to get asset partition: %v", err)
	}

	filter := client.Filter{}
	filter.AddFilter("Name", "eq", "ITdesign Profile Suspend")
	passwordProfile, err := models.GetPasswordRules(sgc, assetPartition, filter)
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
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
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
	fmt.Printf("│ User:      %-25s │\n", updatedUser.Name)
	fmt.Printf("│ Status:    %-25s │\n", passwordStatus.RequestStatus.State)
	fmt.Printf("│ Message:   %-25s │\n", passwordStatus.RequestStatus.Message)
	fmt.Printf("│ Duration:  %-25s │\n", passwordStatus.RequestStatus.TotalDuration)
	fmt.Printf("%s\n", info("╰────────────────────────────────────╯"))

	wg.Done()
}
