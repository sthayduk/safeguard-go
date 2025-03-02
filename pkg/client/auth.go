package client

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// LoginWithPassword authenticates a user using their username and password.
// It first obtains an RSTS token and then exchanges it for a Safeguard token.
// The obtained token is stored in the SafeguardClient's AccessToken field.
// Parameters:
// - username: The username of the user.
// - password: The password of the user.
// Returns an error if the login process fails, otherwise nil.
func (c *SafeguardClient) LoginWithPassword(username, password string) error {
	if c.AccessToken == nil {
		c.AccessToken = &RSTSAuthResponse{
			AuthProvider: AuthProviderLocal,
		}
	} else {
		c.AccessToken.AuthProvider = AuthProviderLocal
	}

	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", username)
	data.Set("password", password)
	data.Set("scope", AuthProviderLocal.String())

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/RSTS/oauth2/token", c.Appliance.getUrl()), strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("RSTS login request failed: %v", err)
	}
	defer resp.Body.Close()

	err = c.handleTokenResponse(resp)
	if err != nil {
		return fmt.Errorf("RSTS login failed: %v", err)
	}

	err = c.exchangeRSTSTokenForSafeguard(c.HttpClient)
	if err != nil {
		return fmt.Errorf("token exchange failed: %v", err)
	}

	c.AccessToken.setUserNamePassword(username, password)
	fmt.Println("✅ Login successful")
	return nil
}
