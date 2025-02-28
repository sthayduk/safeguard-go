package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	safeguard "github.com/sthayduk/safeguard-go"
	"github.com/sthayduk/safeguard-go/client"
)

func main() {
	start := time.Now()

	// Initialize colored output
	success := color.New(color.FgGreen).SprintFunc()

	accessToken := os.Getenv("SAFEGUARD_ACCESS_TOKEN")
	applianceUrl := os.Getenv("SAFEGUARD_HOST_URL")
	apiVersion := os.Getenv("SAFEGUARD_API_VERSION")
	pfxPassword := os.Getenv("SAFEGUARD_PFX_PASSWORD")
	pfxPath := os.Getenv("SAFEGUARD_PFX_PATH")

	var sgc *client.SafeguardClient

	if accessToken == "" {
		sgc = safeguard.SetupClient(applianceUrl, apiVersion, true)
		err := sgc.LoginWithCertificate(pfxPath, pfxPassword, "rsts:sts:primaryproviderid:certificate")
		if err != nil {
			panic(err)
		}
	} else {
		sgc = safeguard.SetupClient(applianceUrl, apiVersion, true)
		sgc.AccessToken = &client.RSTSAuthResponse{
			AccessToken: accessToken,
		}
	}

	err := sgc.ValidateAccessToken()
	if err != nil {
		panic(err)
	}

	// Print total execution time
	duration := time.Since(start)
	fmt.Printf("\n%s Total execution time: %s\n", success("✓"), duration.Round(time.Millisecond))

}
