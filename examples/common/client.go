package common

import (
	"os"

	"github.com/sthayduk/safeguard-go"
)

// InitClient creates and initializes a SafeguardClient using environment variables
func InitClient() (*safeguard.SafeguardClient, error) {
	userToken := os.Getenv("SAFEGUARD_USER_TOKEN")
	applianceUrl := os.Getenv("SAFEGUARD_HOST_URL")
	apiVersion := os.Getenv("SAFEGUARD_API_VERSION")
	pfxPassword := os.Getenv("SAFEGUARD_PFX_PASSWORD")
	pfxPath := os.Getenv("SAFEGUARD_PFX_PATH")

	var sgc *safeguard.SafeguardClient

	if userToken == "" {
		sgc = safeguard.NewClient(applianceUrl, apiVersion, false)
		err := sgc.LoginWithCertificate(pfxPath, pfxPassword)
		if err != nil {
			return nil, err
		}
	} else {
		sgc = safeguard.NewClient(applianceUrl, apiVersion, false)
		sgc.AccessToken = &safeguard.RSTSAuthResponse{
			UserToken: userToken,
		}
	}

	err := sgc.ValidateAccessToken()
	if err != nil {
		return nil, err
	}

	return sgc, nil
}
