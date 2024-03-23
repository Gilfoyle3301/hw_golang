package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	if _, ok := l.items[key]; ok {
		update := l.items[key]
		update.Value = value
		l.queue.MoveToFront(update)
		return true
	}
	if l.capacity == l.queue.Len() {
		l.queue.Remove(l.queue.Back())
		delete(l.items, l.getKey(l.queue.Back()))

	}
	l.items[key] = l.queue.PushFront(value)
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if v, ok := l.items[key]; ok {
		l.queue.MoveToFront(l.items[key])
		return v.Value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.queue = new(list)
	l.items = make(map[Key]*ListItem, l.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (l *lruCache) getKey(value *ListItem) Key {
	for key := range l.items {
		if value == l.items[key] {
			return Key(key)
		}
	}
	return Key("")

}
