package hw02unpackstring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSomeChar(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{input: "я9", expected: "яяяяяяяяя"},
		{input: "สวัสดี", expected: "สวัสดี"},
		{input: "สวัส4ดี", expected: "สวัสสสสดี"},
		{input: "😀0", expected: ""},
		{input: "😀9", expected: "😀😀😀😀😀😀😀😀😀"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run("unpack simbol", func(t *testing.T) {
			unpackString, err := Unpack(tc.input)
			if err != nil {
				require.Error(t, err, "can't unpack string")
			}
			require.Equal(t, tc.expected, unpackString)
		})
	}
}
