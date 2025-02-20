package maps_mine

import "container/list"

type K interface{}
type V int

type Cache struct {
	Key K
	Val V
}

type LRUCache struct {
	size  int
	list  *list.List
	cache map[K]*list.Element
}

func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		size:  size,
		list:  list.New(),
		cache: make(map[K]*list.Element),
	}
}

func (lru *LRUCache) Get(key K) V {
	if e, ok := lru.cache[key]; ok {
		v := e.Value.(*Cache).Val
		lru.list.MoveToFront(e)
		return v

	}
	return V(-1)
}

func (lru *LRUCache) Put(key K, val V) {
	if e, ok := lru.cache[key]; ok {
		e.Value.(*Cache).Val = val
		lru.list.MoveToFront(e)
	} else {
		if lru.list.Len() == lru.size {
			backE := lru.list.Back()
			if backE != nil {
				lru.list.Remove(backE)
				delete(lru.cache, backE.Value.(*Cache).Key)
			}
		}
	}
	newCache := &Cache{
		Key: key,
		Val: val,
	}
	newElement := lru.list.PushFront(newCache)
	lru.cache[key] = newElement
}

func (lru *LRUCache) RecentlyUsed() V {
	if lru.list.Len() == 0 {
		return V(-1)
	}
	return lru.list.Front().Value.(*Cache).Val
}

func (lru *LRUCache) Remove(key K) {
	if e, ok := lru.cache[key]; ok {
		lru.list.Remove(e)
		delete(lru.cache, key)
	}
}

func (lru *LRUCache) Clear() {
	lru.list.Init()
	lru.cache = make(map[K]*list.Element)
}

func (lru *LRUCache) Print() {
	for e := lru.list.Front(); e != nil; e = e.Next() {
		print(e.Value.(*Cache).Key, " ", e.Value.(*Cache).Val, "\n")
	}
}
