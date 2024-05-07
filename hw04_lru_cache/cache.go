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

type itemsCache struct {
	key   Key
	value interface{}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	if _, ok := l.items[key]; ok {
		l.items[key].Value = itemsCache{key, value}
		l.queue.MoveToFront(l.items[key])
		return true
	}
	if l.capacity == l.queue.Len() {
		l.queue.Remove(l.queue.Back())
		delete(l.items, l.queue.Back().Value.(itemsCache).key)
	}

	l.items[key] = l.queue.PushFront(itemsCache{key, value})
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if v, ok := l.items[key]; ok {
		l.queue.MoveToFront(l.items[key])
		return v.Value.(itemsCache).value, true
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
