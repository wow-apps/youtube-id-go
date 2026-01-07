package yid

import "strings"

// Transform specifies case transformation for encoding output.
type Transform int

const (
	// TransformNone applies no transformation (default).
	TransformNone Transform = iota
	// TransformUpper converts output to uppercase.
	TransformUpper
	// TransformLower converts output to lowercase.
	TransformLower
)

// applyTransform applies the case transformation to the value.
func applyTransform(value string, t Transform) string {
	switch t {
	case TransformUpper:
		return strings.ToUpper(value)
	case TransformLower:
		return strings.ToLower(value)
	default:
		return value
	}
}
