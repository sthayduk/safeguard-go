package main

import (
	"os"

	"github.com/sthayduk/safeguard-go/src/client"
)

// Swagger URL:  https://<applianceHost>/service/core/swagger/index.html
// Swagger JSON: https://<applianceHost>/service/core/swagger/v4/swagger.json

var accessToken string
var applianceUrl string
var apiVersion string

func init() {
	accessToken = os.Getenv("SAFEGUARD_ACCESS_TOKEN")
	applianceUrl = os.Getenv("SAFEGUARD_HOST_URL")
	apiVersion = os.Getenv("SAFEGUARD_API_VERSION")
}

func main() {

	var sgc *client.SafeguardClient

	if accessToken == "" {
		sgc = client.New(applianceUrl, apiVersion, true)
		err := sgc.OauthConnect()
		if err != nil {
			panic(err)
		}
	} else {
		sgc = client.New(applianceUrl, apiVersion, true)
		sgc.AccessToken = &client.TokenResponse{
			AccessToken: accessToken,
		}
	}

	err := sgc.ValidateAccessToken()
	if err != nil {
		panic(err)
	}

}
