package cache

import (
	"container/list"
)

type elem[K comparable, V any] struct {
	key K
	val V
}
type LRUCache[K comparable, V any] struct {
	data     map[K]*list.Element
	list     *list.List
	capacity int
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		list:     list.New(),
		data:     make(map[K]*list.Element),
	}
}

func (l *LRUCache[K, V]) Get(key K) (V, bool) {
	node := l.data[key]
	if node == nil {
		var v V
		return v, false
	}
	l.list.MoveToFront(node)
	return node.Value.(elem[K, V]).val, true
}

func (l *LRUCache[K, V]) Put(key K, value V) {
	if node, ok := l.data[key]; ok {
		node.Value = elem[K, V]{key, value}
		l.list.MoveToFront(node)
		return
	}

	l.data[key] = l.list.PushFront(elem[K, V]{key, value})
	if len(l.data) > l.capacity {
		oldest := l.list.Back()
		l.list.Remove(oldest)
		delete(l.data, oldest.Value.(elem[K, V]).key)
	}
}
