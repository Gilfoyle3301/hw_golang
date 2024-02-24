package hw02unpackstring

import (
	"errors"
	"strconv"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var s_builder string
	var s_slice []rune = []rune(s)
	var slh bool
	for i, v := range s_slice {
		switch {
		case i == 0 && unicode.IsDigit(v):
			return "", ErrInvalidString
		case unicode.IsDigit(v) && unicode.IsDigit(s_slice[i-1]):
			return "", ErrInvalidString
		case v == '\\':
			s_builder += string(v)
			slh = true
			continue
		case unicode.IsDigit(v):
			digit, err := strconv.Atoi(string(v))
			if err != nil {
				return "", ErrInvalidString
			}
			if digit == 0 {
				s_builder = s_builder[:len(s_builder)-1]
				continue
			}
			if slh {
				digit = digit*2 - 1
			}
			for j := 0; j < digit-1; j++ {
				if slh {
					s_builder += string(s_builder[len(s_builder)-2])
					continue
				}
				s_builder += string(s_slice[i-1])
			}
			continue
		}
		s_builder += string(v)
	}
	return s_builder, nil
}
