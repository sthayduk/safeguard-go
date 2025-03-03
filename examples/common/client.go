package common

import (
	"os"

	safeguard "github.com/sthayduk/safeguard-go"
	"github.com/sthayduk/safeguard-go/client"
)

// InitClient creates and initializes a SafeguardClient using environment variables
func InitClient() error {
	userToken := os.Getenv("SAFEGUARD_USER_TOKEN")
	applianceUrl := os.Getenv("SAFEGUARD_HOST_URL")
	apiVersion := os.Getenv("SAFEGUARD_API_VERSION")
	pfxPassword := os.Getenv("SAFEGUARD_PFX_PASSWORD")
	pfxPath := os.Getenv("SAFEGUARD_PFX_PATH")

	var sgc *client.SafeguardClient

	if userToken == "" {
		sgc = safeguard.SetupClient(applianceUrl, apiVersion, true)
		err := sgc.LoginWithCertificate(pfxPath, pfxPassword)
		if err != nil {
			return err
		}
	} else {
		sgc = safeguard.SetupClient(applianceUrl, apiVersion, true)
		sgc.AccessToken = &client.RSTSAuthResponse{
			UserToken: userToken,
		}
	}

	err := sgc.ValidateAccessToken()
	if err != nil {
		return err
	}

	return nil
}
