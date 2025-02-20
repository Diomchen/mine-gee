package maps_mine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	lru := NewLRUCache(5)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)
	lru.Put(5, 5)
	lru.Print()

	assert.Equal(t, 5, lru.list.Len())
	assert.Equal(t, V(3), lru.Get(3))
	assert.Equal(t, V(3), lru.RecentlyUsed())
	lru.Put(3, 30)
	assert.Equal(t, V(30), lru.RecentlyUsed())
	lru.Remove(5)
	assert.Equal(t, V(-1), lru.Get(5))
	lru.Clear()
	assert.Equal(t, 0, lru.list.Len())
}
