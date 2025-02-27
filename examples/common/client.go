package common

import (
	"os"

	safeguard "github.com/sthayduk/safeguard-go"
	"github.com/sthayduk/safeguard-go/client"
)

// InitClient creates and initializes a SafeguardClient using environment variables
func InitClient() error {
	accessToken := os.Getenv("SAFEGUARD_ACCESS_TOKEN")
	applianceUrl := os.Getenv("SAFEGUARD_HOST_URL")
	apiVersion := os.Getenv("SAFEGUARD_API_VERSION")

	var sgc *client.SafeguardClient

	if accessToken == "" {
		sgc = safeguard.SetupClient(applianceUrl, apiVersion, true)
		err := sgc.LoginWithOauth()
		if err != nil {
			return err
		}
	} else {
		sgc = safeguard.SetupClient(applianceUrl, apiVersion, true)
		sgc.AccessToken = &client.TokenResponse{
			AccessToken: accessToken,
		}
	}

	err := sgc.ValidateAccessToken()
	if err != nil {
		return err
	}

	return nil
}
