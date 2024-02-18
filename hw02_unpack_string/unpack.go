package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")
var ErrParse = errors.New("invalid parse digit")

func Unpack(s string) (string, error) {
	var next_number bool
	var str_slice []string

	for i, v := range []rune(s) {
		if i < len(s)-1 {
			next_number = !unicode.IsDigit(rune(s[i+1]))
		}
		switch {
		case unicode.IsLetter(v):
			str_slice = append(str_slice, string(v))
		case unicode.IsSpace(v):
			str_slice = append(str_slice, strings.Trim(strconv.QuoteRune(v), "'"))
		case unicode.IsDigit(v) && i != 0 && next_number:
			digit, err := strconv.Atoi(string(v))
			if err != nil {
				return "", ErrParse
			}
			if digit == 0 {
				str_slice = str_slice[:len(str_slice)-1]
			} else {
				str_slice = append(str_slice, strings.Repeat(str_slice[i-1], digit-1))
			}
		default:
			return "", ErrInvalidString
		}
	}
	return strings.Join(str_slice, ""), nil
}
