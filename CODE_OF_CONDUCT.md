# Code of Conduct

This document describes the coding standards and best practices for the project.
All code must pass automated quality checks before merging to main.

## Go

We use **gofmt** for formatting, **go vet** for static analysis, and **golangci-lint** for comprehensive linting.

### Rules

- Go 1.21+ required
- Follow [Effective Go](https://go.dev/doc/effective_go) guidelines
- Use idiomatic Go naming conventions
- All exported functions must have doc comments
- Return errors explicitly, don't panic

### Lint Rules (golangci-lint)

| Linter      | Description                      |
|-------------|----------------------------------|
| errcheck    | Check for unchecked errors       |
| gosimple    | Simplify code                    |
| govet       | Examine Go source code           |
| ineffassign | Detect ineffective assignments   |
| staticcheck | Static analysis                  |
| unused      | Check for unused code            |
| gofmt       | Check formatting                 |
| goimports   | Check import ordering            |
| misspell    | Check for spelling mistakes      |
| unconvert   | Remove unnecessary conversions   |

### Commands

```bash
# Format code
go fmt ./...

# Vet code
go vet ./...

# Lint (comprehensive)
golangci-lint run
```

### Configuration

See `.golangci.yml` for linter configuration.

### Useful Links

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [golangci-lint Documentation](https://golangci-lint.run/)

## Testing

We use the standard **testing** package with race detection and coverage.

### Rules

- Minimum 95% code coverage required
- All public functions must have tests
- Use descriptive test names
- Use table-driven tests where appropriate
- Always run tests with `-race` flag

### Commands

```bash
# Run tests
go test ./...

# Run with race detection
go test -race ./...

# Run with coverage
go test -race -coverprofile=coverage.out ./...
go tool cover -func=coverage.out

# Run specific test
go test -run TestToAlphanumeric ./...
```

### Useful Links

- [Go Testing Documentation](https://pkg.go.dev/testing)
- [Table-Driven Tests](https://go.dev/wiki/TableDrivenTests)

## Markdown

We use **markdownlint** for linting.

### Rules

- Headings must have 1 blank line above
- Multiple headings with same content allowed at different levels
- Standard markdownlint rules apply

### Commands

```bash
# Lint markdown files (if markdownlint-cli installed)
markdownlint '**/*.md' --ignore CLAUDE.md
```

### Useful Links

- [markdownlint Documentation](https://github.com/DavidAnson/markdownlint)
- [Markdown Guide](https://www.markdownguide.org/)

## Git Workflow

### Commit Messages

- Use present tense ("Add feature" not "Added feature")
- Use imperative mood ("Fix bug" not "Fixes bug")
- Keep first line under 72 characters
- Reference issues when applicable

### Branch Naming

- `feature/` - new features
- `fix/` - bug fixes
- `docs/` - documentation updates
- `refactor/` - code refactoring

## Pre-commit Checklist

Before committing, ensure:

- [ ] `go fmt ./...` produces no changes
- [ ] `go vet ./...` passes
- [ ] `golangci-lint run` passes
- [ ] `go test -race ./...` passes with 95%+ coverage
- [ ] All exported functions have doc comments
- [ ] All new functions have tests
