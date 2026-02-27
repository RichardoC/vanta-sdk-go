# vanta-sdk-go

`vanta-sdk-go` is an _unofficial_ idiomatic Go client for the Vanta API, aligned with the public Vanta developer docs and the Postman collection in this repository.

- Import path: `github.com/richardoc/vanta-sdk-go/v1`
- Minimum Go version: `1.26` (from `go.mod`)
- Auth: OAuth client credentials (built-in) or static bearer token
- Surface: generated services and endpoint methods from the collection
- Pagination: explicit `pageCursor/pageSize` params + generic pager helper

## Install

```bash
go get github.com/richardoc/vanta-sdk-go/v1
```

## Quick Start (OAuth)

```go
package main

import (
    "context"
    "fmt"

    vanta "github.com/richardoc/vanta-sdk-go/v1"
)

func main() {
    ctx := context.Background()

    ts, err := vanta.NewOAuthClientCredentialsTokenSource(vanta.OAuthClientCredentialsConfig{
        ClientID:     "YOUR_CLIENT_ID",
        ClientSecret: "YOUR_CLIENT_SECRET",
        Scope:        "controls.self:read controls.self:write",
        // AuthURL defaults to https://api.vanta.com/oauth/token
    })
    if err != nil {
        panic(err)
    }

    client, err := vanta.NewClient(vanta.WithTokenSource(ts))
    if err != nil {
        panic(err)
    }

    pageSize := 10
    resp, err := client.Services.Controls.ListControls(ctx, &vanta.ControlsListControlsParams{
        PageSize: &pageSize,
    })
    if err != nil {
        panic(err)
    }

    fmt.Printf("%+v\n", resp)
}
```

## Type Hints

Generated methods return per-operation typed response structs inferred from response examples and expose typed request body structs for JSON endpoints.

## Authentication Options

- Built-in OAuth client credentials flow via `NewOAuthClientCredentialsTokenSource`.
- Static token mode via `WithTokenSource(vanta.StaticTokenSource("..."))`.

## Pagination

Most list endpoints expose:

- `pageSize`
- `pageCursor`

The SDK also includes generic cursor helpers (`ResultsPage[T]`, `Pager[T]`) for custom typed wrappers.

## Error Handling

Non-2xx responses are returned as `*vanta.APIError` with:

- `StatusCode`
- `Status`
- raw `Body`
- parsed JSON body when available (`ParsedBody`)

## Notes

- Retries are intentionally **not** enabled in-library.
- Multipart endpoints are supported via generated `FormData` fields.
- Base API URL defaults to `https://api.vanta.com`.
