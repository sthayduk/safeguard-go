package common

import (
	"os"

	"github.com/sthayduk/safeguard-go/client"
)

// InitClient creates and initializes a SafeguardClient using environment variables
func InitClient() (*client.SafeguardClient, error) {
	accessToken := os.Getenv("SAFEGUARD_ACCESS_TOKEN")
	applianceUrl := os.Getenv("SAFEGUARD_HOST_URL")
	apiVersion := os.Getenv("SAFEGUARD_API_VERSION")

	var sgc *client.SafeguardClient

	if accessToken == "" {
		sgc = client.New(applianceUrl, apiVersion, false)
		err := sgc.OauthConnect()
		if err != nil {
			return nil, err
		}
	} else {
		sgc = client.New(applianceUrl, apiVersion, false)
		sgc.AccessToken = &client.TokenResponse{
			AccessToken: accessToken,
		}
	}

	err := sgc.ValidateAccessToken()
	if err != nil {
		return nil, err
	}

	return sgc, nil
}
