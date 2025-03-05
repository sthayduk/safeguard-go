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
  - Certificate-based authentication
  - Automatic token refresh
  - Multiple authentication provider support
  - OAuth Connect with callback server
- Client Management
  - Thread-safe appliance URL handling with caching
  - TLS client configuration
  - Cluster leader discovery and management
  - Token expiration tracking
- Access Requests
  - Create single and batch access requests
  - Check out passwords with timeout support
  - Check in access requests
  - Cancel access requests
  - Close access requests based on state
  - Monitor request states (pending, valid, invalid)
  - Support for reason codes and comments
  - Session information tracking
- Me (Current User)
  - Get current user details
  - Get accessible assets and accounts
  - Get actionable requests by role
  - Get account entitlements
  - Request access to accounts
  - Get preferences
  - Get available approval/review requests
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
  - Assets
    - Get assets and details
    - Platform configuration
    - Connection properties
    - Session access properties
    - Directory properties
    - Task scheduling and history
  - Asset Partitions
    - Get partitions and details
    - Get password rules
    - Manage partition owners
  - Asset Groups
    - Get asset groups and details
    - Dynamic grouping rules
    - Tag-based grouping
  - Asset Accounts
    - Create Asset Accounts
    - Get accounts and details
    - Password operations (check/change)
    - SSH key management
    - Account discovery
    - Synchronization groups
    - Task scheduling
    - Enable/disable accounts
    - Suspend/restore accounts
  - Policy Assets
    - Get policy assets
    - Asset policy management
    - SSH host key verification
    - Session access configuration
- Cluster Management
  - Get cluster members
  - Get cluster leader
  - Update cluster leader URL
  - Monitor cluster health
  - Force health checks
  - Network configuration
- Reports
  - Account task schedules
  - Task execution history
- Event Handling
  - Real-time event notifications via SignalR
  - Access Request event monitoring
  - Event data processing
  - Automatic reconnection with backoff
  - Context-based cancellation and shutdown

## Usage

### Authentication and Client Setup

The SafeguardClient handles authentication and API communication:

```go
import (
    safeguard "github.com/sthayduk/safeguard-go"
)

// Create a new client with debug logging
client := safeguard.NewClient("https://your-appliance.domain.com", "v4", true)

// Login with username/password
err := client.LoginWithPassword("username", "password")
if err != nil {
    panic(err)
}

// Login with certificate
err := client.LoginWithCertificate("path/to/cert.pem", "certPassword")
if err != nil {
    panic(err)
}

// Check token expiration
if client.IsTokenExpired() {
    // Handle expired token
}

// Get remaining token time
remainingTime := client.RemainingTokenTime()
```

### Working with Access Requests

```go
// Get access requests with filtering
filter := safeguard.Filter{}
filter.AddFilter("State", "eq", "Available")
requests, err := client.GetAccessRequests(filter)

// Get a specific access request
request, err := client.GetAccessRequest(requestId, nil)

// Create new access requests in batch
responses, err := client.NewAccessRequests(accountEntitlements)

// Check out a password with context and waiting for pending approval
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()
password, err := request.CheckOutPassword(ctx, true)

// Check in a request
updated, err := request.CheckIn()

// Cancel a request
updated, err := request.Cancel()

// Close a request (automatically handles check-in or cancel based on state)
updated, err := request.Close()

// Check request state
if request.IsPending() {
    fmt.Println("Request is pending approval")
}
if request.IsValid() {
    fmt.Println("Request is valid for checkout")
}
if request.IsInvalid() {
    fmt.Println("Request is in invalid state")
}

// Refresh request state from server
updated, err := request.RefreshState()
```

### Working with Users

```go
import (
    safeguard "github.com/sthayduk/safeguard-go"
)

// Get all users
users, err := safeguard.GetUsers(safeguard.Filter{})

// Get a specific user
user, err := safeguard.GetUser(userId, safeguard.Fields{"Name", "Description"})

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
import (
    safeguard "github.com/sthayduk/safeguard-go"
)

// Get all identity providers
providers, err := safeguard.GetIdentityProviders(client)

// Get specific provider
provider, err := safeguard.GetIdentityProvider(providerId)

// Get directory users from provider
users, err := provider.GetDirectoryUsers(safeguard.Filter{})

// Get directory groups from provider
groups, err := provider.GetDirectoryGroups(safeguard.Filter{})
```

### Working with Asset Accounts

```go
import (
    safeguard "github.com/sthayduk/safeguard-go"
)

// Get all asset accounts
accounts, err := safeguard.GetAssetAccounts(safeguard.Filter{})

// Get specific account
account, err := safeguard.GetAssetAccount(accountId, safeguard.Fields{})

// Check password
log, err := account.CheckPassword()

// Change password
log, err := account.ChangePassword()
```

### Working with Current User

