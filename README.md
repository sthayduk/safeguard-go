# Safeguard Go
Safeguard Go is a Go library for interacting with the OneIdentity Safeguard for Privileged Passwords REST API.

## Installation

To install the library, use the following command:

```sh
go get github.com/sthayduk/safeguard-go
```

## Prerequisites

The library requires SSL certificates for OAuth authentication:
- `server.crt` and `server.key` - for the local HTTPS callback server
- `pam.cer` - the Safeguard appliance certificate

## Usage

### Initializing the Client

```go
import (
    "github.com/sthayduk/safeguard-go/client"
)

// Create a new client with debug logging enabled
sgc := client.New("https://your-appliance.domain.com", "v4", true)

// Connect using OAuth2.0
err := sgc.OauthConnect()
if err != nil {
    panic(err)
}
```

### Using Access Tokens

You can also initialize the client with an existing access token:

```go
sgc := client.New(applianceUrl, apiVersion, true)
sgc.AccessToken = &client.TokenResponse{
    AccessToken: "your-access-token",
}

// Validate the token
err := sgc.ValidateAccessToken()
if err != nil {
    panic(err)
}
```

### Working with Asset Partitions

```go
// Get all asset partitions
partitions, err := models.GetAssetPartitions(sgc, client.Filter{})
if err != nil {
    panic(err)
}

// Get a specific asset partition
partition, err := models.GetAssetPartition(sgc, "1", client.Fields{})
if err != nil {
    panic(err)
}

// Get password rules for an asset partition
rules, err := partition.GetPasswordRules(sgc)
if err != nil {
    panic(err)
}
```

### Using Filters

```go
// Create a filter for disabled items
filter := client.Filter{
    Fields: []string{"Disabled", "DisplayName"},
}
filter.AddFilter("Disabled", "eq", "true")

// Use the filter in API calls
results, err := models.GetAssetPartitions(sgc, filter)
```

## Environment Variables

The library supports the following environment variables:
- `SAFEGUARD_ACCESS_TOKEN` - Pre-configured access token
- `SAFEGUARD_HOST_URL` - Safeguard appliance URL
- `SAFEGUARD_API_VERSION` - API version to use

## Authentication Flow

The library implements OAuth2.0 authentication with PKCE (Proof Key for Code Exchange):
1. Generates a code verifier and challenge
2. Starts a local HTTPS server to receive the callback
3. Opens the browser for user authentication
4. Exchanges the authorization code for an access token
5. Converts the rSTS token to a Safeguard access token

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
