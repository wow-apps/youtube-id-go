// Package yid generates YouTube-style short IDs from numbers.
// It provides a lightweight, fast, and reversible base62 encoder with optional obfuscation.
//
// Example:
//
//	encoded, _ := yid.ToAlphanumeric(12345)
//	fmt.Println(encoded) // -> "dnh"
//
//	decoded, _ := yid.ToNumeric("dnh")
//	fmt.Println(decoded) // -> 12345
package yid

import "github.com/wow-apps/youtube-id-go/internal/base62"

// Version is the current version of the package.
const Version = "1.0.0"

// config holds encoding/decoding configuration.
type config struct {
	padUp     int
	secureKey string
	transform Transform
}

// Option configures encoding/decoding behavior.
type Option func(*config)

// MaxPadUp is the maximum safe padUp value to avoid integer overflow.
const MaxPadUp = 11

// WithPadUp sets the padding value for minimum output length.
// Negative values are treated as 0. Values exceeding MaxPadUp (11) are clamped.
func WithPadUp(padUp int) Option {
	return func(c *config) {
		if padUp < 0 {
			padUp = 0
		} else if padUp > MaxPadUp {
			padUp = MaxPadUp
		}
		c.padUp = padUp
	}
}

// WithSecureKey sets the obfuscation key to shuffle the dictionary.
func WithSecureKey(key string) Option {
	return func(c *config) {
		c.secureKey = key
	}
}

// WithTransform sets the case transformation for encoding output.
func WithTransform(t Transform) Option {
	return func(c *config) {
		c.transform = t
	}
}

// defaultConfig returns the default configuration.
func defaultConfig() config {
	return config{
		padUp:     0,
		secureKey: "",
		transform: TransformNone,
	}
}

// ToAlphanumeric converts a number to a short alphanumeric string.
//
// Example:
//
//	yid.ToAlphanumeric(12345)                                    // -> "dnh"
//	yid.ToAlphanumeric(12345, yid.WithSecureKey("secret"))       // -> obfuscated
//	yid.ToAlphanumeric(12345, yid.WithTransform(yid.TransformUpper)) // -> "DNH"
func ToAlphanumeric(number int64, opts ...Option) (string, error) {
	if number < 0 {
		return "", ErrNegativeNumber
	}

	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	dictionary := base62.Dictionary
	if cfg.secureKey != "" {
		dictionary = base62.SecureDictionary(cfg.secureKey)
	}

	result := base62.Encode(number, dictionary, cfg.padUp)
	return applyCaseTransform(result, cfg.transform), nil
}

// ToNumeric converts an alphanumeric string back to a number.
// The input must be the raw (untransformed) value. If you encoded with
// WithTransform, you must decode using the original untransformed value.
// The WithTransform option is ignored by this function.
//
// Example:
//
//	yid.ToNumeric("dnh")                               // -> 12345
//	yid.ToNumeric(encoded, yid.WithSecureKey("secret")) // with same key used for encoding
func ToNumeric(alphanumeric string, opts ...Option) (int64, error) {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	dictionary := base62.Dictionary
	if cfg.secureKey != "" {
		dictionary = base62.SecureDictionary(cfg.secureKey)
	}

	result, err := base62.Decode(alphanumeric, dictionary, cfg.padUp)
	if err != nil {
		return 0, ErrInvalidCharacter
	}
	return result, nil
}
