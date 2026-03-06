# Project Handover Notes

This document is for maintainers (human or Codex) to quickly become productive in this repo and avoid regressions.

## 1) Working Plan For Any New Change

1. Confirm intent against source of truth.
2. Locate impacted runtime files and generated service methods.
3. Make the smallest safe change.
4. Add or update tests for behavior, URL construction, and error handling.
5. Validate docs and examples match runtime defaults.
6. Run verification commands and summarize outcomes.

Use this checklist for every task:

- Identify whether the change is runtime (`v1/client.go`, `v1/auth.go`, helpers) or generated API surface (`v1/generated_services.go`, model files).
- Confirm endpoint/auth semantics against:
  - Vanta public docs
  - `Vanta Postman Env & Collection/Vanta API.postman_collection.json`
  - `Vanta Postman Env & Collection/Vanta Environment.postman_environment.json`
- Preserve existing style:
  - `params *XParams` input (nil-safe)
  - `context.Context` first arg
  - pointer optional query params
  - repeated query params via `query.Add`
  - required path params validated before request
- Update README when defaults, auth flow, or ergonomics change.
- Keep mutations opt-in when writing tools/scripts that inspect data.

## 2) Project Requirements And Intent

- Unofficial idiomatic Go SDK for Vanta API.
- Primary contract is derived from Vanta public developer docs + bundled Postman collection.
- API package path is `github.com/richardoc/vanta-sdk-go/v1`.
- Design goals:
  - predictable typed methods
  - minimal runtime complexity
  - clear error surfaces
  - pagination helpers
  - no implicit retry policy in library

## 3) Repository Map

- `v1/client.go`: HTTP request construction, base URL joining, headers, JSON and multipart execution.
- `v1/auth.go`: token model, static token source, OAuth client credentials source with synchronized caching.
- `v1/options.go`: client options (`WithBaseURL`, `WithHTTPClient`, `WithTokenSource`, `WithUserAgent`).
- `v1/errors.go`: `APIError` and non-2xx body decoding.
- `v1/pagination.go`: generic `ResultsPage[T]`, `Pager[T]`.
- `v1/generated_services.go`: generated services/endpoints (large, primary API surface).
- `v1/people_models.go`: hand-shaped people/person models used by generated methods.
- `scripts/dump-accessible-data.go`: introspection script that calls accessible endpoints and writes JSON.
- `Vanta Postman Env & Collection/`: imported Postman collection and environment.

## 4) Important Runtime Defaults

Current defaults (must stay aligned in code + README + tests):

- Base API URL: `https://api.vanta.com/v1`
- OAuth token URL: `https://api.vanta.com/oauth/token`
- Default user agent: `vanta-sdk-go/0.1`

Auth behavior:

- If using static token, caller provides `WithTokenSource(StaticTokenSource(...))`.
- OAuth token source caches tokens and refreshes with configurable skew.
- OAuth default scope in script tooling is `vanta-api.all:read` unless overridden.

## 5) Generated Surface Conventions

Generated methods follow consistent patterns:

- `func (s *XService) Method(ctx context.Context, params *XMethodParams) (...)`
- `params == nil` handled safely.
- Path params validated and escaped.
- Query pointers map to optional query values.
- Slice filters become repeated query params.
- Multipart endpoints accept `FormData map[string]string`.
- Unknown or unstable payload shapes may return `json.RawMessage` or `map[string]any`.

When editing generated methods manually, keep changes minimal and pattern-compatible. If broad changes are needed, prefer re-generation workflow (not currently automated in-repo).

## 6) Testing And Verification

Primary checks:

```bash
go test ./...
go vet ./...
go fix ./...
```

Tests currently cover:

- URL join/base URL behavior
- OAuth token caching behavior
- error and decode behavior
- pager edge cases (nil page, repeated cursor, missing cursor)

If adding new runtime behavior, add focused tests in `v1/*_test.go`.

## 7) Script Usage For Real API Validation

`scripts/dump-accessible-data.go` is useful for smoke validation with real credentials.

Typical run:

```bash
go run ./scripts/dump-accessible-data.go
```

Important env vars:

- `VANTA_BEARER_TOKEN` or (`VANTA_CLIENT_ID` + `VANTA_CLIENT_SECRET`)
- Optional: `VANTA_DUMP_DIR`, `VANTA_BASE_URL`, `VANTA_SCOPE`, `VANTA_PAGE_SIZE`, `VANTA_INCLUDE_MUTATIONS`

Outputs:

- `index.json` summary
- per-endpoint JSON response files

## 8) Known Risks / Gaps

- No committed in-repo generator pipeline yet (`internal/` currently empty). Large API surface is generated but regeneration process is not codified here.
- The repository’s `go.mod` currently sets `go 1.26.0`; ensure local CI/dev toolchains match or adjust intentionally.
- Postman collection changes upstream can cause drift; periodic reconciliation is required.
- `generated_services.go` is large; targeted edits should be accompanied by tests to prevent subtle regressions.

## 9) Practical Do/Do-Not Notes

Do:

- Keep behavior changes explicit and tested.
- Maintain alignment between runtime defaults, docs, and tests.
- Preserve backward-compatible method signatures when possible.
- Prefer small, reviewable commits.

Do not:

- Silently change auth/base URL defaults without updating README and tests.
- Introduce automatic retries unless explicitly agreed (current design intentionally omits retries).
- Reformat or refactor huge generated files without a concrete reason.

## 10) Fast Start For Next Maintainer

1. Read `README.md`.
2. Read `v1/client.go`, `v1/auth.go`, and `v1/options.go`.
3. Inspect impacted service methods in `v1/generated_services.go`.
4. Compare with relevant Postman endpoint definitions.
5. Implement + test + update docs in one pass.
