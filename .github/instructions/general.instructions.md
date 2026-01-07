---
applyTo: "**/*"
---

# General Project Instructions

## Project Overview

youtube-id-go is a Go library for generating YouTube-style short IDs from numbers.
Lightweight, fast, and reversible base62 encoder with optional obfuscation.

## Tech Stack

- Go 1.21+
- golangci-lint for linting
- Standard testing package with race detection
- markdownlint for markdown linting

## Quality Requirements

All code must pass automated quality checks before merging:

1. `go vet ./...` - no issues
2. `go fmt ./...` - no formatting changes needed
3. `golangci-lint run` - no linting errors
4. `go test -race -coverprofile=coverage.out ./...` - tests pass with 95%+ coverage

## Branch Naming

- `feature/` - new features
- `fix/` - bug fixes
- `docs/` - documentation updates
- `refactor/` - code refactoring
- `test/` - test additions

## PR Title Format

Must start with one of: `feature/`, `fix/`, `docs/`, `refactor/`, `test/`
(case-insensitive)

## Important Notes

- Do NOT suggest changes that conflict with gofmt formatting
- gofmt is the authority for Go code style, not manual preferences
- All exported functions must have doc comments
- Zero external dependencies in production code
