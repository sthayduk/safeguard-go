# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

A Go client library for the OneIdentity **Safeguard for Privileged Passwords** REST API (PAM appliance). Single flat package `safeguard` at the repo root; `module github.com/sthayduk/safeguard-go`, Go 1.24.

## Commands

```sh
go build ./...                      # build the library
go test ./... --race                # full test suite (CI runs exactly this)
go test -run TestName -v            # run a single test
go vet ./...
```

CI (`.github/workflows/go.yml`) builds and runs `go test -v ./... --race` on ubuntu, windows, and macos with Go 1.24. The `--race` flag matters: the client and auth structs use `sync.RWMutex` heavily and concurrency bugs must surface in CI.

### Running examples

`examples/` is a **separate Go module** (`examples/go.mod`) with a `replace github.com/sthayduk/safeguard-go => ../` directive. Run from inside `examples/`:

```sh
cd examples && go run ./me           # or ./accessRequests, ./cluster, etc.
```

Examples are configured entirely via environment variables (see `.vscode/launch.json` for ready-made configs):
- `SAFEGUARD_HOST_URL`, `SAFEGUARD_API_VERSION` (e.g. `v4`)
- Either `SAFEGUARD_USER_TOKEN` (use an existing token) **or** `SAFEGUARD_PFX_PATH` + `SAFEGUARD_PFX_PASSWORD` (certificate login)

`examples/common.InitClient()` reads these env vars and returns a logged-in client. Note the examples module pins older dependency versions than the root — they are intentionally independent.

## Architecture

### Two API surfaces, one source of truth

Every resource operation exists as a method on `*SafeguardClient` (e.g. `client.GetUsers(filter)`, `client.CreateUser(u)`). Entity types *also* expose fluent methods (`user.GetRoles()`, `request.Close()`, `provider.Synchronize()`) — these are thin wrappers that delegate back to the client method. **Put the real logic in the `*SafeguardClient` method; the entity method should just call it.**

> The README's usage snippets sometimes show package-level calls like `safeguard.GetUsers(...)`. That is outdated — the real API is client methods (`client.GetUsers(...)`). Trust the code, not the README, for call signatures.

### How entities call back into the client (`utils.go`)

Entity structs carry an unexported `apiClient *SafeguardClient` field so their fluent methods can issue requests. This back-reference is injected, never set by hand:

- Types implement the `ClientHolder` interface via `func (x Type) SetClient(c *SafeguardClient) any { x.apiClient = c; return x }`.
- After unmarshaling an API response, wrap the result with `addClient(client, v)` (single) or `addClientToSlice(client, v)` (slice) so the returned entities can chain further calls.

When adding a new resource type, implement `SetClient` and route every decoded entity through `addClient`/`addClientToSlice`, or its fluent methods will panic on a nil `apiClient`.

### HTTP layer (`request.go`) — read vs. write routing

All traffic goes through `GetRequest` / `PostRequest` / `PutRequest` / `DeleteRequest` → `sendHttpRequest`. URL selection is deliberate:

- **Reads** use `getReadOnlyRootUrl()` → the **appliance** URL.
- **Writes** (PUT/DELETE) use `getReadWriteRootUrl()` → the **cluster leader** URL (discovered and cached in `client.go`).
- **POST is special**: it tries the read-only (appliance) URL first and falls back to the read-write (cluster leader) URL on failure. This is an intentional workaround for customers who firewall the cluster leader behind a load balancer reachable only via the appliance URL — preserve this fallback.

`sendHttpRequest` treats only 200/201/202 as success; anything else becomes an error containing the response body.

### Cluster awareness (`client.go`)

The client tracks two URLs (`Appliance`, `ClusterLeader`) as `applianceURL` values — thread-safe, self-caching structs (default 3600s TTL, `-1` = infinite). The cluster leader is resolved lazily via `Cluster/Members?filter=IsLeader eq true` and refreshed when its cache expires. Write routing depends on this being correct.

### Authentication & token refresh

`LoginWithPassword` and `LoginWithCertificate` populate `AccessToken *RSTSAuthResponse` (also `sync.RWMutex`-guarded). `NewClient` starts a background `refreshToken` goroutine that blocks on the `authDone` channel until first login, then re-authenticates ~1 minute before expiry using the stored credentials (the original password or cert path/password are kept in `RSTSAuthResponse.credentials` precisely so refresh can re-login). OAuth/Connect flows use a local HTTPS callback server on port 8400.

### TLS trust (`createTLSClient` / `loadCertificates`)

The HTTP client starts from the OS system cert pool (`x509.SystemCertPool`) and **additionally appends every certificate file found in the current working directory** (`.crt`, `.cer`, `.pem`, `.der`, `.p7b`, `.pfx`, etc.). This is why `pam.cer` (the PAM appliance root chain) sits in the repo/examples root — it must be in the process CWD to be trusted. macOS limits appliance host certs to ≤398 days for the SignalR connection (Apple TLS policy).

### Real-time events (`signalr.go`)

`client.NewSignalRClient()` returns an `EventHandler` whose `Run(ctx)` streams Safeguard events onto `EventChannel`, with automatic reconnect/backoff (`cenkalti/backoff`) and context-based shutdown.

### Safe logging (`request.go`)

The package logs requests/responses through a global `slog` logger (`SetLogger`/`GetLogger`). Sensitive data is masked via `slog.LogValuer` wrappers — `SafeHeaders` masks `Authorization` bearer tokens and `SafeResponseBody` masks password-checkout response bodies. **Never log raw headers or raw response bodies for auth/password endpoints; wrap them with `NewSafeHeaders` / `NewSafeResponseBody`.**

### Query building (`filter.go`)

`Filter` and `Fields` build Safeguard's query-string filter syntax. Use the `Op*` operator constants (`OpEqual`, `OpContains`, `OpAnd`, …) and helpers like `AddFilter`, `AddOrderBy`, `AddSearchFilter`, `AddComplexSearchFilter` rather than hand-building query strings.

## Conventions

- Resources are organized one file per domain (`users.go`, `assets.go`, `accessRequest.go`, `idp.go`, …). Match the existing file when extending a resource.
- Exported types and methods carry full GoDoc comments (Parameters/Returns sections). New public API should follow this style.
- Do not commit secrets, certificates, or tokens — `.gitignore` excludes `*.crt`, `*.cer`, `*.pem`, `*.key`, `*.pfx`, `*.env`, and `*.json`. (`.vscode/launch.json` currently contains real-looking demo tokens; do not add more, and don't treat those tokens as valid.)
```
