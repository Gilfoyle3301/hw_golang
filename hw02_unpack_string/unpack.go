package hw02unpackstring

import (
	"errors"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var (
		sBuilder string
		isSlash  bool
	)
	for i, v := range []rune(s) {
		switch {
		case i == 0 && unicode.IsDigit(v):
			return "", ErrInvalidString
		case unicode.IsDigit(v) && unicode.IsDigit([]rune(s)[i-1]):
			return "", ErrInvalidString
		case v == '\\':
			sBuilder += string(v)
			isSlash = true
			continue
		case unicode.IsDigit(v):
			digit := int(v - '0')
			if digit == 0 {
				sBuilder = sBuilder[:len(sBuilder)-1]
				continue
			}
			if isSlash {
				digit = digit*2 - 1
			}
			for j := 0; j < digit-1; j++ {
				if isSlash {
					sBuilder += string(sBuilder[len(sBuilder)-2])
					continue
				}
				sBuilder += string([]rune(s)[i-1])
			}
			continue
		}
		sBuilder += string(v)
	}
	return sBuilder, nil
}
