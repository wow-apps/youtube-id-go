package base62_test

import (
	"testing"

	"github.com/wow-apps/youtube-id-go/internal/base62"
)

func TestEncode_Basic(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{0, "a"},
		{1, "b"},
		{61, "Z"},
		{62, "ba"},
		{12345, "dnh"},
	}
	for _, tt := range tests {
		result := base62.Encode(tt.input, base62.Dictionary, 0)
		if result != tt.expected {
			t.Errorf("Encode(%d) = '%s', want '%s'", tt.input, result, tt.expected)
		}
	}
}

func TestEncode_WithPadUp(t *testing.T) {
	resultNoPad := base62.Encode(1, base62.Dictionary, 0)
	resultWithPad := base62.Encode(1, base62.Dictionary, 3)
	if resultNoPad == resultWithPad {
		t.Error("pad_up should change output")
	}
}

func TestDecode_Basic(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"a", 0},
		{"b", 1},
		{"Z", 61},
		{"ba", 62},
		{"dnh", 12345},
	}
	for _, tt := range tests {
		result, err := base62.Decode(tt.input, base62.Dictionary, 0)
		if err != nil {
			t.Fatalf("Decode('%s') error: %v", tt.input, err)
		}
		if result != tt.expected {
			t.Errorf("Decode('%s') = %d, want %d", tt.input, result, tt.expected)
		}
	}
}

func TestDecode_InvalidCharacter(t *testing.T) {
	_, err := base62.Decode("abc!", base62.Dictionary, 0)
	if err == nil {
		t.Error("expected error for invalid character")
	}
}

func TestDecode_WithPadUp(t *testing.T) {
	encoded := base62.Encode(100, base62.Dictionary, 3)
	decoded, err := base62.Decode(encoded, base62.Dictionary, 3)
	if err != nil {
		t.Fatalf("Decode error: %v", err)
	}
	if decoded != 100 {
		t.Errorf("expected 100, got %d", decoded)
	}
}

func TestSecureDictionary(t *testing.T) {
	dict1 := base62.SecureDictionary("key1")
	dict2 := base62.SecureDictionary("key2")

	// Different keys should produce different dictionaries
	if dict1 == dict2 {
		t.Error("different keys should produce different dictionaries")
	}

	// Same key should produce same dictionary
	dict1Again := base62.SecureDictionary("key1")
	if dict1 != dict1Again {
		t.Error("same key should produce same dictionary")
	}

	// Dictionary should have correct length
	if len(dict1) != base62.DictLen {
		t.Errorf("dictionary length = %d, want %d", len(dict1), base62.DictLen)
	}

	// Dictionary should contain all original characters
	charCount := make(map[byte]int)
	for i := 0; i < len(dict1); i++ {
		charCount[dict1[i]]++
	}
	for i := 0; i < len(base62.Dictionary); i++ {
		if charCount[base62.Dictionary[i]] != 1 {
			t.Errorf("character '%c' appears %d times in shuffled dictionary", base62.Dictionary[i], charCount[base62.Dictionary[i]])
		}
	}
}

func TestRoundtrip(t *testing.T) {
	testNumbers := []int64{0, 1, 10, 100, 1000, 12345, 999999, 1000000000}
	for _, num := range testNumbers {
		encoded := base62.Encode(num, base62.Dictionary, 0)
		decoded, err := base62.Decode(encoded, base62.Dictionary, 0)
		if err != nil {
			t.Fatalf("Decode error for %d: %v", num, err)
		}
		if decoded != num {
			t.Errorf("roundtrip failed: %d -> %s -> %d", num, encoded, decoded)
		}
	}
}

func TestRoundtrip_WithSecureKey(t *testing.T) {
	dict := base62.SecureDictionary("test-key")
	testNumbers := []int64{0, 1, 100, 12345, 999999}
	for _, num := range testNumbers {
		encoded := base62.Encode(num, dict, 0)
		decoded, err := base62.Decode(encoded, dict, 0)
		if err != nil {
			t.Fatalf("Decode error for %d: %v", num, err)
		}
		if decoded != num {
			t.Errorf("roundtrip failed: %d -> %s -> %d", num, encoded, decoded)
		}
	}
}

func TestDictionaryConstant(t *testing.T) {
	if len(base62.Dictionary) != 62 {
		t.Errorf("Dictionary length = %d, want 62", len(base62.Dictionary))
	}
	if base62.DictLen != 62 {
		t.Errorf("DictLen = %d, want 62", base62.DictLen)
	}
}
