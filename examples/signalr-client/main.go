package main

import (
	"context"
	"fmt"
	"os"

	safeguard "github.com/sthayduk/safeguard-go"
)

func main() {
	applianceUrl := os.Getenv("SAFEGUARD_HOST_URL")
	apiVersion := os.Getenv("SAFEGUARD_API_VERSION")
	pfxPassword := os.Getenv("SAFEGUARD_PFX_PASSWORD")
	pfxPath := os.Getenv("SAFEGUARD_PFX_PATH")

	sgc := safeguard.SetupClient(applianceUrl, apiVersion, true)
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

	eventHandler := safeguard.SetupSignalRClient(sgc)

	ctx := context.Background()
	err = eventHandler.Run(ctx)
	if err != nil {
		fmt.Println("Error starting SignalR client:", err)
		return
	}
}
