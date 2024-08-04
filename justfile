format:
    @goimports -l -w cmd internal
    @gofumpt -l -w cmd internal
    @golines -w -m 80 cmd internal

lint:
    @golangci-lint run ./cmd/... ./internal/...
