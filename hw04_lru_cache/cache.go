package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type cacheItem struct {
	key   Key
	value interface{}
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (lruC *lruCache) Set(key Key, value interface{}) bool {
	if _, ok := lruC.items[key]; ok {
		lruC.items[key].Value = cacheItem{key, value}
		lruC.queue.MoveToFront(lruC.items[key])
		return true
	}
	if lruC.queue.Len() == lruC.capacity {
		delete(lruC.items, lruC.queue.Back().Value.(cacheItem).key)
		lruC.queue.Remove(lruC.queue.Back())
	}
	lruC.items[key] = lruC.queue.PushFront(cacheItem{key, value})
	return false
}

func (lruC *lruCache) Get(key Key) (interface{}, bool) {
	if value, ok := lruC.items[key]; ok {
		lruC.queue.MoveToFront(lruC.items[key])
		return value.Value.(cacheItem).value, ok
	}
	return nil, false
}

func (lruC *lruCache) Clear() {
	lruC.items = make(map[Key]*ListItem, lruC.capacity)
	lruC.queue = NewList()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
