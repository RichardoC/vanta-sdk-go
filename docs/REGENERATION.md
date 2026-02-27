# Periodic Regeneration Runbook

This document describes how to regenerate the Vanta Go SDK from the Postman collection and verify the repo is clean.

## Source Inputs (Required)

The generator requires these files in the repository:

- `Vanta Postman Env & Collection/Vanta API.postman_collection.json`
- `Vanta Postman Env & Collection/Vanta Environment.postman_environment.json` (reference only; useful for base/auth URL variables)

Generator entrypoint:

- `cmd/vanta-gen/main.go`

## Generated Output (Produced)

Regeneration rewrites:

- `v1/generated_services.go`

This file is generated code and should not be edited manually.

## Handwritten Runtime Files (Not Generated)

These files are maintained manually and are used by generated code:

- `v1/client.go`
- `v1/options.go`
- `v1/auth.go`
- `v1/errors.go`
- `v1/pagination.go`
- `v1/json.go`
- `v1/helpers.go`
- `v1/doc.go`

## Prerequisites

- Go 1.26+
- (Optional for `-race`) C toolchain available (`gcc`)

## Regeneration Steps

From repository root:

```bash
go run ./cmd/vanta-gen
```

Then format and verify:

```bash
gofmt -w $(find . -name '*.go' -not -path './.tools/*')
go fix ./...
go vet ./...
go test ./...
go test -race ./...
```

## Expected Package Paths

- Module root: `github.com/richardoc/vanta-sdk-go`
- Public SDK import: `github.com/richardoc/vanta-sdk-go/v1`

You can confirm:

```bash
go list ./...
```

Expected entries include:

- `github.com/richardoc/vanta-sdk-go/cmd/vanta-gen`
- `github.com/richardoc/vanta-sdk-go/v1`

## Update Workflow

1. Replace/update Postman collection files under `Vanta Postman Env & Collection/`.
2. Run regeneration and verification commands above.
3. Review diffs in `v1/generated_services.go` for endpoint additions/removals and naming changes.
4. If new endpoint behavior requires runtime support (e.g., new upload/download patterns), update handwritten runtime files in `v1/`.
5. Commit regenerated output and any runtime/doc updates together.

## Troubleshooting

- `go: command not found`:
  Install Go 1.26+ and ensure `go` is on `PATH`.

- `-race requires cgo` or missing `gcc`:
  Install `gcc`/build tools and run with `CGO_ENABLED=1`.

- Generator output changed unexpectedly:
  Check for changes in the Postman collection and re-run tests in `cmd/vanta-gen` to validate naming/dedupe rules.
