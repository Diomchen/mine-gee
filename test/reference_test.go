package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func modifyValue(x int) {
	x = x + 1
}

func modifyValuePtr(x *int) {
	*x = *x + 1
}

func TestReference(t *testing.T) {
	a := 5
	modifyValue(a)
	fmt.Println(a) // 输出 5
	assert.Equal(t, 5, a)

	b := 5
	modifyValuePtr(&b)
	fmt.Println(b) // 输出 6
	assert.Equal(t, 6, b)
}
