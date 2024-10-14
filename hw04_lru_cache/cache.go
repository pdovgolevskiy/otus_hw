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

func (lruC *lruCache) Set(key Key, value interface{}) bool {
	if _, ok := lruC.items[key]; ok {
		lruC.items[key].Value = value
		lruC.queue.MoveToFront(lruC.items[key])
		return true
	}
	lruC.items[key] = lruC.queue.PushFront(value)
	if lruC.queue.Len() > lruC.capacity {
		delete(lruC.items, lruC.queue.Back().Value.(Key))
		lruC.queue.Remove(lruC.queue.Back())
	}
	return false
}

func (lruC *lruCache) Get(key Key) (interface{}, bool) {
	if value, ok := lruC.items[key]; ok {
		return value.Value, ok
	}
	return nil, false
}

func (lruC *lruCache) Clear() {
	lruC.items = make(map[Key]*ListItem)
	lruC.queue = NewList()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
