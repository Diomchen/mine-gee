package otherview

import "fmt"

type MessageQueue[T comparable] struct {
	queue chan T
}

func NewMessageQueue[T comparable](cap int) *MessageQueue[T] {
	return &MessageQueue[T]{
		queue: make(chan T, cap),
	}
}

func (mq *MessageQueue[T]) Enqueue(msg T) error {
	select {
	case mq.queue <- msg:
		return nil
	default:
		return fmt.Errorf("queue is full")
	}
}

func (mq *MessageQueue[T]) Dequeue() (T, error) {
	select {
	case msg := <-mq.queue:
		return msg, nil
	default:
		var empty T
		return empty, fmt.Errorf("queue is empty")
	}
}

func (mq *MessageQueue[T]) Close() {
	close(mq.queue)
}

func (mq *MessageQueue[T]) EnqueueBlocking(msg T) {
	mq.queue <- msg
}

func (mq *MessageQueue[T]) DequeueBlocking() T {
	return <-mq.queue
}
