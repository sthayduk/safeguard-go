# Safeguard Go
A Go library for interacting with the OneIdentity Safeguard for Privileged Passwords REST API.

## Installation

```sh
go get github.com/sthayduk/safeguard-go
```

## Prerequisites

The library requires SSL certificates for OAuth authentication:
- `server.crt` and `server.key` - for the local HTTPS callback server
- `pam.cer` - the Safeguard appliance certificate

## Features

Currently supports the following Safeguard resources:

- Users
  - Get users and user details
  - Get linked accounts
  - Get user roles
  - Get user groups
- Asset Groups
  - Get asset groups and details
- Asset Partitions
  - Get partitions and details
  - Get password rules
- Policy Assets
  - Get policy assets and details
  - Get asset groups
  - Get directory service entries
  - Get policies
- Policy Accounts
  - Get policy accounts and details
- Asset Accounts
  - Get accounts and details
- Reports
  - Get account task schedules
- Roles
  - Get roles and details

## Usage

### Authentication

The client supports two authentication methods:

#### Username/Password Authentication

```go
import (
    "github.com/sthayduk/safeguard-go/client"
)

// Create a new client
sgc := client.New("https://your-appliance.domain.com", "v4", true)

// Login with username/password
err := sgc.LoginWithPassword("username", "password")
if err != nil {
    panic(err)
}
```

#### OAuth Authentication

```go
import (
    "github.com/sthayduk/safeguard-go/client"
)

// Create a new client
sgc := client.New("https://your-appliance.domain.com", "v4", true)

// Login with OAuth (requires SSL certificates in prerequisites)
err := sgc.OauthConnect()
if err != nil {
    panic(err)
}
```

### Working with Users

```go
import "github.com/sthayduk/safeguard-go/models"

// Get all users
users, err := models.GetUsers(sgc, client.Filter{})

// Get a specific user
user, err := models.GetUser(sgc, userId, client.Fields{"Name", "Description"})

// Get user's linked accounts
accounts, err := user.GetLinkedAccounts(sgc)

// Get user's roles
roles, err := user.GetRoles(sgc)
```

### Working with Asset Groups

```go
// Get all asset groups
groups, err := models.GetAssetGroups(sgc, client.Filter{})

// Get specific asset group details
group, err := models.GetAssetGroup(sgc, groupId, client.Fields{"Name", "Description"})
```

## Filters and Fields

The API supports filtering and field selection:

```go
// Create a filter
filter := client.Filter{}
filter.AddFilter("Disabled", "eq", "true")

// Specify fields to return
fields := client.Fields{"Name", "Description", "CreatedDate"}

// Use in API calls
assets, err := models.GetPolicyAssets(sgc, filter)
asset, err := models.GetPolicyAsset(sgc, assetId, fields)
```

## Example Code

See the `examples/` directory for complete usage examples.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.
