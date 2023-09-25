package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRUCache(t *testing.T) {
	l := NewLRUCache(2)
	l.Put(2, 1)
	l.Put(1, 1)
	l.Put(2, 3)
	l.Put(4, 1)
	_, ok := l.Get(1)
	assert.Equal(t, false, ok, "LRUCache.Get(%v)")
	v, _ := l.Get(2)
	assert.Equal(t, 3, v, "LRUCache.Get(%v)")
}
