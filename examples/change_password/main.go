package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/sthayduk/safeguard-go"
	"github.com/sthayduk/safeguard-go/examples/common"
)

func main() {
	start := time.Now()

	// Initialize colored output
	success := color.New(color.FgGreen).SprintFunc()
	info := color.New(color.FgCyan).SprintFunc()
	warning := color.New(color.FgYellow).SprintFunc()

	// Initialize the logger
	logger := log.New(os.Stdout, info("[ChangePassword] "), log.LstdFlags|log.Lshortfile)

	// Initialize the Safeguard client
	logger.Println("Initializing Safeguard client...")
	err := common.InitClient()
	if err != nil {
		logger.Fatalf("%s Failed to initialize client: %v", warning("ERROR:"), err)
	}
	logger.Println("Client initialized successfully.")

	// Retrieve asset account
	assetAccountId := 17
	logger.Printf("Retrieving asset account with ID: %d...", assetAccountId)
	assetAccount, err := safeguard.GetAssetAccount(assetAccountId, safeguard.Fields{})
	if err != nil {
		logger.Fatalf("%s Failed to retrieve asset account: %v", warning("ERROR:"), err)
	}

	// Print asset account details
	fmt.Printf("\n%s\n", info("╭─────────── Asset Account Details ───────────╮"))
	fmt.Printf("│ ID:                %-15d │\n", assetAccount.Id)
	fmt.Printf("│ Name:              %-15s │\n", assetAccount.Name)
	fmt.Printf("│ Asset:             %-15s │\n", assetAccount.Asset.Name)
	fmt.Printf("%s\n\n", info("╰───────────────────────────────────────────╯"))

	// Initiate password change task
	logger.Println("Initiating password change task...")
	changePasswordTask, err := assetAccount.ChangePassword()
	if err != nil {
		logger.Fatalf("%s Failed to initiate password change task: %v", warning("ERROR:"), err)
	}
	logger.Println("Password change task initiated successfully.")

	// Wait for the task to complete
	logger.Println("Waiting for the password change task to complete...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	state, err := changePasswordTask.CheckTaskState(ctx)
	if err != nil {
		logger.Fatalf("%s Failed to check task state: %v", warning("ERROR:"), err)
	}

	// Print task state details
	fmt.Printf("\n%s\n", info("╭─────────── Task State Details ───────────╮"))
	fmt.Printf("│ Status:            %-15t │\n", state)
	fmt.Printf("%s\n\n", info("╰───────────────────---────────────────────╯"))

	// Print total execution time
	duration := time.Since(start)
	fmt.Printf("\n%s Total execution time: %s\n", success("✓"), duration.Round(time.Millisecond))
}
