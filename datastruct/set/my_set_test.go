package set

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	num_s := NewSet[int]()
	num_s.Add(1)
	num_s.Add(2)
	num_s.Add(3)
	assert.Equal(t, true, num_s.Contains(1))
	num_s.Add(3)
	assert.Equal(t, 3, num_s.Size())
	num_s.Print()
	assert.Equal(t, 3, num_s.Size())
	num_s.Remove(2)
	fmt.Println()
	assert.Equal(t, 2, num_s.Size())
	num_s.Print()
}
