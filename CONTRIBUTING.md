# Contributing

Thank you for your interest in contributing to youtube-id-go! This guide will help you get started.

## Getting Started

### Prerequisites

- Go 1.21 or higher
- golangci-lint (optional, for linting)

### Setup

1. Fork and clone the repository:

   ```bash
   git clone https://github.com/your-username/youtube-id-go.git
   cd youtube-id-go
   ```

2. Verify Go installation:

   ```bash
   go version
   ```

3. Download dependencies:

   ```bash
   go mod download
   ```

## Development Workflow

### Making Changes

1. Create a new branch:

   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make your changes following our [Code of Conduct](CODE_OF_CONDUCT.md)

3. Run quality checks:

   ```bash
   # Format code
   go fmt ./...

   # Vet code
   go vet ./...

   # Lint (if golangci-lint installed)
   golangci-lint run

   # Tests with coverage
   go test -v -race -coverprofile=coverage.out ./...
   go tool cover -func=coverage.out
   ```

4. Commit your changes:

   ```bash
   git add .
   git commit -m "Add your descriptive message"
   ```

### Code Quality Requirements

All contributions must:

- Pass `go vet` with no errors
- Pass `go fmt` with no changes needed
- Pass `golangci-lint` checks
- Maintain 95%+ test coverage
- Include tests for new functionality
- Follow Go naming conventions and idioms

## Pull Request Process

1. Update documentation if needed
2. Ensure all checks pass
3. Create a pull request with a clear description
4. Link any related issues

### PR Title Format

- `feat: Add new feature`
- `fix: Fix bug description`
- `docs: Update documentation`
- `refactor: Refactor code`
- `test: Add tests`

## Branch Naming

| Prefix      | Purpose          |
|-------------|------------------|
| `feature/`  | New features     |
| `fix/`      | Bug fixes        |
| `docs/`     | Documentation    |
| `refactor/` | Code refactoring |
| `test/`     | Test additions   |

## Reporting Issues

When reporting issues, please include:

- Go version (`go version`)
- Operating system
- Minimal code example to reproduce
- Expected vs actual behavior

## Questions?

Feel free to open an issue for any questions or suggestions.
