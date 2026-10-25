# Exsited SDK for Go

Go SDK for the Exsited v4 REST API. Handles OAuth2 with a cross-process
token cache, retry with backoff, connection pooling, TLS and proxy configuration,
content-type decoding, an `errors.Is`-classifiable error type, secret redaction and
terminal diagnostics.

**No key is renamed and no case is converted.** The API owns its shapes. A response
comes back as it arrived: `response.JSON()` is the server's own body text byte for byte,
and `response.Map` holds `map[string]any` / `[]any` with JSON numbers kept exact
(`json.Number`), so integer ids never decay to floats. On the way out, a raw JSON string
is sent exactly as you wrote it; a map or a DTO is encoded by the SDK.

A **typed view** sits over that raw path as a convenience, never as a replacement: a body
can be read as a map, as the server's own JSON text, or as a DTO, and a request body can
be written any of those three ways too. With nothing configured, `response.Data` is the
decoded body. See [Responses](#responses) and [Response format](#response-format).

```go
package main

import (
	"context"
	"fmt"
	"log"

	exsited "github.com/exsited/exsited-sdk-go/v4"
	"github.com/exsited/exsited-sdk-go/v4/resources/accounts"
)

func main() {
	client, err := exsited.NewClient(exsited.SdkConfig{
		Auth: exsited.AuthConfig{
			ClientName:   "default",
			BaseURL:      "https://your-instance.exsited.com",
			ClientID:     "client-id",
			ClientSecret: "client-secret",
			RedirectURI:  "https://your-app.example/callback",
		},
		RaiseOnError: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	response, err := client.Accounts().GetAccount(context.Background(), "W13ZGY", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Map) // the body as it arrived
	account := response.DTO().(*accounts.AccountRecord)
	fmt.Println(*account.Name) // the typed view, envelope already unwrapped
}
```

`id`, not `uuid`: a response carries both, and only `id` addresses the record in a path.
The uuid answers `404`; filter on `filter_uuid` instead.

A failure comes back as a response rather than an error, which is right for application
code and wrong for a script that must stop. `RaiseOnError: true` above flips it, so the
`err` after a call is non-nil on any failure. "Failure" is `Succeeded()` being false, which
is a little wider than "non-2xx": a v4 body with `"object": "error"` under a `2xx` counts
too, see [A 2xx that denies itself](#a-2xx-that-denies-itself).

`Close` is **terminal**. It waits for the sends in flight, releases the `LogFile` handle and
the pooled connections of a transport the client built for itself, and every later call on
that client is refused with `exsited.ErrUsage` before any I/O. A transport you passed in
with `exsited.WithTransport(...)` stays yours and stays open. Closing twice is a no-op.

| | |
|---|---|
| Module / import | `github.com/exsited/exsited-sdk-go/v4` / `exsited` |
| Go | 1.24+, the floor `go.mod` declares; built on the `toolchain` line of the same file, currently 1.26.8 |
| Dependencies | none, standard library only |
| Install | `go get github.com/exsited/exsited-sdk-go/v4`, see [Install](#install) |
| Version | `4.0.0` (`exsited.Version()`), read from the build's module metadata |
| Endpoints | 865 in 91 resource families, documented at [exsited.mintlify.app](https://exsited.mintlify.app/introduction) |

## Contents

- [Install](#install)
- [Credentials](#credentials)
- [Configuration](#configuration)
- [Multiple instances](#multiple-instances)
- [Authentication](#authentication)
- [Requests](#requests)
- [Pagination](#pagination)
- [Responses](#responses)
  - [Response format](#response-format)
  - [A 2xx that denies itself](#a-2xx-that-denies-itself)
- [Errors](#errors)
- [Diagnostics](#diagnostics)
- [Retries](#retries)
  - [Retry only what certainly did not arrive](#retry-only-what-certainly-did-not-arrive)
  - [Retrying a status is scoped by method](#retrying-a-status-is-scoped-by-method)
  - [Timeouts](#timeouts)
- [Transport](#transport)
- [Endpoints](#endpoints)
- [Architecture](#architecture)
  - [The layer map](#the-layer-map)
  - [Where the value types live](#where-the-value-types-live)
  - [The object graph](#the-object-graph)
- [Stable formats](#stable-formats)
  - [Cancellation](#cancellation)
  - [Behaviour worth knowing](#behaviour-worth-knowing)
- [Troubleshooting](#troubleshooting)
- [Public API](#public-api)
- [License](#license)

## Install

```sh
go get github.com/exsited/exsited-sdk-go/v4
```

The module path ends in `/v4`, the major version, as Go requires above the first major,
so the `/v4` belongs in the `go get` line and in every import. `go get` without it finds
no 4.x release.

**Prerequisites.** Go 1.24 or newer, the floor `go.mod` declares:

```sh
go version   # go1.24 or newer; go.mod's toolchain line names the release build, 1.26.8
```

With `GOTOOLCHAIN=auto`, the default, an older local Go switches to 1.26.8 for this module,
downloading it once; `GOTOOLCHAIN=local` builds on the installed Go instead. Nothing else
is needed: the SDK depends on the standard library alone.

Import the root package, which re-exports the whole public surface:

```go
import exsited "github.com/exsited/exsited-sdk-go/v4"
```

Every package beneath it is importable too, and each family is its own package
(`.../resources/accounts`, `.../resources/orders`, 91 of them); the module holds 100
packages. Prefer the facade for the core names: it is one import, and it is what the
examples here use.

## Credentials

**The SDK reads no credential from the environment, and never `.env`.** You build an
`AuthConfig` from wherever your credentials live and hand it to `NewClient`, or register it
once for `BuildClient`.

It does read environment variables, none of them a credential: `HTTP_PROXY` /
`HTTPS_PROXY` / `NO_PROXY` when `ProxyURL` is empty (`http.ProxyFromEnvironment`);
`NO_COLOR` and `TERM=dumb` to suppress ANSI colour; the locale variables off Windows, to
choose the console's rule glyph; and `LOCALAPPDATA` / `XDG_CACHE_HOME` / the home
directory to locate the default token cache. And one of its own:
**`EXSITED_RESPONSE_FORMAT`**, a presentation setting, the lowest-priority tier of
[Response format](#response-format).

```go
auth := exsited.AuthConfig{
	ClientName:   "default",
	BaseURL:      os.Getenv("EXSITED_BASE_URL"),
	ClientID:     os.Getenv("EXSITED_CLIENT_ID"),
	ClientSecret: os.Getenv("EXSITED_CLIENT_SECRET"),
	RedirectURI:  os.Getenv("EXSITED_REDIRECT_URI"),
}
client, err := exsited.NewClient(exsited.SdkConfig{Auth: auth})
```

| Field | Required | Default | Notes |
|---|---|---|---|
| `BaseURL` | yes | | Needs an `http` or `https` scheme and a host; trailing `/` stripped. No userinfo, query or fragment |
| `ClientID` | yes | | |
| `ClientSecret` | yes | | Masked by `String()` |
| `RedirectURI` | yes | | Sent in the `client_credentials` token body |
| `ClientName` | yes | | Label in diagnostics; selects the instance by name |
| `GrantType` | | `"client_credentials"` | |
| `TokenCachePath` | | the platform default | Where the token file lives |
| `InMemoryToken` | | `false` | Keep the token for the life of the process, write nothing |

There is no auth string: the v4 credentials grant carries `grant_type`, `client_id`,
`client_secret` and `redirect_uri`, the refresh grant carries `grant_type`, `client_id`,
`client_secret` and `refresh_token`, and no grant takes an auth string.
Validation runs inside `NewClient` and reports every problem at once; classify it with
`errors.Is(err, exsited.ErrConfig)`.

## Configuration

`AuthConfig` is credentials. Behaviour is the rest of `SdkConfig`, and every field has a
default: `SdkConfig{Auth: auth}` is a working client.

| Setting | Default | Notes |
|---|---|---|
| `APIVersion` | `"v4"` | Fills `{api_version}` in every resource path. This SDK serves v4 only; `"4"` is accepted. The token path does not follow it |
| `TimeoutSeconds` | `30.0` | A hard ceiling on **one attempt**, headers and body read together; each retry gets a fresh one. See [Timeouts](#timeouts) |
| `ConnectTimeoutSeconds` | `10.0` | DNS, TCP connect and TLS handshake, inside the attempt ceiling. Configures the client the SDK builds, not one you supply |
| `MaxRetries` | `1` | Total attempts **including the first**, so `1` disables retry |
| `RetryableHTTPCodes` | `[429,500,502,503,504]` | Inert while `MaxRetries` is `1` |
| `UnsafeRetryableHTTPCodes` | `[]` | The second gate every method that is not safe must also pass. Empty, so a write is never replayed on a status |
| `RetryBackoffSeconds` | `0.5` | Base for exponential backoff with full jitter. A `*float64`: `nil` means the default, `exsited.Ptr(0.0)` retries immediately |
| `RetryBackoffMaxSeconds` | `20.0` | Ceiling for one sleep, `Retry-After` included |
| `VerifySsl` | `true` | A `*bool`: `nil` means `true`; `exsited.Ptr(false)` disables certificate checks and emits an `InsecureTransportWarning` |
| `CABundle` | `""` | PEM path for a private CA or TLS interception; refused beside `VerifySsl=false`, and a missing file is refused |
| `ProxyURL` | `""` | Proxy for both schemes, validated at construction; empty falls back to the environment |
| `PoolMaxsize` | `10` | Idle connections per host; a reuse hint, not a ceiling on concurrent requests |
| `MaxResponseBytes` | `67108864` (64 MiB) | Ceiling on one buffered response body; `0` takes the default |
| `DefaultHeaders` | `nil` | Sent on every request, held to the header grammar at construction |
| `RaiseOnError` | `false` | `true` also returns an `*SdkError` for a failed response, keyed on `Succeeded()` |
| `ResponseFormat` | `""` | What `Data` follows: `"map"`, `"json"` or `"dto"`. `""` defers to `EXSITED_RESPONSE_FORMAT`, then `"map"`. An invalid value is refused at construction |
| `UserAgent` | `exsited-go-sdk/<version>` | |
| `LogLevel` | `"WARNING"` | `DEBUG` / `INFO` / `WARNING` / `ERROR` / `CRITICAL`; gates what the log hook and `LogFile` receive |
| `LogBodies` | `false` | A `DEBUG` line with the redacted payload and response |
| `LogFile` | `""` | Also append the log records to this file, `0600`; released by `Close` |
| `PrintApiInfo` | `false` | Client, api version, method, endpoint, params, status, attempt, elapsed |
| `PrintRequestPayload` | `false` | Redacted |
| `PrintResponseBody` | `false` | Redacted |
| `PrintHeaders` | `false` | Redacted |
| `PrintCurl` | `false` | A POSIX-shell curl line, token redacted |
| `PrintOnlyOnError` | `false` | Suppress the block when `Succeeded()` is true |
| `PrintFormat` | `"pretty"` | Or `"compact"`, one line per call |
| `PrintColor` | `false` | ANSI colour, forced off when `NO_COLOR` or `TERM=dumb` is set |
| `PrintMaxChars` | `0` | `0` is unlimited |
| `PrintRequests` | `false` | An alias of `PrintApiInfo`; either turns the block on |
| `DiagnosticWriter` | `nil` | Receives everything the SDK prints; `nil` keeps stdout for exchanges and stderr for warnings |
| `HTTPClient` | `nil` | Sent through instead of a built client, with redirect-following removed from a copy; the TLS, proxy and pool settings are refused beside it |

`NewClient` normalises a copy and validates it before building anything, and
`client.Config()` answers a deep copy of what it resolved. `cfg.Validate()`,
`cfg.Normalised()`, `cfg.DeepCopy()` and `cfg.WithOverrides(map[string]any{...})` are the
same steps on their own.

## Multiple instances

Hold several instances open at once with a `ClientRegistry`. Each entry is an `SdkConfig`,
an `AuthConfig`, a pointer to either, or a mapping of the fields; the first is the default,
and every `ClientName` must be distinct:

```go
registry, err := exsited.NewClientRegistry(production, staging)
defer registry.Close()                           // closes what it built, not an added client

first, _ := registry.GetClient(nil)              // the default
byName, _ := registry.GetClient("staging")
```

Each instance gets its own config, transport and token source. A client is built on first
use, so one never asked for never authenticates.

Or register once for the whole process and build a client in one call:

```go
err := exsited.RegisterInstances(production, staging)   // every block validated
client, err := exsited.BuildClient(nil)                 // the default
client, err = exsited.BuildClient("staging", func(c *exsited.SdkConfig) { c.MaxRetries = 3 })
names, _ := exsited.ListInstanceNames()
blocks, _ := exsited.GetRegisteredInstances()           // deep copies
```

Both enumerators return `ErrUsage` when nothing has been registered. `BuildClient` applies
its overrides to a deep copy of the registered block. `exsited.SdkConfigFromCredentials`
picks one entry out of a list of blocks and applies snake_case overrides, the same keys
`WithOverrides` takes.

## Authentication

Automatic. OAuth2 `client_credentials` against **`POST /api/v4/oauth2/token`** on first
use, cached, and renewed before it expires.

A `401` that means a stale token invalidates that token (compare-and-clear, so a token
another caller already renewed is not thrown away), refreshes, trying the `refresh_token`
grant first and then the credentials, and replays the request once, under the same headers
and the same idempotency key. A `401` that means something else is returned as it is.

The default cache is a file, keyed by a credential fingerprint over the base URL, client
id, client secret and grant type, so one file serves many instances:

| Platform | Default path |
|---|---|
| Windows | `%LOCALAPPDATA%\exsited\exsited_go_token.json` |
| macOS | `~/Library/Caches/exsited/exsited_go_token.json` |
| Linux | `$XDG_CACHE_HOME/exsited/exsited_go_token.json`, else `~/.cache/exsited/exsited_go_token.json` |

`TokenCachePath` names another file, and `InMemoryToken: true` writes nothing. Never the
working tree: the file holds a refresh token. It is written atomically under a
cross-process lock, a sidecar file created `O_CREATE|O_EXCL` and stamped with pid, host,
time and a serial; a leaked lock is broken by asking whether the pid is alive on this host.
When the lock cannot be had at all, the work is refused with `ErrTokenLock` rather than run
without exclusion; set `InMemoryToken` to keep the token without coordinating.

```go
client, _ := exsited.NewClient(cfg, exsited.WithTokenStore(exsited.NewMemoryTokenStore()))
```

## Requests

Every endpoint is a method on its family's type, reached through the client's typed,
cached accessor: a `context.Context` first, then the values
that identify the resource, then the query parameters and the payload where the endpoint
takes them, then `raiseOnError *bool` and at most one trailing response format. A nil
`raiseOnError` and an omitted format defer to the config for that one call.

```go
response, err := client.Accounts().GetAccount(ctx, "W13ZGY", nil)
response, err = client.Accounts().GetAccount(ctx, "W13ZGY", exsited.Ptr(false), exsited.ResponseFormatDTO)

accountsFor := accounts.New(client)                     // bind a family to any requester
response, err = accounts.GetAccount(client, ctx, "W13ZGY", nil)   // the module form

bound, err := client.Resource("accounts")               // by name, errors.Is(err, exsited.ErrAttribute) on a typo
names := client.ListResourceNames()                      // the 91 registered families
```

The generic senders cover everything else, context-first and returning
`(*ApiResponse, error)`:

```go
client.SendGetRequest(ctx, path, options)     client.SendPostRequest(ctx, path, options)
client.SendPutRequest(ctx, path, options)     client.SendPatchRequest(ctx, path, options)
client.SendDeleteRequest(ctx, path, options)  client.SendApiRequest(ctx, method, path, options)
```

`options` is a `*exsited.RequestOptions`; `nil` is the bare authenticated call. Every path
template is exported with a `Path` suffix (`accounts.AccountsPath`, `exsited.AccountsPath`),
so a raw call never restates a path.

| Field | Notes |
|---|---|
| `PathParams` | One per `{name}`; `{api_version}` is filled from config. A missing, extra or blank key is `ErrUsage` before any I/O |
| `QueryParams` | Normalised by the value table below |
| `JSONBody` | A JSON-encodable map, slice or array, a raw JSON string sent byte for byte, or a DTO payload. Not with `FormData` or `Files` |
| `FormData` | A form-encoded body |
| `Files` | A multipart body, `map[string]any{"file": []any{"report.csv", bytes}}`, an optional third element carrying the part's `Content-Type` |
| `Headers` | Beat `DefaultHeaders`; cannot override `Authorization` |
| `Unauthenticated` | `true` skips the token |
| `RaiseOnError` | `*bool`; overrides the config for this call |
| `ResponseFormat` | `"map"`, `"json"` or `"dto"` for this call; `""` defers |
| `DTOShape` | What `DTO()` builds; an endpoint method sets it, a raw call leaves it nil unless you name one (`accounts.AccountRecordShape`) |

**A body may be written three ways, on the same methods**, and all reach the wire as the
same request:

```go
client.Accounts().CancelAccount(ctx, id,
	map[string]any{"account": map[string]any{"effective_date": "2026-08-31"}}, nil)

client.Accounts().CancelAccount(ctx, id, `{"account": {"effective_date": "2026-08-31"}}`, nil)

date := "2026-08-31"
client.Accounts().CancelAccount(ctx, id, &accounts.AccountCancelBody{EffectiveDate: &date}, nil)
```

A raw JSON string is parsed once, locally, only to verify it, and sent as written. An
invalid one, a nil payload (a typed nil included), a nested map in the query string or a
blank path value is refused with `ErrUsage` before the token grant.

**Every `POST` carries an `X-Idempotency-Key`**, generated once per call and kept across its
attempts and the stale-token replay, so v4 collapses a repeated create into one. JSON, form
and multipart writes alike carry it, and v4 refuses a `POST` without one. A key you set in
`Headers` or `DefaultHeaders`, in any spelling, wins; the other verbs carry none.

**Path values** are percent-encoded with an empty safe set, so `"a/b"` cannot open a segment,
and a value that is nothing but dots is refused. An absolute `http(s)://` path on the
configured origin is used as is; one on another origin, or carrying userinfo, is refused,
because every request leaves with the bearer token attached. A `#` in a path is refused.

**Headers** are held to RFC 9110's grammar narrowed to printable ASCII, over the full merged
set: a name is a token, a value is `0x20`-`0x7E` and a tab inside it, with no leading or
trailing space. A refusal names the header and a byte position, never the value.

**Query values** are normalised, names never:

| Go value | On the wire |
|---|---|
| `nil` | dropped |
| `true` / `false` | `"true"` / `"false"` |
| `exsited.SortDirection` | its string value |
| `time.Time` | ISO 8601 with the UTC offset always kept, sub-second only when present |
| any number | as is |
| a slice | the key repeats, one per item |
| a map, or a slice inside a slice | `ErrUsage` |

**Filters and sorting** have helpers, because the v4 grammar is irregular. Equality carries
no operator suffix (`filter_status`), the other operators do (`filter_created_on_gt`), and
`FilterParam` refuses an operator it does not know; `MustFilterParam` panics instead, for a
literal written in source. `FilterValues` renders an `in` list comma-separated in one key,
as v4 splits it. **Sorting is spelled two ways and the wrong one is not an error**: Account,
Accounting Codes, Alert, Integration Connections and Inventory Adjustments take `sort` with a
`-` prefix to descend (`SortParam`); Item, Order, Usage, Purchase Order and Return Merchandise
Authorisation take `order_by` with `direction` (`SortParams`). Send a family the other one
and it answers `200`, unsorted.

## Pagination

```go
options := exsited.PaginateOptions{
	QueryParams: map[string]any{exsited.MustFilterParam("status"): "ACTIVE"},
}
for record, err := range exsited.IterRecords(ctx, client, accounts.AccountsPath, accounts.AccountsKey, options) {
	if err != nil {
		return err
	}
	fmt.Println(record)
}
for page, err := range exsited.IterPages(ctx, client, accounts.AccountsPath, options) {
	_ = page // one *ApiResponse per page
}
```

v4 serves two paging shapes and choosing wrong fails silently, so the pager detects the
shape on every page: the account and item families page by **cursor**, most other listings
by **offset**. `QueryParams` travel with **every** page, because a cursor or an offset is a
position within one filtered, sorted listing. `next_page` at the end of an offset listing is
the literal string `"NULL"` (`exsited.NoNextPage`), and the next offset is computed rather
than a `next_page` URL followed. A repeated cursor is refused rather than followed forever.

Both iterators are lazy and `break` fetches no more. `MaxPages` caps the walk, `PathParams`
fill the template on every page, and `RecordKey` names the key an offset page's records sit
under. A failed page yields an error by default; on `IterPages`, `RaiseOnError:
exsited.Ptr(false)` yields the failed page last instead. The client's own `IterPages` and
`IterRecords` methods are the visitor and collect-all forms.

## Responses

```go
response.StatusCode          // 200
response.Data                // the body under the resolved response format
response.Headers             // map[string]string
response.ContentType         // "application/json"
response.Succeeded()         // 2xx AND a body that does not deny it
response.ErrorType           // exsited.ErrorTypeNotFound, ErrorTypeValidation, ...
response.GetErrorMessages()  // the API's own messages
response.ErrorCode()         // "validation_error", from the v4 error envelope
response.RequestID()         // the API's id for the request; quote it when reporting a problem
response.DocURL()            // the documentation the envelope links
response.FieldErrors()       // []exsited.FieldError
response.RequestInfo         // method, url, params, redacted body and headers, timing
response.RaiseForStatus()    // an *SdkError when Succeeded() is false, else nil
```

**Four members read the body, and three of them ignore configuration:**

```go
response.Data     // follows the resolved response format
response.Map      // ALWAYS the decoded map, list, text or bytes, nil on a 204
response.JSON()   // ALWAYS the server's own body text
response.DTO()    // ALWAYS the typed view, nil where the endpoint names none
```

`Map` is a field rather than a `Map()` method, because the token
provider reads it. It is the only member that can tell a key the API **omitted** from one it
sent as **null**. `JSON()` and `DTO()` are each built on first read and kept, safely under a
concurrent first read. No view stands in for another: asking for a form a body does not have
answers that form's own empty value, and `Data` is `nil` rather than a substitute.

**A DTO converts nothing.** A field declared `*string` reads `nil` when a number arrives;
only an integer widening to a float is allowed. Every field is a pointer, so a JSON null and
an absent key both read `nil`, and `Map` says which. A field's wire name is spelled from its
Go name (`AccountID` is `account_id`, `AddressLine1` is `address_line_1`), and a field whose
key no rule reaches carries a `wire` tag. Every type was measured from real responses; a
field that was always null is left untyped. `exsited.DTOFields(accounts.AccountRecord{})` lists the wire names a type
declares, and `dto.Read[T]` / `exsited.DecodeDto` read any body into one.

### Response format

`Data` follows one setting, resolved from **four tiers, highest first**:

```go
client.Accounts().GetAccounts(ctx, nil, nil, exsited.ResponseFormatDTO)   // 1. per call
exsited.BuildClient(nil, func(c *exsited.SdkConfig) {                      // 2. per client
	c.ResponseFormat = exsited.ResponseFormatDTO
})
// 3. environment:  EXSITED_RESPONSE_FORMAT=dto
// 4. built-in:     "map"
```

The environment sits below code on purpose: a project-wide default must never override a
choice written in source. An invalid value fails at client construction with
`exsited.ErrConfig`, naming all three, rather than falling back. Spellings are trimmed and
case-insensitive at every tier.

### A 2xx that denies itself

**A 2xx alone does not mean success.** A v4 body whose `object` member is `"error"` is a
failure whatever the status says, so `Succeeded()` is false, `RaiseOnError` fires on it, and
the error envelope's accessors read it. The rule can move a 2xx from success to failure and
never the reverse. A denied 2xx keeps the status that arrived and classifies
`ErrorTypeUnknown`.

Decoding follows the content type: JSON to `map[string]any` / `[]any` with numbers kept as
`json.Number`, `text/*` to a `string` decoded per its charset, anything else to `[]byte`
(the PDF endpoints), a 204 or empty body to `nil`.

| Status | `ErrorType` |
|---|---|
| 2xx | `ErrorTypeNone`, or `ErrorTypeUnknown` when the body denies success |
| 3xx | `ErrorTypeRedirect`; not followed, read `Location` off `Headers` |
| 400, 422 | `ErrorTypeValidation` |
| 401 / 403 | `ErrorTypeAuthentication` / `ErrorTypeAuthorization` |
| 404 / 405 / 409 / 429 | `ErrorTypeNotFound` / `ErrorTypeMethodNotAllowed` / `ErrorTypeConflict` / `ErrorTypeRateLimit` |
| 5xx | `ErrorTypeServer` |
| timeout / connection failure | `ErrorTypeTimeout` / `ErrorTypeNetwork` |
| a cancelled context | `ErrorTypeUnknown`, see [Cancellation](#cancellation) |

## Errors

A failed call does **not** produce an error by default. It comes back as an `*ApiResponse`
with `Succeeded()` false, so an expected failure is ordinary control flow:

```go
if !response.Succeeded() {
	fmt.Println(response.StatusCode, response.ErrorCode(), response.FieldErrors())
}
```

Opt in to a real `error` config-wide with `RaiseOnError: true`, or per call with
`exsited.Ptr(true)`. The response still comes back beside the `*SdkError`:

```go
response, err := client.Accounts().GetAccount(ctx, "missing", exsited.Ptr(true))
var sdkError *exsited.SdkError
if errors.As(err, &sdkError) {
	fmt.Println(sdkError.StatusCode, sdkError.ErrorType) // 404 not_found
	fmt.Println(sdkError.Error())      // 404 GET /api/{api_version}/accounts/{account_id} - not_found response (1 reported)
	fmt.Println(sdkError.ErrorCode, sdkError.RequestID)
	fmt.Println(sdkError.Explain())
}
```

`SdkError` is the one error type the SDK returns. Classify it with `errors.Is`:

| Sentinel | Matches |
|---|---|
| `exsited.ErrConfig` | credentials or settings missing or invalid, from `NewClient` |
| `exsited.ErrUsage` | a caller mistake caught before any I/O |
| `exsited.ErrAttribute` | a name that does not exist on a client or resource; it also matches `ErrUsage` |
| `exsited.ErrAuth` | credentials rejected, or a `401` surviving one refresh |
| `exsited.ErrTokenCache`, `exsited.ErrTokenLock` | the token cache file or its lock |

`ErrorType` is the second, independent classification, by status; its string values are
stable. `InsecureTransportWarning` is a value, not an error: construction still succeeds,
and the client prints it once per message when `VerifySsl` is off or the base URL is plain
`http://` to a host that is not loopback.

**`Error()` is one structural line, and it is the form to retain.** It carries the status,
the method, the **endpoint template** and the SDK's own summary, nothing a server or a
customer supplied. The template, not the resolved path, because a resolved path names a
record; where there is no template the fallback is the path's shape, `<4-segment path>`. A
transport failure is named by its type. Nothing is lost: the API's messages are on
`GetErrorMessages()`, the body on `RawResponse`, the envelope on `ErrorCode`, `RequestID`,
`DocURL` and `FieldErrors`, the resolved path and query on `RequestInfo`, the underlying
failure on `Cause` (and redacted on `CauseText`), and all of it together on `Explain()`.

**`Explain()` is a diagnostic, not a log line.** It renders the resolved path, the query,
the cause and up to 500 characters of the body, and redaction masks by secret-named key, so
a customer's name or email prints verbatim. Send it to a terminal a person is looking at,
and keep `Error()`, `StatusCode` and `ErrorType` for anything retained.

## Diagnostics

```go
client, _ := exsited.NewClient(exsited.SdkConfig{
	Auth:                auth,
	PrintApiInfo:        true,
	PrintRequestPayload: true,
	PrintResponseBody:   true,
	PrintHeaders:        true,
	PrintCurl:           true,
	PrintOnlyOnError:    true,                      // only when a call fails
	PrintFormat:         exsited.PrintFormatPretty, // or exsited.PrintFormatCompact
})
```

The pretty format is a plain title line, `label: value` rows, `-- section ----` rules and raw
JSON blocks. A failed v4 exchange adds its `errors` and
`request id` rows:

```
Exsited  GET  200  0.0231s
client      : default
api version : v4
method      : GET
endpoint    : https://your-instance.exsited.com/api/v4/accounts
params      : {"limit":1}
status      : 200
attempt     : 1 / 1
elapsed     : 0.0231s
```

The compact format is one line per call:
`exsited | default | GET | https://.../api/v4/accounts | 200 | 1/1 | 0.0231s`.
`PrintMaxChars` cuts the block, `PrintColor` colours it, and `PrintCurl` emits a POSIX-shell
curl line with the token redacted. The block goes to stdout and warnings to stderr, both to
`DiagnosticWriter` when it is set. Every value taken from outside the SDK is neutralised
first, so an escape sequence in a body cannot move the cursor, and a rune the console cannot
carry is spelled as an escape.

Logging is separate. `LogLevel` gates the records a log hook receives: an `INFO` line per
exchange and per attempt, a `WARN` line for a `401` that is refreshed and replayed, an `ERROR`
line for a request that got no response, and with `LogBodies` a `DEBUG` line holding the
redacted payload and response. `LogFile` appends the same records, timestamped and prefixed
`exsited.<client>`:

```go
client, _ := exsited.NewClient(cfg, exsited.WithLogger(func(line string) { log.Println(line) }))
```

Everything printed or logged passes through one redactor: a key whose name carries one of
the eighteen sensitive markers (`secret`, `password`, `token`, `bearer`, `jwt`,
`authorization`, `apikey`, `cookie`, `session`, `signature` and the rest) prints as `***`,
matched ignoring case and separators.

## Retries

Off by default: `MaxRetries` counts the first attempt.

```go
exsited.NewClient(exsited.SdkConfig{Auth: auth, MaxRetries: 3})
```

Exponential backoff with full jitter, capped by `RetryBackoffMaxSeconds`. A `Retry-After`
header, in the delay-seconds form or as an HTTP date, beats the computed jitter but is
clamped by the same ceiling. `400` is not retryable, and neither is `401`: the retry loop
sits inside the transport, beneath the client's refresh-and-replay, so retrying a `401` would
spend every attempt on a token already known to be dead.

### Retry only what certainly did not arrive

| Failure | Proves | Retried on |
|---|---|---|
| DNS resolution, connect timeout, connection refused, TLS handshake | the request was **never sent** | any method |
| read timeout, a reset while awaiting the response, a truncated body | delivery is **unknown** | the safe set plus `PUT` and `DELETE`, **never** `POST` or `PATCH` |
| caller cancellation, anything that is not a transport failure | | nothing |

A `POST` that timed out awaiting its response is not retried, because the instance may
already have created the record. The idempotency key makes a replay safe only where the
server honours it, and this rule does not assume that it does.

### Retrying a status is scoped by method

| | Retried on |
|---|---|
| `GET` `HEAD` `OPTIONS` `TRACE`, RFC 9110's safe set | `RetryableHTTPCodes` |
| every other method | `UnsafeRetryableHTTPCodes` (default **empty**) as well |

The safe set is an allow-list, so a method this SDK has never heard of is not replayed. A
status must be in **both** lists to be retried for a method that is not safe. Each list is
replaced wholesale when you set it.

### Timeouts

| | Bounded by |
|---|---|
| connect: DNS, TCP connect, TLS handshake | `ConnectTimeoutSeconds`, and the attempt ceiling above it |
| read: headers and body together | the remainder of the attempt ceiling |
| one attempt | `TimeoutSeconds` |
| one call | up to `MaxRetries` times `TimeoutSeconds`, plus backoff |
| preparation: encode, multipart assembly, token grant | your own `context` |

`ConnectTimeoutSeconds`, `VerifySsl`, `CABundle`, `ProxyURL` and `PoolMaxsize` configure the
client the SDK builds; a client supplied through `HTTPClient` or `exsited.WithHTTPClient`
keeps its own. The wait for the token-cache lock sits outside all of it.

## Transport

```go
exsited.NewClient(exsited.SdkConfig{
	Auth:             auth,
	ProxyURL:         "http://proxy.internal:8080",
	PoolMaxsize:      10,
	MaxResponseBytes: 256 << 20,
	VerifySsl:        exsited.Ptr(true), // nil already means true
	CABundle:         "",                // PEM path, the answer to TLS interception
})
```

One `net/http` client per SDK client, so two with different TLS or proxy settings never
contend. `PoolMaxsize` bounds idle connections per host and caps nothing about concurrency.

**Bodies are bounded, and every ceiling is explicit:**

| Ceiling | Value | What it bounds |
|---|---|---|
| `exsited.MaxRequestBodyBytes` | 64 MiB | the final wire body of any request |
| `exsited.MaxMultipartBytes` | 64 MiB | the multipart body's exact size, judged before a byte is copied |
| `MaxResponseBytes` | 64 MiB, configurable | one buffered response body, on the declared length and then the running total, decompressed |
| `exsited.MaxCABundleBytes` | 4 MiB | the `CABundle` PEM file |
| `auth.MaxTokenCacheBytes` | 1 MiB | the token cache file |

An over-sized response fails naming the cap and the field, classified `ErrorTypeUnknown`,
and is not retried. A multipart body is assembled once, with a boundary checked against
every byte of its parts, and every attempt sends that one document.

**Redirects are not followed.** A `3xx` comes back like any other non-2xx, with `Location`
on `Headers`, so following one is a request you make deliberately rather than something the
SDK does with your `Authorization`.

Supply your own transport with `exsited.WithTransport(...)`, and you keep it: `Close` leaves
an injected transport open, so one transport can back several clients.

## Endpoints

**865 endpoint methods over 91 families**, one package per family under `resources/`, each
reached through its accessor (`client.Accounts()`, `client.Orders()`, ...) or bound with its
`New`. Every endpoint, with its request and response bodies, is documented at
[exsited.mintlify.app](https://exsited.mintlify.app/introduction).

What the API does that the code cannot tell you:

- **v4 creates an account at `POST /api/v4/accounts`**, answered `201`, which
  `CreateAccount` sends; `POST /api/v4/express` with the account embedded also creates one.
- **Reads are nested, writes are flat.** `currency` reads back as an object and is written as
  `"AUD"`. Posting a response body back unchanged is a `400`.
- **A locked account's delete is a `409` with code `internal_error`**, a business rule wearing
  a server fault's name. Not retryable.

## Architecture

The SDK packages, with `exsited.go` as their facade.

### The layer map

| Package | Contents |
|---|---|
| `auth/` | the token provider, the memory and file token stores, the sidecar lock, the credential fingerprint |
| `client/` | `exsited_client.go`, `client_registry.go`, `console.go`, `pagination.go`, `resource_binding.go`, and `glyphs*.go`, the console-encoding probe |
| `config/` | `auth_config.go` (credentials), `sdk_config.go` (behaviour), `select_credentials.go` |
| `http/` | `http_transport.go`, `transport.go`, `url_builder.go`, `query_params.go`, `retry_policy.go`, and the shared value types `api_response.go`, `error_type.go`, `request_info.go`, `request_options.go` |
| `resources/<family>/` | the 91 families, four files each |
| `errors/`, `redact/`, `dto/` | one error family, one redactor, one typed view; all three dependency leaves |

The facade re-exports the public surface as type aliases and ordinary functions, so the
ordinary consumer imports one package. Nothing callable on it is a variable holding a
function, so no importer can replace the SDK's behaviour for another; the error sentinels and
`FilterOperators` are the only package-level variables.

### Where the value types live

`api_response.go`, `request_info.go` and `request_options.go` sit in `http/` rather than
beside the client that hands them to a caller, because Go forbids import cycles and `client`
sits at the top of the graph:

```
client    ──> http       (builds the default HTTPTransport)
client    ──> auth       (builds the default TokenProvider and TokenStore)
client    ──> resources  (registers every family, and a typed accessor for each)
auth      ──> http       (the token provider calls the transport)
resources ──> http       (ApiResponse, RequestOptions, Requester)
```

`http.Requester` is the subset of the client a family calls, defined below both `client`
and `resources`, so a family depends on the client's behaviour without importing it.
`ErrorType` is declared in `errors`, the leaf, and aliased from `http`, so
`SdkError.ErrorType` is a concrete type.

### The object graph

```
NewClient(SdkConfig)
    │
    ├── AuthConfig ─────────── credentials
    └── SdkConfig ──────────── behaviour, validated together at construction
    │
    ▼
ExsitedClient ──────────────── validation, path and query params, header merge,
    │                          the idempotency key, printing and logging
    │   SendGetRequest · SendPostRequest · SendPutRequest
    │   SendPatchRequest · SendDeleteRequest · SendApiRequest
    │
    ├── Accounts(), Orders(), ... one cached accessor per family
    │
    ├── TokenProvider ──────── bearer token, renewed before it expires
    │       └── TokenStore ─── MemoryTokenStore | FileTokenStore, under a lock
    │
    ▼
HTTPTransport ───────────── net/http, TLS, proxy, pooling, retries, decoding
    │
    ▼
ApiResponse ──────────────── StatusCode, Data, Map, Headers, ErrorType, RequestInfo
```

Retry sits inside the transport, beneath the token refresh, which is why `401` is not a
retryable status. Validation sits above the transport, so a caller mistake is reported as
`ErrUsage` where it was written, before a socket opens.

## Stable formats

Formats a change to which is a breaking change, each documented where it is implemented:
the redaction markers (`redact/redaction.go`), the credential fingerprint
(`auth/token_store.go`, `BuildCredentialsNamespace`), the query-parameter value table
(`http/query_params.go`) and the `Retry-After` grammar (`http/retry_policy.go`).

- The **credential fingerprint** is `v4`-labelled, each value fed in with its length in
  code points, and must stay byte-identical across releases, or an existing cache entry is
  no longer found. The token file is BOM-less UTF-8 JSON.
- The default cache file is `exsited_go_token.json` in the platform's per-user `exsited/`
  directory.
- `Version()` and `SDKVersion()` are functions, not variables, so no importer can restamp
  what every request announces. They read the version of
  `github.com/exsited/exsited-sdk-go/v4` from the build's module metadata, which carries the
  tag `go get` resolved, so a build on 4.0.0 reports `4.0.0`, and a build that carries no
  module version reports the same release number. The default `User-Agent` is
  `exsited-go-sdk/<version>`, `exsited-go-sdk/4.0.0` for this release.
- The `ErrorType` string values.

### Cancellation

A cancelled context is wrapped in an `*SdkError`, and the identity survives the wrap:
`Unwrap` returns the cause, so `errors.Is(err, context.Canceled)` answers it. The wrapper's
`ErrorType` is `ErrorTypeUnknown`. `context.DeadlineExceeded` is not a cancellation and
classifies as `ErrorTypeTimeout`. A cancellation is never retried.

### Behaviour worth knowing

| Topic | Detail |
|---|---|
| `Map` is a field | Not a `Map()` method, because the token provider reads it |
| Pointers everywhere on a DTO | A JSON null and an absent key both read `nil`; `Map` tells them apart |
| The token path | `/api/v4/oauth2/token`, for the `client_credentials` and `refresh_token` grants alike |
| The idempotency key | One per call on `POST`, kept across attempts and the replay; a caller's own key wins |
| Sorting spellings | `sort` for five families, `order_by` with `direction` for five others; the wrong one is ignored with a `200` |
| Two method sets | The safe set gates a **status** retry; the safe set plus `PUT` and `DELETE` gates a **transport-failure** retry |
| Default retry sets | `exsited.DefaultRetryableHTTPCodes()` and `DefaultUnsafeRetryableHTTPCodes()` return copies, so no caller can edit the default for another |
| Described, not echoed | A refusal names a field and describes its value (`<an https URL of 31 characters>`); the value rides `CauseText` for `Explain()` |
| The glyph probe | A rule or ellipsis above ASCII comes from `ResolveBoxGlyphs`: the console output code page on Windows, the locale charset elsewhere |

## Troubleshooting

| Symptom | Fix |
|---|---|
| `go get github.com/exsited/exsited-sdk-go` fails, or finds no 4.x release | The module path ends in `/v4`: `go get github.com/exsited/exsited-sdk-go/v4`, and import it with the `/v4` |
| `errors.Is(err, exsited.ErrConfig)` naming blank fields | A credential arrived blank. Check `BaseURL`, `ClientID`, `ClientSecret`, `RedirectURI`, `ClientName` |
| `errors.Is(err, exsited.ErrAuth)` on the first call | Credentials rejected by the token endpoint. Check `ClientID`, `ClientSecret` and `RedirectURI` against the instance |
| A listing comes back unsorted with a `200` | The family takes the other sort spelling: `SortParam` or `SortParams` |
| `filter_status_eq` is refused | Equality carries no suffix: `filter_status` |
| A walk returns one page and reports nothing | The pager was bypassed; `IterPages` detects cursor or offset per page |
| A DTO field reads `nil` for a value plainly in the body | The declared type does not match what arrived. Nothing is coerced; read `Map` |
| Want a fresh token | Delete the cache file, or set `InMemoryToken: true` |

## Public API

Everything below is reachable from the single root import `exsited`, and each name is also
reachable on its own package path.

| Group | Names |
|---|---|
| Entry points | `NewClient` `NewExsitedClient` `ExsitedClient` `ClientRegistry` `NewClientRegistry` `RegisterInstances` `BuildClient` `ListInstanceNames` `GetRegisteredInstances` |
| Client options | `ClientOption` `WithTransport` `WithTokenStore` `WithTokenProvider` `WithRetryPolicy` `WithLogger` |
| Configuration | `AuthConfig` `SdkConfig` `NewSdkConfig` `APIVersion` `V4` `ResponseFormat` `ResponseFormatMap` `ResponseFormatJSON` `ResponseFormatDTO` `DefaultResponseFormat` `ResponseFormatEnvVar` `IsResponseFormat` `ResponseFormatNames` `ParseResponseFormat` `FormatMap` `FormatJSON` `FormatDto` `PrintFormatPretty` `PrintFormatCompact` `SelectCredentials` `ReadClientName` `LabelCredentials` `SdkConfigFromCredentials` `DefaultRetryableHTTPCodes` `DefaultUnsafeRetryableHTTPCodes` |
| Requests | `RequestOptions` `TransportRequest` `Requester` `DTOPayload` `Ptr` `SnapshotJSONValue` `IdempotencyHeader` |
| Query | `BuildQueryParams` `QueryPair` `EncodeQueryParams` `SortDirection` `SortAsc` `SortDesc` `Ascending` `Descending` `FilterOperators` `FilterParam` `MustFilterParam` `SortParam` `MustSortParam` `SortParams` `MustSortParams` |
| Responses | `ApiResponse` (`Data` `Map` `JSON()` `DTO()`) `RequestInfo` `NewRequestInfo` `DTOShape` |
| Errors | `SdkError` `FieldError` `ErrorType` `RequestContext` `EndpointNamer` `ErrConfig` `ErrUsage` `ErrAttribute` `ErrAuth` `ErrTokenCache` `ErrTokenLock` `NewAttributeError` `NewConfigError` `NewUsageError` `NewAuthError` `NewAuthResponseError` `NewInsecureTransportWarning` `InsecureTransportWarning` `ClassifyErrorType` `ErrorTypeNone` ... `ErrorTypeUnknown` |
| Pagination | `IterPages` `IterRecords` `PaginateOptions` `PageOptions` `ReadNextCursor` `ReadNextOffset` `DefaultRecordKey` `NoNextPage` |
| Resources | per family: the family type and `New<Family>`, its views, `<View>Shape` functions, `...Path` and `...Key` constants; `RegisterResourceModule` `RegisteredResourceNames` `ResolveResourceModule` `NewBoundResource` `BoundResource` `ResourceFactory` |
| DTO base | `Dto` `Enveloped` `BesideEnvelope` `TopLevelFields` `DTOFields` `DecodeDto` `EncodeDto` |
| Tokens | `TokenProvider` `NewTokenProvider` `TokenRecord` `TokenStore` `MemoryTokenStore` `NewMemoryTokenStore` `FileTokenStore` `NewFileTokenStore` `BuildCredentialsNamespace` `DefaultTokenCachePath` |
| HTTP | `HTTPTransport` `NewHTTPTransport` `TransportOption` `WithHTTPClient` `WithProgress` `WithTransportRetryPolicy` `RetryPolicy` `NewRetryPolicy` `BuildURL` `SubstitutePathParams` `FindPathParamNames` `MaxMultipartBytes` `MaxRequestBodyBytes` `MaxCABundleBytes` `ReadCappedFile` |
| Console | `RenderExchange` `BuildCurlCommand` `RenderJSONValue` `ResolveBoxGlyphs` `EncodeForConsole` `Neutralise` `NeutraliseKeepNewlines` |
| Redaction | `ShouldRedact` `RedactValue` `RedactMapping` `RedactText` `SensitiveKeyMarkers` `Redacted` `TooDeep` `Circular` `MaxRedactionDepth` |
| Version | `SDKVersion()` `Version()` |

## License

MIT. See [LICENSE](LICENSE).
