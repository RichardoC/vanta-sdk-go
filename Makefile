.PHONY: generate
generate:
	go run ./cmd/vanta-gen

.PHONY: generate-node
generate-node:
	node ./tools/generate.js

.PHONY: verify
verify:
	go test ./...
