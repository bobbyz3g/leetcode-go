package cache

import (
	"container/list"
)

type elem struct {
	key int
	val int
}
type LRUCache struct {
	data     map[int]*list.Element
	list     *list.List
	capacity int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		list:     list.New(),
		data:     make(map[int]*list.Element),
	}
}
func (l *LRUCache) Get(key int) (int, bool) {
	node := l.data[key]
	if node == nil {
		return 0, false
	}
	l.list.MoveToFront(node)
	return node.Value.(elem).val, true
}

func (l *LRUCache) Put(key, value int) {
	if node, ok := l.data[key]; ok {
		node.Value = elem{key, value}
		l.list.MoveToFront(node)
		return
	}

	l.data[key] = l.list.PushFront(elem{key, value})
	if len(l.data) > l.capacity {
		oldest := l.list.Back()
		l.list.Remove(oldest)
		delete(l.data, oldest.Value.(elem).key)
	}
}
