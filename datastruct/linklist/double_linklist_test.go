package linklist

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestDouble(t *testing.T) {
	head := NewDoubleLinkList(0)

	head.InsertHead(-1)
	head.InsertTail(1)
	head.InsertTail(2)
	head.InsertHead(-2)
	head.InsertTail(3)

	head.Print()

	popHead, _ := head.PopHead()
	popTail, _ := head.PopTail()
	assert.Equal(t, popHead, -2)
	assert.Equal(t, popTail, 3)

}
