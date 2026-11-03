# Exsited SDK for Go

Go SDK for the Exsited v4 REST API. It handles OAuth2 with a token cache, retry with
backoff, connection pooling, TLS and proxy configuration, a classifiable error type and
secret redaction. It has no dependencies beyond the Go standard library.

A response comes back as the API sent it: `response.JSON()` is the body text,
`response.Map` is the decoded body as `map[string]any`, and `response.DTO()` is a typed
view of the same body. Keys are never renamed or converted.

## Requirements

- Go 1.24 or newer

## Install

```sh
go get github.com/exsited/exsited-sdk-go/v4
```

The module path ends in `/v4`, the major version, so the `/v4` belongs in the `go get`
line and in every import.

## Configure credentials

The SDK never reads credentials on its own. Build an `AuthConfig` from wherever your
credentials live, for example environment variables, and pass it to `NewClient`:

```go
auth := exsited.AuthConfig{
	ClientName:   "default",
	BaseURL:      os.Getenv("EXSITED_BASE_URL"),      // https://your-instance.example.com
	ClientID:     os.Getenv("EXSITED_CLIENT_ID"),     // your-client-id
	ClientSecret: os.Getenv("EXSITED_CLIENT_SECRET"), // your-client-secret
	RedirectURI:  os.Getenv("EXSITED_REDIRECT_URI"),  // https://your-app.example/callback
}
client, err := exsited.NewClient(exsited.SdkConfig{Auth: auth})
```

All five fields are required. Every other setting of `SdkConfig` has a default, so
`SdkConfig{Auth: auth}` is a working client.

## First call

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
			BaseURL:      "https://your-instance.example.com",
			ClientID:     "your-client-id",
			ClientSecret: "your-client-secret",
			RedirectURI:  "https://your-app.example/callback",
		},
		RaiseOnError: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	response, err := client.Accounts().GetAccount(context.Background(), "your-account-id", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Map) // the body as it arrived
	account := response.DTO().(*accounts.AccountRecord)
	fmt.Println(*account.Name) // the typed view
}
```

Every endpoint is a method on a resource family of the client, such as
`client.Accounts()` or `client.Orders()`. The accessors and methods are listed in the
[documentation](https://exsited.mintlify.app).

## Errors

By default a failed call returns a response whose `Succeeded()` is false. With
`RaiseOnError: true`, as above, the returned `err` is also non-nil on any failure. It is
an `*exsited.SdkError`, which you can classify with `errors.Is` or `errors.As`:

```go
var sdkError *exsited.SdkError
if errors.As(err, &sdkError) {
	fmt.Println(sdkError.StatusCode, sdkError.ErrorType)
}
```

## Pagination

List endpoints page by cursor or by offset. `IterRecords` detects the shape for you and
yields every record lazily:

```go
for record, err := range exsited.IterRecords(ctx, client, accounts.AccountsPath, accounts.AccountsKey, exsited.PaginateOptions{}) {
	if err != nil {
		return err
	}
	fmt.Println(record)
}
```

## Documentation

Guides and the full endpoint reference are at https://exsited.mintlify.app.

## License

MIT. See [LICENSE](LICENSE).
