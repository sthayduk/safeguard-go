package main

import (
	"fmt"
	"os"

	"itdesign.at/safeguard-go/client"
	"itdesign.at/safeguard-go/models"
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

	filter := client.Filter{
		Fields: []string{"Disabled", "DisplayName"},
	}

	filter.AddFilter("Disabled", "eq", "true")

	me, err := models.GetUsers(sgc, filter)
	if err != nil {
		panic(err)
	}

	for _, user := range me {
		jsonStr, err := user.ToJson()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", jsonStr)
	}
}
