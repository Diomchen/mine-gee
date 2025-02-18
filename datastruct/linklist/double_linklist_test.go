package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDouble(t *testing.T) {
	dll := NewDoubleLinkList()

	dll.InsertHead(-1)
	dll.InsertTail(1)
	dll.InsertTail(2)
	dll.InsertHead(-2)
	dll.InsertTail(3)

	// -2 -1 1 2 3
	dll.Print()

	popHead, _ := dll.PopHead()
	popTail, _ := dll.PopTail()
	assert.Equal(t, popHead.data, -2)
	assert.Equal(t, popTail.data, 3)

	// -1 1 2
	dll.Print()

	indexNode, _ := dll.GetNode(2)
	assert.Equal(t, indexNode.data, 2)

	indexErrNode, _ := dll.GetNode(3)
	assert.Nil(t, indexErrNode)

	assert.Equal(t, dll.Size(), 3)
	assert.Equal(t, dll.IsEmpty(), false)

	dll.Clear()
	assert.Equal(t, dll.Size(), 0)
	assert.Equal(t, dll.IsEmpty(), true)

}
