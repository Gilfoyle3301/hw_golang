package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})
	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()
	wg.Wait()
}

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

		for _, caseDo := range testCase {
			cache.Set(Key(caseDo.key), caseDo.value)
		}

		valueGet, inCache := cache.Get("twoEl")
		require.False(t, inCache)
		require.Nil(t, valueGet)
	})
}
