package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if len(str) == 0 {
		return "", nil
	}
	builder := strings.Builder{}
	chars := []rune(str)
	if unicode.IsNumber(chars[0]) || !unicode.IsLetter(chars[0]) {
		return "", ErrInvalidString
	}
	for i := 0; i < len(chars)-1; i++ {
		if unicode.IsNumber(chars[i+1]) {
			if i+2 < len(chars) {
				if unicode.IsNumber(chars[i+2]) {
					return "", ErrInvalidString
				}
			}
			cnt := int(chars[i+1] - '0')
			builder.WriteString(strings.Repeat(string(chars[i]), cnt))
		} else {
			if !unicode.IsNumber(chars[i]) {
				builder.WriteRune(chars[i])
			}
		}
	}

	if unicode.IsNumber(chars[len(chars)-1]) {
		cnt := int(chars[len(chars)-1] - '0')
		builder.WriteString(strings.Repeat(string(chars[len(chars)-2]), cnt-1))
	} else {
		builder.WriteRune(chars[len(chars)-1])
	}

	return builder.String(), nil
}
