package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/sthayduk/safeguard-go"
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

	var sgc *safeguard.SafeguardClient

	if accessToken == "" {
		sgc = safeguard.NewClient(applianceUrl, apiVersion, true)
		err := sgc.LoginWithCertificate(pfxPath, pfxPassword)
		if err != nil {
			panic(err)
		}
	} else {
		sgc = safeguard.NewClient(applianceUrl, apiVersion, true)
		sgc.AccessToken = &safeguard.RSTSAuthResponse{
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
