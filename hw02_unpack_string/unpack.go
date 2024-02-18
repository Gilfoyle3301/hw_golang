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
	var nextNumber bool
	var strSlice []string

	for i, v := range []rune(s) {
		if i < len(s)-1 {
			nextNumber = !unicode.IsDigit(rune(s[i+1]))
		}
		switch {
		case unicode.IsLetter(v):
			strSlice = append(strSlice, string(v))
		case unicode.IsSpace(v):
			strSlice = append(strSlice, strings.Trim(strconv.QuoteRune(v), "'"))
		case unicode.IsDigit(v) && i != 0 && nextNumber:
			digit, err := strconv.Atoi(string(v))
			if err != nil {
				return "", ErrParse
			}
			if digit == 0 {
				strSlice = strSlice[:len(strSlice)-1]
			} else {
				strSlice = append(strSlice, strings.Repeat(strSlice[i-1], digit-1))
			}
		default:
			return "", ErrInvalidString
		}
	}
	return strings.Join(strSlice, ""), nil
}
