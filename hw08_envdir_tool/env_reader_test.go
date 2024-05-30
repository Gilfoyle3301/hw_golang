package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("standart", func(t *testing.T) {
		tc := Environment{
			"UNSET": {Value: "", NeedRemove: true},
			"HELLO": {Value: "\"hello\"", NeedRemove: false},
			"FOO":   {Value: "   foo\nwith new line", NeedRemove: false},
			"EMPTY": {Value: "", NeedRemove: false},
			"BAR":   {Value: "bar", NeedRemove: false},
		}
		testsData, err := ReadDir("./testdata/env")
		require.NoError(t, err)
		require.Equal(t, tc, testsData)
	})
	t.Run("empty dir", func(t *testing.T) {
		dir, err := os.MkdirTemp(".", "")
		require.NoError(t, err)
		env, err := ReadDir(dir)
		require.NoError(t, err)
		require.Len(t, env, 0)
	})
	t.Run("not exist dir", func(t *testing.T) {
		_, err := ReadDir("notExistDir")
		require.Error(t, os.ErrNotExist, err)
	})
}
