package yid

import "errors"

// ErrInvalidCharacter is returned when decoding encounters an invalid character.
var ErrInvalidCharacter = errors.New("yid: invalid character in input")

// ErrNegativeNumber is returned when encoding a negative number.
var ErrNegativeNumber = errors.New("yid: negative numbers are not supported")
