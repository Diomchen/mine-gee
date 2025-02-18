package linklist

import (
	"errors"
	"fmt"
)

type Node struct {
	prev *Node
	next *Node
	data interface{}
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewDoubleLinkList() *LinkedList {
	return &LinkedList{}
}

// 在最前面插入数值
func (l *LinkedList) InsertHead(data interface{}) error {
	newNode := &Node{nil, l.head, data}
	if l.head != nil {
		l.head.prev = newNode
	} else {
		l.tail = newNode
	}
	l.head = newNode
	l.size++
	return nil
}

// 在最后面插入数值
func (l *LinkedList) InsertTail(data interface{}) error {
	newNode := &Node{l.tail, nil, data}
	if l.tail != nil {
		l.tail.next = newNode
	} else {
		l.head = newNode
	}
	l.tail = newNode
	l.size++
	return nil
}

// 取出最前面数值
func (l *LinkedList) PopHead() (*Node, error) {
	if l.head == nil {
		return nil, errors.New("List is empty.")
	}
	popNode := l.head
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	} else {
		l.tail = nil
	}
	l.size--
	return popNode, nil
}

// 取出最后面数值
func (l *LinkedList) PopTail() (*Node, error) {
	if l.tail == nil {
		return nil, errors.New("List is empty.")
	}
	popNode := l.tail
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = nil
	}
	l.size--
	return popNode, nil
}

func (l *LinkedList) Print() {
	if l.head == nil {
		fmt.Println("List is empty.")
		return
	}

	next := l.head
	fmt.Println()
	for next != nil {
		fmt.Printf("%v ", next.data)
		next = next.next
	}
	fmt.Printf("\n double linklist size is %v", l.size)
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *LinkedList) GetNode(index int) (*Node, error) {
	if index < 0 || index >= l.size {
		return nil, errors.New("Index out of range.")
	}
	counter := 0
	next := l.head
	for next != nil {
		if counter == index {
			break
		}
		next = next.next
		counter++
	}
	return next, nil
}
