---
applyTo: "**/*.go"
---

# Go Code Instructions

## Formatting Authority

**gofmt is the single source of truth for code formatting.**

- Do NOT suggest formatting changes that conflict with `gofmt`
- Do NOT suggest style changes based on personal preferences
- Run `go fmt ./...` to verify formatting before suggesting changes
- Run `go vet ./...` to verify code quality before suggesting changes

## Code Style

- Go 1.21+ required
- Follow [Effective Go](https://go.dev/doc/effective_go) guidelines
- Use idiomatic Go naming conventions (camelCase for unexported, PascalCase for exported)
- Use functional options pattern for optional parameters
- Return errors explicitly, don't panic

## Linting Rules (golangci-lint)

Enabled linters:

- `errcheck` - check for unchecked errors
- `gosimple` - simplify code
- `govet` - examine Go source code
- `ineffassign` - detect ineffective assignments
- `staticcheck` - static analysis
- `unused` - check for unused code
- `gofmt` - check formatting
- `goimports` - check import ordering
- `misspell` - check for spelling mistakes
- `unconvert` - remove unnecessary conversions

## Error Handling

- Use sentinel errors for well-known error conditions
- Wrap errors with context using `fmt.Errorf("context: %w", err)`
- Check all errors - never ignore them
- Use `errors.Is()` and `errors.As()` for error comparison

## Testing

- Use standard `testing` package
- Use table-driven tests where appropriate
- Minimum 95% code coverage required
- All public functions must have tests
- Use `t.Run()` for subtests
- Use `-race` flag for race detection

## Commands

```bash
# Format code
go fmt ./...

# Vet code
go vet ./...

# Run linter
golangci-lint run

# Run tests
go test -v -race ./...

# Run tests with coverage
go test -v -race -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```
