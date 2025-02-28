package common

import (
	"fmt"
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
		sgc = safeguard.SetupClient(applianceUrl, apiVersion, false)
		err := sgc.LoginWithOauth()
		if err != nil {
			return err
		}
	} else {
		sgc = safeguard.SetupClient(applianceUrl, apiVersion, false)
		sgc.AccessToken = &client.RSTSAuthResponse{
			AccessToken: accessToken,
		}
	}

	err := sgc.ValidateAccessToken()
	if err != nil {
		return err
	}

	// Get Cluster Leader
	clusterLeader, err := safeguard.GetClusterLeader()
	if err != nil {
		fmt.Println("Failed to get cluster leader:", err)
		return nil
	}

	sgc.SetClusterLeader(clusterLeader.Name)

	return nil
}
