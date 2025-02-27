package pkg

import "github.com/sthayduk/safeguard-go/client"

var c *client.SafeguardClient

func SetupClient(url, version string, insecure bool) *client.SafeguardClient {
	c = client.New(url, version, insecure)
	return c
}
