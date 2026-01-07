package yid

import "github.com/wow-apps/youtube-id-go/internal/base62"

// Encoder provides reusable encoding/decoding with preset options.
//
// Example:
//
//	enc := yid.New(yid.WithSecureKey("secret"), yid.WithTransform(yid.TransformUpper))
//	enc.Encode(12345)    // -> "HQJ" (transformed for display)
//	enc.EncodeRaw(12345) // -> "hqj" (raw for storage/decoding)
//	enc.Decode("hqj")    // -> 12345
type Encoder struct {
	padUp      int
	transform  Transform
	dictionary string
}

// New creates a new Encoder with the given options.
//
// Example:
//
//	enc := yid.New(yid.WithSecureKey("my-secret"), yid.WithTransform(yid.TransformUpper))
//	enc.Encode(12345) // -> "HQJ"
func New(opts ...Option) *Encoder {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	dictionary := base62.Dictionary
	if cfg.secureKey != "" {
		dictionary = base62.SecureDictionary(cfg.secureKey)
	}

	return &Encoder{
		padUp:      cfg.padUp,
		transform:  cfg.transform,
		dictionary: dictionary,
	}
}

// Encode converts a number to an alphanumeric string with transformation applied.
// Returns an error if number is negative.
func (e *Encoder) Encode(number int64) (string, error) {
	if number < 0 {
		return "", ErrNegativeNumber
	}
	result := base62.Encode(number, e.dictionary, e.padUp)
	return applyTransform(result, e.transform), nil
}

// EncodeRaw converts a number to an alphanumeric string without transformation.
// Use this when you need the raw value for storage or decoding.
// Returns an error if number is negative.
func (e *Encoder) EncodeRaw(number int64) (string, error) {
	if number < 0 {
		return "", ErrNegativeNumber
	}
	return base62.Encode(number, e.dictionary, e.padUp), nil
}

// Decode converts an alphanumeric string back to a number.
// Expects the raw (non-transformed) value from EncodeRaw().
func (e *Encoder) Decode(alphanumeric string) (int64, error) {
	result, err := base62.Decode(alphanumeric, e.dictionary, e.padUp)
	if err != nil {
		return 0, ErrInvalidCharacter
	}
	return result, nil
}
