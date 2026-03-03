# Scripts

## `dump-accessible-data.go`

Dumps responses from callable Vanta SDK endpoints into JSON files so you can inspect what your current credentials can access.

By default it calls read-only methods (`Get*`, `List*`) across all generated services, writes one file per response page, and creates an `index.json` summary.

### Prerequisites

- Go toolchain installed
- Valid Vanta credentials using one of:
  - `VANTA_BEARER_TOKEN`
  - `VANTA_CLIENT_ID` + `VANTA_CLIENT_SECRET`

### Run

```bash
go run ./scripts/dump-accessible-data.go
```

### Environment Variables

- `VANTA_BEARER_TOKEN`: Static bearer token (preferred if already available)
- `VANTA_CLIENT_ID`: OAuth client ID (used when bearer token is not set)
- `VANTA_CLIENT_SECRET`: OAuth client secret (used when bearer token is not set)
- `VANTA_SCOPE`: OAuth scope; default is `vanta-api.all:read`
- `VANTA_BASE_URL`: Optional API base URL override
- `VANTA_DUMP_DIR`: Output directory; default is `vanta-api-dump`
- `VANTA_PAGE_SIZE`: Page size for list endpoints; default `100`, max `100`
- `VANTA_INCLUDE_MUTATIONS`: Set to `1` or `true` to also call non-read methods

### Output

The script writes:

- `index.json`: per-endpoint call status
- `<Service>_<Method>.json`: non-list endpoint responses
- `<Service>_<Method>_page_###.json`: paginated list responses

Each `index.json` entry includes:

- `service`
- `method`
- `status` (`ok` or `error`)
- `file` (first output file when written)
- `pages` (for list methods)
- `error` (when a call fails)

### Notes

- Each endpoint call uses a 45 second timeout.
- List endpoints follow `results.pageInfo.hasNextPage/endCursor` pagination.
- File names are sanitized to safe characters.
