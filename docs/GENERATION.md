# Generation Architecture

This repository uses a generated API surface with a handwritten runtime.

## Inputs

- `Vanta Postman Env & Collection/Vanta API.postman_collection.json`
- `developer.vanta.com/reference/*` (manual parity checks and drift review)

## Outputs

- `v1/generated_services.go`

## Handwritten Runtime

- `v1/client.go`: request construction, auth header injection, JSON and multipart request support
- `v1/auth.go`: token source contracts and OAuth client-credentials implementation
- `v1/errors.go`: typed API errors
- `v1/pagination.go`: cursor pagination primitives
- `v1/options.go`: client configuration
- `v1/json.go`: helper for explicit typed decoding

## Generation Steps

1. Parse Postman folder items recursively.
2. Extract operation metadata:
   - HTTP method
   - normalized path
   - path params
   - query params and inferred scalar types
   - request body mode (JSON vs multipart)
3. Group operations by top-level path segment into service structs.
4. Emit endpoint methods and parameter structs.

## Response Typing

The Postman collection provides broad endpoint definitions and examples but not a stable complete schema contract equivalent to a full OpenAPI document. The generated layer therefore:

- gives typed method signatures and typed parameter fields,
- generates top-level typed response structs inferred from examples,
- and falls back to `json.RawMessage` only when no structured response example is available.

This keeps the generated surface complete while remaining safe and forward-compatible.
