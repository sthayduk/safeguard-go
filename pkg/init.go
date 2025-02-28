package pkg

import "github.com/sthayduk/safeguard-go/client"

var c *client.SafeguardClient

// SetupClient initializes a new SafeguardClient with the provided configuration.
//
// Parameters:
//   - url: The base URL of the Safeguard API endpoint
//   - version: The API version to use (e.g., "v3")
//   - insecure: If true, skips SSL certificate validation
//
// Returns:
//   - *client.SafeguardClient: The initialized client instance that will be used globally
func SetupClient(url, version string, insecure bool) *client.SafeguardClient {
	c = client.New(url, version, insecure)
	return c
}
