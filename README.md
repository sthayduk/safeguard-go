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

- Authentication
  - Username/Password authentication
  - OAuth authentication
  - Multiple authentication provider support
- Access Requests
  - Create single and batch access requests
  - Check out passwords
  - Check in access requests
  - Cancel access requests
  - Close access requests based on state
  - Handle emergency access
  - Review and approve requests
  - Monitor request states and sessions
- Me (Current User)
  - Get current user details
  - Get accessible assets and accounts
  - Get actionable requests by role
  - Get account entitlements
  - Request access to accounts
- Users
  - Get users and user details
  - Get linked accounts
  - Get user roles and groups
  - Get user preferences
  - Delete users
  - Link and Unlink PolicyAccounts
- Identity Providers
  - Get providers and details
  - Get directory users
  - Get directory groups
  - Support for multiple provider types (LDAP, RADIUS, etc.)
- User Groups
  - Get groups and details
  - Directory properties support
- Asset Management
  - Asset Partitions
    - Get partitions and details
    - Get password rules
  - Asset Groups
    - Get asset groups and details
    - Dynamic grouping rules
  - Asset Accounts
    - Get accounts and details
    - Password operations (check/change)
  - Policy Assets
    - Get policy assets
    - Asset policy management
- Reports
  - Account task schedules
  - Task execution history

## Usage

### Authentication

The client supports multiple authentication methods:

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

// Or login with OAuth
err := sgc.OauthConnect()
if err != nil {
    panic(err)
}
```

### Working with Users

```go
import (
    "github.com/sthayduk/safeguard-go/models"
    "github.com/sthayduk/safeguard-go/client"
)

// Get all users
users, err := models.GetUsers(sgc, client.Filter{})

// Get a specific user
user, err := models.GetUser(sgc, userId, client.Fields{"Name", "Description"})

// Get user's linked accounts
accounts, err := user.GetLinkedAccounts()

// Get user's roles
roles, err := user.GetRoles()

// Get user's groups
groups, err := user.GetGroups()

// Delete a user
err = user.Delete()
```

### Working with Identity Providers

```go
// Get all identity providers
providers, err := models.GetIdentityProviders(sgc)

// Get specific provider
provider, err := models.GetIdentityProvider(sgc, providerId)

// Get directory users from provider
users, err := provider.GetDirectoryUsers(client.Filter{})

// Get directory groups from provider
groups, err := provider.GetDirectoryGroups(client.Filter{})
```

### Working with Asset Accounts

```go
// Get all asset accounts
accounts, err := models.GetAssetAccounts(sgc, client.Filter{})

// Get specific account
account, err := models.GetAssetAccount(sgc, accountId, client.Fields{})

// Check password
log, err := account.CheckPassword()

// Change password
log, err := account.ChangePassword()
```

### Working with Current User

```go
// Get current user's actionable requests
requests, err := models.GetMeActionableRequests(sgc, client.Filter{})

// Get requests for specific role
requests, err := models.GetMeActionableRequestsByRole(sgc, models.ApproverRole, client.Filter{})

// Get detailed actionable requests with helper methods
result, err := models.GetMeActionableRequestsDetailed(sgc, client.Filter{})

// Get pending requests
pending := result.GetPendingRequests()

// Filter requests by state
available := result.FilterRequestsByState(models.StateRequestAvailable)

// Get account entitlements
entitlements, err := models.GetMeAccountEntitlements(sgc, 
    models.PasswordEntitlement,
    true,  // includeActiveRequests
    false, // filterByCredential
    client.Filter{})

// Get accessible assets
assets, err := models.GetMeAccessRequestAssets(sgc, client.Filter{})
```

## Query Parameters

The API supports filtering and field selection:

```go
// Create a filter
filter := client.Filter{}
filter.AddFilter("Disabled", "eq", "true")
filter.AddFilter("Name", "like", "admin")

// Specify fields to return
fields := client.Fields{"Name", "Description", "CreatedDate"}

// Use in API calls
users, err := models.GetUsers(sgc, filter)
user, err := models.GetUser(sgc, userId, fields)
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.
