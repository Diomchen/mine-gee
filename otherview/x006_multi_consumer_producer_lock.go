package otherview

import "sync"

type ThreadSafeQueue[T comparable] struct {
	items    []T
	capacity int
	mu       sync.Mutex
	notEmpty *sync.Cond
	notFull  *sync.Cond
}