```go
import (
    safeguard "github.com/sthayduk/safeguard-go"
)

// Get current user's actionable requests
requests, err := safeguard.GetMeActionableRequests(safeguard.Filter{})

// Get requests for specific role
requests, err := safeguard.GetMeActionableRequestsByRole(safeguard.ApproverRole, safeguard.Filter{})

// Get detailed actionable requests with helper methods
result, err := safeguard.GetMeActionableRequestsDetailed(safeguard.Filter{})

// Get pending requests
pending := result.GetPendingRequests()

// Filter requests by state
available := result.FilterRequestsByState(safeguard.StateRequestAvailable)

// Get account entitlements
entitlements, err := safeguard.GetMeAccountEntitlements(
    safeguard.PasswordEntitlement,
    true,  // includeActiveRequests
    false, // filterByCredential
    safeguard.Filter{})

// Get accessible assets
assets, err := safeguard.GetMeAccessRequestAssets(safeguard.Filter{})
```

### Working with Real-time Events

The library provides real-time event notification capabilities using SignalR:

```go
import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "syscall"
    
    safeguard "github.com/sthayduk/safeguard-go"
)

// Create a new event handler
eventHandler := client.NewSignalRClient()

// Create a context with cancellation
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// Set up signal handling for graceful shutdown
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

// Start the event handler in a goroutine
go func() {
    if err := eventHandler.Run(ctx); err != nil {
        fmt.Printf("SignalR error: %v\n", err)
    }
}()

// Process events as they arrive
for {
    select {
    case event := <-eventHandler.EventChannel:
        fmt.Printf("Received event: %s\n", event.Message)
        fmt.Printf("Access Request Type: %+v\n", event.Data.AccessRequestType)
        // Process different event types
        switch event.Data.EventName {
        case "AccessRequestCreated":
            // Handle new access request
        case "AccessRequestStatusChanged":
            // Handle status change
        }
    case sig := <-sigChan:
        fmt.Printf("Received signal: %v\n", sig)
        cancel() // Gracefully shut down
        return
    case <-ctx.Done():
        fmt.Println("Context cancelled, shutting down...")
        return
    }
}
```

The `SignalREvent` structure provides detailed information about events:

```go
type SignalREvent struct {
    ApplianceId string
    Name        string
    Time        time.Time
    Message     string
    AuditLogUri *string
    Data        EventData // Contains detailed event information
}

type EventData struct {
    AccessRequestType       AccessRequestType
    AccountName             string
    AssetName               string
    Requester               string
    RequestId               string
    EventName               string
    EventTimestamp          time.Time
    EventUserDisplayName    string
    // Many more fields available
}
```

# Safeguard Filter Examples

## Overview

The Safeguard API uses a query string-based filtering system to refine API requests. 
The `safeguard.Filter` type provides a convenient way to construct these filters.

## Examples

The `main.go` file demonstrates:

1. **Basic Filters**: Creating filters with fields, ordering, and count options
2. **Filter Conditions**: Adding equality, comparison, and text search conditions
3. **Complex Search**: Creating complex search filters across multiple fields
4. **API Integration**: Using filters with the Safeguard API
5. **Asset Filtering**: Searching for assets by name patterns

## Key Filter Features

- Field selection (`filter.AddField()`)
- Sorting (`filter.AddOrderBy()`)
- Simple conditions (`filter.AddFilter()`)
- Complex searches (`filter.AddComplexSearchFilter()`)
- Standard search patterns (`filter.AddSearchFilter()`)

## Running the Examples

Ensure you have set up your client configuration in the common package, then run:

```bash
go run main.go
```

## Filter Operators

The library provides constants for all supported filter operators:

- `OpEqual`, `OpNotEqual` - Equality operators (eq, ne)
- `OpGreaterThan`, `OpGreaterThanOrEqual`, `OpLessThan`, `OpLessThanOrEqual` - Comparison operators (gt, ge, lt, le)
- `OpContains`, `OpIContains` - Text search operators (contains, icontains)
- `OpStartsWith`, `OpIStartsWith`, `OpEndsWith`, `OpIEndsWith` - Pattern matching (sw, isw, ew, iew)
- `OpAnd`, `OpOr`, `OpNot` - Logical operators (and, or, not)
- `OpIn` - Collection operator (in)

## Sample Query Output

A basic filter might produce a query string like:
```
?fields=Name,Description&filter=Name eq 'Administrator'&count=true&orderby=-CreatedDate
```

A complex search filter might produce:
```
?filter=(Name contains 'server' or NetworkAddress contains 'server' or Description contains 'server')
```


## Example Access Request Workflow

```go
import (
    safeguard "github.com/sthayduk/safeguard-go"
)

// Get user information
me, err := client.GetMe()
if err != nil {
    panic(err)
}
fmt.Printf("Logged in as: %s\n", me.Name)

// Get account entitlements
entitlements, err := client.GetMeAccountEntitlements()
if err != nil {
    panic(err)
}

// Create a new access request
responses, err := client.NewAccessRequests(entitlements)
if err != nil {
    panic(err)
}

// Check out password from the first successful request
for _, response := range responses {
    if !response.hasError() {
        password, err := response.AccessRequest.CheckOutPassword(context.Background(), true)
        if err != nil {
            fmt.Println("Error checking out password:", err)
            continue
        }
        fmt.Println("Password:", password)
        
        // Close the request when done
        _, err = response.AccessRequest.Close()
        if err != nil {
            fmt.Println("Error closing request:", err)
        }
        break
    }
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.
