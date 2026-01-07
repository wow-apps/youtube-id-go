package yid_test

import (
	"testing"

	"github.com/wow-apps/youtube-id-go"
)

// TestToAlphanumeric_Basic tests basic number to alphanumeric conversion.
func TestToAlphanumeric_Basic(t *testing.T) {
	result, err := yid.ToAlphanumeric(12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "dnh" {
		t.Errorf("expected 'dnh', got '%s'", result)
	}
}

// TestToAlphanumeric_Zero tests conversion of zero.
func TestToAlphanumeric_Zero(t *testing.T) {
	result, err := yid.ToAlphanumeric(0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "a" {
		t.Errorf("expected 'a', got '%s'", result)
	}
}

// TestToAlphanumeric_One tests conversion of one.
func TestToAlphanumeric_One(t *testing.T) {
	result, err := yid.ToAlphanumeric(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "b" {
		t.Errorf("expected 'b', got '%s'", result)
	}
}

// TestToAlphanumeric_Boundaries tests boundary of single character output.
func TestToAlphanumeric_Boundaries(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{61, "Z"},
		{62, "ba"},
	}
	for _, tt := range tests {
		result, err := yid.ToAlphanumeric(tt.input)
		if err != nil {
			t.Fatalf("unexpected error for %d: %v", tt.input, err)
		}
		if result != tt.expected {
			t.Errorf("for %d: expected '%s', got '%s'", tt.input, tt.expected, result)
		}
	}
}

// TestToAlphanumeric_LargeNumber tests conversion of large numbers.
func TestToAlphanumeric_LargeNumber(t *testing.T) {
	result, err := yid.ToAlphanumeric(999999999)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) == 0 {
		t.Error("expected non-empty result")
	}
}

// TestToAlphanumeric_WithSecureKey tests conversion with secure key.
func TestToAlphanumeric_WithSecureKey(t *testing.T) {
	withKey, err := yid.ToAlphanumeric(12345, yid.WithSecureKey("my-secret"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	withoutKey, err := yid.ToAlphanumeric(12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if withKey == withoutKey {
		t.Error("secure key should produce different output")
	}
}

// TestToAlphanumeric_DifferentSecureKeys tests that different keys produce different outputs.
func TestToAlphanumeric_DifferentSecureKeys(t *testing.T) {
	result1, _ := yid.ToAlphanumeric(12345, yid.WithSecureKey("key1"))
	result2, _ := yid.ToAlphanumeric(12345, yid.WithSecureKey("key2"))
	if result1 == result2 {
		t.Error("different keys should produce different outputs")
	}
}

// TestToAlphanumeric_TransformUpper tests uppercase transformation.
func TestToAlphanumeric_TransformUpper(t *testing.T) {
	result, err := yid.ToAlphanumeric(12345, yid.WithTransform(yid.TransformUpper))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "DNH" {
		t.Errorf("expected 'DNH', got '%s'", result)
	}
}

// TestToAlphanumeric_TransformLower tests lowercase transformation.
func TestToAlphanumeric_TransformLower(t *testing.T) {
	result, err := yid.ToAlphanumeric(12345, yid.WithTransform(yid.TransformLower))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "dnh" {
		t.Errorf("expected 'dnh', got '%s'", result)
	}
}

// TestToAlphanumeric_TransformNone tests no transformation.
func TestToAlphanumeric_TransformNone(t *testing.T) {
	result, err := yid.ToAlphanumeric(12345, yid.WithTransform(yid.TransformNone))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "dnh" {
		t.Errorf("expected 'dnh', got '%s'", result)
	}
}

// TestToAlphanumeric_WithPadUp tests conversion with pad_up parameter.
func TestToAlphanumeric_WithPadUp(t *testing.T) {
	resultNoPad, _ := yid.ToAlphanumeric(1)
	resultWithPad, _ := yid.ToAlphanumeric(1, yid.WithPadUp(3))
	if resultNoPad == resultWithPad {
		t.Error("pad_up should change output")
	}
}

// TestToAlphanumeric_CombinedOptions tests conversion with multiple options.
func TestToAlphanumeric_CombinedOptions(t *testing.T) {
	result, err := yid.ToAlphanumeric(12345, yid.WithSecureKey("secret"), yid.WithTransform(yid.TransformUpper))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Check it's uppercase
	for _, c := range result {
		if c >= 'a' && c <= 'z' {
			t.Errorf("expected uppercase, got lowercase in '%s'", result)
			break
		}
	}
}

// TestToAlphanumeric_NegativeNumber tests that negative numbers return error.
func TestToAlphanumeric_NegativeNumber(t *testing.T) {
	_, err := yid.ToAlphanumeric(-1)
	if err == nil {
		t.Error("expected error for negative number")
	}
	if err != yid.ErrNegativeNumber {
		t.Errorf("expected ErrNegativeNumber, got %v", err)
	}
}

// TestToNumeric_Basic tests basic alphanumeric to number conversion.
func TestToNumeric_Basic(t *testing.T) {
	result, err := yid.ToNumeric("dnh")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 12345 {
		t.Errorf("expected 12345, got %d", result)
	}
}

// TestToNumeric_Zero tests conversion of 'a' back to zero.
func TestToNumeric_Zero(t *testing.T) {
	result, err := yid.ToNumeric("a")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 0 {
		t.Errorf("expected 0, got %d", result)
	}
}

// TestToNumeric_One tests conversion of 'b' back to one.
func TestToNumeric_One(t *testing.T) {
	result, err := yid.ToNumeric("b")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 1 {
		t.Errorf("expected 1, got %d", result)
	}
}

// TestToNumeric_Boundaries tests boundary values.
func TestToNumeric_Boundaries(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"Z", 61},
		{"ba", 62},
	}
	for _, tt := range tests {
		result, err := yid.ToNumeric(tt.input)
		if err != nil {
			t.Fatalf("unexpected error for '%s': %v", tt.input, err)
		}
		if result != tt.expected {
			t.Errorf("for '%s': expected %d, got %d", tt.input, tt.expected, result)
		}
	}
}

// TestToNumeric_WithSecureKey tests conversion with secure key.
func TestToNumeric_WithSecureKey(t *testing.T) {
	encoded, _ := yid.ToAlphanumeric(12345, yid.WithSecureKey("my-secret"))
	decoded, err := yid.ToNumeric(encoded, yid.WithSecureKey("my-secret"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if decoded != 12345 {
		t.Errorf("expected 12345, got %d", decoded)
	}
}

// TestToNumeric_WrongSecureKey tests that wrong secure key produces wrong result.
func TestToNumeric_WrongSecureKey(t *testing.T) {
	encoded, _ := yid.ToAlphanumeric(12345, yid.WithSecureKey("correct-key"))
	decoded, _ := yid.ToNumeric(encoded, yid.WithSecureKey("wrong-key"))
	if decoded == 12345 {
		t.Error("wrong key should produce wrong result")
	}
}

// TestToNumeric_WithPadUp tests conversion with pad_up parameter.
func TestToNumeric_WithPadUp(t *testing.T) {
	encoded, _ := yid.ToAlphanumeric(100, yid.WithPadUp(3))
	decoded, err := yid.ToNumeric(encoded, yid.WithPadUp(3))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if decoded != 100 {
		t.Errorf("expected 100, got %d", decoded)
	}
}

// TestToNumeric_InvalidCharacter tests that invalid characters return error.
func TestToNumeric_InvalidCharacter(t *testing.T) {
	_, err := yid.ToNumeric("abc!")
	if err == nil {
		t.Error("expected error for invalid character")
	}
	if err != yid.ErrInvalidCharacter {
		t.Errorf("expected ErrInvalidCharacter, got %v", err)
	}
}

// TestRoundtrip tests encoding and decoding produces original value.
func TestRoundtrip(t *testing.T) {
	testNumbers := []int64{0, 1, 10, 100, 1000, 12345, 999999}
	for _, num := range testNumbers {
		encoded, err := yid.ToAlphanumeric(num)
		if err != nil {
			t.Fatalf("encoding %d: %v", num, err)
		}
		decoded, err := yid.ToNumeric(encoded)
		if err != nil {
			t.Fatalf("decoding %s: %v", encoded, err)
		}
		if decoded != num {
			t.Errorf("roundtrip failed: %d -> %s -> %d", num, encoded, decoded)
		}
	}
}

// TestRoundtrip_WithSecureKey tests roundtrip with secure key.
func TestRoundtrip_WithSecureKey(t *testing.T) {
	testNumbers := []int64{0, 1, 100, 12345, 999999}
	for _, num := range testNumbers {
		encoded, _ := yid.ToAlphanumeric(num, yid.WithSecureKey("test-key"))
		decoded, _ := yid.ToNumeric(encoded, yid.WithSecureKey("test-key"))
		if decoded != num {
			t.Errorf("roundtrip failed: %d -> %s -> %d", num, encoded, decoded)
		}
	}
}

// TestTransform_ValuesExist tests that all transform values exist.
func TestTransform_ValuesExist(t *testing.T) {
	transforms := []yid.Transform{yid.TransformNone, yid.TransformUpper, yid.TransformLower}
	for i, tr := range transforms {
		if int(tr) != i {
			t.Errorf("unexpected transform value: %d", tr)
		}
	}
}

// TestTransform_ValuesUnique tests that transform values are unique.
func TestTransform_ValuesUnique(t *testing.T) {
	seen := make(map[yid.Transform]bool)
	transforms := []yid.Transform{yid.TransformNone, yid.TransformUpper, yid.TransformLower}
	for _, tr := range transforms {
		if seen[tr] {
			t.Errorf("duplicate transform value: %d", tr)
		}
		seen[tr] = true
	}
}

// TestEncoder_BasicEncode tests basic encoding with Encoder.
func TestEncoder_BasicEncode(t *testing.T) {
	enc := yid.New()
	result, err := enc.Encode(12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "dnh" {
		t.Errorf("expected 'dnh', got '%s'", result)
	}
}

// TestEncoder_BasicDecode tests basic decoding with Encoder.
func TestEncoder_BasicDecode(t *testing.T) {
	enc := yid.New()
	result, err := enc.Decode("dnh")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 12345 {
		t.Errorf("expected 12345, got %d", result)
	}
}

// TestEncoder_EncodeRaw tests encode_raw method.
func TestEncoder_EncodeRaw(t *testing.T) {
	enc := yid.New(yid.WithTransform(yid.TransformUpper))
	raw, err := enc.EncodeRaw(12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	display, err := enc.Encode(12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if raw != "dnh" {
		t.Errorf("expected raw 'dnh', got '%s'", raw)
	}
	if display != "DNH" {
		t.Errorf("expected display 'DNH', got '%s'", display)
	}
}

// TestEncoder_WithSecureKey tests Encoder with secure key.
func TestEncoder_WithSecureKey(t *testing.T) {
	enc := yid.New(yid.WithSecureKey("secret"))
	encoded, err := enc.Encode(12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	decoded, err := enc.Decode(encoded)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if decoded != 12345 {
		t.Errorf("expected 12345, got %d", decoded)
	}
}

// TestEncoder_WithTransform tests Encoder with transform.
func TestEncoder_WithTransform(t *testing.T) {
	enc := yid.New(yid.WithTransform(yid.TransformUpper))
	result, err := enc.Encode(12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, c := range result {
		if c >= 'a' && c <= 'z' {
			t.Errorf("expected uppercase, got lowercase in '%s'", result)
			break
		}
	}
}

// TestEncoder_WithPadUp tests Encoder with pad_up.
func TestEncoder_WithPadUp(t *testing.T) {
	enc := yid.New(yid.WithPadUp(3))
	encoded, err := enc.Encode(100)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	decoded, err := enc.Decode(encoded)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if decoded != 100 {
		t.Errorf("expected 100, got %d", decoded)
	}
}

// TestEncoder_Roundtrip tests Encoder roundtrip.
func TestEncoder_Roundtrip(t *testing.T) {
	enc := yid.New(yid.WithSecureKey("test"), yid.WithPadUp(2))
	testNumbers := []int64{0, 1, 100, 12345}
	for _, num := range testNumbers {
		raw, err := enc.EncodeRaw(num)
		if err != nil {
			t.Fatalf("encoding %d: %v", num, err)
		}
		decoded, err := enc.Decode(raw)
		if err != nil {
			t.Fatalf("decoding %s: %v", raw, err)
		}
		if decoded != num {
			t.Errorf("roundtrip failed: %d -> %s -> %d", num, raw, decoded)
		}
	}
}

// TestEncoder_MultipleEncodersIndependence tests that multiple encoders work independently.
func TestEncoder_MultipleEncodersIndependence(t *testing.T) {
	enc1 := yid.New(yid.WithSecureKey("key1"))
	enc2 := yid.New(yid.WithSecureKey("key2"))
	result1, _ := enc1.Encode(12345)
	result2, _ := enc2.Encode(12345)
	if result1 == result2 {
		t.Error("different encoders with different keys should produce different outputs")
	}
}

// TestEncoder_NegativeNumber tests that negative numbers return error.
func TestEncoder_NegativeNumber(t *testing.T) {
	enc := yid.New()
	_, err := enc.Encode(-1)
	if err == nil {
		t.Error("expected error for negative number")
	}
	_, err = enc.EncodeRaw(-1)
	if err == nil {
		t.Error("expected error for negative number in EncodeRaw")
	}
}

// TestEncoder_InvalidCharacter tests that invalid characters return error.
func TestEncoder_InvalidCharacter(t *testing.T) {
	enc := yid.New()
	_, err := enc.Decode("abc!")
	if err == nil {
		t.Error("expected error for invalid character")
	}
}

// TestVersion tests that Version constant exists.
func TestVersion(t *testing.T) {
	if yid.Version == "" {
		t.Error("Version should not be empty")
	}
}

// TestEdgeCases_VeryLargeNumber tests with very large numbers.
func TestEdgeCases_VeryLargeNumber(t *testing.T) {
	largeNum := int64(1000000000000000) // 10^15
	encoded, err := yid.ToAlphanumeric(largeNum)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	decoded, err := yid.ToNumeric(encoded)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if decoded != largeNum {
		t.Errorf("expected %d, got %d", largeNum, decoded)
	}
}

// TestEdgeCases_EmptySecureKey tests that empty string secure key works like no key.
func TestEdgeCases_EmptySecureKey(t *testing.T) {
	resultNoKey, _ := yid.ToAlphanumeric(12345)
	resultEmptyKey, _ := yid.ToAlphanumeric(12345, yid.WithSecureKey(""))
	if resultNoKey != resultEmptyKey {
		t.Error("empty secure key should be same as no key")
	}
}

// TestEdgeCases_SpecialCharactersInSecureKey tests secure key with special characters.
func TestEdgeCases_SpecialCharactersInSecureKey(t *testing.T) {
	encoded, _ := yid.ToAlphanumeric(12345, yid.WithSecureKey("!@#$%^&*()"))
	decoded, _ := yid.ToNumeric(encoded, yid.WithSecureKey("!@#$%^&*()"))
	if decoded != 12345 {
		t.Errorf("expected 12345, got %d", decoded)
	}
}

// TestEdgeCases_UnicodeSecureKey tests secure key with unicode characters.
func TestEdgeCases_UnicodeSecureKey(t *testing.T) {
	encoded, _ := yid.ToAlphanumeric(12345, yid.WithSecureKey("ключ"))
	decoded, _ := yid.ToNumeric(encoded, yid.WithSecureKey("ключ"))
	if decoded != 12345 {
		t.Errorf("expected 12345, got %d", decoded)
	}
}

// TestEdgeCases_LongSecureKey tests with very long secure key.
func TestEdgeCases_LongSecureKey(t *testing.T) {
	longKey := ""
	for i := 0; i < 1000; i++ {
		longKey += "a"
	}
	encoded, _ := yid.ToAlphanumeric(12345, yid.WithSecureKey(longKey))
	decoded, _ := yid.ToNumeric(encoded, yid.WithSecureKey(longKey))
	if decoded != 12345 {
		t.Errorf("expected 12345, got %d", decoded)
	}
}

// TestEdgeCases_ConsecutiveNumbers tests that consecutive numbers produce different outputs.
func TestEdgeCases_ConsecutiveNumbers(t *testing.T) {
	seen := make(map[string]bool)
	for i := int64(0); i < 100; i++ {
		result, _ := yid.ToAlphanumeric(i)
		if seen[result] {
			t.Errorf("duplicate output for %d: %s", i, result)
		}
		seen[result] = true
	}
}

// TestEdgeCases_PadUpZero tests pad_up with zero value.
func TestEdgeCases_PadUpZero(t *testing.T) {
	result, _ := yid.ToAlphanumeric(12345, yid.WithPadUp(0))
	if result != "dnh" {
		t.Errorf("expected 'dnh', got '%s'", result)
	}
}

// TestEdgeCases_PadUpOne tests pad_up with one (should be same as zero).
func TestEdgeCases_PadUpOne(t *testing.T) {
	result0, _ := yid.ToAlphanumeric(12345, yid.WithPadUp(0))
	result1, _ := yid.ToAlphanumeric(12345, yid.WithPadUp(1))
	if result0 != result1 {
		t.Errorf("pad_up=0 and pad_up=1 should be same, got '%s' and '%s'", result0, result1)
	}
}

// TestEdgeCases_DictionaryCharacters tests that output only contains valid dictionary characters.
func TestEdgeCases_DictionaryCharacters(t *testing.T) {
	validChars := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	isValid := func(c rune) bool {
		for _, vc := range validChars {
			if c == vc {
				return true
			}
		}
		return false
	}

	testNumbers := []int64{0, 1, 62, 100, 12345, 999999}
	for _, num := range testNumbers {
		result, _ := yid.ToAlphanumeric(num)
		for _, c := range result {
			if !isValid(c) {
				t.Errorf("invalid character '%c' in output for %d", c, num)
			}
		}
	}
}
