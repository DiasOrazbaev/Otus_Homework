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
	if _, status := l.items[key]; status {
		node := l.items[key]
		node.Value = value
		l.queue.MoveToFront(node)
		return true
	} else if l.queue.Len()+1 == l.capacity {
		l.queue.Back().Value = nil
		l.queue.Remove(l.queue.Back())
	}
	node := l.queue.PushFront(value)
	l.items[key] = node
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	node, ok := l.items[key]
	if !ok || node.Value == nil {
		return nil, false
	}
	l.queue.MoveToFront(node)
	return node.Value, true
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
