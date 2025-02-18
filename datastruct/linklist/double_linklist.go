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

func NewDoubleLinkList(data interface{}) *Node {
	return &Node{nil, nil, data}
}

// 在最前面插入数值
func (n *Node) InsertHead(data interface{}) error {
	if n.prev != nil {
		return errors.New("Current node is not the head of the list.")
	}
	newNode := &Node{nil, n, data}
	n.prev = newNode
	return nil
}

// 在最后面插入数值
func (n *Node) InsertTail(data interface{}) error {
	for {
		if n.next == nil {
			break
		}
		n = n.next
	}

	newNode := &Node{n, nil, data}
	n.next = newNode
	return nil
}

// 取出最前面数值
func (n *Node) PopHead() (*Node, error) {
	if n == nil {
		return nil, errors.New("List is empty.")
	}
	if n.prev == nil {
		popNode := n
		n = n.next
		n.prev = nil
		popNode.next = nil
		return popNode, nil
	}
	return nil, nil
}

// 取出最后面数值
func (n *Node) PopTail() (*Node, error) {
	if n == nil {
		return nil, errors.New("List is empty.")
	}
	for {
		if n.next == nil {
			break
		}
		n = n.next
	}

	popNode := n
	n = n.prev
	n.next = nil
	popNode.prev = nil
	return popNode, nil
}

func (n *Node) Print() {
	if n == nil {
		fmt.Println("List is empty.")
		return
	}
	for {
		if n.next == nil {
			break
		}
		fmt.Printf("%v ", n.data)
		n = n.next
	}
	fmt.Printf("%v\n", n.data)
}
