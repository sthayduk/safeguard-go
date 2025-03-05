package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sthayduk/safeguard-go"
)

func main() {
	applianceUrl := os.Getenv("SAFEGUARD_HOST_URL")
	apiVersion := os.Getenv("SAFEGUARD_API_VERSION")
	pfxPassword := os.Getenv("SAFEGUARD_PFX_PASSWORD")
	pfxPath := os.Getenv("SAFEGUARD_PFX_PATH")

	sgc := safeguard.NewClient(applianceUrl, apiVersion, false)
	err := sgc.LoginWithCertificate(pfxPath, pfxPassword)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = sgc.ValidateAccessToken()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	eventHandler := sgc.NewSignalRClient()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start event handler in a goroutine
	go func() {
		if err := eventHandler.Run(ctx); err != nil {
			fmt.Printf("SignalR error: %v\n", err)
		}
	}()

	// Event processing loop
	for {
		select {
		case event := <-eventHandler.EventChannel:
			fmt.Printf("Received event     : %s\n", event.Message)
			fmt.Printf("Access Request Type: %+v\n", event.Data.AccessRequestType)
		case sig := <-sigChan:
			fmt.Printf("Received signal: %v\n", sig)
			cancel()
			return
		case <-ctx.Done():
			fmt.Println("Context cancelled, shutting down...")
			return
		}
	}
}
