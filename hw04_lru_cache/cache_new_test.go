package hw04lrucache

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestNewCache(t *testing.T) {
	t.Run("simple check work method Set", func(t *testing.T) {
		cache := NewCache(3)
		wasInCache := cache.Set("one", 111)
		require.False(t, wasInCache)
		wasInCache = cache.Set("two", 222)
		require.False(t, wasInCache)
		wasInCache = cache.Set("three", 333)
		require.False(t, wasInCache)
		wasInCache = cache.Set("one", 111)
		require.True(t, wasInCache)
	})

	t.Run("simple check method Get", func(t *testing.T) {
		cache := NewCache(2)
		emptyGet, ok := cache.Get("")
		require.Empty(t, emptyGet)
		require.False(t, ok)

		cache.Set("one", 111)
		cache.Set("two", 222)

		varGet, ok := cache.Get("one")

		require.Equal(t, 111, varGet)
		require.True(t, ok)

	})

	t.Run("purge logic", func(t *testing.T) {
		cache := NewCache(3)
		testCase := []struct {
			key   string
			value interface{}
		}{
			{key: "oneEl", value: 122},
			{key: "twoEl", value: 546},
			{key: "threeEl", value: 5421},
			{key: "fourEl", value: 542},
		}

		for _, case_do := range testCase {
			cache.Set(Key(case_do.key), case_do.value)
		}
		valueGet, inCache := cache.Get("twoEl")
		require.False(t, inCache)
		require.Nil(t, valueGet)
	})
}
