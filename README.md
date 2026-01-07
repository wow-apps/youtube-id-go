# YouTubeID for Go

Generate YouTube-style short IDs from numbers. Lightweight, fast, and reversible base62 encoder with optional obfuscation.

[![Go Reference](https://pkg.go.dev/badge/github.com/wow-apps/youtube-id-go.svg)](https://pkg.go.dev/github.com/wow-apps/youtube-id-go)
[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Coverage](https://img.shields.io/badge/coverage-95%25-brightgreen.svg)](https://github.com/wow-apps/youtube-id-go)

## Other programming languages

- **PHP**: [kvz/youtube-id](https://github.com/kvz/youtube-id)
- **Python**: [wow-apps/youtube-id-py](https://github.com/wow-apps/youtube-id-py/)
- **TypeScript**: [wow-apps/youtube-id-ts](https://github.com/wow-apps/youtube-id-ts/)
- **Go**: [wow-apps/youtube-id-go](https://github.com/wow-apps/youtube-id-go/)

## Features

- **Lightweight** - Zero dependencies, pure Go
- **Fast** - Simple base62 encoding/decoding
- **Reversible** - Encode and decode without data loss
- **Obfuscation** - Optional secure key to shuffle the dictionary
- **Type-safe** - Idiomatic Go with proper error handling

## Installation

```bash
go get github.com/wow-apps/youtube-id-go
```

## Quick Start

```go
package main

import (
    "fmt"
    yid "github.com/wow-apps/youtube-id-go"
)

func main() {
    // Encode a number to a short string
    encoded, _ := yid.ToAlphanumeric(12345)
    fmt.Println(encoded) // -> "dnh"

    // Decode back to number
    decoded, _ := yid.ToNumeric("dnh")
    fmt.Println(decoded) // -> 12345
}
```

## Usage

### Basic Encoding/Decoding

```go
import yid "github.com/wow-apps/youtube-id-go"

// Number to alphanumeric
yid.ToAlphanumeric(0)        // -> "a", nil
yid.ToAlphanumeric(61)       // -> "Z", nil
yid.ToAlphanumeric(62)       // -> "ba", nil
yid.ToAlphanumeric(12345)    // -> "dnh", nil
yid.ToAlphanumeric(999999)   // -> "eGGf", nil

// Alphanumeric to number
yid.ToNumeric("dnh")         // -> 12345, nil
```

### With Secure Key (Obfuscation)

Use a secure key to shuffle the dictionary, making IDs harder to predict:

```go
import yid "github.com/wow-apps/youtube-id-go"

// Without key
yid.ToAlphanumeric(12345)                              // -> "dnh"

// With secure key (different output)
yid.ToAlphanumeric(12345, yid.WithSecureKey("secret")) // -> different value

// Decode with the same key
yid.ToNumeric(encoded, yid.WithSecureKey("secret"))    // -> 12345
```

### Case Transformation

```go
import yid "github.com/wow-apps/youtube-id-go"

yid.ToAlphanumeric(12345, yid.WithTransform(yid.TransformUpper)) // -> "DNH"
yid.ToAlphanumeric(12345, yid.WithTransform(yid.TransformLower)) // -> "dnh"
```

### Encoder for Repeated Operations

For repeated operations with the same settings, use the `Encoder`:

```go
import yid "github.com/wow-apps/youtube-id-go"

// Create encoder with preset options
enc := yid.New(yid.WithSecureKey("my-secret"), yid.WithTransform(yid.TransformUpper))

// Encode
enc.Encode(12345)     // -> "HQJ" (transformed for display)
enc.EncodeRaw(12345)  // -> "hqj" (raw for storage/decoding)

// Decode
enc.Decode("hqj")     // -> 12345
```

## API Reference

### Functions

#### `ToAlphanumeric(number int64, opts ...Option) (string, error)`

Convert a number to a short alphanumeric string.

#### `ToNumeric(alphanumeric string, opts ...Option) (int64, error)`

Convert an alphanumeric string back to a number.

#### `New(opts ...Option) *Encoder`

Create a reusable `Encoder` instance with preset options.

### Options

| Option                    | Description               |
|---------------------------|---------------------------|
| `WithPadUp(int)`          | Padding value             |
| `WithSecureKey(string)`   | Key to shuffle dictionary |
| `WithTransform(Transform)`| Case transformation       |

### Encoder Methods

| Method                 | Description                                     |
|------------------------|-------------------------------------------------|
| `Encode(number)`       | Convert number to alphanumeric (with transform) |
| `EncodeRaw(number)`    | Convert number to alphanumeric (no transform)   |
| `Decode(alphanumeric)` | Convert alphanumeric to number                  |

### Transform Constants

| Constant         | Description          |
|------------------|----------------------|
| `TransformNone`  | No transformation    |
| `TransformUpper` | Uppercase output     |
| `TransformLower` | Lowercase output     |

### Errors

| Error                 | Description                          |
|-----------------------|--------------------------------------|
| `ErrNegativeNumber`   | Input number is negative             |
| `ErrInvalidCharacter` | Input contains invalid character     |

## Use Cases

- **URL shorteners** - Convert database IDs to short URLs
- **Public IDs** - Hide sequential database IDs from users
- **Share codes** - Generate readable codes for sharing
- **Invite links** - Create short invitation tokens

## Performance

The library uses base62 encoding (a-z, 0-9, A-Z) which provides:

| Number Range         | Output Length |
|----------------------|---------------|
| 0 - 61               | 1 character   |
| 62 - 3,843           | 2 characters  |
| 3,844 - 238,327      | 3 characters  |
| 238,328 - 14,776,335 | 4 characters  |

## Contributing

Contributions are welcome! Please read our [Contributing Guidelines](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCT.md).

```bash
# Clone and setup
git clone https://github.com/wow-apps/youtube-id-go.git
cd youtube-id-go

# Run tests
go test -v -race ./...

# Run tests with coverage
go test -v -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Credits

A Go port of the YouTube-style ID generator originally created by [Kevin van Zonneveld](https://github.com/kvz) and contributors.

## License

[MIT](LICENSE) (c) Oleksii Samara, Kevin van Zonneveld
