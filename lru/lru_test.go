package lru

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
	assert.Equal(t, -1, l.Get(1), "LRUCache.Get(%v)")
	assert.Equal(t, 3, l.Get(2), "LRUCache.Get(%v)")
}
