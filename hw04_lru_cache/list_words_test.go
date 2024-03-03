package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListWords(t *testing.T) {
	t.Run("lenght list", func(t *testing.T) {
		list := NewList()
		require.Equal(t, 0, list.Len())
	})
	t.Run("Push words", func(t *testing.T) {
		list := NewList()

		list.PushFront("one")
		require.Equal(t, 1, list.Len())

		for i, v := range []string{"back", "future", "pop", "yakumi", "socks"} {
			if i%2 == 0 {
				list.PushBack(v)
			} else {
				list.PushFront(v)
			}
		}

		require.Equal(t, 6, list.Len())

		list.MoveToFront(list.Back())
		list.MoveToFront(list.Back())

		var wordsList []string
		for i := list.Front(); i != nil; i = i.Next {
			wordsList = append(wordsList, i.Value.(string))
		}

		require.Equal(t, []string{"pop", "socks", "yakumi", "future", "one", "back"}, wordsList)

	})
}
