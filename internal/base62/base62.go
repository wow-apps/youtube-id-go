// Package base62 provides core base62 encoding/decoding algorithms.
package base62

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"sort"
	"strings"
)

// Dictionary: a-z + 0-9 + A-Z (62 characters)
const Dictionary = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// DictLen is the length of the dictionary (62).
const DictLen = 62

// ErrInvalidCharacter is returned when decoding encounters an invalid character.
var ErrInvalidCharacter = errors.New("base62: invalid character in input")

// pow calculates base^exp using integer arithmetic.
func pow(base, exp int) int64 {
	result := int64(1)
	b := int64(base)
	for i := 0; i < exp; i++ {
		result *= b
	}
	return result
}

// Encode converts a number to a base62 string using the given dictionary.
func Encode(number int64, dictionary string, padUp int) string {
	if padUp > 1 {
		number += pow(DictLen, padUp-1)
	}

	if number == 0 {
		return string(dictionary[0])
	}

	// Find the highest power of 62 that fits in the number
	t := 0
	temp := number
	for temp >= DictLen {
		temp /= DictLen
		t++
	}

	var result strings.Builder
	for t >= 0 {
		bcp := pow(DictLen, t)
		index := (number / bcp) % int64(DictLen)
		result.WriteByte(dictionary[index])
		number -= index * bcp
		t--
	}

	return result.String()
}

// Decode converts a base62 string back to a number.
func Decode(alphanumeric, dictionary string, padUp int) (int64, error) {
	var result int64
	length := len(alphanumeric)
	multiplier := int64(1)

	for i := 0; i < length; i++ {
		char := alphanumeric[length-1-i]
		index := strings.IndexByte(dictionary, char)
		if index == -1 {
			return 0, ErrInvalidCharacter
		}
		result += int64(index) * multiplier
		multiplier *= int64(DictLen)
	}

	if padUp > 1 {
		result -= pow(DictLen, padUp-1)
	}

	return result, nil
}

// charPair holds a hash character and its corresponding dictionary character.
type charPair struct {
	hashChar byte
	dictChar byte
}

// SecureDictionary shuffles the dictionary based on a secure key.
// This makes it harder to calculate the corresponding numeric ID without knowing the key.
func SecureDictionary(secureKey string) string {
	hash := sha256.Sum256([]byte(secureKey))
	hashHex := hex.EncodeToString(hash[:])

	// Use SHA512 if SHA256 hex is too short (SHA256 hex is 64 chars, so this won't happen,
	// but we match Python logic for consistency)
	if len(hashHex) < DictLen {
		hash512 := sha512.Sum512([]byte(secureKey))
		hashHex = hex.EncodeToString(hash512[:])
	}

	// Create pairs of hash char and dictionary char
	pairs := make([]charPair, DictLen)
	for i := 0; i < DictLen; i++ {
		pairs[i] = charPair{
			hashChar: hashHex[i],
			dictChar: Dictionary[i],
		}
	}

	// Sort by hash char in descending order
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].hashChar > pairs[j].hashChar
	})

	// Build result
	result := make([]byte, DictLen)
	for i, p := range pairs {
		result[i] = p.dictChar
	}

	return string(result)
}
